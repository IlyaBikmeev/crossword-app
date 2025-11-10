# Word Datasets

This folder contains sample word lists for testing the crossword generator.

## Available Datasets

| File | Words | Best For |
|------|-------|----------|
| `words35.txt` | 35 | Quick testing and debugging |
| `words500.txt` | 500 | Medium-sized crosswords |
| `words1000.txt` | 1000 | Larger crosswords |
| `words2k.txt` | 2000 | Extra large grids |
| `words5k.txt` | 5000 | Very large grids (slow) |
| `words10k.txt` | 10000 | Maximum size (very slow) |

## Language Support

The generator supports **any language**, including:
- English (Latin alphabet)
- Russian (Cyrillic alphabet)
- Any other Unicode characters

## Creating Your Own Dataset

Create a text file with one word per line:

```
HELLO
WORLD
CROSSWORD
```

Or in Russian:

```
ПРИВЕТ
МИР
КРОССВОРД
```

### Tips for Good Datasets

1. **Word Length**: Mix of different lengths (4-12 letters work best)
2. **Common Letters**: Include words with frequently used letters for better intersections
3. **Avoid**: Very short words (< 3 letters) or very long words (> 15 letters)
4. **Size**: 50-500 words is optimal for reasonable execution time

## External Sources for Russian Words

You can download Russian word lists from:
- [Russian Word Frequency List](https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists/Russian)
- [OpenCorpora](http://opencorpora.org/)
- [Russian National Corpus](https://ruscorpora.ru/)

## Usage

```bash
# Use a dataset
./bin/crossword datasets/words500.txt

# With options
./bin/crossword --max=5 --mqt=2.0 datasets/words500.txt
```

