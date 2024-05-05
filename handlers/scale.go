package handlers

import (
	"context"
	"deploynginx/utils"
	"fmt"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ScaleDeployment is the handler of scale command which scales the deployment
// if number of replicas provided then it takes that or else takes 1 as default
func ScaleDeployment(cmd *cobra.Command) {
	config, err := cmd.Flags().GetString("kubeconfig")
	if err != nil {
		panic(err)
	}

	replicas, err := cmd.Flags().GetInt("replicas")
	if err != nil {
		panic(err)
	}

	if replicas < 0 {
		panic("Invalid number of replicas")
	}

	clientSet := utils.ConnectToK8s(config)
	deploymentClient := clientSet.AppsV1().Deployments("default")

	deployment, err := deploymentClient.Get(
		context.TODO(),
		utils.DeploymentName,
		v1.GetOptions{})

	if err != nil {
		panic(err)
	}

	r := int32(replicas)
	deployment.Spec.Replicas = &r

	_, err = deploymentClient.Update(context.TODO(), deployment, v1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully scaled")
}
