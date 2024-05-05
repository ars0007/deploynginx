package utils

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const DeploymentName = "nginx-deployment"

func getPath() string {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}

	return filepath.Join(home, ".kube", "config")
}

func ConnectToK8s(configPath string) *kubernetes.Clientset {
	path := configPath
	if configPath == "" {
		path = getPath()
	}

	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		panic("failed to create K8s config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("Failed to create K8s clientset")
	}

	return clientset
}
