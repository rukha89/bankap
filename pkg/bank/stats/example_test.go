package stats

import (
	"github.com/coursar/bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000_00, 
		},
		{
			Amount: 5_000_00, 
		},
		{
			Amount: 6_000_00, 
		},
	}
	fmt.Println(Avg(payments))
	// Output: 700000
}
