package morestring

// declaring another file in the same package, this can be accessed directly using 
// packagename.ConcatString 

func ConcatString(str1 string, str2 string) string {

	ans := []rune(str1)


	// slice = append(slice, anotherSlice...), cannot directly append slice to another(it will try to append [] to [] which will be like [[]], a type mismatch)
	//  so we need to unpack the slice using ... which is equivalent to replacing with s[0], s[1], s[2]...
	ans = append(ans, []rune(str2)... )

	return string(ans)
}