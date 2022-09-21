package main

import (
	"fmt"
	"sort"
)

type SortingAlgorithm interface {
	Sort(items []int)
}

type BubbleSort struct {
}

func (b *BubbleSort) Sort(items []int) {
	fmt.Println("Sorting with bubble sort")
	n := len(items)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}

type QuickSort struct {
}

func (q *QuickSort) Sort(items []int) {
	fmt.Println("Sorting with quick sort")
	// Use built in pattern-defeating quicksort or implement your own
	sort.Ints(items)
}

type Sorter struct {
	algorithm SortingAlgorithm
}

func (s *Sorter) setSortingAlgorithm(algorithm SortingAlgorithm) {
	s.algorithm = algorithm
}

func (s *Sorter) sort(items []int) {
	s.algorithm.Sort(items)
}

func main() {
	var sorter Sorter

	items := []int{1, 5, 4, 2, 8}
	sorter.setSortingAlgorithm(&BubbleSort{})
	sorter.sort(items)
	fmt.Println("Items are sorted:", items)

	sorter.setSortingAlgorithm(&QuickSort{})
	items2 := []int{1, 5, 4, 2, 8}
	sorter.sort(items2)
	fmt.Println("Items are sorted:", items2)
}
