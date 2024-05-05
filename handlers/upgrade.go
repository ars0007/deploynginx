package handlers

import (
	"context"
	"deploynginx/utils"
	"errors"
	"time"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func UpdateVersion(cmd *cobra.Command) error {
	// Get the kube configs flag if provided
	config, err := cmd.Flags().GetString("kubeconfig")
	if err != nil {
		return err
	}

	// Get the image version (ie: nginx:latest) this will be passed
	// as flag by the user
	version, err := cmd.Flags().GetString("version")
	if err != nil {
		return err
	}

	// If the version is not provided, just throw error
	if version == "" {
		return errors.New("invalid version provided")
	}

	// Take out the deployment client
	clientSet := utils.ConnectToK8s(config)
	deploymentClient := clientSet.AppsV1().Deployments("default")

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFn()

	// Get the deployment by name
	deployment, err := deploymentClient.Get(ctx, utils.DeploymentName, v1.GetOptions{})
	if err != nil {
		return err
	}

	// There is only one container for this deployment
	deployment.Spec.Template.Spec.Containers[0].Image = version

	// Update the deployment
	_, err = deploymentClient.Update(ctx, deployment, v1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}
