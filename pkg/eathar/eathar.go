package eathar

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//This needs to be exported to work with the JSON marshalling
// omitempty thing is there as container won't always be relevant (e.g. hostPID)
type Finding struct {
	Check     string
	Namespace string
	Pod       string
	Container string `json:",omitempty"`
}

func Hostnet(kubeconfig string, jsonrep bool) {
	var hostnetcont []Finding
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
			p := Finding{Check: "hostnet", Namespace: pod.Namespace, Pod: pod.Name, Container: ""}
			//fmt.Printf("Namespace %s - Pod %s is using Host networking\n", p.namespace, p.pod)
			hostnetcont = append(hostnetcont, p)
		}
	}
	report(hostnetcont, jsonrep)
}

func Hostpid(kubeconfig string, jsonrep bool) {
	var hostpidcont []Finding
	clientset := connectToCluster(kubeconfig)
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	//Debugging command
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {

		if pod.Spec.HostPID {
			p := Finding{Check: "hostpid", Namespace: pod.Namespace, Pod: pod.Name, Container: ""}
			//fmt.Printf("Namespace %s - Pod %s is using Host PID\n", p.namespace, p.pod)
			hostpidcont = append(hostpidcont, p)
		}
	}
	report(hostpidcont, jsonrep)
}

func AllowPrivEsc(kubeconfig string, jsonrep bool) {
	var allowprivesccont []Finding
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
				p := Finding{Check: "allowprivesc", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name}
				//fmt.Printf("Namespace: %s - Pod: %s - Container: %s does not block privilege escalation\n", p.namespace, p.pod, p.container)
				allowprivesccont = append(allowprivesccont, p)
			}
		}
	}
	report(allowprivesccont, jsonrep)
}

func Privileged(kubeconfig string, jsonrep bool) {
	var privcont []Finding
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
				p := Finding{Check: "privileged", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name}
				//fmt.Printf("Namespace: %s - Pod: %s - Container  : %s is running as privileged \n", p.namespace, p.pod, p.container)
				//And we append it to our slice of all our privileged containers
				privcont = append(privcont, p)
			}
		}
	}
	// Just to prove our slice is working
	report(privcont, jsonrep)
}

func AddedCapabilities(kubeconfig string, jsonrep bool) {
	var capadded []Finding
	clientset := connectToCluster(kubeconfig)
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			cap_added := container.SecurityContext != nil && container.SecurityContext.Capabilities.Add != nil
			if cap_added {
				p := Finding{Check: "Added Capabilities", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name}
				capadded = append(capadded, p)
			}
		}
	}
	report(capadded, jsonrep)
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

func report(f []Finding, jsonrep bool) {
	if !jsonrep {
		fmt.Printf("Findings for the %s check\n", f[0].Check)
		for _, i := range f {
			if i.Container == "" {
				fmt.Printf("namespace %s : pod %s\n", i.Namespace, i.Pod)
			} else {
				fmt.Printf("namespace %s : pod %s : container %s\n", i.Namespace, i.Pod, i.Container)
			}
		}
	} else {

		js, err := json.MarshalIndent(f, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(js))
	}

}
