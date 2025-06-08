package watcher

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)


func TestLabelNamespace(t *testing.T) {
	testEnv := &envtest.Environment{} // define the test env variable object
	cfg, err := testEnv.Start() // start the test env, which will create a new Kubernetes cluster for testing purposes, cfg represents the kubeconfig
	if err != nil {
		t.Fatalf("Failed to start envtes: %v", err)
	}
	defer testEnv.Stop() // stop the test env after the test is done
	

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		t.Fatalf("Failed to create clientset: %v", err)
	} // create a new clientset for the test env

	ns := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "test-ns"}}
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Fialed to create namespace: %v", err)
	}

	labels := map[string]string{"foo": "bar"}
	err = LabelNamespace(clientset, ns.Name, labels)
	if err != nil {
		t.Fatalf("failed to NsLabeled: %v", err)
	}

	got, err := clientset.CoreV1().Namespaces().Get(context.TODO(), "test-ns", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get namespace: %v", err)
	}

	if got.Labels["foo"] != "bar" {
		t.Errorf("Expected label foo=bar, got %v:", got.Labels)
	}
}