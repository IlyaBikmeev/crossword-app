# Crossword Generator

A high-performance crossword puzzle generator written in Go that automatically creates optimal crossword layouts from a list of words.

> Implementation of the **criss-cross problem** from Charles Wetherell's ["Etudes for Programmers"](https://archive.org/details/etudesforprogram0000weth).

[![Live Demo](https://img.shields.io/badge/🎮_Live_Demo-Try_Now-blue?style=for-the-badge)](https://ilyabikmeev.github.io/crossword-app/)

Try the interactive web visualization with sample crosswords!

## Features

- 🚀 **Fast Algorithm**: Uses depth-first search with pruning to efficiently explore the solution space
- 📊 **Quality Metrics**: Multiple evaluation metrics (density, intersection count) to find optimal layouts
- 🎯 **Smart Word Placement**: Automatic detection of valid word intersections
- 🔧 **Configurable**: Adjustable quality thresholds and solution limits
- 📈 **Scalable**: Handles word lists from dozens to thousands of words

## Installation

### Prerequisites

- Go 1.20 or higher
- Make (optional, for convenient build commands)

### Build from source

#### Using Make (recommended)

```bash
git clone https://github.com/ilyabikmeev/crossword-app.git
cd crossword-app
make build
```

The binary will be created in `bin/crossword`.

#### Using Go directly

```bash
git clone https://github.com/ilyabikmeev/crossword-app.git
cd crossword-app
go build -o crossword cmd/crossword/main.go
```

### Quick Start

```bash
# Build and run with default settings
make run
```

## Usage

### Basic Usage

```bash
# Using the binary directly
./bin/crossword datasets/words35.txt

# Or using Make (uses datasets/words35.txt by default)
make run
```

### With Options

```bash
# Run with custom parameters
./bin/crossword --max=10 --mqt=2.0 datasets/words500.txt > solution.txt
```

### Command-line Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--debug` | bool | false | Enable debug mode with detailed logging |
| `--parallel` | bool | false | ⚠️ **[In Development]** Enable parallel processing using multiple CPU cores |
| `--max` | int | 1 | Maximum number of solutions to find |
| `--mqt` | float64 | 4.4 | Minimum quality threshold (lower = more strict pruning) |

## Input Format

Create a text file with one word per line:

```
HELLO
WORLD
CROSSWORD
PUZZLE
GENERATOR
```

The generator **supports any language** including Cyrillic (Russian) characters:

```
ПРИВЕТ
МИР
КРОССВОРД
```

### Sample Datasets

The repository includes several sample word lists in the `datasets/` folder:

| File | Words | Description |
|------|-------|-------------|
| `words35.txt` | 35 | Small dataset for quick testing |
| `words500.txt` | 500 | Medium dataset |
| `words1000.txt` | 1000 | Large dataset |
| `words2k.txt` | 2000 | Extra large dataset |
| `words5k.txt` | 5000 | Very large dataset |
| `words10k.txt` | 10000 | Maximum size dataset |

You can use your own word list or download Russian word lists from various sources.

## Output

The program outputs the crossword grid(s) with quality scores:

```
=== Solution #1 with score 45.23 ===
H.E.L.L.O.|.O.......|W.R.L.D.|.O.......|.S.....P.|.S......|W.O.R.D.|
#HELLO,WORLD,CROSS,WORD
```

### Web Visualization

For better visualization, save the output to `solution.txt` and open `index.html` in your browser:

```bash
# Generate crossword and save to file
./bin/crossword datasets/words35.txt > solution.txt

# Open index.html in your browser
open index.html  # macOS
# or
xdg-open index.html  # Linux
# or just double-click index.html on Windows
```

The HTML viewer provides:
- 🎨 Beautiful grid visualization
- 📊 Interactive interface
- 📱 Responsive design for mobile devices
- ✨ Modern UI with gradient backgrounds

[![Crossword Demo Screenshot](https://img.shields.io/badge/Click_to_open-Live_Demo-brightgreen?style=for-the-badge&logo=github)](https://ilyabikmeev.github.io/crossword-app/)

## Architecture

### Core Components

- **`Grid`**: Represents the crossword grid with word placement logic
- **`Solver`**: Implements the search algorithm (DFS with pruning)
- **`Metric`**: Evaluates the quality of a crossword layout
- **`Renderer`**: Outputs the crossword in various formats

### Algorithm

1. **Preprocessing**: Words are sorted by "crossability" score based on letter frequency
2. **Initial Placement**: First word is placed horizontally at origin
3. **DFS Search**: Recursively try to place remaining words at valid intersection points
4. **Pruning**: Branch cutting based on quality threshold to avoid poor solutions
5. **Evaluation**: Each complete solution is scored using configurable metrics

### Quality Metrics

#### Density Metric
Measures how compact the crossword is:
```
density = (filled cells) / (total bounding box area)
```

#### Density & Intersection Metric
Combines grid density with intersection count:
```
score = (density_weight × density) + (intersection_weight × intersection_ratio)
```

## Performance

### Benchmark Results

Execution time on different word list sizes (single-threaded):

| Word Count | Execution Time |
|------------|----------------|
| 35 words   | ~0.5s |
| 500 words  | ~66s |
| 2000 words | ~1070s |

### Optimization Tips

1. **Adjust Quality Threshold**: Lower `--mqt` for faster execution (more pruning)
2. **Limit Solutions**: Set reasonable `--max` value to stop early
3. **Word List Size**: Smaller word lists (50-500 words) work best
4. **⚠️ Parallel Mode**: Currently in development

## Make Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the project |
| `make run` | Build and run with default settings (words35.txt) |
| `make clean` | Remove build artifacts |

## Examples

### Generate a simple crossword

```bash
make run
# Or: ./bin/crossword datasets/words35.txt
```

### Generate and visualize in browser

```bash
# Generate crossword
./bin/crossword datasets/words35.txt > solution.txt

# Open visualization (macOS)
open index.html
```

### Find multiple high-quality solutions

```bash
./bin/crossword --max=5 --mqt=1.0 datasets/words500.txt
```

### Debug mode for development

```bash
./bin/crossword --debug datasets/words500.txt
```


## Project Structure

```
crossword-app/
├── cmd/
│   └── crossword/
│       └── main.go          # CLI entry point
├── internal/
│   └── core/
│       ├── grid.go          # Grid data structure and placement logic
│       ├── solver.go        # DFS solver algorithm
│       ├── quality.go       # Quality metrics
│       ├── render.go        # Output rendering
│       └── util.go          # Helper functions
├── datasets/                # Sample word lists
│   ├── words35.txt
│   ├── words500.txt
│   ├── words1000.txt
│   ├── words2k.txt
│   ├── words5k.txt
│   └── words10k.txt
├── bin/                     # Build output directory
├── index.html               # Simple web-based visualization
├── Makefile                 # Build automation
├── go.mod                   # Go module definition
└── README.md
```

## Algorithm Details

### Word Placement Rules

1. **First Word**: Placed horizontally at origin (0, 0)
2. **Subsequent Words**: Must intersect with at least one existing word
3. **Valid Intersection**: Same letter at crossing point
4. **No Adjacent Words**: Words cannot touch without intersecting
5. **No Duplicate States**: Grid hashing prevents revisiting same configurations

### Crossability Score

Words are sorted by how likely they are to form intersections:
```go
score = Σ(frequency of each letter in the word list)
```

Words with common letters (E, A, R, etc.) are placed first to maximize intersection opportunities.

## Performance Profiling

The application includes built-in CPU profiling:

```bash
# Run with profiling (cpu.prof will be created automatically)
./bin/crossword datasets/words500.txt

# Analyze the profile
go tool pprof cpu.prof
```

## Contributing

Contributions are welcome! Areas for improvement:

- **Parallel solver implementation** (currently in development)
- Additional quality metrics
- Different search strategies (A*, beam search)
- Output formats (PDF, SVG, HTML)
- Interactive mode
- Constraint satisfaction solver

## Live Demo

Experience the interactive crossword visualization:

### [🎮 Try it now: ilyabikmeev.github.io/crossword-app](https://ilyabikmeev.github.io/crossword-app/)

**Features:**
- Navigate through multiple crossword solutions
- Use ← → arrow keys or click buttons
- Responsive design works on mobile
- Beautiful gradient UI

---

### Deploy Your Own

Enable GitHub Pages in Settings → Pages (Branch: `main`, Folder: `/root`) and push your `solution.txt`:

```bash
./bin/crossword --max=5 datasets/words35.txt > solution.txt
git add solution.txt index.html
git commit -m "Update demo"
git push
```

Your demo will be available at `https://[your-username].github.io/crossword-app/`

## License

MIT License - feel free to use this project for any purpose.

## Author

Ilya Bikmeev

## Acknowledgments

- Based on the **criss-cross problem** from ["Etudes for Programmers"](https://archive.org/details/etudesforprogram0000weth) by Charles Wetherell
- Inspired by traditional crossword puzzle construction techniques
- Uses depth-first search with intelligent pruning

