package stats

import (
	"github.com/usmon1983/myDZ/pkg/bank/stats"
	"fmt"
)

func ExampleAvg()  {
	payments := []types.Payment {
		{
			Amount: 10_000_00,
			Category: "Рестораны",
		},
		{
			Amount: 20_000_00,
			Category: "Рестораны",
		},
		{
			Amount: 30_000_00,
			Category: "Рестораны",
		},
	}
	PaymentAvg := Avg(payments)

	for _, payment := range PaymentAvg {
		fmt.Println(payment)
	}

	//Output: 20_000_00
}