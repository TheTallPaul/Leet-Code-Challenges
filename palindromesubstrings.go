// 647. Palindromic Substrings
// Given a string, your task is to count how many palindromic substrings in this string.
// The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.
// https://leetcode.com/problems/palindromic-substrings/

// countSubstrings counts how many palindromic substrings are in the provided string
func countSubstrings(s string) int {
    total  := 0
    
    for i := 0; i < 2 * len(s) - 1; i++ {
        left  := i / 2
        right := left + i % 2
        
        for left >= 0 && right < len(s) && s[left] == s[right] {
            total = total + 1
            left  = left - 1
            right = right + 1
        }
    }
    
    return total
}
