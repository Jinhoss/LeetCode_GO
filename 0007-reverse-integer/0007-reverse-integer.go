func reverseInt(x int) int {
	result := int(0)
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	return result
}

func reverse(x int) int {
    reversed := reverseInt(x)
    if reversed > math.MaxInt32 || reversed < math.MinInt32 {
        return 0} else{
        return reversed
    }
}