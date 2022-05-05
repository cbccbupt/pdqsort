package sortlearning

import (
	"math/bits"

	"golang.org/x/exp/constraints"
)

func PDQsortV1[T constraints.Ordered](v []T) {
	recurseV1(v, bits.Len(uint(len(v))))
}

func recurseV1[T constraints.Ordered](v []T, limit int) {
	const maxInsertion = 24 // slices of up to this length get sorted using insertion sort.

	var (
		// True if the last partitioning was reasonably balanced.
		wasBalanced = true
	)

	for {
		length := len(v)

		// Very short slices get sorted using insertion sort.
		if length <= maxInsertion {
			InsertionSort(v)
			return
		}

		// If too many bad pivot choices were made, simply fall back to heapsort in order to
		// guarantee `O(n log n)` worst-case.
		if limit == 0 {
			HeapSort(v)
			return
		}

		if !wasBalanced {
			limit--
		}

		// Choose a pivot and try guessing whether the slice is already sorted.
		pivotidx := choosePivotV1(v)

		// Partition the slice.
		mid := partitionv1(v, pivotidx)

		left, right := v[:mid], v[mid+1:]
		if len(left) < len(right) {
			wasBalanced = len(left) >= len(v)/8
			recurseV1(left, limit)
			v = right
		} else {
			wasBalanced = len(right) >= len(v)/8
			recurseV1(right, limit)
			v = left
		}
	}
}

func partitionv1[T constraints.Ordered](v []T, pivotidx int) int {
	pivot := v[pivotidx]
	v[0], v[pivotidx] = v[pivotidx], v[0]
	i, j := 1, len(v)-1

	for {
		for i <= j && v[i] < pivot {
			i++
		}
		for i <= j && v[j] >= pivot {
			j--
		}
		if i > j {
			break
		}
		v[i], v[j] = v[j], v[i]
		i++
		j--
	}
	v[j], v[0] = v[0], v[j]
	return j
}
