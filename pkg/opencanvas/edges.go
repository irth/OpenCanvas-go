package opencanvas

type Side string

const (
	Left   Side = "left"
	Right  Side = "right"
	Top    Side = "top"
	Bottom Side = "bottom"
)

var ValidSides = []Side{
	Left,
	Right,
	Top,
	Bottom,
}

type EndType string

const (
	Start EndType = "start"
	End   EndType = "end"
)

var ValidEndTypes = []EndType{
	Start,
	End,
}

// An Edge is a line that connects one node to another.
type Edge struct {
	// ID is a unique ID for the edge
	ID string `json:"id"`

	// FromNode is the node `id` where the connection starts.
	FromNode string `json:"fromNode"`
	// FromSide (optional) is the side where this edge starts.
	FromSide Side `json:"fromSide"`
	// FromEnd (optional) is the shape of the endpoint at the edge start.
	FromEnd EndType `json:"fromEnd"`

	// ToNode is the node `id` where the connection ends.
	ToNode string `json:"toNode"`
	// ToSide (optional) is the side where this edge ends.
	ToSide Side `json:"toSide"`
	// ToEnd (optional) is the shape of the endpoint at the edge end.
	ToEnd EndType `json:"toEnd"`

	// Color (optional) is the color of the edge, either in hex or a number from 1 to 6.
	Color string `json:"color,omitempty"`

	// Label (optional) is a text label for the edge.
	Label string `json:"label,omitempty"`
}
