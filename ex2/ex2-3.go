package ex2

import (
	"context"
	"fmt"
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
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

	name := "nginx-deployment"
	deploymentClient := client.AppsV1().Deployments(apiv1.NamespaceDefault)
	deleteDeployment(deploymentClient, name)
}

func deleteDeployment(deploymentClient appsv1.DeploymentInterface, name string) {
	err := deploymentClient.Delete(context.Background(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Fatal("Error in deleting the deployment", err)
	} else {
		fmt.Println("Succesfully deleted the deployment nginx-deployment")
	}
}

// If you want to list all the deployments in the default namespace.
func getDeploymentList(deploymentClient appsv1.DeploymentInterface) {
	deploymentList, err := deploymentClient.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error in getting the deploymente in the deafult namespace", err)
	}

	if len(deploymentList.Items) == 0 {
		log.Fatal("No deployments in the default namespace")
	}

	for _, deployment := range deploymentList.Items {
		fmt.Println("Deployment ", deployment.Name)
	}
}
