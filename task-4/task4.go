package main
import "fmt"


func main(){
	var n int
	fmt.Scanf("%d\n", &n)
	nums := make([][]int, n)
	for i :=0; i<n; i++ {
		nums[i] = make([]int, n)
	}

	for i:=0; i<n;i++{
		for j:=0; j<n;j++{
			fmt.Scanf("%d",&nums[i][j])
			
		}
		fmt.Scanf("\n")
	}

	flag:= true
	for i:=0; i<n;i++{
		for j:=0; j <n; j++{
			if(nums[i][j] != nums[n-1-i][n-1-j]){
				flag = false
				break
			}
		}
		if !flag{
			fmt.Printf("NO")
			break
		}
		
	}	
	if flag{
		fmt.Printf("YES")
	}

}
