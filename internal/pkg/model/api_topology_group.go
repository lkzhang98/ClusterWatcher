package model

type TopologyGroup struct {
	Id    string        `json:"group_id"`
	Nodes NodeSummaries `json:"nodes,omitempty"`
	//Adjacency report.IDList       `json:"adjacency"`
}

type APITopologyGroup struct {
	GroupList   []string                 `json:"group_list"`
	SubTopology map[string]TopologyGroup `json:"sub_topology"`
}

func NewAPITopologyGroup() *APITopologyGroup {
	return &APITopologyGroup{
		GroupList:   make([]string, 0),
		SubTopology: make(map[string]TopologyGroup),
	}
}
