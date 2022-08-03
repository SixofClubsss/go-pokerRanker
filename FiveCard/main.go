/*
	dReam Tables Poker Hand Ranker
	Originally written in C++ and used in dReam Tables Dero Poker Tables
	https://dreamtables.net
	Translated to go
*/

package main

import (
	"fmt"

	"github.com/fatih/color"
)

var control int
var confirm string
var skip bool = true

func handToText(rank int) string { /// Hand rank text
	var handRankText string
	switch rank {
	case 1:
		handRankText = "Royal Flush"
	case 2:
		handRankText = "Straight Flush"
	case 3:
		handRankText = "Four of a Kind"
	case 4:
		handRankText = "Full House"
	case 5:
		handRankText = "Flush"
	case 6:
		handRankText = "Straight"
	case 7:
		handRankText = "Three of a Kind"
	case 8:
		handRankText = "Two Pair"
	case 9:
		handRankText = "Pair"
	case 10:
		handRankText = "High Card"
	}
	return handRankText
}

func cardEquiv(card int) string { /// Card #s to unicode cards
	var cardEquiv string
	switch card {
	case 1:
		cardEquiv = "🂡"
	case 2:
		cardEquiv = "🂢"
	case 3:
		cardEquiv = "🂣"
	case 4:
		cardEquiv = "🂤"
	case 5:
		cardEquiv = "🂥"
	case 6:
		cardEquiv = "🂦"
	case 7:
		cardEquiv = "🂧"
	case 8:
		cardEquiv = "🂨"
	case 9:
		cardEquiv = "🂩"
	case 10:
		cardEquiv = "🂪"
	case 11:
		cardEquiv = "🂫"
	case 12:
		cardEquiv = "🂭"
	case 13:
		cardEquiv = "🂮"
	case 14:
		cardEquiv = "🂱"
	case 15:
		cardEquiv = "🂲"
	case 16:
		cardEquiv = "🂳"
	case 17:
		cardEquiv = "🂴"
	case 18:
		cardEquiv = "🂵"
	case 19:
		cardEquiv = "🂶"
	case 20:
		cardEquiv = "🂷"
	case 21:
		cardEquiv = "🂸"
	case 22:
		cardEquiv = "🂹"
	case 23:
		cardEquiv = "🂺"
	case 24:
		cardEquiv = "🂻"
	case 25:
		cardEquiv = "🂽"
	case 26:
		cardEquiv = "🂾"
	case 27:
		cardEquiv = "🃑"
	case 28:
		cardEquiv = "🃒"
	case 29:
		cardEquiv = "🃓"
	case 30:
		cardEquiv = "🃔"
	case 31:
		cardEquiv = "🃕"
	case 32:
		cardEquiv = "🃖"
	case 33:
		cardEquiv = "🃗"
	case 34:
		cardEquiv = "🃘"
	case 35:
		cardEquiv = "🃙"
	case 36:
		cardEquiv = "🃚"
	case 37:
		cardEquiv = "🃛"
	case 38:
		cardEquiv = "🃝"
	case 39:
		cardEquiv = "🃞"
	case 40:
		cardEquiv = "🃁"
	case 41:
		cardEquiv = "🃂"
	case 42:
		cardEquiv = "🃃"
	case 43:
		cardEquiv = "🃄"
	case 44:
		cardEquiv = "🃅"
	case 45:
		cardEquiv = "🃆"
	case 46:
		cardEquiv = "🃇"
	case 47:
		cardEquiv = "🃈"
	case 48:
		cardEquiv = "🃉"
	case 49:
		cardEquiv = "🃊"
	case 50:
		cardEquiv = "🃋"
	case 51:
		cardEquiv = "🃍"
	case 52:
		cardEquiv = "🃎"
	}
	return cardEquiv
}

func showMenu() { /// Start Menu
	dReamText := color.New(color.Bold, color.BgBlack).PrintlnFunc()
	optionTextBlack := color.New(color.BgWhite, color.FgBlack).PrintFunc()
	optionTextRed := color.New(color.BgWhite, color.FgRed).PrintFunc()
	fmt.Println()
	dReamText("  dReam Tables Poker Hand Ranker  ")
	///fmt.Println()
	for {
		optionTextBlack("♠")
		optionTextRed("♥")
		optionTextBlack("♣")
		optionTextRed("♦")
		optionTextBlack("     Select Option #      ")
		optionTextRed("♦")
		optionTextBlack("♣")
		optionTextRed("♥")
		optionTextBlack("♠\n")
		fmt.Println()
		fmt.Println("1. Compare Hands")
		fmt.Println("9. Exit")
		fmt.Print("\nOption: ")
		fmt.Scan(&control)
		if control == 1 || control == 9 {
			break
		}
		fmt.Println("Invalid Option.")
		fmt.Println()
	}
}

func another() {

	if !skip {
		fmt.Print("Compare more hands? ")
		optionTextGreen := color.New(color.FgGreen).PrintFunc()
		optionTextRed := color.New(color.FgRed).PrintlnFunc()
		optionTextGreen("y")
		fmt.Print("/")
		optionTextRed("n")
		for {
			fmt.Scan(&confirm)

			if confirm == "y" || confirm == "Y" {
				control = 1
				break

			} else if confirm == "n" || confirm == "N" {
				control = 9
				break
			} else {
				fmt.Println("Enter y / n")
			}

		}
	}
}

func main() {
	showMenu()
	for {
		another()
		skip = false
		if control == 9 { /// Check for exit
			sign := color.New(color.Bold, color.BgWhite, color.FgRed).PrintlnFunc()
			sign("Exiting... 🃖 ")
			break
		}
		control = 0
		clearHands()
		for { /// User inputs how many hands
			fmt.Print("How many hands to Compare: ")
			fmt.Scan(&totalHands)
			if totalHands > 1 && totalHands < 7 {
				break
			}
			fmt.Println("Enter a Number between 2 and 6") /// If invalid
		}

		fmt.Println("\nCard values are 1 through 52, in suited order.")
		fmt.Println("Spades = 1(A) to 13(K), Hearts = 14(A) to 26(K), \nClubs = 27(A) to 39(K), Diamonds = 40(A) to 52(K). \n\nExamples of input: ")
		fmt.Println("Royal Flush (Spades) 1 13 12 11 10")
		fmt.Println("Full House (Aces over two's) 1 14 27 2 15")
		fmt.Println("Hand combinations can be input in any order")
		fmt.Println()
		getHands(totalHands)
		fmt.Println()
		compareAll()

	}

}
