// DupFile выводит имена файлов в которых есть повторяющиеся строки
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filesWithDuplicates := make(map[string]bool)

	if len(os.Args) == 1 {
		// Обработка stdin
		counts := make(map[string]int)
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			counts[input.Text()]++
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("stdin: line '%s' appears %d times\n", line, n)
				filesWithDuplicates["stdin"] = true
			}
		}
	} else {
		// Обработка файлов
		for _, filename := range os.Args[1:] {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			counts := make(map[string]int)
			input := bufio.NewScanner(file)
			for input.Scan() {
				counts[input.Text()]++
			}
			file.Close()

			hasDuplicates := false
			for line, n := range counts {
				if n > 1 {
					if !hasDuplicates {
						fmt.Printf("File: %s\n", filename)
						hasDuplicates = true
						filesWithDuplicates[filename] = true
					}
					fmt.Printf("  Line '%s' appears %d times\n", line, n)
				}
			}
		}
	}

	// Итоговый вывод
	if len(filesWithDuplicates) > 0 {
		fmt.Println("\nFiles with duplicates:")
		for filename := range filesWithDuplicates {
			fmt.Printf("  %s\n", filename)
		}
	} else {
		fmt.Println("No duplicates found in any file")
	}
}
