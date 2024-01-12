package repository

func FindCardByNumber(cardNumber string) (Card, error) {
	return Card{
		Number: "1234123412341234",
		CVV: "123",
		ExpirationDate: "10/22",
	}, nil
}

func SaveCard(card Card) error {
	return nil
}