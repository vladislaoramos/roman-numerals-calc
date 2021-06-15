package main

import (
	"bufio"
	"fmt"
	"os"
	"roman-calculator/solution"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	calc := solution.Calculator{}
	for sc.Scan() {
		expr := sc.Text()
		calc.Init(expr)
		res, err := calc.Calculate()
		if err == nil {
			fmt.Println(res)
		} else {
			fmt.Println(err)
		}
	}
}

//1: (V + I) * (X - I) + (M - C * L - D / X + Z) // OK
//2: (()) // FAIL
//3: (() // FAIL
//4: ()) // FAIL
//5: x+i // FAIL (NOT CAPS)
//6: (X + I) * (M + Z) - (C / D + I * M * X) // OK
//7: ((X + I) * (M + Z) - (C / D + I * M * X)) // OK
//8: () * (X + L) // FAIL
//9: (Z) / (X + L) // OK
//10: I + I // OK
//11: (X+M)*I // OK
//12: (X+M)*V // FAIL: >3999
//13: XX/(-VI) // FAIL
//14: XX/-VI // FAIL
