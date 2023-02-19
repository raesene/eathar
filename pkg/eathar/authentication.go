package eathar

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Creates a list of users defined in cluster role binding RBAC rules for the cluster
func PrincipalList(options *pflag.FlagSet, principal string) []string {
	principalList := make(map[string]bool)
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}

	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		for _, subject := range clusterRoleBinding.Subjects {
			if subject.Kind == principal {
				if principal == "ServiceAccount" {
					subject.Name = subject.Namespace + "/" + subject.Name
				}
				principalList[subject.Name] = true
			}
		}
	}

	principalListSlice := make([]string, 0, len(principalList))
	for key := range principalList {
		principalListSlice = append(principalListSlice, key)
	}
	return principalListSlice
}
