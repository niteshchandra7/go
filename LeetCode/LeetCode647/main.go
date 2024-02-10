package main

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += getCountEvenAndOdd(&s, i)
	}
	return count
}

func getCountEvenAndOdd(s *string, idx int) int {
	i, j := idx, idx
	count := getCount(s, i, j)
	i, j = idx-1, idx
	if i >= 0 {
		count += getCount(s, i, j)
	}
	return count
}

func getCount(s *string, i int, j int) int {
	count := 0
	for i >= 0 && j < len(*s) {
		if (*s)[i] != (*s)[j] {
			break
		}
		count++
		i--
		j++
	}
	return count
}

func main() {

}
