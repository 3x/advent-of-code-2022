package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// ---------- Part 1 ----------
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Close file before program exits
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Run for each line in `file`
	for scanner.Scan() {
		// Append the current line to `lines`.
		// The `append()` function returns a new slice with the second argument
		// added, so we set `lines` equal to that result in order to keep using
		// the same slice.
		lines = append(lines, scanner.Text())
	}

	var running_total int = 0
	var elf_totals []int
	for _, line := range lines {
		// temp_slice = append(temp_slice, )
		if line == "" {
			elf_totals = append(elf_totals, running_total)

			// Reset the running total, as a new elf's list starts next
			running_total = 0
		} else {
			line_value, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			running_total += line_value
		}
	}

	// Sort the elf totals from smallest to greatest
	// https://stackoverflow.com/a/43714927
	sort.Ints(elf_totals)

	fmt.Fprintln(os.Stdout, "The elf carrying the most calories is carrying", elf_totals[len(elf_totals)-1], "calories.")

	// ---------- Part 2 ----------

	var top_three_total int = 0
	for i := 1; i <= 3; i++ {
		top_three_total += elf_totals[len(elf_totals)-i]
	}

	fmt.Fprintln(os.Stdout, "The three elves carrying the most calories share", top_three_total, "calories between themselves.")
}
