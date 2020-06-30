/**
 * 23. Merge k Sorted Lists
 * 
 * Merge k sorted linked lists and return it as one sorted list.
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/
 */

// mergeKLists merges a slice of linked lists into one sorted link list
func mergeKLists(lists []*ListNode) *ListNode {
    headStruct := ListNode{-1, nil}
    currentStruct := ListNode{math.MaxInt64, nil}
    head, current := &headStruct, &currentStruct
    unterminatedLists := true
    currentIndex := 0
    
    for unterminatedLists {
        unterminatedLists = false
        current = &currentStruct
        for i, list := range lists {
            if list != nil && list.Val < current.Val {
                unterminatedLists = true
                current = list
                currentIndex = i
            }
        }
        if unterminatedLists {
            head.Next = current
            head = head.Next
            lists[currentIndex] = lists[currentIndex].Next
        }
    }
    
    return headStruct.Next
}
