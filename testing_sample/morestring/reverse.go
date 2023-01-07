package morestring


// Function to reverse a string, this function will be exported and can be used in other packages using morestring
func ReverseString(s string) string{

	ans := []rune{}

	strarr := []rune(s) // s is collection of characters and each char is int32 or rune

	for i := len(strarr)-1 ; i>=0; i-- {
		ans = append(ans, strarr[i])
	}

	return string(ans)
}