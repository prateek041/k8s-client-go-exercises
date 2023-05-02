package ex1

import (
	"fmt"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	rules := clientcmd.NewDefaultClientConfigLoadingRules() // This returns some default loading rules for the kubeconfig file.
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		log.Fatal("Error finding the kubeconfig file")
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Error in building the clientset for K8s")
	}

	fmt.Println(clientSet)

}
