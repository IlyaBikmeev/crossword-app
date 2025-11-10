package core

import (
	"sort"
	"strings"
)

const (
	Horiz Direction = iota
	Vert
)

type Direction int
type Point struct{ X, Y int }
type Placement struct {
	Word      string
	Point     Point
	Direction Direction
}

type Grid struct {
	data  map[Point]rune
	words []Placement
	score float64
}

func NewGrid() *Grid {
	return &Grid{
		data: make(map[Point]rune),
	}
}

func (g *Grid) Equals(other *Grid) bool {
	if len(g.data) != len(other.data) {
		return false
	}
	for p, ch := range g.data {
		if ch2, ok := other.data[p]; !ok || ch != ch2 {
			return false
		}
	}
	return true
}

func (g *Grid) Hash() string {
	if len(g.data) == 0 {
		return ""
	}

	norm := g.Normalize()

	_, maxX, _, maxY := norm.bounds()

	var b strings.Builder

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if ch, ok := norm.data[Point{x, y}]; ok {
				b.WriteRune(ch)
			} else {
				b.WriteRune('.')
			}
		}
		b.WriteRune('|') // разделитель строк
	}

	// Часть 2: Отсортированный список слов
	if len(norm.words) > 0 {
		words := make([]string, len(norm.words))
		for i, pl := range norm.words {
			words[i] = pl.Word
		}
		sort.Strings(words)

		b.WriteRune('#') // разделитель между сеткой и словами
		b.WriteString(strings.Join(words, ","))
	}

	return b.String()
}

func (g *Grid) Normalize() *Grid {
	if len(g.data) == 0 {
		return g.Copy()
	}

	minX, _, minY, _ := g.bounds()

	normalized := NewGrid()
	for p, ch := range g.data {
		normalized.data[Point{p.X - minX, p.Y - minY}] = ch
	}

	for _, w := range g.words {
		normalized.words = append(normalized.words, Placement{
			Word:      w.Word,
			Point:     Point{w.Point.X - minX, w.Point.Y - minY},
			Direction: w.Direction,
		})
	}

	return normalized
}

func (g *Grid) Copy() *Grid {
	newGrid := &Grid{
		data:  make(map[Point]rune, len(g.data)),
		words: make([]Placement, len(g.words)),
		score: g.score,
	}

	for p, r := range g.data {
		newGrid.data[p] = r
	}

	copy(newGrid.words, g.words)

	return newGrid
}

func (g *Grid) Evaluate(metric Metric) float64 {
	g.score = metric.Evaluate(g)
	return g.score
}

func (g *Grid) Area() int {
	minX, maxX, minY, maxY := g.bounds()

	width := maxX - minX + 1
	height := maxY - minY + 1
	return width * height
}

func (g *Grid) Density() float64 {
	if len(g.data) == 0 {
		return 0.0
	}
	return float64(len(g.data)) / float64(g.Area())
}

func (g *Grid) Intersections() int {
	count := 0
	for p := range g.data {
		hasHoriz := g.has(p.X-1, p.Y) || g.has(p.X+1, p.Y)
		hasVert := g.has(p.X, p.Y-1) || g.has(p.X, p.Y+1)
		if hasHoriz && hasVert {
			count++
		}
	}
	return count
}

func (g *Grid) PlaceWord(placement Placement) bool {
	if !g.canPlace(placement) {
		return false
	}
	dx, dy := g.dirDelta(placement.Direction)

	runes := []rune(placement.Word)

	for i, ch := range runes {
		p := Point{placement.Point.X + i*dx, placement.Point.Y + i*dy}
		g.data[p] = ch
	}
	return true
}

func (g *Grid) RemoveWord(placement Placement) {
	dx, dy := g.dirDelta(placement.Direction)
	runes := []rune(placement.Word)

	for i, _ := range runes {
		p := Point{placement.Point.X + i*dx, placement.Point.Y + i*dy}
		g.data[p] = '.'
	}
}

