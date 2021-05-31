package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}
type Fake struct {
}

func (wallet *Wallet) Pay(amount int) error {
	if wallet.Cash < amount {
		return fmt.Errorf("No money")
	}
	wallet.Cash -= amount
	return nil
}

func Buy(payer Payer) {
	switch payer.(type) {
	case *Wallet:
		fmt.Print("Наличка")
	case *Card:
		c, ok := payer.(*Card)
		if !ok {
			fmt.Println("неизвестный тип карты")
		}

		fmt.Println(c.CardHolder)
	default:
		fmt.Println("новый тип оплаты")
	}
	result := payer.Pay(100)
	if result != nil {
		panic(result)
	}
	fmt.Println("Success")
}

type Card struct {
	Balance    int
	ValidUntil string
	CardHolder string
	CVV        string
	Number     string
}

func (card *Card) Pay(amount int) error {
	if card.Balance < amount {
		return fmt.Errorf("No money")
	}
	card.Balance -= amount
	return nil
}

type ApplePay struct {
	Money   int
	AppleId string
}

func (ap *ApplePay) Pay(amount int) error {
	if ap.Money < amount {
		return fmt.Errorf("No money")
	}
	ap.Money -= amount
	return nil
}

type Ringer interface {
	Ring(string) error
}

type NFCPhone interface {
	Ringer
	Payer
}

type Phone struct {
	Money int
}

func (ap *Phone) Pay(amount int) error {
	if ap.Money < amount {
		return fmt.Errorf("No money")
	}
	ap.Money -= amount
	return nil
}

func (ap *Phone) Ring(num string) error {
	fmt.Println("Call number ", num)
	return nil
}

func PayInMerto(phone NFCPhone) {
	err := phone.Pay(40)
	if err != nil {
		fmt.Println("ошибка при оплате телефоном")
	}
	fmt.Println("OK METRO")
}

func main() {
	userWallet := Wallet{Cash: 500}
	Buy(&userWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 600, CardHolder: "Sberbank"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 700}
	Buy(myMoney)

	phone := Phone{Money: 50}
	PayInMerto(&phone)
}
