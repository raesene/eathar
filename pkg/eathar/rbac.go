package eathar

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetClusterAdminUsers(options *pflag.FlagSet) {
	clientset, err := initKubeClient()

	//Make a list of ClusterRoleBindings to return
	var clusterAdminRoleBindingList v1.ClusterRoleBindingList
	if err != nil {
		log.Print(err)
	}

	// Get all the ClusterRoleBindings
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	// Get all the ClusterRoles
	//clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})

	//Print out the ClusterRoleBindings
	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		//fmt.Println(clusterRoleBinding.Name)
		//Get bindings for cluster-admin
		if clusterRoleBinding.RoleRef.Name == "cluster-admin" {
			//fmt.Printf("Binding name %s, Referenced role %s, Subjects %s\n", clusterRoleBinding.Name, clusterRoleBinding.RoleRef.Name, clusterRoleBinding.Subjects)
			clusterAdminRoleBindingList.Items = append(clusterAdminRoleBindingList.Items, clusterRoleBinding)
		}
	}
	reportRBAC(clusterAdminRoleBindingList, options, "Cluster Admin Users")
}
