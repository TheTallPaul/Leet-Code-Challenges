/*
 * 674. Longest Continuous Increasing Subsequence
 *
 * Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).
 * 
 * https://leetcode.com/problems/longest-continuous-increasing-subsequence/
**/

import "math"

// findLengthOfLCIS finds the length of the longest continuous increasing subsequence
func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    longestLength := 1
    currentLength := 1
    prevNum := math.MaxInt64
    
    for _, num := range nums {
        if num > prevNum {
            currentLength = currentLength + 1
            if currentLength > longestLength {
                longestLength = currentLength
            }
        } else {
            currentLength = 1
        }
        prevNum = num
    }
    
    return longestLength
}
