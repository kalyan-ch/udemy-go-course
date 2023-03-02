package main

func main() {
    cards := deck{newCard(), "AD"}
    cards = append(cards, "6S")
    cards.print()
}

func newCard() string {
    return "AS"
}
