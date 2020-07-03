/*
* 268. Missing Number
*
* Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
*
* https://leetcode.com/problems/missing-number/
*/

function missingNumber(nums: number[]): number {
    var total = nums.reduce(function(a, b){ return a + b; });
    
    return (nums.length + 1)*nums.length/2 - total;
};
