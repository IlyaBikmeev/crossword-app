package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ilyabikmeev/crossword-app/internal/core"
	"os"
	"sort"
)

func main() {
	// Флаги
	debug := flag.Bool("debug", false, "Включить режим отладки")
	minQualityThreshold := flag.Float64("mqt", 4.4, "Минимальный уровень качества, ниже которого ветка отбрасывается")
	maxSolutions := flag.Int("max", 1, "Maximum number of solutions")
	parallel := flag.Bool("parallel", false, "Parallel flag")
	flag.Parse()

	// Проверка аргументов
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Использование: crossword-cli [--debug] [--all] <файл_со_словами>")
		os.Exit(1)
	}
	inputFile := args[0]

	// Чтение слов из файла
	words := []string{}
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}

	if *debug {
		fmt.Printf("DEBUG: прочитано слов: %d\n", len(words))
	}

	sort.Slice(words, func(i, j int) bool { return len(words[i]) > len(words[j]) })

	// Создание решателя
	solver := core.NewSolver(
		words,
		*maxSolutions,
		*minQualityThreshold,
		core.DensityMetric{})

	if *debug {
		fmt.Println("DEBUG: запускаем поиск решений...")
	}

	if *parallel {
		solver.FindSolutionsParallel()
	} else {
		solver.FindSolutions()
	}

	if *debug {
		fmt.Printf("DEBUG: найдено решений: %d\n", len(solver.Solutions))
		fmt.Printf("DEBUG: лучший результат: %f\n", solver.BestGrid.Evaluate(core.NewDensityAndIntersectionMetric(100, 100)))
	}

	for i, grid := range solver.Solutions {
		fmt.Printf("=== Решение #%d ===\n", i+1)
		fmt.Println(grid.Hash())
	}
}
