package main
import "fmt"

func main(){
	var x int
	
	fmt.Scan(&x)

	if x == 17 {
		fmt.Println("YAY")
	} else {
		fmt.Println("NO")
	}
}