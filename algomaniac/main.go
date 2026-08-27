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

// Algomaniac Finals
// Algomaniac finals, a part of Convolution Fest of Jadavpur University, will be held on March 1717.
// Shreyan can only go to Jadavpur University on March XX.
// Print YAY if he can attend the Algomaniac finals and NO if he cannot.
// Input Format
// The first and only line of input contains one integer, XX, the day of march Shreyan can go to Jadavpur University. 
