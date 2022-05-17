package heap

type Heap []int

// 0 1 2 3 4 5 6 7 8
//
// head:    0
//      1       2
//    3   4   5   6
//   7 8

// 0.left  -> 1   2*0+1
// 0.right -> 2   2*1
// 1.left  -> 3   2*1+1
// 1.right -> 4   2*2
// left = 2*i+1
// right = 2*(i+1)

// 1.parent -> 0  2/2-1=0
// 2.parent -> 0  3/2-1=0
// 3.parent -> 1  4/2-1=1
// 4.parent -> 1  5/2-1=1
// (i+1)/2-1

func left(x int) int {
	return 2*x + 1
}

func right(x int) int {
	return 2 * (x + 1)
}

func parent(x int) int {
	return (x+1)/2 - 1
}

func (h Heap) sink(k int) {
	// sink
	for left(k) < len(h) {
		j := left(k)
		if j < len(h) && h[j] < h[j+1] {
			j++ // point to right
		}
		if h[k] >= h[j] {
			break
		} else {
			h[k], h[j] = h[j], h[k]
		}
		k = j
	}
}

func (h Heap) swim(k int) {
	for k > 0 && h[k] > h[parent(k)] {
		h[k], h[parent(k)] = h[parent(k)], h[k]
		k = parent(k)
	}
}

func (h *Heap) Pop() int {
	// sink
	return 0
}

func (h *Heap) Push(x int) {

	// swim
	// return
}
