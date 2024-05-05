package handlers

import (
	"context"
	"deploynginx/utils"
	"errors"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ScaleDeployment is the handler of scale command which scales the deployment
// if number of replicas provided then it takes that or else takes 1 as default
func ScaleDeployment(cmd *cobra.Command) error {
	config, err := cmd.Flags().GetString("kubeconfig")
	if err != nil {
		return err
	}

	replicas, err := cmd.Flags().GetInt("replicas")
	if err != nil {
		return err
	}

	if replicas < 0 {
		return errors.New("invalid number of replicas")
	}

	clientSet := utils.ConnectToK8s(config)
	deploymentClient := clientSet.AppsV1().Deployments("default")

	deployment, err := deploymentClient.Get(
		context.TODO(),
		utils.DeploymentName,
		v1.GetOptions{})

	if err != nil {
		return err
	}

	r := int32(replicas)
	deployment.Spec.Replicas = &r

	_, err = deploymentClient.Update(context.TODO(), deployment, v1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}
