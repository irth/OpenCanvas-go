// Package opencanvas handles parsing the JSON Canvas 1.0 file format.
// See https://jsoncanvas.com for more information.
package opencanvas

type Canvas struct {
	// Nodes (optional) is an array of nodes.
	Nodes Nodes `json:"nodes,omitempty"`

	// Edges (optional) is an array of edges.
	Edges []Edge `json:"edges"`
}

// Edges
