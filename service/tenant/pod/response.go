package pod

import v1 "k8s.io/api/core/v1"

// Response holds the data returned from the ListPods function.
type Response struct {
	Pods []v1.Pod `json:"pods"`
}
