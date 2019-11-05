package util

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
	// appsv1 "k8s.io/api/apps/v1"
	// corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	// dynclient "sigs.k8s.io/controller-runtime/pkg/client"

	// "github.com/openshift/sriov-network-operator/pkg/apis"
	// netattdefv1 "github.com/openshift/sriov-network-operator/pkg/apis/k8s/v1"
	sriovnetworkv1 "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1"
)

var (
	RetryInterval        = time.Second * 1
	ApiTimeout           = time.Second * 10
	Timeout              = time.Second * 60
	CleanupRetryInterval = time.Second * 1
	CleanupTimeout       = time.Second * 5
)

func WaitForNamespacedObject(obj runtime.Object, client framework.FrameworkClient, namespace, name string, retryInterval, timeout time.Duration) error {

	err := wait.PollImmediate(retryInterval, timeout, func() (done bool, err error) {
		ctx, cancel := goctx.WithTimeout(goctx.Background(), ApiTimeout)
		defer cancel()
		err = client.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, obj)
		if err != nil {
			if errors.IsNotFound(err) {
				return false, nil
			}
			return false, err
		}
		return true, nil
	})
	if err != nil {
		fmt.Printf("failed to wait for obj %s/%s to exist: %v", namespace, name, err)
		return err
	}

	return nil
}

func NotExsitNamespacedObject(obj runtime.Object, client framework.FrameworkClient, namespace, name string, retryInterval, timeout time.Duration) error {

	err := wait.PollImmediate(retryInterval, timeout, func() (done bool, err error) {
		ctx, cancel := goctx.WithTimeout(goctx.Background(), ApiTimeout)
		defer cancel()
		err = client.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, obj)
		if err != nil {
			if errors.IsNotFound(err) {
				return true, nil
			}
			return false, err
		}
		return false, nil
	})
	if err != nil {
		fmt.Printf("failed to wait for obj %s/%s to not exist: %v", namespace, name, err)
		return err
	}

	return nil
}

func GenerateSriovNetworkCRs(namespace string, specs map[string]sriovnetworkv1.SriovNetworkSpec) []*sriovnetworkv1.SriovNetwork {
	var crs []*sriovnetworkv1.SriovNetwork

	for k, v := range specs {
		crs = append(crs, &sriovnetworkv1.SriovNetwork{
			TypeMeta: metav1.TypeMeta{
				Kind:       "SriovNetwork",
				APIVersion: "sriovnetwork.openshift.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      k,
				Namespace: namespace,
			},
			Spec: v,
		})
	}
	return crs
}
