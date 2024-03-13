package opencanvas

import (
	"encoding/json"
	"fmt"
)

type NodeType string

const (
	Text  NodeType = "text"
	File  NodeType = "file"
	Link  NodeType = "link"
	Group NodeType = "group"
)

type Node interface {
	NodeType() NodeType
	Generic() GenericNode
}

type Nodes []Node

func (ns *Nodes) UnmarshalJSON(b []byte) error {
	// split the JSON into separate nodes
	var raw []json.RawMessage
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	// for each node, determine the type then unmarshal it and append to ourselves
	for i, raw := range raw {
		var partialNode struct {
			Type NodeType `json:"type"`
		}
		err := json.Unmarshal(raw, &partialNode)
		if err != nil {
			return fmt.Errorf("failed to unmarshal node %d: %w", i, err)
		}

		var realNode Node
		switch partialNode.Type {
		case Text:
			realNode = &TextNode{}
		case File:
			realNode = &FileNode{}
		case Link:
			realNode = &LinkNode{}
		case Group:
			realNode = &GroupNode{}
		default:
			return fmt.Errorf("unknown node type for node %d: %s", i, partialNode.Type)
		}

		err = json.Unmarshal(raw, realNode)
		if err != nil {
			return fmt.Errorf("failed to unmarshal node %d: %w", i, err)
		}
		*ns = append(*ns, realNode)
	}

	return nil
}

// GenericNode contains attributes included by all node types.
type GenericNode struct {
	// ID is a unique identifier for the node.
	ID string `json:"id"`
	// Type is the type of the node.
	Type string `json:"type"`

	Rectangle `json:",inline"`

	// Color is the color of the node, either in hex or a number from 1 to 6
	// representing a predefined color.
	Color string `json:"color,omitempty"`
}

func (gn GenericNode) NodeType() NodeType {
	return NodeType(gn.Type)
}

func (gn GenericNode) Generic() GenericNode {
	return gn
}

// TextNode stores text.
type TextNode struct {
	GenericNode `json:",inline"`

	// Text contains plain text with Markdown syntax.
	Text string `json:"text"`
}

// FileNode references another file or attachment.
type FileNode struct {
	GenericNode `json:",inline"`

	// Path to the file within the system.
	File string `json:"file"`

	// Subpath (optional) is a subpath that may link to a heading or a block. Always starts with a `#`.
	Subpath string `json:"subpath,omitempty"`
}

// LinkNode references an URL.
type LinkNode struct {
	GenericNode `json:",inline"`

	URL string `json:"url"`
}

type BackgroundStyle string

const (
	Cover  BackgroundStyle = "cover"
	Ratio  BackgroundStyle = "ratio"
	Repeat BackgroundStyle = "repeat"
)

var ValidBackgroundStyles = []BackgroundStyle{
	Cover,
	Ratio,
	Repeat,
}

// GroupNode is used as a visual container for nodes within it.
type GroupNode struct {
	Node `json:",inline"`

	// Label (optional) is a text label for the group.
	Label string `json:"label,omitempty"`
	// Background (optional) is the path to the background image.
	Background string `json:"background,omitempty"`
	// BackgroundStyle (optional) is the rendering style of the background image.
	BackgroundStyle BackgroundStyle `json:"backgroundStyle,omitempty"`
}
