package model

type APITopology map[string]APIHostTopology

type APIHostTopology struct {
	NodeSummary
	Children map[string]APIPodTopology `json:"children"`
}

type APIPodTopology struct {
	NodeSummary
	Children NodeSummaries `json:"children"`
}
