package stats

import (
	"github.com/usmon1983/myDZ/pkg/bank/stats"
	"fmt"
)

func ExampleAvg()  {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 10_000_00,
			Category: "Рестораны",
		},
		{
			ID: 2,
			Amount: 20_000_00,
			Category: "Рестораны",
		},
		{
			ID: 3,
			Amount: 30_000_00,
			Category: "Рестораны",
		},
	}
	PaymentAvg := Avg(payments)

	fmt.Println(PaymentAvg)
	}

	//Output: 20_000_00
}