package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources(){
	fmt.Println(PaymentSources([]types.Card{
		{
			PAN: `56555554444444478965`,
			Balance: 1_000_00,
			Active: true,
		},
	}))

	
	//Output
	//56555554444444478965

}
