package searcher

type NodeInfo struct {
	labels map[string]string
}

type ClusterInfo struct {
	nodes []NodeInfo
}
