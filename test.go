package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree(node []string) *TreeNode {
	if len(node) == 0 || node[0] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(node[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	index := 1
	for index < len(node) && len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if index < len(node) {
			if node[index] != "null" {
				left_Val, _ := strconv.Atoi(node[index])
				n.Left = &TreeNode{Val: left_Val}
				queue = append(queue, n.Left)
			}
			index++
		}
		if index < len(node) {
			if node[index] != "null" {
				right_Val, _ := strconv.Atoi(node[index])
				n.Right = &TreeNode{Val: right_Val}
				queue = append(queue, n.Right)
			}
			index++
		}
	}
	return root
}

func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("null")
	}
	queue := []*TreeNode{root}
	ans := []int{}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			ans = append(ans, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	fmt.Println(ans)
}

func sortedArrayToBST(nums []string) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := n / 2
	val, _ := strconv.Atoi(nums[mid])
	root := &TreeNode{Val: val}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	ans := []int{}
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		if e.Value == nil {
			e = stack.Back()
			stack.Remove(e)
			node := e.Value.(*TreeNode)
			ans = append(ans, node.Val)
			continue
		}
		node := e.Value.(*TreeNode)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		stack.PushBack(node)
		stack.PushBack(nil)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return ans
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numLine := strings.Fields(scanner.Text())
	nums := []string{}
	for i := 0; i < len(numLine); i++ {
		nums = append(nums, numLine[i])
	}
	root := buildTree(nums)
	var targetSum int
	fmt.Scan(&targetSum)
	fmt.Println(pathSum(root, targetSum))
}
