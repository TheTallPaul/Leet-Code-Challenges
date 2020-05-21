# First Unique Character in a String
# Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
# https://leetcode.com/problems/first-unique-character-in-a-string/

class Solution:
    def firstUniqChar(self, s: str) -> int:
        """
        :type s: str
        :rtype: int
        """

        charMap = {}
        
        for char in s:
            if char in charMap:
                charMap[char] = charMap[char] + 1
            else:
                charMap[char] = 1
        
        for i, char in enumerate(s):
            if charMap[char] == 1:
                return i
            
        return -1
