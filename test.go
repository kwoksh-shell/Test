package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummy.Next
}

func buildInterSectList(interSectval, skipA, skipB int, l1, l2 []int) (*ListNode, *ListNode) {
	var commonHead *ListNode
	if interSectval != 0 {
		commonNum := l1[skipA:]
		commonHead = buildList(commonNum)
	}
	headA := buildList(l1[:skipA])
	if headA != nil && commonHead != nil {
		curA := headA
		for curA.Next != nil {
			curA = curA.Next
		}
		curA.Next = commonHead
	} else if commonHead != nil {
		headA = commonHead
	}
	headB := buildList(l2[skipB:])
	if headB != nil && commonHead != nil {
		curB := headB
		for curB.Next != nil {
			curB = curB.Next
		}
		curB.Next = commonHead
	} else if commonHead != nil {
		headB = commonHead
	}
	return headA, headB
}

func Check(l1, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = l2
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = l1
		}
	}
	return p1
}

func PrintList(head *ListNode) {
	outPut := []string{}
	for head != nil {
		outPut = append(outPut, strconv.Itoa(head.Val))
		head = head.Next
	}
	fmt.Println(strings.Join(outPut, "->") + "->nil")
}

func buildCycleList(head *ListNode, pos int) *ListNode {
	if head == nil || pos < 0 {
		return head
	}
	lastNode := head
	listLen := 1
	for lastNode.Next != nil {
		lastNode = lastNode.Next
		listLen++
	}
	if pos >= listLen {
		return head
	}
	cycleEntry := head
	for i := 0; i < pos; i++ {
		cycleEntry = cycleEntry.Next
	}
	lastNode.Next = cycleEntry
	return head
}

func HasCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	hasCycle := false
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}
	if hasCycle == false {
		return nil
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numLine := strings.Fields(scanner.Text())
	nums := []int{}
	for i := 0; i < len(numLine); i++ {
		num, _ := strconv.Atoi(numLine[i])
		nums = append(nums, num)
	}
	head := buildList(nums)
	PrintList(head)
	var pos int
	fmt.Scan(&pos)
	l1 := buildCycleList(head, pos)
	fmt.Println(HasCycle(l1).Val)

	/* 相交链表
	scanner.Scan()
	numLine2 := strings.Fields(scanner.Text())
	nums2 := []int{}
	for i := 0; i < len(numLine2); i++ {
		num, _ := strconv.Atoi(numLine2[i])
		nums2 = append(nums2, num)
	}

	var interSectval, skipA, skipB int
	fmt.Scan(&interSectval, &skipA, &skipB)
	l1, l2 := buildInterSectList(interSectval, skipA, skipB, nums, nums2)
	PrintList(Check(l1, l2))
	*/

}
