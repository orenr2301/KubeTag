package watcher

import (
	"context"
	"encoding/json"


	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)


type NamespaceConfig struct{
	Name string 			`json:"name"`
	Labels map[string]string `json:"labels"`
}

func LabelNamespace(clientset *kubernetes.Clientset, namespace string, labels map[string]string) error {


	var patch []map[string]interface{}
	for key, value := range labels {
		patch = append(patch, map[string]interface{}{
				"op": "add",
				"path": "/metadata/labels/" + key,
				"value": value,
			})
	}

	patchBytes, err := json.Marshal(patch)
	if err != nil {
		return err
	}

	_, err = clientset.CoreV1().Namespaces().Patch(context.TODO(), namespace, types.JSONPatchType, patchBytes, metav1.PatchOptions{})
	return err
	}
	
