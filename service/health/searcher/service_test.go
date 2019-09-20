package searcher

import (
	"context"
	"testing"

	apiextensionsfake "github.com/giantswarm/apiextensions/pkg/clientset/versioned/fake"
	"github.com/giantswarm/health-service/flag"
	"github.com/spf13/viper"

	"github.com/giantswarm/micrologger/microloggertest"
)

func Test_Service(t *testing.T) {
	testCases := []struct {
		RequestFactory  func() Request
		ExpectedFactory func() Response
		ErrorMatcher    func(error) bool
	}{
		// Case 1. Having an empty request given there should be one master with
		// proper defaults having set.
		{
			RequestFactory: func() Request {
				request := Request{}
				return request
			},
			ExpectedFactory: func() Response {
				request := Response{
					ClusterHealth: "green",
				}
				return request
			},
			ErrorMatcher: nil,
		},
	}

	var err error
	var newService *Service
	{
		f := flag.New()
		v := viper.New()
		v.Set(f.Service.Provider.Kind, "aws")

		newConfig := Config{
			G8sClient: apiextensionsfake.NewSimpleClientset(),
			Logger:    microloggertest.New(),
			Provider:  "aws",
		}
		newService, err = New(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}

	for i, testCase := range testCases {
		request, err := newService.Search(context.TODO(), testCase.RequestFactory())
		if err != nil && testCase.ErrorMatcher == nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ExpectedFactory().ClusterHealth != request.ClusterHealth {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
	}
}
