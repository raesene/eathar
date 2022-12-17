package eathar

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetClusterAdminUsers(options *pflag.FlagSet) v1.ClusterRoleBindingList {
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
	return clusterAdminRoleBindingList

}

func GetSecretsUsers(options *pflag.FlagSet) v1.ClusterRoleBindingList {
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
	return getSecretsUsersList

}

func CreatePVUsers(options *pflag.FlagSet) v1.ClusterRoleBindingList {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}

	var createPVClusterRoles v1.ClusterRoleList
	clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRole := range clusterRoles.Items {
		for _, policy := range clusterRole.Rules {
			for _, resource := range policy.Resources {
				if resource == "persistentvolumes" {
					for _, verb := range policy.Verbs {
						if verb == "create" || verb == "*" {
							createPVClusterRoles.Items = append(createPVClusterRoles.Items, clusterRole)
							//We don't want to have this in multiple times if it lists multiple verbs
							break
						}
					}
				}
			}
		}
	}
	var createPVUsersList v1.ClusterRoleBindingList
	//Get all the ClusterRoleBindings
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		for _, clusterRole := range createPVClusterRoles.Items {
			if clusterRoleBinding.RoleRef.Name == clusterRole.Name {
				createPVUsersList.Items = append(createPVUsersList.Items, clusterRoleBinding)
			}
		}
	}
	return createPVUsersList

}

//Function to get a list of users with access to the escalate verb
func EscalateUsers(options *pflag.FlagSet) v1.ClusterRoleBindingList {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}

	var escalateClusterRoles v1.ClusterRoleList
	clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	//TODO: This isn't quite right as it will also pick up users with access to the escalate verb on other resources
	for _, clusterRole := range clusterRoles.Items {
		for _, policy := range clusterRole.Rules {
			for _, verb := range policy.Verbs {
				if verb == "escalate" {
					escalateClusterRoles.Items = append(escalateClusterRoles.Items, clusterRole)
				}
			}
		}
	}
	var escalateUsersList v1.ClusterRoleBindingList
	//Get all the ClusterRoleBindings
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		for _, clusterRole := range escalateClusterRoles.Items {
			if clusterRoleBinding.RoleRef.Name == clusterRole.Name {
				escalateUsersList.Items = append(escalateUsersList.Items, clusterRoleBinding)
			}
		}
	}
	return escalateUsersList

}

//Function to list users with access to the impersonate verb
func ImpersonateUsers(options *pflag.FlagSet) v1.ClusterRoleBindingList {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}

	var impersonateClusterRoles v1.ClusterRoleList
	clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRole := range clusterRoles.Items {
		for _, policy := range clusterRole.Rules {
			for _, verb := range policy.Verbs {
				if verb == "impersonate" {
					impersonateClusterRoles.Items = append(impersonateClusterRoles.Items, clusterRole)
				}
			}
		}
	}
	var impersonateUsersList v1.ClusterRoleBindingList
	//Get all the ClusterRoleBindings
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		for _, clusterRole := range impersonateClusterRoles.Items {
			if clusterRoleBinding.RoleRef.Name == clusterRole.Name {
				impersonateUsersList.Items = append(impersonateUsersList.Items, clusterRoleBinding)
			}
		}
	}
	return impersonateUsersList

}

//Function to list users with access to the bind verb
func BindUsers(options *pflag.FlagSet) v1.ClusterRoleBindingList {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}

	var bindClusterRoles v1.ClusterRoleList
	clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRole := range clusterRoles.Items {
		for _, policy := range clusterRole.Rules {
			for _, verb := range policy.Verbs {
				if verb == "bind" {
					bindClusterRoles.Items = append(bindClusterRoles.Items, clusterRole)
				}
			}
		}
	}
	var bindUsersList v1.ClusterRoleBindingList
	//Get all the ClusterRoleBindings
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		for _, clusterRole := range bindClusterRoles.Items {
			if clusterRoleBinding.RoleRef.Name == clusterRole.Name {
				bindUsersList.Items = append(bindUsersList.Items, clusterRoleBinding)
			}
		}
	}
	return bindUsersList

}

//Function to list users who can create or modify validatingadmissionwebhookconfigurations
func ValidatingWebhookUsers(options *pflag.FlagSet) v1.ClusterRoleBindingList {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}

	var validatingWebhookClusterRoles v1.ClusterRoleList
	clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRole := range clusterRoles.Items {
		for _, policy := range clusterRole.Rules {
			for _, resource := range policy.Resources {
				if resource == "validatingadmissionwebhookconfigurations" {
					for _, verb := range policy.Verbs {
						if verb == "create" || verb == "update" || verb == "patch" || verb == "delete" || verb == "*" {
							validatingWebhookClusterRoles.Items = append(validatingWebhookClusterRoles.Items, clusterRole)
							//We don't want to have this in multiple times if it lists multiple verbs
							break
						}
					}
				}
			}
		}
	}
	var validatingWebhookUsersList v1.ClusterRoleBindingList
	//Get all the ClusterRoleBindings
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		for _, clusterRole := range validatingWebhookClusterRoles.Items {
			if clusterRoleBinding.RoleRef.Name == clusterRole.Name {
				validatingWebhookUsersList.Items = append(validatingWebhookUsersList.Items, clusterRoleBinding)
			}
		}
	}
	return validatingWebhookUsersList
}

//Function to list users who can create or modify mutatingadmissionwebhookconfigurations
func MutatingWebhookUsers(options *pflag.FlagSet) v1.ClusterRoleBindingList {
	clientset, err := initKubeClient()
	if err != nil {
		log.Print(err)
	}
	var mutatingWebhookClusterRoles v1.ClusterRoleList
	clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRole := range clusterRoles.Items {
		for _, policy := range clusterRole.Rules {
			for _, resource := range policy.Resources {
				if resource == "mutatingadmissionwebhookconfigurations" {
					for _, verb := range policy.Verbs {
						if verb == "create" || verb == "update" || verb == "patch" || verb == "delete" || verb == "*" {
							mutatingWebhookClusterRoles.Items = append(mutatingWebhookClusterRoles.Items, clusterRole)
							//We don't want to have this in multiple times if it lists multiple verbs
							break
						}
					}
				}
			}
		}
	}
	var mutatingWebhookUsersList v1.ClusterRoleBindingList
	//Get all the ClusterRoleBindings
	clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}
	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		for _, clusterRole := range mutatingWebhookClusterRoles.Items {
			if clusterRoleBinding.RoleRef.Name == clusterRole.Name {
				mutatingWebhookUsersList.Items = append(mutatingWebhookUsersList.Items, clusterRoleBinding)
			}
		}
	}
	return mutatingWebhookUsersList

}
