// Package opencanvas handles parsing the JSON Canvas 1.0 file format.
// See https://jsoncanvas.com for more information.
package opencanvas

import "math"

type Canvas struct {
	// Nodes (optional) is an array of nodes.
	Nodes Nodes `json:"nodes,omitempty"`

	// Edges (optional) is an array of edges.
	Edges []Edge `json:"edges"`
}

type Rectangle struct {
	X int
	Y int
	W int
	H int
}

func (c Canvas) BoundingBox() Rectangle {
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt

	for _, node := range c.Nodes {
		g := node.Generic()

		if g.X < minX {
			minX = g.X
		}

		if g.X+g.Width > maxX {
			maxX = g.X + g.Width
		}

		if g.Y < minY {
			minY = g.Y
		}

		if g.Y+g.Height > maxY {
			maxY = g.Y + g.Height
		}
	}

	// TODO: take edges into account?

	return Rectangle{
		X: minX,
		Y: minY,
		W: maxX - minX,
		H: maxY - minY,
	}
}
