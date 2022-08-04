/*
	dReam Tables Poker Hand Ranker
	Originally written in C++ and used in dReam Tables Dero Poker Tables
	https://dreamtables.net
	Translated to go
*/

package main

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
)

var errColor = color.New(color.Bold, color.FgRed).PrintlnFunc()

var totalHands int
var communityCard [5]int
var p1HandRaw [2]int
var p2HandRaw [2]int
var p3HandRaw [2]int
var p4HandRaw [2]int
var p5HandRaw [2]int
var p6HandRaw [2]int
var arrSplit [2]int

var pc1 []int
var pc2 []int
var cc1 [2]int
var cc2 [2]int
var cc3 [2]int
var cc4 [2]int
var cc5 [2]int
var p1HighPair int
var p2HighPair int
var p3HighPair int
var p4HighPair int
var p5HighPair int
var p6HighPair int
var p1Rank int
var p2Rank int
var p3Rank int
var p4Rank int
var p5Rank int
var p6Rank int
var fHighCardArr []int
var p1HighCardArr [5]int
var p2HighCardArr [5]int
var p3HighCardArr [5]int
var p4HighCardArr [5]int
var p5HighCardArr [5]int
var p6HighCardArr [5]int

func clearHands() { /// Clears hand arrays before new input
	p1HandRaw = [2]int{0, 0}
	p2HandRaw = [2]int{0, 0}
	p3HandRaw = [2]int{0, 0}
	p4HandRaw = [2]int{0, 0}
	p5HandRaw = [2]int{0, 0}
	p6HandRaw = [2]int{0, 0}
	communityCard = [5]int{0, 0, 0, 0, 0}
}

func getHands(totalHands int) { /// Gets card values for totalHands selected
	p1Rank = 100
	p1HighPair = 0
	p2Rank = 100
	p2HighPair = 0
	p3Rank = 100
	p3HighPair = 0
	p4Rank = 100
	p4HighPair = 0
	p5Rank = 100
	p5HighPair = 0
	p6Rank = 100
	p6HighPair = 0

	switch totalHands {
	case 2:
		getFlop()
		p1Rank = getHand1()
		p2Rank = getHand2()
	case 3:
		getFlop()
		p1Rank = getHand1()
		p2Rank = getHand2()
		p3Rank = getHand3()
	case 4:
		getFlop()
		p1Rank = getHand1()
		p2Rank = getHand2()
		p3Rank = getHand3()
		p4Rank = getHand4()
	case 5:
		getFlop()
		p1Rank = getHand1()
		p2Rank = getHand2()
		p3Rank = getHand3()
		p4Rank = getHand4()
		p5Rank = getHand5()
	case 6:
		getFlop()
		p1Rank = getHand1()
		p2Rank = getHand2()
		p3Rank = getHand3()
		p4Rank = getHand4()
		p5Rank = getHand5()
		p6Rank = getHand6()
	}

}

