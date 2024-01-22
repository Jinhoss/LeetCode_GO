func convert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
		return s
	}

	result := make([]string, numRows)
	index, step := 0, 1

	for _, char := range s {
		result[index] += string(char)

		if index == 0 {
			step = 1
		} else if index == numRows-1 {
			step = -1
		}

		index += step
	}
    
    convertResult:="";
    for  _, subStr := range(result){
        convertResult+=subStr
    }
	return convertResult
}