package transaction

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

func New(sender, receiver string, amount float64) Transaction {
	return Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}