func getFlop() { /// Gets community cards
	for {
		fmt.Println("Input community cards: ")

		_, err := fmt.Scanln(&communityCard[0], &communityCard[1], &communityCard[2], &communityCard[3], &communityCard[4])
		if err != nil {
			if communityCard[4] == 0 { /// Did not input 5 cards
				errColor("Incorrect Format, Try Again")
			}

			if communityCard[4] > 0 { /// Input had extra card, removed
				errColor("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if communityCard[0] > 0 && communityCard[0] < 53 && communityCard[1] > 0 && communityCard[1] < 53 && communityCard[2] > 0 &&
			communityCard[2] < 53 && communityCard[3] > 0 && communityCard[3] < 53 && communityCard[4] > 0 && communityCard[4] < 53 { /// Ensure correct input format
			break
		} else {
			errColor("Enter numbers 1 - 52")
		}
	}

	suitSplit(communityCard[0])
	cc1 = [2]int{arrSplit[0], arrSplit[1]}

	suitSplit(communityCard[1])
	cc2 = [2]int{arrSplit[0], arrSplit[1]}

	suitSplit(communityCard[2])
	cc3 = [2]int{arrSplit[0], arrSplit[1]}

	suitSplit(communityCard[3])
	cc4 = [2]int{arrSplit[0], arrSplit[1]}

	suitSplit(communityCard[4])
	cc5 = [2]int{arrSplit[0], arrSplit[1]}
}

func getHand1() int { /// Splitting card value and suit for making individual hand, highcard value and rank
	for {
		fmt.Println("Input the card values for Player 1: ")

		_, err := fmt.Scanln(&p1HandRaw[0], &p1HandRaw[1])
		if err != nil {
			if p1HandRaw[1] == 0 { /// Did not input 5 cards
				errColor("Incorrect Format, Try Again")
			}

			if p1HandRaw[1] > 0 { /// Input had extra card, removed
				errColor("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p1HandRaw[0] > 0 && p1HandRaw[0] < 53 && p1HandRaw[1] > 0 && p1HandRaw[1] < 53 { /// Ensure correct input format
			break
		} else {
			errColor("Enter numbers 1 - 52")
		}
	}

	suitSplit(p1HandRaw[0])
	pc1 = []int{arrSplit[0], arrSplit[1]}

	suitSplit(p1HandRaw[1])
	pc2 = []int{arrSplit[0], arrSplit[1]}

	Rank := compareMine()
	copy(p1HighCardArr[:], fHighCardArr)

	if p1HighCardArr[0] == 1 {
		p1HighCardArr[0] = 14
	}

	if p1HighCardArr[1] == 1 {
		p1HighCardArr[1] = 14
	}

	if p1HighCardArr[2] == 1 {
		p1HighCardArr[2] = 14
	}

	if p1HighCardArr[3] == 1 {
		p1HighCardArr[3] = 14
	}

	if p1HighCardArr[4] == 1 {
		p1HighCardArr[4] = 14
	}

	p1HighPair = getHighPair(p1HighCardArr)

	return Rank
}

func getHand2() int {
	for {
		fmt.Println("Input the card values for Player 2: ")

		_, err := fmt.Scanln(&p2HandRaw[0], &p2HandRaw[1])
		if err != nil {
			if p2HandRaw[1] == 0 { /// Did not input 5 cards
				fmt.Println("Incorrect Format, Try Again")
			}

			if p2HandRaw[1] > 0 { /// Input had extra card, removed
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p2HandRaw[0] > 0 && p2HandRaw[0] < 53 && p2HandRaw[1] > 0 && p2HandRaw[1] < 53 { /// Ensure correct input format
			break
		} else {
			errColor("Enter numbers 1 - 52")
		}
	}

	suitSplit(p2HandRaw[0])
	pc1 = []int{arrSplit[0], arrSplit[1]}

	suitSplit(p2HandRaw[1])
	pc2 = []int{arrSplit[0], arrSplit[1]}

	Rank := compareMine()
	copy(p2HighCardArr[:], fHighCardArr)

	if p2HighCardArr[0] == 1 {
		p2HighCardArr[0] = 14
	}

	if p2HighCardArr[1] == 1 {
		p2HighCardArr[1] = 14
	}

	if p2HighCardArr[2] == 1 {
		p2HighCardArr[2] = 14
	}

	if p2HighCardArr[3] == 1 {
		p2HighCardArr[3] = 14
	}

	if p2HighCardArr[4] == 1 {
		p2HighCardArr[4] = 14
	}

	p2HighPair = getHighPair(p2HighCardArr)

	return Rank
}

func getHand3() int {
	for {
		fmt.Println("Input the card values for Player 3: ")

		_, err := fmt.Scanln(&p3HandRaw[0], &p3HandRaw[1])
		if err != nil {
			if p3HandRaw[1] == 0 { /// Did not input 5 cards
				fmt.Println("Incorrect Format, Try Again")
			}

			if p3HandRaw[1] > 0 { /// Input had extra card, removed
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p3HandRaw[0] > 0 && p3HandRaw[0] < 53 && p3HandRaw[1] > 0 && p3HandRaw[1] < 53 { /// Ensure correct input format
			break
		} else {
			errColor("Enter numbers 1 - 52")
		}
	}

	suitSplit(p3HandRaw[0])
	pc1 = []int{arrSplit[0], arrSplit[1]}

	suitSplit(p3HandRaw[1])
	pc2 = []int{arrSplit[0], arrSplit[1]}

	Rank := compareMine()
	copy(p3HighCardArr[:], fHighCardArr)

	if p3HighCardArr[0] == 1 {
		p3HighCardArr[0] = 14
	}

	if p3HighCardArr[1] == 1 {
		p3HighCardArr[1] = 14
	}

	if p3HighCardArr[2] == 1 {
		p3HighCardArr[2] = 14
	}

	if p3HighCardArr[3] == 1 {
		p3HighCardArr[3] = 14
	}

	if p3HighCardArr[4] == 1 {
		p3HighCardArr[4] = 14
	}

	p3HighPair = getHighPair(p3HighCardArr)

	return Rank
}

func getHand4() int {
	for {
		fmt.Println("Input the card values for Player 4: ")

		_, err := fmt.Scanln(&p4HandRaw[0], &p4HandRaw[1])
		if err != nil {
			if p4HandRaw[1] == 0 { /// Did not input 5 cards
				fmt.Println("Incorrect Format, Try Again")
			}

			if p4HandRaw[1] > 0 { /// Input had extra card, removed
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p4HandRaw[0] > 0 && p4HandRaw[0] < 53 && p4HandRaw[1] > 0 && p4HandRaw[1] < 53 { /// Ensure correct input format
			break
		} else {
			errColor("Enter numbers 1 - 52")
		}
	}

	suitSplit(p4HandRaw[0])
	pc1 = []int{arrSplit[0], arrSplit[1]}

	suitSplit(p4HandRaw[1])
	pc2 = []int{arrSplit[0], arrSplit[1]}

	Rank := compareMine()
	copy(p4HighCardArr[:], fHighCardArr)

	if p4HighCardArr[0] == 1 {
		p4HighCardArr[0] = 14
	}

	if p4HighCardArr[1] == 1 {
		p4HighCardArr[1] = 14
	}

	if p4HighCardArr[2] == 1 {
		p4HighCardArr[2] = 14
	}

	if p4HighCardArr[3] == 1 {
		p4HighCardArr[3] = 14
	}

	if p4HighCardArr[4] == 1 {
		p4HighCardArr[4] = 14
	}

	p4HighPair = getHighPair(p4HighCardArr)

	return Rank
}

func getHand5() int {
	for {
		fmt.Println("Input the card values for Player 5: ")

		_, err := fmt.Scanln(&p5HandRaw[0], &p5HandRaw[1])
		if err != nil {
			if p5HandRaw[1] == 0 { /// Did not input 5 cards
				fmt.Println("Incorrect Format, Try Again")
			}

			if p5HandRaw[1] > 0 { /// Input had extra card, removed
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p5HandRaw[0] > 0 && p5HandRaw[0] < 53 && p5HandRaw[1] > 0 && p5HandRaw[1] < 53 { /// Ensure correct input format
			break
		} else {
			errColor("Enter numbers 1 - 52")
		}
	}

	suitSplit(p5HandRaw[0])
	pc1 = []int{arrSplit[0], arrSplit[1]}

	suitSplit(p5HandRaw[1])
	pc2 = []int{arrSplit[0], arrSplit[1]}

	Rank := compareMine()
	copy(p5HighCardArr[:], fHighCardArr)

	if p5HighCardArr[0] == 1 {
		p5HighCardArr[0] = 14
	}

	if p5HighCardArr[1] == 1 {
		p5HighCardArr[1] = 14
	}

	if p5HighCardArr[2] == 1 {
		p5HighCardArr[2] = 14
	}

	if p5HighCardArr[3] == 1 {
		p5HighCardArr[3] = 14
	}

	if p5HighCardArr[4] == 1 {
		p5HighCardArr[4] = 14
	}

	p5HighPair = getHighPair(p5HighCardArr)

	return Rank
}

func getHand6() int {
	for {
		fmt.Println("Input the card values for Player 6: ")

		_, err := fmt.Scanln(&p6HandRaw[0], &p6HandRaw[1])
		if err != nil {
			if p6HandRaw[1] == 0 { /// Did not input 5 cards
				fmt.Println("Incorrect Format, Try Again")
			}

			if p6HandRaw[1] > 0 { /// Input had extra card, removed
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p6HandRaw[0] > 0 && p6HandRaw[0] < 53 && p6HandRaw[1] > 0 && p6HandRaw[1] < 53 { /// Ensure correct input format
			break
		} else {
			errColor("Enter numbers 1 - 52")
		}
	}

	suitSplit(p6HandRaw[0])
	pc1 = []int{arrSplit[0], arrSplit[1]}

	suitSplit(p6HandRaw[1])
	pc2 = []int{arrSplit[0], arrSplit[1]}

	Rank := compareMine()
	copy(p6HighCardArr[:], fHighCardArr)

	if p6HighCardArr[0] == 1 {
		p6HighCardArr[0] = 14
	}

	if p6HighCardArr[1] == 1 {
		p6HighCardArr[1] = 14
	}

	if p6HighCardArr[2] == 1 {
		p6HighCardArr[2] = 14
	}

	if p6HighCardArr[3] == 1 {
		p6HighCardArr[3] = 14
	}

	if p6HighCardArr[4] == 1 {
		p6HighCardArr[4] = 14
	}

	p6HighPair = getHighPair(p6HighCardArr)

	return Rank
}

func makeHand(h, s []int) int { /// Determines hand rank after suit slipt

	pHand := h
	pSuits := s

	sort.Ints(pHand)

	/// Royal flush
	if pHand[0] == 1 && pHand[1] == 10 && pHand[2] == 11 && pHand[3] == 12 && pHand[4] == 13 &&
		pSuits[0] == pSuits[1] && pSuits[0] == pSuits[2] && pSuits[0] == pSuits[3] && pSuits[0] == pSuits[4] {

		return 1

	}

	/// Straight flush
	if pHand[0]+1 == pHand[1] && pHand[1]+1 == pHand[2] && pHand[2]+1 == pHand[3] && pHand[3]+1 == pHand[4] && pHand[0]+4 == pHand[4] &&
		pSuits[0] == pSuits[1] && pSuits[0] == pSuits[2] && pSuits[0] == pSuits[3] && pSuits[0] == pSuits[4] {

		return 2
	}

	/// Four of a Kind
	if (pHand[0] == pHand[1] && pHand[1] == pHand[2] && pHand[2] == pHand[3]) ||
		(pHand[1] == pHand[2] && pHand[2] == pHand[3] && pHand[3] == pHand[4]) {

		return 3
	}

	/// Full House
	if (pHand[0] == pHand[1] && pHand[1] == pHand[2] && pHand[3] == pHand[4]) ||
		(pHand[0] == pHand[1] && pHand[2] == pHand[3] && pHand[3] == pHand[4]) {

		return 4
	}

	/// Flush
	if pSuits[0] == pSuits[1] && pSuits[0] == pSuits[2] && pSuits[0] == pSuits[3] && pSuits[0] == pSuits[4] {

		return 5
	}

	/// Straight
	if pHand[0]+1 == pHand[1] && pHand[1]+1 == pHand[2] && pHand[2]+1 == pHand[3] && pHand[3]+1 == pHand[4] && pHand[0]+4 == pHand[4] ||
		pHand[0] == 1 && pHand[1] == 10 && pHand[2] == 11 && pHand[3] == 12 && pHand[4] == 13 {
		return 6
	}

	/// Three of a Kind
	if (pHand[0] == pHand[1] && pHand[1] == pHand[2]) ||
		(pHand[1] == pHand[2] && pHand[2] == pHand[3]) ||
		(pHand[2] == pHand[3] && pHand[3] == pHand[4]) {
		return 7
	}

	/// Two Pair
	if (pHand[0] == pHand[1] && pHand[2] == pHand[3]) ||
		(pHand[1] == pHand[2] && pHand[3] == pHand[4]) ||
		(pHand[0] == pHand[1] && pHand[3] == pHand[4]) {
		return 8
	}

	/// Pair
	if pHand[0] == pHand[1] || pHand[0] == pHand[2] || pHand[0] == pHand[3] || pHand[0] == pHand[4] ||
		pHand[1] == pHand[2] || pHand[1] == pHand[3] || pHand[1] == pHand[4] ||
		pHand[2] == pHand[3] || pHand[2] == pHand[4] ||
		pHand[3] == pHand[4] {
		return 9
	} else {

		return 10
	}
}

func getHighPair(h [5]int) int { /// Gets high pair from hand

	var highPair int

	for i := 0; i < 4; i++ {
		if h[i] == h[i+1] {
			if h[i] > highPair {
				highPair = h[i]
			}
		}
	}

	return highPair
}

func isPaired(h []int) int { /// Gets high pair slice

	var isPaired int

	for i := 0; i < 4; i++ {
		if h[i] == h[i+1] {
			if h[i] > isPaired {
				isPaired = h[i]
			}
		}
	}

	return isPaired
}

func searchAllRanks(r int, h, s []int) []int { /// Rank specific outcome in search

	var swap = h
	var swapSuit = s
	hole := []int{pc1[0], pc2[0]}
	holeSuit := []int{pc1[1], pc2[1]}
	paired := isPaired(swap)
	/// Debug
	// red := color.New(color.FgHiRed).PrintlnFunc()
	// red("Comm: %s", swap)
	// cyan := color.New(color.FgHiCyan).PrintlnFunc()
	// cyan("Hole: %s", hole)
	// green := color.New(color.FgHiGreen).PrintlnFunc()
	// green("Rank: %s", r)

	/// If flush
	if r == 5 {
		if hole[0] > swap[4] && hole[1] > hole[0] && holeSuit[0] == swapSuit[0] && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[2], swap[3], swap[4], hole[0], hole[1]}
		} else if hole[0] > swap[2] && hole[1] > swap[3] && holeSuit[0] == swapSuit[0] && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[0], swap[1], hole[0], hole[1], swap[4]}
		} else if hole[0] > swap[1] && hole[1] > swap[2] && holeSuit[0] == swapSuit[0] && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[0], hole[0], hole[1], swap[3], swap[4]}
		} else if hole[0] > swap[0] && hole[1] > swap[1] && holeSuit[0] == swapSuit[0] && holeSuit[1] == swapSuit[0] {
			swap = []int{hole[0], hole[1], swap[2], swap[3], swap[4]}

		} else if hole[1] > swap[4] && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[1], swap[2], swap[3], swap[4], hole[1]}
		} else if hole[1] > swap[3] && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[1], swap[2], swap[3], hole[1], swap[4]}
		} else if hole[1] > swap[2] && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[1], swap[2], hole[1], swap[3], swap[4]}
		} else if hole[1] > swap[1] && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[1], hole[1], swap[2], swap[3], swap[4]}
		} else if hole[1] > swap[0] && holeSuit[1] == swapSuit[0] {
			swap = []int{hole[1], swap[1], swap[2], swap[3], swap[4]}
		} else if hole[1] == 1 && swap[4] <= 13 && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[1], swap[2], swap[3], swap[4], hole[1]}
		} else if hole[1] == 1 && swap[3] <= 13 && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[1], swap[2], swap[3], hole[1], swap[4]}
		} else if hole[1] == 1 && swap[2] <= 13 && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[1], swap[2], hole[1], swap[3], swap[4]}
		} else if hole[1] == 1 && swap[1] <= 13 && holeSuit[1] == swapSuit[0] {
			swap = []int{swap[1], hole[1], swap[2], swap[3], swap[4]}
		} else if hole[1] == 1 && swap[0] <= 13 && holeSuit[1] == swapSuit[0] {
			swap = []int{hole[1], swap[1], swap[2], swap[3], swap[4]}
		}

		if hole[0] > swap[4] && holeSuit[0] == swapSuit[0] {
			swap = []int{swap[1], swap[2], swap[3], swap[4], hole[0]}
		} else if hole[0] > swap[3] && holeSuit[0] == swapSuit[0] {
			swap = []int{swap[1], swap[2], swap[3], hole[0], swap[4]}
		} else if hole[0] > swap[2] && holeSuit[0] == swapSuit[0] {
			swap = []int{swap[1], swap[2], hole[0], swap[3], swap[4]}
		} else if hole[0] > swap[1] && holeSuit[0] == swapSuit[0] {
			swap = []int{swap[1], hole[0], swap[2], swap[3], swap[4]}
		} else if hole[0] > swap[0] && holeSuit[0] == swapSuit[0] {
			swap = []int{hole[0], swap[1], swap[2], swap[3], swap[4]}
		} else if hole[0] == 1 && swap[4] <= 13 && holeSuit[0] == swapSuit[0] {
			swap = []int{swap[1], swap[2], swap[3], swap[4], hole[0]}
		} else if hole[0] == 1 && swap[3] <= 13 && holeSuit[0] == swapSuit[0] {
			swap = []int{swap[1], swap[2], swap[3], hole[0], swap[4]}
		} else if hole[0] == 1 && swap[2] <= 13 && holeSuit[0] == swapSuit[0] {
			swap = []int{swap[1], swap[2], hole[0], swap[3], swap[4]}
		} else if hole[0] == 1 && swap[1] <= 13 && holeSuit[0] == swapSuit[0] {
			swap = []int{swap[1], hole[0], swap[2], swap[3], swap[4]}
		} else if hole[0] == 1 && swap[0] <= 13 && holeSuit[0] == swapSuit[0] {
			swap = []int{hole[0], swap[1], swap[2], swap[3], swap[4]}
		}
	}

	sort.Ints(hole)

	/// If high card
	if r == 10 {
		if hole[0] > swap[0] && swap[0] != 1 {
			swap[0] = hole[0]
		} else if hole[0] > swap[1] {
			swap[1] = hole[0]
		} else if hole[0] > swap[2] {
			swap[2] = hole[0]
		} else if hole[0] > swap[3] {
			swap[3] = hole[0]
		} else if hole[0] > swap[4] {
			swap[4] = hole[0]
		} else if hole[0] == 1 && swap[0] <= 13 {
			swap[0] = hole[0]
		} else if hole[0] == 1 && swap[1] <= 13 {
			swap[1] = hole[0]
		} else if hole[0] == 1 && swap[2] <= 13 {
			swap[2] = hole[0]
		} else if hole[0] == 1 && swap[3] <= 13 {
			swap[3] = hole[0]
		} else if hole[0] == 1 && swap[4] <= 13 {
			swap[4] = hole[0]
		}

		if hole[1] > swap[0] && swap[0] != 1 {
			swap[0] = hole[1]
		} else if hole[1] > swap[1] && swap[1] != hole[0] {
			swap[1] = hole[1]
		} else if hole[1] > swap[2] && swap[2] != hole[0] {
			swap[2] = hole[1]
		} else if hole[1] > swap[3] && swap[3] != hole[0] {
			swap[3] = hole[1]
		} else if hole[1] > swap[4] {
			swap[4] = hole[1]
		} else if hole[1] > swap[4] && swap[4] == hole[0] {
			swap = []int{swap[2], swap[3], swap[4], hole[0], hole[1]}
		} else if hole[1] == 1 && swap[0] <= 13 && swap[0] != 1 {
			swap[0] = hole[1]
		} else if hole[1] == 1 && swap[1] <= 13 && swap[1] != 1 {
			swap[1] = hole[1]
		} else if hole[1] == 1 && swap[2] <= 13 && swap[2] != 1 {
			swap[2] = hole[1]
		} else if hole[1] == 1 && swap[3] <= 13 && swap[3] != 1 {
			swap[3] = hole[1]
		} else if hole[1] == 1 && swap[4] <= 13 && swap[4] != 1 {
			swap[4] = hole[1]
		}

	}
	/// If pair, two pair, three of a kind, four of a kind
	if r == 9 || r == 8 || r == 7 || r == 3 {

		if hole[0] == hole[1] {
			if hole[0] > swap[0] && swap[0] == swap[1] && swap[1] != swap[2] {
				swap = []int{hole[0], hole[1], swap[2], swap[3], swap[4]}
			} else if hole[0] > swap[1] && swap[1] == swap[2] && swap[2] != swap[3] {
				swap = []int{swap[0], hole[0], hole[1], swap[3], swap[4]}
			} else if hole[0] > swap[2] && swap[2] == swap[3] && swap[3] != swap[4] {
				swap = []int{swap[0], swap[1], hole[0], hole[1], swap[4]}
			} else if hole[0] > swap[3] && swap[3] == swap[4] {
				swap = []int{swap[0], swap[1], swap[2], hole[0], hole[1]}
			}

		} else {
			if hole[0] > swap[4] && swap[4] != swap[3] && swap[4] != swap[2] && swap[4] != swap[1] && swap[4] != swap[0] {
				swap[4] = hole[0]
			} else if hole[0] > swap[3] && swap[3] != swap[4] && swap[3] != swap[2] && swap[3] != swap[1] && swap[3] != swap[0] {
				swap[3] = hole[0]
			} else if hole[0] > swap[2] && swap[2] != swap[4] && swap[2] != swap[3] && swap[2] != swap[1] && swap[2] != swap[0] {
				swap[2] = hole[0]
			} else if hole[0] > swap[1] && swap[1] != swap[4] && swap[1] != swap[3] && swap[1] != swap[2] && swap[1] != swap[0] {
				swap[1] = hole[0]
			} else if hole[0] > swap[0] && swap[0] != swap[4] && swap[0] != swap[3] && swap[0] != swap[2] && swap[0] != swap[1] && swap[0] != 1 {
				swap[0] = hole[0]
			} else if hole[0] == 1 && swap[4] <= 13 && swap[4] != swap[3] && swap[4] != swap[2] && swap[4] != swap[1] && swap[4] != swap[0] {
				swap[4] = hole[0]
			} else if hole[0] == 1 && swap[3] <= 13 && swap[3] != swap[4] && swap[3] != swap[2] && swap[3] != swap[1] && swap[3] != swap[0] {
				swap[3] = hole[0]
			} else if hole[0] == 1 && swap[2] <= 13 && swap[2] != swap[4] && swap[2] != swap[3] && swap[2] != swap[1] && swap[2] != swap[0] {
				swap[2] = hole[0]
			} else if hole[0] == 1 && swap[1] <= 13 && swap[1] != swap[4] && swap[1] != swap[3] && swap[1] != swap[2] && swap[1] != swap[0] {
				swap[1] = hole[0]
			} else if hole[0] == 1 && swap[0] <= 13 && swap[0] != swap[4] && swap[0] != swap[3] && swap[0] != swap[2] && swap[0] != swap[1] {
				swap[0] = hole[0]
			}

			if hole[1] > swap[4] && swap[4] != swap[3] && swap[4] != swap[2] && swap[4] != swap[1] && swap[4] != swap[0] && swap[4] != 1 && hole[1] != paired {
				swap[4] = hole[1]
			} else if hole[1] > swap[3] && swap[3] != swap[4] && swap[3] != swap[2] && swap[3] != swap[1] && swap[3] != swap[0] && swap[3] != 1 && hole[1] != paired {
				swap[3] = hole[1]
			} else if hole[1] > swap[2] && swap[2] != swap[4] && swap[2] != swap[3] && swap[2] != swap[1] && swap[2] != swap[0] && swap[2] != 1 && hole[1] != paired {
				swap[2] = hole[1]
			} else if hole[1] > swap[1] && swap[1] != swap[4] && swap[1] != swap[3] && swap[1] != swap[2] && swap[1] != swap[0] && swap[1] != 1 && hole[1] != paired {
				swap[1] = hole[1]
			} else if hole[1] > swap[0] && swap[0] != swap[4] && swap[0] != swap[3] && swap[0] != swap[2] && swap[0] != swap[1] && swap[0] != 1 && hole[1] != paired {
				swap[0] = hole[1]
			} else if hole[1] == 1 && swap[4] <= 13 && swap[4] != swap[3] && swap[4] != swap[2] && swap[4] != swap[1] && swap[4] != swap[0] && swap[4] != 1 {
				swap[4] = hole[1]
			} else if hole[1] == 1 && swap[3] <= 13 && swap[3] != swap[4] && swap[3] != swap[2] && swap[3] != swap[1] && swap[3] != swap[0] && swap[3] != 1 {
				swap[3] = hole[1]
			} else if hole[1] == 1 && swap[2] <= 13 && swap[2] != swap[4] && swap[2] != swap[3] && swap[2] != swap[1] && swap[2] != swap[0] && swap[2] != 1 {
				swap[2] = hole[1]
			} else if hole[1] == 1 && swap[1] <= 13 && swap[1] != swap[4] && swap[1] != swap[3] && swap[1] != swap[2] && swap[1] != swap[0] && swap[1] != 1 {
				swap[1] = hole[1]
			} else if hole[1] == 1 && swap[0] <= 13 && swap[0] != swap[4] && swap[0] != swap[3] && swap[0] != swap[2] && swap[0] != swap[1] && swap[0] != 1 {
				swap[0] = hole[1]
			}
		}

	}

	/// If straight or straight flush
	if r == 6 || r == 2 {
		if hole[0] == swap[4]+1 && hole[1] == swap[4]+2 {
			swap = []int{swap[2], swap[3], swap[4], hole[0], hole[1]}
		} else if hole[0] == swap[4]+1 {
			swap = []int{swap[1], swap[2], swap[3], swap[4], hole[0]}
		} else if hole[1] == swap[4]+1 {
			swap = []int{swap[1], swap[2], swap[3], swap[4], hole[1]}
		} else if swap[0] == 8 && swap[1] == 9 && swap[2] == 10 && swap[3] == 11 && swap[4] == 12 && hole[0] == swap[4]+1 {
			swap = []int{swap[1], swap[2], swap[3], swap[4], hole[0]}
		} else if swap[0] == 8 && swap[1] == 9 && swap[2] == 10 && swap[3] == 11 && swap[4] == 12 && hole[1] == swap[4]+1 {
			swap = []int{swap[1], swap[2], swap[3], swap[4], hole[1]}
		} else if swap[0] == 8 && swap[1] == 9 && swap[2] == 10 && swap[3] == 11 && swap[4] == 12 && hole[0] == 1 && hole[1] == swap[4]+2 {
			swap = []int{swap[2], swap[3], swap[4], hole[1], hole[0]}
		}

	}

	/// If full house
	if r == 4 && hole[0] == hole[1] {
		if hole[0] > swap[4] && hole[1] > swap[3] && swap[4] != swap[2] {
			swap = []int{swap[0], swap[1], swap[2], hole[1], hole[0]}
		} else if hole[0] > swap[0] && hole[1] > swap[1] && swap[0] != swap[2] {
			swap = []int{hole[0], hole[1], swap[2], swap[3], swap[4]}
		}
	}

	return swap[:]

}

