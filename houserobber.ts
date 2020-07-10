/**
 * 198. House Robber
 *
 * You are a professional robber planning to rob houses along a street. 
 * Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
 * Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.
 *
 * https://leetcode.com/problems/house-robber/
**/

function rob(nums: number[]): number {
    let totalA = 0;
    let totalB = 0;
    
    for (let i = 0; i < nums.length; i++) {
        if (i%2 == 0) { // even
            totalA = Math.max(totalA + nums[i], totalB);
        } else { // odd
            totalB = Math.max(totalA, totalB + nums[i]);
        }
    }
    
    return Math.max(totalA, totalB);
};
