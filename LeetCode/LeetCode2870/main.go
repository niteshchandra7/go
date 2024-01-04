package main

import "math"

func minOperations(nums []int) int {
	mp :=make(map[int]int)
	cnt:=0
	ans:=0
	for i:=0;i<len(nums);i++{
		mp[nums[i]]+=1
		cnt = max(cnt,mp[nums[i]])
	}
	dp :=make([]int,cnt+1)
	for i:=0;i<len(dp);i++{
		dp[i] = math.MaxInt
	}
	for _,v:= range mp{
		if v==1 {
			return -1
		}
		temp := minOperationByCount(v,dp)
		ans+=temp
	}
	return ans   
}

func minOperationByCount(cnt int, dp []int) int{
	if cnt==1 {
		return math.MaxInt-1
	} else if cnt==0{
		return 0
	} else if cnt==2 || cnt==3{
		return 1
	} else if dp[cnt]!=math.MaxInt{
		return dp[cnt]
	}
	dp[cnt] =  min(1+minOperationByCount(cnt-2,dp),1+minOperationByCount(cnt-3,dp))
	return dp[cnt]
}


func main(){

}