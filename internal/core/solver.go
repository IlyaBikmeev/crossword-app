package core

import (
	"fmt"
	"sort"
	"strings"
)

type Solver struct {
	Words               []string
	MinQualityThreshold float64
	MaxSolutions        int
	BestGrid            *Grid
	Solutions           []*Grid
	Metric              Metric
}

func NewSolver(Words []string, MaxSolutions int, MinQualityThreshold float64, metric Metric) *Solver {
	return &Solver{
		Words:               preprocessWords(Words),
		MaxSolutions:        MaxSolutions,
		MinQualityThreshold: MinQualityThreshold,
		Metric:              metric,
	}
}

func (s *Solver) FindSolutions() {
	var helper func(grid *Grid, index, depth int)
	seen := make(map[string]bool)
	var bestScore = 0.0

	helper = func(grid *Grid, index, depth int) {
		if len(s.Solutions) >= s.MaxSolutions {
			return
		}

		hash := grid.Hash()
		if _, exists := seen[hash]; exists {
			return
		}

		seen[hash] = true
		currentScore := grid.Evaluate(s.Metric)

		if index > 0 {
			if bestScore-currentScore > s.MinQualityThreshold {
				//fmt.Printf("Слишком хуевая схема: %f\n", currentScore)
				return
			}
		}

		if index >= len(s.Words) {
			s.Solutions = append(s.Solutions, grid)

			if currentScore > bestScore {
				fmt.Printf("🏆Found a better solution! New high score: %f\n", currentScore)
				bestScore = currentScore
			}

			return
		}

		for _, placement := range grid.PositionsList(s.Words[index]) {
			newGrid := grid.Copy()
			newGrid.PlaceWord(placement)
			helper(newGrid, index+1, depth+1)
		}
	}

	helper(NewGrid(), 0, 0)

	sort.Slice(s.Solutions, func(i, j int) bool {
		return s.Solutions[i].score > s.Solutions[j].score
	})
	s.BestGrid = s.Solutions[0]

}

//func (s *Solver) FindSolutionsParallel() {
//	type Job struct {
//		grid  *Grid
//		index int
//	}
//
//	seen := make(map[string]bool)
//	bestScore := 0.0
//
//	var worker func(job Job, wg *sync.WaitGroup)
//
//	worker = func(job Job, wg *sync.WaitGroup) {
//		defer wg.Done()
//
//		index := job.index
//		grid := job.grid
//
//		if len(s.Solutions) >= s.MaxSolutions {
//			return
//		}
//
//		hash := grid.Hash()
//		if _, exists := seen[hash]; exists {
//			return
//		}
//
//		seen[hash] = true
//		currentScore := grid.Evaluate(s.Metric)
//
//		if index > 0 {
//			if bestScore-currentScore > s.MinQualityThreshold {
//				//fmt.Printf("Слишком хуевая схема: %f\n", currentScore)
//				return
//			}
//		}
//
//		if index >= len(s.Words) {
//			s.Solutions = append(s.Solutions, grid)
//
//			if currentScore > bestScore {
//				fmt.Printf("🏆Found a better solution! New high score: %f\n", currentScore)
//				bestScore = currentScore
//			}
//
//			return
//		}
//
//		for _, placement := range grid.PositionsList(s.Words[index]) {
//			newGrid := grid.Copy()
//			newGrid.PlaceWord(placement)
//			worker(newGrid, index+1)
//		}
//
//	}
//}

func preprocessWords(words []string) []string {
	var processed []string
	for _, word := range words {
		processed = append(processed, preprocessWord(word))
	}
	sortByCrossability(processed)
	return processed
}

func preprocessWord(word string) string {
	return strings.TrimSpace(
		strings.ToUpper(word),
	)
}

func sortByCrossability(words []string) []string {
	freq := computeLetterFrequency(words)
	type ws struct {
		word  string
		score int
	}
	scored := make([]ws, len(words))
	for i, w := range words {
		scored[i] = ws{w, wordCrossScore(w, freq)}
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].score > scored[j].score
	})

	res := make([]string, len(words))
	for i, s := range scored {
		res[i] = s.word
	}
	return res
}

func computeLetterFrequency(words []string) map[rune]int {
	freq := make(map[rune]int)
	for _, word := range words {
		for _, ch := range word {
			freq[ch]++
		}
	}
	return freq
}

func wordCrossScore(word string, freq map[rune]int) int {
	score := 0
	runes := []rune(word)
	for _, ch := range runes {
		score += freq[ch] - 1
	}
	return score
}
