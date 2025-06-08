package watcher

import (
	"reflect"
	"testing"
)


func MergedLabels(defaultLabels, nsLabels map[string]string) map[string]string {
	merged := make(map[string]string)
	for k, v := range defaultLabels {
		merged[k] = v
	}

	for k, v := range nsLabels {
		merged[k] = v
	}

	return merged
}

func TestMergedLabels(t *testing.T) {
	defaultLabels := map[string]string{"env": "production", "team": "devops"}
	ns := map[string]string{"n": "3", "space": "test", "memebr": "one"}
	expectedOrder := map[string]string{"env": "production", "team": "devops", "n": "3", "space": "test", "memebr": "one"}
	
	
	result := MergedLabels(defaultLabels, ns)
	if !reflect.DeepEqual(result, expectedOrder) {
		t.Errorf("Expected %v, got %v", expectedOrder, result)
	}
}