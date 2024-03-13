package opencanvas_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/irth/opencanvas-go/pkg/opencanvas"
	"github.com/stretchr/testify/assert"
)

func GetSampleDataPath(path ...string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path = append([]string{wd, "..", "..", "sample-data"}, path...)
	return filepath.Join(path...)
}

func OpenSampleData(path ...string) *os.File {
	file, err := os.Open(GetSampleDataPath(path...))
	if err != nil {
		panic(err)
	}
	return file
}

func TestUnmarshalingJSONCanvasSample(t *testing.T) {
	f := OpenSampleData("jsoncanvas", "sample.canvas")
	defer f.Close()

	var canvas opencanvas.Canvas
	err := json.NewDecoder(f).Decode(&canvas)
	assert.NoError(t, err)

	expected := opencanvas.Canvas{
		Nodes: opencanvas.Nodes{
			&opencanvas.FileNode{
				GenericNode: opencanvas.GenericNode{
					ID:     "8132d4d894c80022",
					Type:   "file",
					X:      -280,
					Y:      -200,
					Width:  570,
					Height: 560,
					Color:  "6",
				},
				File: "readme.md",
			},
			&opencanvas.FileNode{
				GenericNode: opencanvas.GenericNode{
					ID:     "7efdbbe0c4742315",
					Type:   "file",
					X:      -280,
					Y:      -440,
					Width:  217,
					Height: 80,
				},
				File: "_site/logo.svg",
			},
			&opencanvas.TextNode{
				GenericNode: opencanvas.GenericNode{
					ID:     "59e896bc8da20699",
					Type:   "text",
					X:      40,
					Y:      -440,
					Width:  250,
					Height: 160,
				},
				Text: "Learn more:\n\n- [Readme](readme.md)\n- [Spec](version/1.0.md)\n- [Github](https://github.com/obsidianmd/jsoncanvas)",
			},
			&opencanvas.FileNode{
				GenericNode: opencanvas.GenericNode{
					ID:     "0ba565e7f30e0652",
					Type:   "file",
					X:      360,
					Y:      -400,
					Width:  400,
					Height: 400,
				},
				File: "spec/1.0.md",
			},
		},
		Edges: []opencanvas.Edge{
			{
				ID:       "6fa11ab87f90b8af",
				FromNode: "7efdbbe0c4742315",
				FromSide: opencanvas.Right,
				ToNode:   "59e896bc8da20699",
				ToSide:   opencanvas.Left,
			},
		},
	}
	assert.ElementsMatch(t, expected.Nodes, canvas.Nodes)
	assert.ElementsMatch(t, expected.Edges, canvas.Edges)
}
