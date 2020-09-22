package types

type Money int64

type Currency string

//Коды валют
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
)

type PAN string

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

//Category редставляет собой категорию, в который был совершен платеж (авто, аптеки, рестораны и т.д.).
type Category string

//Payment предяставляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}

type PaymentSource struct {
	Type    string // 'card'
	Number  string // номер вида '5058 xxxx xxxx 8888'
	Balance Money  // баланс в дирамах
}
