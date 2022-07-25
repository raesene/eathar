package eathar

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type finding struct {
	check     string
	namespace string
	pod       string
	container string
}

func Hostnet(kubeconfig string) {
	var hostnetcont []finding
	clientset := connectToCluster(kubeconfig)
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	//Debugging command
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {

		if pod.Spec.HostNetwork {
			// We set the container to blank as there's no container for this finding
			p := finding{check: "hostnet", namespace: pod.Namespace, pod: pod.Name, container: ""}
			//fmt.Printf("Namespace %s - Pod %s is using Host networking\n", p.namespace, p.pod)
			hostnetcont = append(hostnetcont, p)
		}
	}
	report(hostnetcont)
}

func Hostpid(kubeconfig string) {
	var hostpidcont []finding
	clientset := connectToCluster(kubeconfig)
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	//Debugging command
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {

		if pod.Spec.HostPID {
			p := finding{check: "hostpid", namespace: pod.Namespace, pod: pod.Name, container: ""}
			//fmt.Printf("Namespace %s - Pod %s is using Host PID\n", p.namespace, p.pod)
			hostpidcont = append(hostpidcont, p)
		}
	}
	report(hostpidcont)
}

func AllowPrivEsc(kubeconfig string) {
	var allowprivesccont []finding
	clientset := connectToCluster(kubeconfig)
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	//Debugging command
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			// Logic here is if there's no security context, or there is a security context and no mention of allow privilege escalation then the default is true
			// We don't catch the case of someone explicitly setting it to true, but that seems unlikely
			allowPrivilegeEscalation := (container.SecurityContext == nil) || (container.SecurityContext != nil && container.SecurityContext.AllowPrivilegeEscalation == nil)
			if allowPrivilegeEscalation {
				p := finding{check: "allowprivesc", namespace: pod.Namespace, pod: pod.Name, container: container.Name}
				//fmt.Printf("Namespace: %s - Pod: %s - Container: %s does not block privilege escalation\n", p.namespace, p.pod, p.container)
				allowprivesccont = append(allowprivesccont, p)
			}
		}
	}
	report(allowprivesccont)
}

func Privileged(kubeconfig string) {
	var privcont []finding
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
				p := finding{check: "privileged", namespace: pod.Namespace, pod: pod.Name, container: container.Name}
				//fmt.Printf("Namespace: %s - Pod: %s - Container  : %s is running as privileged \n", p.namespace, p.pod, p.container)
				//And we append it to our slice of all our privileged containers
				privcont = append(privcont, p)
			}
		}
	}
	// Just to prove our slice is working
	report(privcont)
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

func report(f []finding) {
	fmt.Printf("Findings for the %s check\n", f[0].check)
	for _, i := range f {
		if i.container == "" {
			fmt.Printf("namespace %s : pod %s\n", i.namespace, i.pod)
		} else {
			fmt.Printf("namespace %s : pod %s : container %s\n", i.namespace, i.pod, i.container)
		}
	}
}
