package model

import (
	"strings"

	"github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

func getKeycloakExtensionsInitContainerEnv(cr *v1alpha1.Keycloak) []v1.EnvVar {
	env := []v1.EnvVar{
		{
			Name:  KeycloakExtensionEnvVar,
			Value: strings.Join(cr.Spec.Extensions, ","),
		},
	}

	if len(cr.Spec.KeycloakDeploymentSpec.Experimental.Env) > 0 {
		// We override Keycloak pre-defined envs with what user specified. Not the other way around.
		env = MergeEnvs(cr.Spec.KeycloakDeploymentSpec.Experimental.Env, env)
	}

	return env
}

func KeycloakExtensionsInitContainers(cr *v1alpha1.Keycloak) []v1.Container {

	return []v1.Container{
		{
			Name:  "extensions-init",
			Image: Profiles.GetInitContainerImage(cr),
			Env:   getKeycloakExtensionsInitContainerEnv(cr),
			VolumeMounts: []v1.VolumeMount{
				{
					Name:      "keycloak-extensions",
					ReadOnly:  false,
					MountPath: KeycloakExtensionsInitContainerPath,
				},
			},
			TerminationMessagePath:   "/dev/termination-log",
			TerminationMessagePolicy: "File",
			ImagePullPolicy:          "Always",
		},
	}
}
