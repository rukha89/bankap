package card

import (
	"bank/pkg/bank/types"
)



func PaymentSources(cards []types.Card) []types.PaymentSource {
	var payment []types.PaymentSource
	for _, card := range cards{
		if card.Active == true && card.Balance > 0 {
			PAN := string(card.PAN)

			payment = append(payment, types.PaymentSource{
				Type: `card`,
				Number: PAN,
				Balance: card.Balance,
			})

		}
	}
	return payment
}
