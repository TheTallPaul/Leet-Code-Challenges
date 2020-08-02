/*
 * 537. Complex Number Multiplication
 *
 * Given two strings representing two complex numbers.
 * You need to return a string representing their multiplication. Note i2 = -1 according to the definition.
 *
 * https://leetcode.com/problems/complex-number-multiplication/
**/

import "strconv"

// complexNumberMultiply takes two strings of the form "a+bi" and returns the result in that form
func complexNumberMultiply(complexNum1 string, complexNum2 string) string {
    realString1, imaginaryString1, _ := splitStringAtByte(complexNum1, '+')
    realString2, imaginaryString2, _ := splitStringAtByte(complexNum2, '+')
        
    // Extract the integers from the split strings
    realNum1, _ := strconv.Atoi(realString1)
    realNum2, _ := strconv.Atoi(realString2)
    imaginaryNum1, _ := strconv.Atoi(imaginaryString1[:len(imaginaryString1)-1])
    imaginaryNum2, _ := strconv.Atoi(imaginaryString2[:len(imaginaryString2)-1])
    
    // Solve for the two halves of the multiplication
    realAnswer := realNum1 * realNum2 - imaginaryNum1 * imaginaryNum2
    imaginaryAnswer := realNum1 * imaginaryNum2 + realNum2 * imaginaryNum1
    
    return strconv.Itoa(realAnswer) + "+" + strconv.Itoa(imaginaryAnswer) + "i"
}

// splitStringAtByte returns two substrings of the provided string to the before and after of the provided byte.
// If the byte is not present in the string, the entire string is returned.
// A bool indicating whether the split was a success or not is also returned.
func splitStringAtByte(fullString string, keyChar byte) (string, string, bool) {
    for index, char := range fullString {
        if char == rune(keyChar) {
            return fullString[:index], fullString[index+1:], true
        }
    }
    
    return fullString, fullString, false
}
