package main

import (
	"fmt"
	"os"
)

func main() {
	var numbers [5]float64
	var sum float64

	fmt.Println("لطفاً 5 عدد را وارد کنید:")

	for i := 0; i < 5; i++ {
		fmt.Printf("عدد %d: ", i+1)
		_, err := fmt.Scan(&numbers[i])
		if err != nil {
			fmt.Println("خطا در خواندن عدد:", err)
			os.Exit(1)
		}
		sum += numbers[i]
	}

	average := sum / 5
	fmt.Printf("\nمیانگین اعداد: %.2f\n", average)
}

