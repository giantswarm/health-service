package node

import (
	"strconv"
	"testing"

	"github.com/giantswarm/micrologger/microloggertest"
	"github.com/giantswarm/tenantcluster/tenantclustertest"
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
				Logger:        microloggertest.New(),
				TenantCluster: tenantclustertest.New(tenantclustertest.Config{}),
			},
			errorMatcher: nil,
		},
		{
			name: "case 1: invalidConfigError returned when logger is missing",
			inputObj: Config{
				TenantCluster: tenantclustertest.New(tenantclustertest.Config{}),
			},
			errorMatcher: IsInvalidConfig,
		},
		{
			name: "case 2: invalidConfigError returned when tenantcluster is missing",
			inputObj: Config{
				Logger: microloggertest.New(),
			},
			errorMatcher: IsInvalidConfig,
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