func (g *Grid) PositionsList(word string) []Placement {
	runes := []rune(word)
	var placements []Placement

	if len(g.data) == 0 {
		placements = append(placements, Placement{Word: word, Point: Point{0, 0}, Direction: Horiz})
		return placements
	}

	keys := make([]Point, 0)
	for k, _ := range g.data {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].X != keys[j].X {
			return keys[i].X < keys[j].X
		}
		return keys[i].Y < keys[j].Y
	})

	for _, k := range keys {
		gridPoint := k

		for i := range runes {
			horizPlacement := Placement{
				Word:      word,
				Point:     Point{X: gridPoint.X - i, Y: gridPoint.Y},
				Direction: Horiz,
			}
			if g.canPlace(horizPlacement) {
				placements = append(placements, horizPlacement)
			}

			vertPlacement := Placement{
				Word:      word,
				Point:     Point{X: gridPoint.X, Y: gridPoint.Y - i},
				Direction: Vert,
			}

			if g.canPlace(vertPlacement) {
				placements = append(placements, vertPlacement)
			}
		}
	}

	sort.Slice(placements, func(i, j int) bool {
		return g.positionScore(placements[i]) > g.positionScore(placements[j])
	})

	return placements
}

func (g *Grid) Render(renderer GridRenderer) {
	if len(g.data) == 0 {
		renderer.Finish()
		return
	}

	minX, maxX, minY, maxY := g.bounds()

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			ch, ok := g.data[Point{x, y}]
			if !ok {
				ch = '.'
			}
			renderer.DrawCell(x, y, ch)
		}
		renderer.DrawCell(maxX+1, y, '\n')
	}

	renderer.Finish()

}

func (g *Grid) center() (int, int) {
	minX, maxX, minY, maxY := g.bounds()
	return (maxX - minX) / 2, (maxY - minY) / 2
}

func (g *Grid) has(x, y int) bool {
	_, ok := g.data[Point{x, y}]
	return ok
}

func (g *Grid) positionScore(p Placement) int {
	score := 0
	dx, dy := g.dirDelta(p.Direction)

	runes := []rune(p.Word)
	for i := range runes {
		x := p.Point.X + i*dx
		y := p.Point.Y + i*dy
		if g.has(x, y) {
			score += 10
		}
	}

	cx, cy := g.center()
	dist := abs(p.Point.X-cx) + abs(p.Point.Y-cy)
	score -= dist
	return score
}

func (g *Grid) canPlace(placement Placement) bool {
	if g.isOccupiedBeforeOrAfter(placement) {
		return false
	}

	hasIntersection := false
	gridNotEmpty := len(g.data) > 0

	runes := []rune(placement.Word)
	dx, dy := g.dirDelta(placement.Direction)
	for i, ch := range runes {
		p := Point{placement.Point.X + i*dx, placement.Point.Y + i*dy}
		if occupied, ok := g.data[p]; ok {
			if occupied != ch || g.parallelConflict(p, dx, dy) {
				return false
			}
			hasIntersection = true
			continue
		}

		if g.perpendicularConflict(p, placement.Direction) {
			return false
		}
	}

	if gridNotEmpty && !hasIntersection {
		return false
	}

	return true
}

func (g *Grid) isOccupiedBeforeOrAfter(p Placement) bool {
	dx, dy := g.dirDelta(p.Direction)
	runes := []rune(p.Word)
	before := Point{p.Point.X - dx, p.Point.Y - dy}
	after := Point{p.Point.X + dx*len(runes), p.Point.Y + dy*len(runes)}
	_, okBefore := g.data[before]
	_, okAfter := g.data[after]
	return okBefore || okAfter
}

func (g *Grid) parallelConflict(p Point, dx, dy int) bool {
	before := Point{p.X - dx, p.Y - dy}
	after := Point{p.X + dx, p.Y + dy}
	_, conflictBefore := g.data[before]
	_, conflictAfter := g.data[after]
	return conflictBefore || conflictAfter
}

func (g *Grid) perpendicularConflict(pt Point, dir Direction) bool {
	var neighbors []Point
	if dir == Horiz {
		neighbors = []Point{{0, -1}, {0, 1}}
	} else {
		neighbors = []Point{{-1, 0}, {1, 0}}
	}

	for _, n := range neighbors {
		np := Point{pt.X + n.X, pt.Y + n.Y}
		if _, ok := g.data[np]; ok {
			return true
		}
	}
	return false
}

func (g *Grid) dirDelta(dir Direction) (int, int) {
	if dir == Horiz {
		return 1, 0
	}
	return 0, 1
}

func (g *Grid) bounds() (minX, maxX, minY, maxY int) {
	if len(g.data) == 0 {
		return 0, 0, 0, 0
	}

	minX, maxX, minY, maxY = 999999, -999999, 999999, -999999
	for p := range g.data {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return
}
