package eathar

import "github.com/spf13/pflag"

//Creates a list of images in use in the cluster
func ImageList(options *pflag.FlagSet) {
	var imageList []Finding
	pods := connectWithPods()
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			imageList = append(imageList, Finding{Check: "Image List", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name, Image: container.Image})
		}
	}
	report(imageList, options, "Image List")
}
