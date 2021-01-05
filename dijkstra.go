package dijkstra

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func (g *Graph) setup() error {
	if len(g.Vertices) == 0 {
		return errors.New("please create vertex first.")
	}

	length := len(g.Vertices)
	table := make([][]info, length)
	for idx, vertex := range g.Vertices {
		row := make([]info, length)
		for i := 0; i < length; i++ {
			if vertex.number == i {
				row[i] = info{0, "self"}
				continue
			}

			if weight, ok := vertex.weights[i]; ok {
				// row[i] = info{weight, []string{vertex.ID, g.conversionTable[i]}}
				row[i] = info{weight, vertex.ID + "->" + g.conversionTable[i]}
			} else {
				// 沒有連通
				row[i] = info{-1, "X"}
			}
		}
		table[idx] = row
	}
	g.table = table
	return nil
}

func (g *Graph) DrawTableWithPath() {
	_ = g.setup()

	l := len(g.Vertices)
	header := []string{""}
	data := [][]string{}

	for idx, v := range g.Vertices {
		header = append(header, v.ID)

		row := []string{}
		row = append(row, strings.ToUpper(g.conversionTable[idx]))
		for j := 0; j < l; j++ {
			row = append(row, g.table[idx][j].path)
		}
		data = append(data, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetRowLine(true)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func (g *Graph) DrawTableWithDistance() {
	_ = g.setup()

	l := len(g.Vertices)
	header := []string{""}
	data := [][]string{}

	for idx, v := range g.Vertices {
		header = append(header, v.ID)

		row := []string{}
		row = append(row, strings.ToUpper(g.conversionTable[idx]))
		for j := 0; j < l; j++ {
			row = append(row, strconv.FormatInt(g.table[idx][j].distance, 10))
		}
		data = append(data, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetRowLine(true)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
