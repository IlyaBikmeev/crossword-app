package core

import (
	"fmt"
	"testing"
)

func TestCanPlaceWord_EmptyGrid(t *testing.T) {
	g := NewGrid()

	word1 := Placement{
		Word:      "ROOM",
		Point:     Point{0, 0},
		Direction: Horiz,
	}

	if !g.canPlace(word1) {
		t.Errorf("Expected the word %s to be placeable on an empty grid", word1.Word)
	}
}

func TestCanPlaceWord_LetterConflict(t *testing.T) {
	g := NewGrid()
	g.PlaceWord(Placement{Word: "ALL", Point: Point{0, 0}, Direction: Vert})

	word := Placement{Word: "ROOM", Point: Point{0, 0}, Direction: Horiz}

	// conflict with letter 'R' in ROOM
	if g.canPlace(word) {
		t.Errorf("Letter conflict: the word '%s' should not be placeable", word.Word)
	}
}

func TestCanPlaceWord_CorrectIntersection(t *testing.T) {
	g := NewGrid()
	g.PlaceWord(Placement{Word: "OMEGA", Point: Point{0, 0}, Direction: Horiz}) // intersection with letter 'O'
	word := Placement{Word: "ROOM", Point: Point{0, -1}, Direction: Vert}

	if !g.canPlace(word) {
		t.Errorf("Correct intersection: expected the word '%s' to be placeable", word.Word)
	}
}

func TestCanPlaceWord_PerpendicularConflict(t *testing.T) {
	g := NewGrid()
	// neighbor below vertically
	g.PlaceWord(Placement{Word: "XENON", Point: Point{0, 1}, Direction: Horiz})

	word := Placement{Word: "ROOM", Point: Point{0, 0}, Direction: Horiz}

	if g.canPlace(word) {
		t.Errorf("Perpendicular conflict: the word '%s' should not be placeable", word.Word)
	}
}

func TestCanPlaceWord_BeforeOccupied(t *testing.T) {
	g := NewGrid()
	g.PlaceWord(Placement{Word: "X", Point: Point{-1, 0}, Direction: Horiz})

	word := Placement{Word: "ROOM", Point: Point{0, 0}, Direction: Horiz}

	if g.canPlace(word) {
		t.Errorf("Cell before the word is occupied: the word '%s' should not be placeable", word.Word)
	}
}

func TestCanPlaceWord_AfterOccupied(t *testing.T) {
	g := NewGrid()
	// place a word after the test word
	g.PlaceWord(Placement{Word: "X", Point: Point{4, 0}, Direction: Horiz})

	word := Placement{Word: "ROOM", Point: Point{0, 0}, Direction: Horiz}

	if g.canPlace(word) {
		t.Errorf("Cell after the word is occupied: the word '%s' should not be placeable", word.Word)
	}
}

func TestHashNormalized(t *testing.T) {
	g1 := NewGrid()
	g1.PlaceWord(Placement{Word: "ATTRIBUTE", Point: Point{20, 20}, Direction: Vert})
	g1.PlaceWord(Placement{Word: "ATTITUDE", Point: Point{20, 20}, Direction: Horiz})

	g2 := g1.Normalize()

	hash1 := g1.Hash()
	hash2 := g2.Hash()

	fmt.Println("Hash 1:", hash1)
	fmt.Println("Hash 2:", hash2)

	if hash1 != hash2 {
		t.Error("❌ Хеши НЕ совпадают, нужно проверять Normalize()")
	}
}
