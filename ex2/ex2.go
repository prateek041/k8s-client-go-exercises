package ex2

import (
	"context"
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetPods() {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{}).ClientConfig()

	if err != nil {
		log.Fatal("Error in getting the Kubeconfig file")
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Error in creatin the clientset for k8s")
	}

	podName := "my-first-pod"
	namespace := "default"

	pod, err := clientSet.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", podName, namespace)
		} else {
			log.Fatalf("Error getting pod %s in namespace %s: %v", podName, namespace, err)
		}
	} else {
		fmt.Printf("Pod %s in namespace %s exists\n", podName, namespace)
	}
	_ = pod
}
