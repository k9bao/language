package main

import (
	"fmt"

	"./cgoG"
)

func main() {
	l := []int{3, 2, 1, 5, 6, 4}
	m := l[2:2]
	fmt.Println(m)

	// fmt.Println(max_m(l, 2))
	fmt.Println("main")
	//cprocessDemo()
	x, y := 1, 2
	fmt.Println(x, y)
	//http.ListenAndServe(":12345", nil)
}

func cprocessDemo() {
	cgoG.Go2C()
}

/*
def heapify(arr, n, i):
    largest = i # Initialize largest as root
    l = 2 * i + 1     # left = 2*i + 1
    r = 2 * i + 2     # right = 2*i + 2

    # See if left child of root exists and is
    # greater than root
    if l < n and arr[i] < arr[l]:
        largest = l

    # See if right child of root exists and is
    # greater than root
    if r < n and arr[largest] < arr[r]:
        largest = r

    # Change root, if needed
    if largest != i:
        arr[i],arr[largest] = arr[largest],arr[i] # swap

        # Heapify the root.
        heapify(arr, n, largest)
*/
func heapify(arr []int, n, i int) {
	largest := i // Initialize largest as root
	l := 2*i + 1 // left = 2*i + 1
	r := 2*i + 2 // right = 2*i + 2

	// See if left child of root exists and is
	// greater than root
	if l < n && arr[i] < arr[l] {
		largest = l
	}

	// See if right child of root exists and is
	// greater than root
	if r < n && arr[largest] < arr[r] {
		largest = r
	}

	// Change root, if needed
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // swap
		// Heapify the root.
		heapify(arr, n, largest)
	}

}

/*
def max_m(arr, m):
    n = len(arr)

    # Build a maxheap.
    for i in range(n, -1, -1):
        heapify(arr, n, i)

    # One by one extract elements
    for i in range(n-1, n-m-1, -1):
        arr[i], arr[0] = arr[0], arr[i] # swap
        heapify(arr, i, 0)

	return arr[n-1:n-m-1:-1]
*/

func max_m(arr []int, m int) []int {
	n := len(arr)

	// Build a maxheap.
	for i := n; i > 0; i-- {
		heapify(arr, n, i)
	}

	// One by one extract elements
	for i := n - 1; i > n-m-1; i-- {
		arr[i], arr[0] = arr[0], arr[i] // swap
		heapify(arr, i, 0)
	}
	return arr[n-m : n]
}
