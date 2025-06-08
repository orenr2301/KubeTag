package watcher

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)


func Connector() (*kubernetes.Clientset, error) {
	
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())

	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset, nil
}



func HandleNamespaceEvent(eventType watch.EventType, namespace *v1.Namespace) {
    log.Printf("Handling event: %s for Namespace: %s\n", eventType, namespace.Name)

	ns := namespace
	if eventType == watch.Added {
		log.Printf("New Namesapce : %v, calling function to handle %v",ns, eventType)
	}
}

func WatchNS(clientset *kubernetes.Clientset, defaultLabels map[string]string) {

	watchInterface, err := clientset.CoreV1().Namespaces().Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for event := range watchInterface.ResultChan() {
        namespace, ok := event.Object.(*v1.Namespace)
        if !ok {
            continue
        }

		switch event.Type {
		case watch.Added:
			fmt.Printf("New Namesapce added: $s\n", namespace.Name)
			if len(defaultLabels) > 0 {
				err := LabelNamespace(clientset, namespace.Name, defaultLabels)
				if err != nil {
					fmt.Printf("Error labeling new namesapce %s: %v", namespace.Name, err)
				} else {
					fmt.Printf("Labels added to namesapce %s successfully. \n", namespace.Name)
				}
			}
		}

	}
	
}
