package ex2

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getPodsInCluster() {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		log.Fatal("Error in creting the config file", err)
	}

	// creating a clinet to communicate with kubernetes API
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Error in creating the client")
	}

	podClient := client.CoreV1().Pods("kube-system")
	pods, err := podClient.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error in getting the podlist", err)
	}

	fmt.Println("Pods in the kube-system namespace kube-system")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
}