func compareMine() int { /// Search thorugh all hand combinations to find best

	/// Two Hole cards
	e1Hand := []int{cc1[0], cc2[0], cc3[0], pc1[0], pc2[0]}
	e1Suits := []int{cc1[1], cc2[1], cc3[1], pc1[1], pc2[1]}
	fRank := makeHand(e1Hand, e1Suits)
	fHighCardArr = e1Hand

	e2Hand := []int{cc1[0], cc2[0], pc1[0], cc4[0], pc2[0]}
	e2Suits := []int{cc1[1], cc2[1], pc1[1], cc4[1], pc2[1]}
	nRank := makeHand(e2Hand, e2Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e2Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e2Hand, e2Suits)

	}

	e3Hand := []int{cc1[0], pc1[0], cc3[0], cc4[0], pc2[0]}
	e3Suits := []int{cc1[1], pc1[1], cc3[1], cc4[1], pc2[1]}
	nRank = makeHand(e3Hand, e3Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e3Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e3Hand, e3Suits)

	}

	e4Hand := []int{pc1[0], cc2[0], cc3[0], cc4[0], pc2[0]}
	e4Suits := []int{pc1[1], cc2[1], cc3[1], cc4[1], pc2[1]}
	nRank = makeHand(e4Hand, e4Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e4Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e4Hand, e4Suits)

	}

	e5Hand := []int{cc1[0], cc2[0], pc1[0], pc2[0], cc5[0]}
	e5Suits := []int{cc1[1], cc2[1], pc1[1], pc2[1], cc5[1]}
	nRank = makeHand(e5Hand, e5Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e5Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e5Hand, e5Suits)

	}

	e6Hand := []int{cc1[0], pc1[0], cc3[0], pc2[0], cc5[0]}
	e6Suits := []int{cc1[1], pc1[1], cc3[1], pc2[1], cc5[1]}
	nRank = makeHand(e6Hand, e6Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e6Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e6Hand, e6Suits)

	}

	e7Hand := []int{pc1[0], cc2[0], cc3[0], pc2[0], cc5[0]}
	e7Suits := []int{pc1[1], cc2[1], cc3[1], pc2[1], cc5[1]}
	nRank = makeHand(e7Hand, e7Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e7Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e7Hand, e7Suits)

	}

	e8Hand := []int{cc1[0], pc1[0], pc2[0], cc4[0], cc5[0]}
	e8Suits := []int{cc1[1], pc1[1], pc2[1], cc4[1], cc5[1]}
	nRank = makeHand(e8Hand, e8Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e8Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e8Hand, e8Suits)

	}

	e9Hand := []int{pc1[0], cc2[0], pc2[0], cc4[0], cc5[0]}
	e9Suits := []int{pc1[1], cc2[1], pc2[1], cc4[1], cc5[1]}
	nRank = makeHand(e9Hand, e9Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e9Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e9Hand, e9Suits)

	}

	e10Hand := []int{pc1[0], pc2[0], cc3[0], cc4[0], cc5[0]}
	e10Suits := []int{pc1[1], pc2[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(e10Hand, e10Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e10Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e10Hand, e10Suits)

	}

	/// First Hole card
	e11Hand := []int{cc1[0], cc2[0], cc3[0], cc4[0], pc1[0]}
	e11Suits := []int{cc1[1], cc2[1], cc3[1], cc4[1], pc1[1]}
	nRank = makeHand(e11Hand, e11Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e11Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e11Hand, e11Suits)

	}

	e12Hand := []int{cc1[0], cc2[0], cc3[0], pc1[0], cc5[0]}
	e12Suits := []int{cc1[1], cc2[1], cc3[1], pc1[1], cc5[1]}
	nRank = makeHand(e12Hand, e12Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e12Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e12Hand, e12Suits)

	}

	e13Hand := []int{cc1[0], cc2[0], pc1[0], cc4[0], cc5[0]}
	e13Suits := []int{cc1[1], cc2[1], pc1[1], cc4[1], cc5[1]}
	nRank = makeHand(e13Hand, e13Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e13Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e13Hand, e13Suits)

	}

	e14Hand := []int{cc1[0], pc1[0], cc3[0], cc4[0], cc5[0]}
	e14Suits := []int{cc1[1], pc1[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(e14Hand, e14Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e14Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e14Hand, e14Suits)

	}

	e15Hand := []int{pc1[0], cc2[0], cc3[0], cc4[0], cc5[0]}
	e15Suits := []int{pc1[1], cc2[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(e15Hand, e15Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e15Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e15Hand, e15Suits)

	}

	/// Second Hole card
	e16Hand := []int{cc1[0], cc2[0], cc3[0], cc4[0], pc2[0]}
	e16Suits := []int{cc1[1], cc2[1], cc3[1], cc4[1], pc2[1]}
	nRank = makeHand(e16Hand, e16Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e16Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e16Hand, e16Suits)

	}

	e17Hand := []int{cc1[0], cc2[0], cc3[0], pc2[0], cc5[0]}
	e17Suits := []int{cc1[1], cc2[1], cc3[1], pc2[1], cc5[1]}
	nRank = makeHand(e17Hand, e17Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e17Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e17Hand, e17Suits)

	}

	e18Hand := []int{cc1[0], cc2[0], pc2[0], cc4[0], cc5[0]}
	e18Suits := []int{cc1[1], cc2[1], pc2[1], cc4[1], cc5[1]}
	nRank = makeHand(e18Hand, e18Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e18Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e18Hand, e18Suits)

	}

	e19Hand := []int{cc1[0], pc2[0], cc3[0], cc4[0], cc5[0]}
	e19Suits := []int{cc1[1], pc2[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(e19Hand, e19Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e19Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e19Hand, e19Suits)

	}

	e20Hand := []int{pc2[0], cc2[0], cc3[0], cc4[0], cc5[0]}
	e20Suits := []int{pc2[1], cc2[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(e20Hand, e20Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e20Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e20Hand, e20Suits)

	}

	/// All community
	e0Hand := []int{cc1[0], cc2[0], cc3[0], cc4[0], cc5[0]}
	e0Suits := []int{cc1[1], cc2[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(e0Hand, e0Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = e0Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = searchAllRanks(fRank, e0Hand, e0Suits)

	}

	return fRank
}
