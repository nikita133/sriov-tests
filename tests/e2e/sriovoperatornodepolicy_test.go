package e2e

import (
	goctx "context"
	// "encoding/json"
	"fmt"
	// "reflect"
	// "strings"
	// "testing"
	"time"

	// dptypes "github.com/intel/sriov-network-device-plugin/pkg/types"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	// "github.com/operator-framework/operator-sdk/pkg/test/e2eutil"
	// admv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/runtime"
	// "k8s.io/apimachinery/pkg/types"
	// "k8s.io/apimachinery/pkg/util/wait"
	dynclient "sigs.k8s.io/controller-runtime/pkg/client"

	// "github.com/openshift/sriov-network-operator/pkg/apis"
	// netattdefv1 "github.com/openshift/sriov-network-operator/pkg/apis/k8s/v1"
	sriovnetworkv1 "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/openshift/sriov-tests/pkg/util"
)

var _ = Describe("Operator", func() {

	// BeforeEach(func() {
	// 	// get global framework variables
	// 	f := framework.Global
	// 	var err error

	// 	// Turn off Operator Webhook
	// 	// config := &sriovnetworkv1.SriovOperatorConfig{}
	// 	// err = WaitForNamespacedObject(config, f.Client, namespace, "default", RetryInterval, Timeout)
	// 	// Expect(err).NotTo(HaveOccurred())

	// 	// *config.Spec.EnableInjector = false
	// 	// err = f.Client.Update(goctx.TODO(), config)
	// 	// Expect(err).NotTo(HaveOccurred())

	// 	// daemonSet := &appsv1.DaemonSet{}
	// 	// err = WaitForNamespacedObjectDeleted(daemonSet, f.Client, namespace, "network-resources-injector", RetryInterval, Timeout)
	// 	// Expect(err).NotTo(HaveOccurred())

	// 	// mutateCfg := &admv1beta1.MutatingWebhookConfiguration{}
	// 	// err = WaitForNamespacedObjectDeleted(mutateCfg, f.Client, namespace, "network-resources-injector-config", RetryInterval, Timeout)
	// 	// Expect(err).NotTo(HaveOccurred())
	// })

	Context("with single policy", func() {
		policy := &sriovnetworkv1.SriovNetworkNodePolicy{
			TypeMeta: metav1.TypeMeta{
				Kind:       "SriovNetworkNodePolicy",
				APIVersion: "sriovnetwork.openshift.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "policy-1",
				Namespace: namespace,
			},
			Spec: sriovnetworkv1.SriovNetworkNodePolicySpec{
				ResourceName: "resource_1",
				NodeSelector: map[string]string{
					"feature.node.kubernetes.io/network-sriov.capable": "true",
				},
				Priority: 99,
				Mtu:      9000,
				NumVfs:   6,
				NicSelector: sriovnetworkv1.SriovNetworkNicSelector{
					Vendor:      "8086",
					RootDevices: []string{"0000:86:00.1"},
				},
				DeviceType: "vfio-pci",
			},
		}

		It("should config sriov", func() {
			// get global framework variables
			f := framework.Global
			var err error

			By("generate the config for device plugin")
			err = f.Client.Create(goctx.TODO(), policy, &framework.CleanupOptions{TestContext: &oprctx, Timeout: ApiTimeout, RetryInterval: RetryInterval})
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(3 * time.Second)
			config := &corev1.ConfigMap{}
			err = WaitForNamespacedObject(config, f.Client, namespace, "device-plugin-config", RetryInterval, Timeout)
			Expect(err).NotTo(HaveOccurred())

			err = ValidateDevicePluginConfig(policy, config.Data["config.json"])

			By("provision the cni and device plugin daemonsets")
			cniDaemonSet := &appsv1.DaemonSet{}
			err = WaitForDaemonSetReady(cniDaemonSet, f.Client, namespace, "sriov-cni", RetryInterval, Timeout)
			Expect(err).NotTo(HaveOccurred())

			dpDaemonSet := &appsv1.DaemonSet{}
			err = WaitForDaemonSetReady(dpDaemonSet, f.Client, namespace, "sriov-device-plugin", RetryInterval, Timeout)
			Expect(err).NotTo(HaveOccurred())

			By("update the spec of SriovNetworkNodeState CR")
			nodeList := &corev1.NodeList{}
			lo := &dynclient.MatchingLabels{
				"feature.node.kubernetes.io/network-sriov.capable": "true",
			}
			err = f.Client.List(goctx.TODO(), nodeList, lo)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(nodeList.Items)).To(Equal(1))

			name := nodeList.Items[0].GetName()

			nodeState := &sriovnetworkv1.SriovNetworkNodeState{}
			err = WaitForSriovNetworkNodeStateReady(nodeState, policy, f.Client, namespace, name, RetryInterval, Timeout)
			Expect(err).NotTo(HaveOccurred())

			fmt.Fprintf(GinkgoWriter, "nodeState: %v\n\n", nodeState)

			found := false
			for _, address := range policy.Spec.NicSelector.RootDevices {
				for _, iface := range nodeState.Spec.Interfaces {
					if iface.PciAddress == address {
						found = true
						Expect(iface.NumVfs).To(Equal(policy.Spec.NumVfs))
						Expect(iface.Mtu).To(Equal(policy.Spec.Mtu))
						Expect(iface.DeviceType).To(Equal(policy.Spec.DeviceType))
					}
				}
			}
			Expect(found).To(BeTrue())

			By("update the status of SriovNetworkNodeState CR")
			found = false
			for _, address := range policy.Spec.NicSelector.RootDevices {
				for _, iface := range nodeState.Status.Interfaces {
					if iface.PciAddress == address {
						found = true
						Expect(iface.NumVfs).To(Equal(policy.Spec.NumVfs))
						Expect(iface.Mtu).To(Equal(policy.Spec.Mtu))
						Expect(len(iface.VFs)).To(Equal(policy.Spec.NumVfs))
						for _, vf := range iface.VFs {
							if policy.Spec.DeviceType == "netdevice" || policy.Spec.DeviceType == ""{
								Expect(vf.Mtu).To(Equal(policy.Spec.Mtu))
							}
							Expect(vf.Driver).To(Equal(policy.Spec.DeviceType))
						}
						break
					}
				}
			}
			Expect(found).To(BeTrue())
		})
	})

})
