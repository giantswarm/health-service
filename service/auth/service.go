package auth

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	authorizationv1 "k8s.io/api/authorization/v1"
	"k8s.io/client-go/kubernetes"
)

// Config represents the configuration used to create a new service object.
type Config struct {
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger

	Provider string
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	// Dependencies.
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "k8sclient must not be empty")
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "logger must not be empty")
	}
	if config.Provider == "" {
		return nil, microerror.Maskf(invalidConfigError, "provider must not be empty")
	}

	newService := &Service{
		Config: config,
	}

	return newService, nil
}

type Service struct {
	Config
}

// Search searches for the cluster information.
// It try to find cluster information in CR and fallback to storage service when nothing is found.
func (s *Service) CheckAuthorization(ctx context.Context, request Request) (*Response, error) {
	var clusterResource string
	{
		switch s.Provider {
		case "aws":
			clusterResource = "AWSConfig"
		case "kvm":
			clusterResource = "KVMConfig"
		case "azure":
			clusterResource = "AzureConfig"
		}
	}
	sarRequest := authorizationv1.SubjectAccessReview{
		Spec: authorizationv1.SubjectAccessReviewSpec{
			User:   request.User,
			Groups: request.Groups,
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Verb:     "get",
				Group:    "core.giantswarm.io", // this should be a constant,
				Version:  "v1alpha1",           // this should be a constant
				Resource: clusterResource,
				Name:     request.ClusterID,
			},
		},
	}
	sarResponse, err := s.K8sClient.AuthorizationV1().SubjectAccessReviews().Create(&sarRequest)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	response := Response{
		Allowed: sarResponse.Status.Allowed,
	}

	return &response, nil
}
