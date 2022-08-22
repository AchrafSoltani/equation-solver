package lib

import "fmt"

func Monomial() {
	var op string
	var a, b, c, y int
	fmt.Println("=======================")
	fmt.Println("= Monomial solver 0.1 =")
	fmt.Println("=======================")
	fmt.Println("Please select an operation:")
	fmt.Println("(1) Slope-intercept form: y=ax+b")
	fmt.Println("(2) Standard form: ax+by=c")
	fmt.Print(">> ")
	fmt.Scanf("%s", &op)
	switch op {
	case "1":
		fmt.Print(">> a = ")
		fmt.Scanf("%d", &a)
		fmt.Print(">> b = ")
		fmt.Scanf("%d", &b)
		fmt.Print(">> y = ")
		fmt.Scanf("%d", &y)
		solveSlopeInterceptForm(a, b, y)
	case "2":
		fmt.Print(">> a = ")
		fmt.Scanf("%d", &a)
		fmt.Print(">> b = ")
		fmt.Scanf("%d", &b)
		fmt.Print(">> c = ")
		fmt.Scanf("%d", &c)
		solveStandardForm(a, b, c)
	}

}

func solveSlopeInterceptForm(a int, b int, y int) {
	var x float64 = float64((y - b) / a)
	fmt.Printf("the solution to %dx + %d = %d is x = %.2f \n", a, b, y, x)
}

func solveStandardForm(a int, b int, c int) {
	var x, y float64
	x = float64(c) / float64(a)
	y = float64(c) / float64(b)
	fmt.Printf("the solution to %dx + %dy = %d is x = %.2f and y = %.2f \n", a, b, c, x, y)
}
