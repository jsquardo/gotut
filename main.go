package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six od Spades")

	cards.print()
}

func newCard() string {
	return "Five od Diamonds"
}
