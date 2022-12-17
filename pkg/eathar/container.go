package eathar

import "github.com/spf13/pflag"

//Creates a list of images in use in the cluster
func ImageList(options *pflag.FlagSet) []string {
	imageList := make(map[string]bool)
	pods := connectWithPods()
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			imageList[container.Image] = true
		}
	}

	imageListSlice := make([]string, 0, len(imageList))

	for key := range imageList {
		imageListSlice = append(imageListSlice, key)
	}

	return imageListSlice
	//reportImage(imageListSlice, options, "Image List")
}
