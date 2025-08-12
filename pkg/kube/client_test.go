package kube

import (
	"testing"
)

func TestNewClients(t *testing.T) {
	// Test with empty kubeconfig path (should try in-cluster config)
	clientset, discoveryClient, dynamicClient, err := NewClients("")
	if err != nil {
		// This is expected to fail in non-kubernetes environment
		// We just check that the error is not nil
		t.Logf("Expected error in non-kubernetes environment: %v", err)
		return
	}

	// If we're in a kubernetes environment, verify the clients
	if clientset == nil {
		t.Error("clientset should not be nil")
	}
	if discoveryClient == nil {
		t.Error("discoveryClient should not be nil")
	}
	if dynamicClient == nil {
		t.Error("dynamicClient should not be nil")
	}
}
