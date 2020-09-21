package stats

import (
	"github.com/usmon1983/myDZ/pkg/bank/types"
)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	//payments_in :=  []types.Payment{}
	sum := 0
	next := 0
	for _, payment := range payments {
		sum += int(payment.Amount)
		next += 1
	}
	AVG := sum/next
	return types.Money(AVG)
}