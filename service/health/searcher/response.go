package searcher

type Response struct {
	ClusterHealth string `json:"clusterHealth"`
}

func DefaultResponse() Response {
	return Response{
		ClusterHealth: "",
	}
}
