package main
import "fmt"

func main(){

	var n int

	fmt.Scanf("%d", &n)
	nums:=make([]int, n)

	for i:=0; i<n; i++ {
		fmt.Scanf("%d",&nums[i])
	}
	lastElem := nums[n-1]
	for i := n - 1; i > 0; i-- {
		nums[i] = nums[i-1]
	} 
	nums[0] = lastElem

	for i:=0; i<n; i++{
		fmt.Printf("%d ",nums[i])
	}
}