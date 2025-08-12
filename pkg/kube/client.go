package kube

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// NewClients creates Kubernetes clients for both in-cluster and out-of-cluster execution
func NewClients(kubeconfigPath string) (kubernetes.Interface, discovery.DiscoveryInterface, dynamic.Interface, error) {
	var config *rest.Config
	var err error

	if kubeconfigPath != "" {
		// Use provided kubeconfig file
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			return nil, nil, nil, err
		}
	} else {
		// Try in-cluster config first, fallback to default kubeconfig
		config, err = rest.InClusterConfig()
		if err != nil {
			// Fallback to default kubeconfig location
			config, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
			if err != nil {
				return nil, nil, nil, err
			}
		}
	}

	// Create clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	// Create discovery client
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	// Create dynamic client
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	return clientset, discoveryClient, dynamicClient, nil
}
