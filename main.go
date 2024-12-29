package main	
import "fmt"
func getsum(num int)int{
	if(num==0){
		return 0
	}
	return num+getsum(num-1)
}
func main(){
	res:=getsum(5)
	fmt.Println(res)
}
