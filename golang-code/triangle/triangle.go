package main

const lens int = 10

func YangHuiTriangle(){
	nums :=[]int{}
	var value int
	for i:=0;i<lens;i++{
		for j:=0;j<lens-i;j++{
			print(" ")
		}
		for j:=0;j<i+1;j++{
			length := len(nums)
			if j==0||j==i{
				value = 1
			}else {
				value = nums[length-i]+nums[length-i-1]
			}
			nums = append(nums,value)
			print(value," ")
		}
		println("")
	}
}
func main(){
	YangHuiTriangle()
}
