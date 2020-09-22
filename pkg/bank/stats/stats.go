package stats

import (
	"github.com/coursar/bank/pkg/bank/types"
)



func Avg(payments []types.Payment) types.Money{
	avg := types.Money(0)
	sum := types.Money(0)
	cnt := types.Money(0)
	for _, payment := range payments {
		sum += payment.Amount
		cnt += types.Money(1)
	}
	avg = sum / cnt
	return avg
	
}