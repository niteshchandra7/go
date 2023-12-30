package main

func makeEqual(words []string) bool {
	n:=len(words)
	mp := make(map[byte]int)
	for i:=0;i<len(words);i++{
		for j:=0;j<len(words[i]);j++{
			mp[words[i][j]]+=1
		}
	}
	for _,v:= range mp {
		if v%n!=0{
			return false
		}
	}
	return true
}

func main(){


}