package eathar

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//Containers are privileged not pods
type privileged struct {
	namespace string
	pod       string
	container string
}

func Hostnet(kubeconfig string) {

	clientset := connectToCluster(kubeconfig)
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	//Debugging command
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {

		if pod.Spec.HostNetwork {
			fmt.Printf("Pod %s is using Host networking\n", pod.Name)
		}
	}
}

func Hostpid(kubeconfig string) {
	clientset := connectToCluster(kubeconfig)
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	//Debugging command
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {

		if pod.Spec.HostPID {
			fmt.Printf("Pod %s is using Host PID\n", pod.Name)
		}
	}
}

func AllowPrivEsc(kubeconfig string) {
	clientset := connectToCluster(kubeconfig)
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	//Debugging command
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		for i, container := range pod.Spec.Containers {
			// Logic here is if there's no security context, or there is a security context and no mention of allow privilege escalation then the default is true
			// We don't catch the case of someone explicitly setting it to true, but that seems unlikely
			allowPrivilegeEscalation := (container.SecurityContext == nil) || (container.SecurityContext != nil && container.SecurityContext.AllowPrivilegeEscalation == nil)
			if allowPrivilegeEscalation {
				fmt.Printf("Pod: %s - Container %d : %s does not block privilege escalation\n", pod.Name, i+1, container.Name)
			}
		}
	}
}

func Privileged(kubeconfig string) {
	var privcont []privileged
	clientset := connectToCluster(kubeconfig)
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	//Debugging command
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			// if you try to check privileged for nil on it's own, it doesn't work you need to check security context too
			privileged_container := container.SecurityContext != nil && container.SecurityContext.Privileged != nil && *container.SecurityContext.Privileged
			if privileged_container {
				// So we create a new privileged struct from our matching container
				p := privileged{namespace: pod.Namespace, pod: pod.Name, container: container.Name}
				fmt.Printf("Namespace: %s - Pod: %s - Container  : %s is running as privileged \n", p.namespace, p.pod, p.container)
				//And we append it to our slice of all our privileged containers
				privcont = append(privcont, p)
			}
		}
	}
	// Just to prove our slice is working
	fmt.Printf("we have %d privileged containers\n", len(privcont))
}

// This is our function for connecting to the cluster
func connectToCluster(kubeconfig string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return clientset
}
