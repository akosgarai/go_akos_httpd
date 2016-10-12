package game

import (
	"log"
	"os"
	"strings"
)

type Game struct {
	rows   int
	cols   int
	logger *log.Logger
}

// New returns an uninitialized game service.
func New(rows, cols int) *Game {
	return &Game{
		rows:   rows,
		cols:   cols,
		logger: log.New(os.Stderr, "[Game] ", log.LstdFlags),
	}
}

func (g *Game) Render() string {
	g.logger.Println("Page rendering")
	var content []string
	content = append(content, "<html><head></head><body><table border=\"1\">")
	for j := 0; j < g.rows; j++ {
		content = append(content, "<tr>")
		for k := 0; k < g.cols; k++ {
			content = append(content, "<td style=\"width: 20px; height: 20px;\">&nbsp;</td>")
		}
		content = append(content, "</tr>")
	}
	content = append(content, "</table></body></html>")
	return strings.Join(content, "")
}
