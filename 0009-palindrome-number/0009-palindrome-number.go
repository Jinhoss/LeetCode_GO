func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(x int) bool {
    if x<0{
        return false
    } else {
        stringNum := strconv.Itoa(x)
        reversedNum := reverseString(stringNum)
        return stringNum == reversedNum
    }
    
}