package stats

import (
	"github.com/coursar/bank/pkg/bank/types"
	"fmt"
)



func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount: 10_000_00,
			Category: `A`, 
		},
		{
			Amount: 5_000_00, 
			Category: `B`,
		},
		{
			Amount: 6_000_00, 
			Category: `A`,
		},
	}
	fmt.Println(TotalInCategory(payments, `A`))
	// Output: 1600000
}
