func lengthOfLongestSubstring(s string) int {
	subStr := ""
	maxLength := 0

	for _, char := range s {
		if index := strings.IndexRune(subStr, char); index != -1 {
			subStr = subStr[index+1:]
		}
		subStr += string(char)
		if maxLength < len(subStr) {
			maxLength = len(subStr)
		}
	}
	return maxLength
}