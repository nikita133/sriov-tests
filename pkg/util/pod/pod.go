package pod

import (
	"bytes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/utils/pointer"
	"os"
	"strings"

	testclient "github.com/openshift/sriov-tests/pkg/util/client"
	"github.com/openshift/sriov-tests/pkg/util/namespaces"
)

func getDefinition() *corev1.Pod {
	podName := "testpod" + rand.String(12)
	podObject := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: podName,
			Namespace: namespaces.Test},
		Spec: corev1.PodSpec{
			TerminationGracePeriodSeconds: pointer.Int64Ptr(0),
			Containers: []corev1.Container{{Name: "test",
				Image:   "quay.io/schseba/utility-container:latest",
				Command: []string{"/bin/bash", "-c", "sleep INF"}}}}}

	return podObject
}

func DefineWithNetworks(networks []string) *corev1.Pod {
	podObject := getDefinition()
	podObject.Annotations = map[string]string{"k8s.v1.cni.cncf.io/networks": strings.Join(networks, ",")}

	return podObject
}

func DefineWithHostNetwork() *corev1.Pod {
	podObject := getDefinition()
	podObject.Spec.HostNetwork = true

	return podObject
}

// ExecCommand runs command in the pod and returns buffer output
func ExecCommand(cs *testclient.ClientSet, pod *corev1.Pod, command ...string) (string, string, error) {
	var buf, errbuf bytes.Buffer
	req := cs.CoreV1Interface.RESTClient().
		Post().
		Namespace(pod.Namespace).
		Resource("pods").
		Name(pod.Name).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: pod.Spec.Containers[0].Name,
			Command:   command,
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(cs.Config, "POST", req.URL())
	if err != nil {
		return buf.String(), errbuf.String(), err
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: &buf,
		Stderr: &errbuf,
		Tty:    true,
	})
	if err != nil {
		return buf.String(), errbuf.String(), err
	}

	return buf.String(), errbuf.String(), nil
}
