package host

import (
	"strconv"
	"testing"

	apiextensionsfake "github.com/giantswarm/apiextensions/pkg/clientset/versioned/fake"
	"github.com/giantswarm/micrologger/microloggertest"

	"github.com/giantswarm/health-service/pkg/errors"
)

func Test_Health_New(t *testing.T) {
	testCases := []struct {
		name         string
		inputObj     Config
		errorMatcher func(err error) bool
	}{
		{
			name: "case 0: a service is successfully created",
			inputObj: Config{
				G8sClient: apiextensionsfake.NewSimpleClientset(),
				Logger:    microloggertest.New(),
				Provider:  "aws",
			},
			errorMatcher: nil,
		},
		{
			name: "case 1: invalidConfigError returned when provider is missing",
			inputObj: Config{
				G8sClient: apiextensionsfake.NewSimpleClientset(),
				Logger:    microloggertest.New(),
			},
			errorMatcher: errors.IsInvalidConfig,
		},
		{
			name: "case 2: invalidConfigError returned when provider is not a known provider",
			inputObj: Config{
				G8sClient: apiextensionsfake.NewSimpleClientset(),
				Logger:    microloggertest.New(),
				Provider:  "invalid",
			},
			errorMatcher: errors.IsInvalidConfig,
		},
		{
			name: "case 3: invalidConfigError returned when logger is missing",
			inputObj: Config{
				G8sClient: apiextensionsfake.NewSimpleClientset(),
				Provider:  "aws",
			},
			errorMatcher: errors.IsInvalidConfig,
		},
		{
			name: "case 4: invalidConfigError returned when g8sclient is missing",
			inputObj: Config{
				Logger:   microloggertest.New(),
				Provider: "aws",
			},
			errorMatcher: errors.IsInvalidConfig,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, err := New(tc.inputObj)
			switch {
			case err == nil && tc.errorMatcher == nil:
				// correct; carry on
			case err != nil && tc.errorMatcher == nil:
				t.Fatalf("error == %#v, want nil", err)
			case err == nil && tc.errorMatcher != nil:
				t.Fatalf("error == nil, want non-nil")
			case !tc.errorMatcher(err):
				t.Fatalf("error == %#v, want matching", err)
			}
		})
	}
}
