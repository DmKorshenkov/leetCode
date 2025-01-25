package main

//Given an integer x, return true if x is apalindrome, and false otherwise.

/*Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/
import "fmt"
func main(){
}

func isPalindrom(num int) bool{
	if num < 0 {
		return false
	}

	var arr []int
	arr = make([]int, 0, 11)


	for num != 0 {
		arr = append(arr, num % 10)
		if len(arr) == 1 && arr[0] == 0 {
			return false
		}
		//num = (num - (num % 10)) / 10
		num /= 10
	}
	
	for i:= 0; i < len(arr) / 2 ;i++{
		if arr[i] != arr[len(arr) - 1 - i]{
			return false
		}
	}
	return true
}



















//
