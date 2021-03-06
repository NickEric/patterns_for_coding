package top_k_elements

import (
	"container/heap"
	"testing"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// https://leetcode-cn.com/problems/smallest-k-lcci/
func smallestK(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	if k == 0 {
		return []int{}
	}
	mh := new(Heap)
	for i := 0; i < k; i++ {
		heap.Push(mh, arr[i])
	}

	min := (*mh)[0]
	for i := k; i < len(arr); i++ {
		if min > arr[i] {
			heap.Pop(mh)
			heap.Push(mh, arr[i])
			min = (*mh)[0]
		}
	}

	return *mh
}

// https://leetcode-cn.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	mh := new(FreHeap)
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if _, ok := m[n]; ok {
			m[n] += 1
		} else {
			m[n] = 1
		}
	}

	var cur = 0
	for num, fre := range m {
		if cur < k {
			cur++
			heap.Push(mh, Fre{
				num: num,
				fre: fre,
			})
		} else {
			top := (*mh)[0]
			if fre > top.fre {
				heap.Pop(mh)
				heap.Push(mh, Fre{
					num: num,
					fre: fre,
				})
			}
		}
	}

	var result []int
	for _, f := range *mh {
		result = append(result, f.num)
	}
	return result
}

type Fre struct {
	num int
	fre int
}

type FreHeap []Fre

func (h FreHeap) Len() int {
	return len(h)
}

func (h FreHeap) Less(i, j int) bool {
	return h[i].fre < h[j].fre
}

func (h FreHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FreHeap) Push(x interface{}) {
	*h = append(*h, x.(Fre))
}

func (h *FreHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestTopKFrequent(t *testing.T) {
	topKFrequent([]int{5, 3, 1, 1, 1, 3, 5, 73, 1}, 3)
}

// quick select 的方式进行查找
func findKthLargest(nums []int, k int) int {
	index := quick_select(nums, 0, len(nums)-1, k-1)
	return nums[index]
}

func quick_select(arr []int, start, end int, k int) int {
	pivot := partition(arr, start, end)
	switch {
	case pivot > k:
		return quick_select(arr, start, pivot-1, k)
	case pivot < k:
		return quick_select(arr, pivot+1, end, k)
	default:
		return pivot
	}
}

func partition(arr []int, start, end int) int {
	pivot := start
	index := pivot
	for i := index; i <= end; i++ {
		if arr[i] > arr[pivot] {
			index += 1
			swap(arr, i, index)
		}
	}
	swap(arr, pivot, index)
	return index
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func TestFindKthLargest(t *testing.T) {
	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
}
