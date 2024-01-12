package repository

import "C0lliNN/better-handlers-blog/processor"

type CardRepository struct {}

func NewCardRepository() *CardRepository {
	return &CardRepository{}
}

func (*CardRepository) FindCardByNumber(cardNumber string) (processor.Card, error) {
	return processor.Card{
		Number: "1234123412341234",
		CVV: "123",
		ExpirationDate: "10/22",
	}, nil
}

func (*CardRepository) SaveCard(card processor.Card) error {
	return nil
}