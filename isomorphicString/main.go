package main

func main() {
	print(isIsomorphic("ab", "bb"))

}
func isIsomorphic(s string, t string) bool {

	for now := 0; now < len(s); now++ {
		for next := now + 1; next < len(s); next++ {
			if s[now] == s[next] && t[now] != t[next] ||
				s[now] != s[next] && t[now] == t[next] {
				return false
			}
		}
	}
	//	fmt.Println(checkS, checkT)
	return true
}
