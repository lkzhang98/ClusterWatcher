package model

import "github.com/weaveworks/scope/report"

type Parent struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	TopologyID string `json:"topologyId"`
}

type NodeSummary struct {
	ID         string               `json:"id"`
	Label      string               `json:"label"`
	LabelMinor string               `json:"labelMinor"`
	Rank       string               `json:"rank"`
	Shape      string               `json:"shape,omitempty"`
	Tag        string               `json:"tag,omitempty"`
	Stack      bool                 `json:"stack,omitempty"`
	Pseudo     bool                 `json:"pseudo,omitempty"`
	Metadata   []report.MetadataRow `json:"metadata,omitempty"`
	Parents    []Parent             `json:"parents,omitempty"`
	Metrics    []report.MetricRow   `json:"metrics,omitempty"`
	Tables     []report.Table       `json:"tables,omitempty"`
	Adjacency  report.IDList        `json:"adjacency,omitempty"`
}

type NodeSummaries map[string]NodeSummary
