// Package opencanvas handles parsing the JSON Canvas 1.0 file format.
// See https://jsoncanvas.org for more information.
package opencanvas

import "math"

type Canvas struct {
	// Nodes (optional) is an array of nodes.
	Nodes Nodes `json:"nodes,omitempty"`

	// Edges (optional) is an array of edges.
	Edges []Edge `json:"edges"`
}

type Rectangle struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"width"`
	H int `json:"height"`
}

type BoundingBoxable interface {
	BoundingBox() Rectangle
}

func (r Rectangle) BoundingBox() Rectangle {
	return r
}

func (r Rectangle) AsPercentOf(boxable BoundingBoxable) PercentRectangle {
	boundingBox := boxable.BoundingBox()

	relative := Rectangle{
		X: r.X - boundingBox.X,
		Y: r.Y - boundingBox.Y,
		W: r.W,
		H: r.H,
	}

	return PercentRectangle{
		X: float64(relative.X) / float64(boundingBox.W) * 100,
		Y: float64(relative.Y) / float64(boundingBox.H) * 100,
		W: float64(relative.W) / float64(boundingBox.W) * 100,
		H: float64(relative.H) / float64(boundingBox.H) * 100,
	}
}

type PercentRectangle struct {
	X float64
	Y float64
	W float64
	H float64
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

		if g.X+g.W > maxX {
			maxX = g.X + g.W
		}

		if g.Y < minY {
			minY = g.Y
		}

		if g.Y+g.H > maxY {
			maxY = g.Y + g.H
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
