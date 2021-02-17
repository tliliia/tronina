func reverse(s []int) {
	fmt.Println(s)
	if len(s) > 1 {
		var i int
		var j int
		j = len(s) - 1
		for i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
}
func SolutionRotate(A []int, k int) []int {
	/*
	   reverse(0, i-1)   /* cba|defgh
	   reverse(i, n-1)   /* cba|hgfed
	   reverse(0, n-1)   /* defgh|abc
	*/

	if k > len(A) {
		k = k % len(A)
	}
	reverse(A[:len(A)-k])
	reverse(A[len(A)-k:])
	reverse(A)
	return A
}