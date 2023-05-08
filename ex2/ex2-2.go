package ex2

import (
	"context"
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/pointer"
)

func CreateDeployment() {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		log.Fatal("Error in finding the Kubeconfig", err)
	}

	clientSet, err := kubernetes.NewForConfig(config) // This is the client set for everything.
	if err != nil {
		log.Fatal("Error in creating a clientset", err)
	}

	deploymentClientSet := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault) // a clientset for deployment

	// try putting typeMeta to use as well
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "first",
					"tier": "backend",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "first",
						"tier": "backend",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "nginx",
							Image: "nginx:latest",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// now create the deployment
	fmt.Println("Creating the deployment")
	result, err := deploymentClientSet.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Fatal("Error in creating the deployment", err)
	}

	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}
