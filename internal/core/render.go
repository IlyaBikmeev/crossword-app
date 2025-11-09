package core

import (
	"fmt"
	"io"
)

type GridRenderer interface {
	DrawCell(x, y int, ch rune)
	Finish()
}

type SimpleGridRenderer struct {
	Out io.Writer
}

func NewSimpleGridRenderer(out io.Writer) *SimpleGridRenderer {
	return &SimpleGridRenderer{Out: out}
}

func (r *SimpleGridRenderer) DrawCell(x, y int, ch rune) {
	fmt.Fprintf(r.Out, "%c", ch)
}

func (r *SimpleGridRenderer) Finish() {
	fmt.Println()
}
