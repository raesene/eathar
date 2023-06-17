package eathar

/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/

import (
	"context"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func initKubeClient() (*kubernetes.Clientset, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		log.Printf("initKubeClient: failed creating ClientConfig with", err)
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("initKubeClient: failed creating Clientset with", err)
		return nil, err
	}
	return clientset, nil
}

func connectWithPods(options *pflag.FlagSet) *corev1.PodList {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}

	exclude, err := options.GetString("exclude")

	var excludeList []string
	if exclude != "" {
		excludeList = strings.Split(exclude, ",")
	}
	if err != nil {
		log.Print(err)
	}
	filteredPods := &corev1.PodList{}
	for _, pod := range pods.Items {
		if len(excludeList) > 0 {
			excluded := false
			for _, s := range excludeList {
				if strings.Contains(pod.Namespace, s) {
					excluded = true
					break
				}
			}
			if excluded {
				continue
			}
		}
		filteredPods.Items = append(filteredPods.Items, pod)
	}

	return filteredPods
}
