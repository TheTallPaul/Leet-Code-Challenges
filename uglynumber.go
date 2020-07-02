/* 
* 264. Ugly Number II
*
* Write a program to find the n-th ugly number.
*
* https://leetcode.com/problems/ugly-number-ii/
*/

// insertNum inserts an int into a sorted list if the number has not been inserted yet
func insertNum(list []int, num int) []int {
    // Find index of where number should go 
    index := sort.Search(len(list), func(i int) bool { return list[i] >= num })
    // Nil or empty slice or after last element
    if index >= len(list) { 
        list =  append(list, num)
    } else if list[index] != num { // Match not found at position
        list = append(list[:index+1], list[index:]...)
        list[index] = num
    }
    
    return list
}

// nthUglyNumber finds the n'th ugly number
func nthUglyNumber(n int) int {
    unusedNums := []int{1}
    ugliestNum := 1
    uglyNumCount := 0
    
    for uglyNumCount < n {
        // Pop off the smallest unused number
        ugliestNum = unusedNums[0]
        unusedNums = unusedNums[1:]
        uglyNumCount = uglyNumCount + 1
        
        // Add all of the products of the new number to the unused list
        unusedNums = insertNum(unusedNums, 2*ugliestNum)
        unusedNums = insertNum(unusedNums, 3*ugliestNum)
        unusedNums = insertNum(unusedNums, 5*ugliestNum)
    }
    
    return ugliestNum
}
