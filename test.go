package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func patition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		Swap(nums, i, j)
	}
	Swap(nums, left, i)
	return i
}

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := left + rand.Intn(right-left+1)
	Swap(nums, pivotIndex, left)
	pivot := patition(nums, left, right)
	QuickSort(nums, left, pivot-1)
	QuickSort(nums, pivot+1, right)
}

func mergeSort(nums []int) []int {
	n := len(nums)
	if n == 0 || n == 1 {
		return nums
	}
	mid := n / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(l1, l2 []int) []int {
	res := []int{}
	for len(l1) != 0 && len(l2) != 0 {
		if l1[0] < l2[0] {
			res = append(res, l1[0])
			l1 = l1[1:]
		} else {
			res = append(res, l2[0])
			l2 = l2[1:]
		}
	}
	res = append(res, l1...)
	res = append(res, l2...)
	return res
}

func HeapSort(nums []int) []int {
	heapSize := len(nums)
	buildHeap(nums, heapSize)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapSize--
		maxHeapfy(nums, 0, heapSize)
	}
	return nums
}

func buildHeap(nums []int, heapSize int) {
	for i := heapSize/2 - 1; i >= 0; i-- {
		maxHeapfy(nums, i, heapSize)
	}
}

func maxHeapfy(nums []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapfy(nums, largest, heapSize)
	}
}

func main() {
	nums := []int{}
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numLine := strings.Fields(scanner.Text())
	for i := 0; i < len(numLine); i++ {
		num, _ := strconv.Atoi(numLine[i])
		nums = append(nums, num)
	}
	fmt.Println(HeapSort(nums))
}
