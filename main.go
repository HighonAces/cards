package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getDeck(w http.ResponseWriter, r *http.Request) {
	cards := newDeck()
	cards.shuffle()
	fmt.Printf("got request to getDeck \n")
	io.WriteString(w, cards.toString())
}

func Deal(w http.ResponseWriter, r *http.Request) {
	cards := newDeck()
	cards.shuffle()
	dealtCards, _ := deal(cards, 2)
	fmt.Printf("got request to Deal \n")
	io.WriteString(w, dealtCards.toString())
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Nothing on /, use /deal or /getdeck\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/deal", Deal)
	http.HandleFunc("/getdeck", getDeck)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()
}
