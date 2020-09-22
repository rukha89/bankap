package stats

import (
	"github.com/coursar/bank/pkg/bank/types"
)


// TotalInCategory находит сумму покупок в определенной категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}
