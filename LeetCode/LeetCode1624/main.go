package main

func maxLengthBetweenEqualCharacters(s string) int {
	mp := make(map[byte]int)
	ans := -1
	for i := 0; i < len(s); i++ {
		val, ok := mp[s[i]]
		if !ok {
			mp[s[i]] = i
		} else {
			if i-val-1 > ans {
				ans = i - val - 1
			}
		}
	}
	return ans
}
func main() {

}
