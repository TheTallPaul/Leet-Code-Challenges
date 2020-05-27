/**
21. Merge Two Sorted Lists

Merge two sorted linked lists and return it as a new list. 
The new list should be made by splicing together the nodes of the first two lists.

https://leetcode.com/problems/merge-two-sorted-lists/
**/

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function(l1, l2) {
    let merged_list = new ListNode()
    let merged_list_head = merged_list
    
    while (l1 || l2) {
        if (!l2 || (l1 && l1.val < l2.val)) {
            merged_list.next = l1
            l1 = l1.next
        } else {
            merged_list.next = l2
            l2 = l2.next
        }
        
        merged_list = merged_list.next
    }
    
    return merged_list_head.next
};
