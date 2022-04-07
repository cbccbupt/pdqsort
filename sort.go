package sortlearning

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](v []T) {
	for cur := 1; cur < len(v); cur++ {
		for j := cur; j > 0 && v[j] < v[j-1]; j-- {
			v[j], v[j-1] = v[j-1], v[j]
		}
	}
}

func QuickSort[T constraints.Ordered](v []T) {
	if len(v) > 1 {
		p := partition(v)
		QuickSort(v[:p])
		QuickSort(v[p+1:])
	}
}

func partition[T constraints.Ordered](v []T) int {
	pivot := v[0]
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

func HeapSort[T constraints.Ordered](v []T) {
	// Build heap with greatest element at top.
	for i := (len(v) - 1) / 2; i >= 0; i-- {
		siftDown(v, i)
	}

	// Pop elements into end of v.
	for i := len(v) - 1; i >= 1; i-- {
		v[0], v[i] = v[i], v[0]
		siftDown(v[:i], 0)
	}
}

func siftDown[T constraints.Ordered](v []T, node int) {
	for {
		child := 2*node + 1
		if child >= len(v) {
			break
		}
		if child+1 < len(v) && v[child] < v[child+1] {
			child++
		}
		if v[node] >= v[child] {
			return
		}
		v[node], v[child] = v[child], v[node]
		node = child
	}
}

const sss = "sdjhuisahdisuajhdijuashdkjashdkjoashdkjashdkashdksajhd"

func qqq(c byte) bool {
	return sss[c] != 0
}
