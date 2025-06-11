package watcher

import (
	"context"
	"encoding/json"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)


type NamespaceConfig struct{
	Name string 			`json:"name"`
	Labels map[string]string `json:"labels"`
}

func LabelNamespace(clientset *kubernetes.Clientset, namespace string, labels map[string]string) error {

	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")

	ns, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if ns.Annotations == nil || ns.Annotations["app.kubetag.opt/managed"] != "true" {
		log.Printf("%s: Namespace is managed by Helm standard pre-built labels. e.g app.kubernetes.io/managed-by: Helm or helm.sh.\nTo allow KubeTag to manage labels, add annotation app.kubetag.opt/managed: 'true'", formattedTime)
		return nil
	}

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


