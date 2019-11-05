package operator

import (
	goctx "context"
	// "encoding/json"
	"fmt"
	// "reflect"
	"strings"
	// "testing"
	// "time"

	// dptypes "github.com/intel/sriov-network-device-plugin/pkg/types"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	// "github.com/operator-framework/operator-sdk/pkg/test/e2eutil"
	appsv1 "k8s.io/api/apps/v1"
	// corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/runtime"
	// "k8s.io/apimachinery/pkg/types"
	// "k8s.io/apimachinery/pkg/util/wait"
	// dynclient "sigs.k8s.io/controller-runtime/pkg/client"

	// "github.com/openshift/sriov-network-operator/pkg/apis"
	netattdefv1 "github.com/openshift/sriov-network-operator/pkg/apis/k8s/v1"
	sriovnetworkv1 "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/openshift/sriov-tests/pkg/util"
)

var _ = Describe("Operator", func() {
	BeforeEach(func() {
		// get global framework variables
		f := framework.Global
		// wait for sriov-network-operator to be ready
		deploy := &appsv1.Deployment{}
		err := WaitForNamespacedObject(deploy, f.Client, namespace, "sriov-network-operator", RetryInterval, Timeout)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("DaemonSets up by default", func() {
		DescribeTable("should be possible to create daemonSet",
			func(dsName string) {
				// get global framework variables
				f := framework.Global
				// wait for sriov-network-operator to be ready
				daemonSet := &appsv1.DaemonSet{}
				err := WaitForNamespacedObject(daemonSet, f.Client, namespace, dsName, RetryInterval, Timeout)
				Expect(err).NotTo(HaveOccurred())
			},
			Entry("operator-webhook", "operator-webhook"),
			Entry("network-resources-injector", "network-resources-injector"),
			Entry("sriov-network-config-daemon", "sriov-network-config-daemon"),
		)
	})

	Describe("with SriovNetwork", func() {
		specs := map[string]sriovnetworkv1.SriovNetworkSpec{
			"test-0": {
				ResourceName: "resource_1",
				IPAM:         `{"type":"host-local","subnet":"10.56.217.0/24","rangeStart":"10.56.217.171","rangeEnd":"10.56.217.181","routes":[{"dst":"0.0.0.0/0"}],"gateway":"10.56.217.1"}`,
				Vlan:         100,
			},
			"test-1": {
				ResourceName:     "resource_1",
				IPAM:             `{"type":"host-local","subnet":"10.56.217.0/24","rangeStart":"10.56.217.171","rangeEnd":"10.56.217.181","routes":[{"dst":"0.0.0.0/0"}],"gateway":"10.56.217.1"}`,
				NetworkNamespace: "default",
			},
			"test-2": {
				ResourceName: "resource_1",
				IPAM:         `{"type":"host-local","subnet":"10.56.217.0/24","rangeStart":"10.56.217.171","rangeEnd":"10.56.217.181","routes":[{"dst":"0.0.0.0/0"}],"gateway":"10.56.217.1"}`,
				SpoofChk:     "on",
			},
			"test-3": {
				ResourceName: "resource_1",
				IPAM:         `{"type":"host-local","subnet":"10.56.217.0/24","rangeStart":"10.56.217.171","rangeEnd":"10.56.217.181","routes":[{"dst":"0.0.0.0/0"}],"gateway":"10.56.217.1"}`,
				Trust:        "on",
			},
		}
		sriovnets := GenerateSriovNetworkCRs(namespace, specs)
		DescribeTable("should be possible to create net-att-def",
			func(cr *sriovnetworkv1.SriovNetwork) {
				var err error
				spoofchk := ""
				trust := ""
				state := ""

				if cr.Spec.Trust == "on" {
					trust = `"trust":"on",`
				} else if cr.Spec.Trust == "off" {
					trust = `"trust":"off",`
				}

				if cr.Spec.SpoofChk == "on" {
					spoofchk = `"spoofchk":"on",`
				} else if cr.Spec.SpoofChk == "off" {
					spoofchk = `"spoofchk":"off",`
				}

				if cr.Spec.LinkState == "auto" {
					state = `"link_state":"auto",`
				} else if cr.Spec.LinkState == "enable" {
					state = `"link_state":"enable",`
				} else if cr.Spec.LinkState == "disable" {
					state = `"link_state":"disable",`
				}

				vlanQoS := cr.Spec.VlanQoS

				expect := fmt.Sprintf(`{ "cniVersion":"0.3.1", "name":"sriov-net", "type":"sriov", "vlan":%d,%s%s%s"vlanQoS":%d,"ipam":%s }`, cr.Spec.Vlan, spoofchk, trust, state, vlanQoS, cr.Spec.IPAM)

				// get global framework variables
				f := framework.Global
				err = f.Client.Create(goctx.TODO(), cr, &framework.CleanupOptions{TestContext: &oprctx, Timeout: ApiTimeout, RetryInterval: RetryInterval})
				Expect(err).NotTo(HaveOccurred())
				ns := namespace
				if cr.Spec.NetworkNamespace != "" {
					ns = cr.Spec.NetworkNamespace
				}
				netAttDef := &netattdefv1.NetworkAttachmentDefinition{}
				err = WaitForNamespacedObject(netAttDef, f.Client, ns, cr.GetName(), RetryInterval, Timeout)
				Expect(err).NotTo(HaveOccurred())
				anno := netAttDef.GetAnnotations()

				Expect(anno["k8s.v1.cni.cncf.io/resourceName"]).To(Equal("openshift.io/" + cr.Spec.ResourceName))
				Expect(strings.TrimSpace(netAttDef.Spec.Config)).To(Equal(expect))

			},
			Entry("with vlan flag", sriovnets[0]),
			Entry("with networkNamespace flag", sriovnets[1]),
			Entry("with SpoofChk flag on", sriovnets[2]),
			Entry("with Trust flag on", sriovnets[3]),
		)
	})
})
