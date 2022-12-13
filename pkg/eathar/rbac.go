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
	if err != nil {
		log.Print(err)
	}
	//Make a list of ClusterRoleBindings to return
	var clusterAdminRoleBindingList v1.ClusterRoleBindingList

	// Get all the ClusterRoleBindings
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	// Get all the ClusterRoles
	//clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})

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

func GetSecretsUsers(options *pflag.FlagSet) {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}

	//var GetSecretsUsersList v1.ClusterRoleBindingList
	// Get all the ClusterRoleBindings
	//clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	//if err != nil {
	//	log.Print(err)
	//}
	var getSecretsClusterRoles v1.ClusterRoleList
	clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRole := range clusterRoles.Items {
		for _, policy := range clusterRole.Rules {
			for _, resource := range policy.Resources {
				//We include list here as listing secrets gives you the contents of the secret
				if resource == "secrets" {
					for _, verb := range policy.Verbs {
						if verb == "get" || verb == "list" || verb == "*" {
							getSecretsClusterRoles.Items = append(getSecretsClusterRoles.Items, clusterRole)
							//We don't want to have this in multiple times if it lists multiple verbs
							break
						}
					}
				}
			}
		}
	}
	var getSecretsUsersList v1.ClusterRoleBindingList
	//Get all the ClusterRoleBindings
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		for _, clusterRole := range getSecretsClusterRoles.Items {
			if clusterRoleBinding.RoleRef.Name == clusterRole.Name {
				getSecretsUsersList.Items = append(getSecretsUsersList.Items, clusterRoleBinding)
			}
		}
	}
	reportRBAC(getSecretsUsersList, options, "Users with access to secrets")
}
