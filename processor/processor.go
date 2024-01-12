package processor

import "fmt"

type cardRepository interface {
	FindCardByNumber(cardNumber string) (Card, error)
	SaveCard(card Card) error
}

type Processor struct {
	cardRepository cardRepository
}

func NewProcessor(cardRepository cardRepository) *Processor {
	return &Processor{
		cardRepository: cardRepository,
	}
}

type TransactionRequest struct {
	CardNumber string `json:"card_number"`
	CVV string `json:"cvv"`
	ExpirationDate string `json:"expiration_date"`
	Amount int `json:"amount"`
}

func (p *Processor) ProcessTransaction(req TransactionRequest) error {
	card, err := p.cardRepository.FindCardByNumber(req.CardNumber)
	if err != nil {
		return err
	}

	if card.CVV != req.CVV || card.ExpirationDate != req.ExpirationDate  {
		return fmt.Errorf("invalid cvv or exp date")
	}

	if card.AvaliableLimit < req.Amount {
		return fmt.Errorf("insufficient funds")
	}

	card.AvaliableLimit -= req.Amount
	if err := p.cardRepository.SaveCard(card); err != nil {
		return err
	}

	return nil
}
