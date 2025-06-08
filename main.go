package main

import (
	
	"encoding/json"
	"fmt"
	"github.com/orenr2301/KubeTag/pkg/watcher"
	"os"

	
	
)


func main() {

	namespaceEnv := os.Getenv("NAMESPACES")
	var namespaces []watcher.NamespaceConfig
	if err := json.Unmarshal([]byte(namespaceEnv), &namespaces); err != nil {
		panic(err.Error())
	} 

	

	// Get Deafult Variables from envinronemnt vairables // need to handled the error in case there no mapping under that default vairables
	defdefaultLabelsEnv := os.Getenv("DEFAULT_LABELS") 	
	var defaultLabels map[string]string
	if err := json.Unmarshal([]byte(defdefaultLabelsEnv), &defaultLabels); err != nil {
		panic(err)
	}
		//Applying label to all namespaces


	clientset, err := watcher.Connector()
	if err != nil {
		fmt.Errorf("Couldnt connect cluster or getting apiserver: %v", err)

	}

	//Iterating over each namesapce and apply labels
	for _, ns := range namespaces {
		mergedLabels := make(map[string]string) //Mapping object to store the default labels 
		if len(defaultLabels) > 0 {
			for key, value := range defaultLabels {
				mergedLabels[key] = value //setting the default labels key value pairs
				// if ns.Labels == nil {
				// 	ns.Labels = make(map[string]string) //if labels are not set, create a new map
				// }
			}
		}
		for key, value := range ns.Labels {
			mergedLabels[key] = value 
		}

		err = watcher.LabelNamespace(clientset, ns.Name, mergedLabels)
		if err != nil {
			fmt.Printf("Error labeling namesapce %s: %v\n", ns.Name, err)
		} else {
			fmt. Printf("Namespace %s labels successfuly.\n", ns.Name)
		}
	}

	watcher.WatchNS(clientset, defaultLabels)
	
}
