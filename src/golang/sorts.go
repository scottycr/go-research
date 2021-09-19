package golang

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const MAXSET = 10000

// Handles any "fun" errors that may pop up.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initSet(set *[MAXSET]int) error {
	file, err := os.Open("sampleset.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	var i int = 0
	var val int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err = strconv.Atoi(scanner.Text())
		check(err)
		set[i] = val
		i++
	}

	return scanner.Err()
}

func printElements(set *[MAXSET]int) {
	fmt.Println("Index\tElement:")
	for i := 0; i < MAXSET/1000; i++ {
		fmt.Println(i*1000, "\t", set[i*1000])
	}
	fmt.Println()
}

func printExecutionTime(execution time.Duration) {
	fmt.Printf("Execution time: %v\n", execution)
}

func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func bubbleSort(set *[MAXSET]int) {
	for i := 0; i < MAXSET-1; i++ {
		for j := 0; j < MAXSET-i-1; j++ {
			if set[j] > set[j+1] {
				swap(&set[j], &set[j+1])
			}
		}
	}
}

func partition(set *[MAXSET]int, low int, high int) int {
	pivot := set[high]
	i := (low - 1)
	for j := low; j <= high-1; j++ {
		if set[j] < pivot {
			i++
			swap(&set[i], &set[j])
		}
	}
	swap(&set[i+1], &set[high])
	return (i + 1)
}

func quickSort(set *[MAXSET]int, start int, end int) {
	if start < end {
		pi := partition(set, start, end)

		quickSort(set, start, pi-1)
		quickSort(set, pi+1, end)
	}
}

func Sorts() {
	var set [MAXSET]int
	err := initSet((*[10000]int)(&set))
	check(err)

	fmt.Println("[Bubble Sort]")
	fmt.Println("\nUnsorted:")
	printElements((*[10000]int)(&set))

	fmt.Println("Sorted:")
	start := time.Now()
	bubbleSort((*[10000]int)(&set))
	execution := time.Since(start)
	printElements((*[10000]int)(&set))
	printExecutionTime(execution)

	// The set is changed directly in the functions to help speed up execution times.
	// Therefore the array is initialized again after each sort.
	err = initSet((*[10000]int)(&set))
	check(err)
	fmt.Println("\n[Quick Sort]")
	fmt.Println("\nUnsorted:")
	printElements((*[10000]int)(&set))

	fmt.Println("\nSorted:")
	start = time.Now()
	quickSort((*[10000]int)(&set), 0, MAXSET-1)
	execution = time.Since(start)
	printElements((*[10000]int)(&set))
	printExecutionTime(execution)
}
