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
)

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

func getFlop() { /// Splitting card value and suit for making individual hand, highcard value and rank
	for {
		fmt.Println("Input community cards: ")

		_, err := fmt.Scanln(&communityCard[0], &communityCard[1], &communityCard[2], &communityCard[3], &communityCard[4])
		if err != nil {
			if communityCard[4] == 0 { /// Did not input 5 cards
				fmt.Println("Incorrect Format, Try Again")
			}

			if communityCard[4] > 0 { /// Input had extra card, removed
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if communityCard[0] > 0 && communityCard[0] < 53 && communityCard[1] > 0 && communityCard[1] < 53 && communityCard[2] > 0 &&
			communityCard[2] < 53 && communityCard[3] > 0 && communityCard[3] < 53 && communityCard[4] > 0 && communityCard[4] < 53 { /// Ensure correct input format
			break
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
				fmt.Println("Incorrect Format, Try Again")
			}

			if p1HandRaw[1] > 0 { /// Input had extra card, removed
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p1HandRaw[0] > 0 && p1HandRaw[0] < 53 && p1HandRaw[1] > 0 && p1HandRaw[1] < 53 { /// Ensure correct input format
			break
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

	for i := 0; i < 4; i++ {
		if p1HighCardArr[i] == p1HighCardArr[i+1] {
			if p1HighCardArr[i] > p1HighPair {
				p1HighPair = p1HighCardArr[i]
			}
		}
	}

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

	for i := 0; i < 4; i++ {
		if p2HighCardArr[i] == p2HighCardArr[i+1] {
			if p2HighCardArr[i] > p2HighPair {
				p2HighPair = p2HighCardArr[i]
			}
		}
	}

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

	for i := 0; i < 4; i++ {
		if p3HighCardArr[i] == p3HighCardArr[i+1] {
			if p3HighCardArr[i] > p3HighPair {
				p3HighPair = p3HighCardArr[i]
			}
		}
	}

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

	for i := 0; i < 4; i++ {
		if p4HighCardArr[i] == p4HighCardArr[i+1] {
			if p4HighCardArr[i] > p4HighPair {
				p4HighPair = p4HighCardArr[i]
			}
		}
	}

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

	for i := 0; i < 4; i++ {
		if p5HighCardArr[i] == p5HighCardArr[i+1] {
			if p5HighCardArr[i] > p5HighPair {
				p5HighPair = p5HighCardArr[i]
			}
		}
	}

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

	for i := 0; i < 4; i++ {
		if p6HighCardArr[i] == p6HighCardArr[i+1] {
			if p6HighCardArr[i] > p6HighPair {
				p6HighPair = p6HighCardArr[i]
			}
		}
	}

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

func makeCommHighCard(r int, s []int) []int {

	var swap = s
	hole := []int{pc1[0], pc2[0]}
	sort.Ints(hole)
	// red := color.New(color.FgHiRed).PrintlnFunc()
	// red("Comm: %s", swap)
	// cyan := color.New(color.FgHiCyan).PrintlnFunc()
	// cyan("Hole: %s", hole)
	// green := color.New(color.FgHiGreen).PrintlnFunc()
	// green("Rank: %s", r)
	/// If high card
	if r == 10 {
		if hole[0] > swap[0] {
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
		} else if hole[1] > swap[1] && swap[1] != 1 {
			swap[1] = hole[1]
		} else if hole[1] > swap[2] && swap[2] != 1 {
			swap[2] = hole[1]
		} else if hole[1] > swap[3] && swap[3] != 1 {
			swap[3] = hole[1]
		} else if hole[1] > swap[4] && swap[4] != 1 {
			swap[4] = hole[1]
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
		if hole[0] > swap[4] && swap[4] != swap[3] && swap[4] != swap[2] && swap[4] != swap[1] && swap[4] != swap[0] {
			swap[4] = hole[0]
		} else if hole[0] > swap[3] && swap[3] != swap[4] && swap[3] != swap[2] && swap[3] != swap[1] && swap[3] != swap[0] {
			swap[3] = hole[0]
		} else if hole[0] > swap[2] && swap[2] != swap[4] && swap[2] != swap[3] && swap[2] != swap[1] && swap[2] != swap[0] {
			swap[2] = hole[0]
		} else if hole[0] > swap[1] && swap[1] != swap[4] && swap[1] != swap[3] && swap[1] != swap[2] && swap[1] != swap[0] {
			swap[1] = hole[0]
		} else if hole[0] > swap[0] && swap[0] != swap[4] && swap[0] != swap[3] && swap[0] != swap[2] && swap[0] != swap[1] {
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

		if hole[1] > swap[4] && swap[4] != swap[3] && swap[4] != swap[2] && swap[4] != swap[1] && swap[4] != swap[0] && swap[4] != 1 {
			swap[4] = hole[1]
		} else if hole[1] > swap[3] && swap[3] != swap[4] && swap[3] != swap[2] && swap[3] != swap[1] && swap[3] != swap[0] && swap[3] != 1 {
			swap[3] = hole[1]
		} else if hole[1] > swap[2] && swap[2] != swap[4] && swap[2] != swap[3] && swap[2] != swap[1] && swap[2] != swap[0] && swap[2] != 1 {
			swap[2] = hole[1]
		} else if hole[1] > swap[1] && swap[1] != swap[4] && swap[1] != swap[3] && swap[1] != swap[2] && swap[1] != swap[0] && swap[1] != 1 {
			swap[1] = hole[1]
		} else if hole[1] > swap[0] && swap[0] != swap[4] && swap[0] != swap[3] && swap[0] != swap[2] && swap[0] != swap[1] && swap[0] != 1 {
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
	/// If straight, flush or straight flush
	if r == 6 || r == 5 || r == 2 {
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

	if r == 4 && hole[0] == hole[1] {
		if hole[0] > swap[4] && hole[1] > swap[3] && swap[4] != swap[2] {
			swap = []int{swap[0], swap[1], swap[2], hole[1], hole[0]}
		} else if hole[0] > swap[0] && hole[1] > swap[1] && swap[0] != swap[2] {
			swap = []int{hole[0], hole[1], swap[2], swap[3], swap[4]}
		}
	}

	return swap[:]

}

func compareMine() int {

	/// Two Hole cards
	p1Hand := []int{cc1[0], cc2[0], cc3[0], pc1[0], pc2[0]}
	p1Suits := []int{cc1[1], cc2[1], cc3[1], pc1[1], pc2[1]}
	fRank := makeHand(p1Hand, p1Suits)
	fHighCardArr = p1Hand

	p2Hand := []int{cc1[0], cc2[0], pc1[0], cc4[0], pc2[0]}
	p2Suits := []int{cc1[1], cc2[1], pc1[1], cc4[1], pc2[1]}
	nRank := makeHand(p2Hand, p2Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p2Hand
	}

	p3Hand := []int{cc1[0], pc1[0], cc3[0], cc4[0], pc2[0]}
	p3Suits := []int{cc1[1], pc1[1], cc3[1], cc4[1], pc2[1]}
	nRank = makeHand(p3Hand, p3Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p3Hand
	}

	p4Hand := []int{pc1[0], cc2[0], cc3[0], cc4[0], pc2[0]}
	p4Suits := []int{pc1[1], cc2[1], cc3[1], cc4[1], pc2[1]}
	nRank = makeHand(p4Hand, p4Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p4Hand
	}

	p5Hand := []int{cc1[0], cc2[0], pc1[0], pc2[0], cc5[0]}
	p5Suits := []int{cc1[1], cc2[1], pc1[1], pc2[1], cc5[1]}
	nRank = makeHand(p5Hand, p5Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p5Hand
	}

	p6Hand := []int{cc1[0], pc1[0], cc3[0], pc2[0], cc5[0]}
	p6Suits := []int{cc1[1], pc1[1], cc3[1], pc2[1], cc5[1]}
	nRank = makeHand(p6Hand, p6Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p6Hand
	}

	p7Hand := []int{pc1[0], cc2[0], cc3[0], pc2[0], cc5[0]}
	p7Suits := []int{pc1[1], cc2[1], cc3[1], pc2[1], cc5[1]}
	nRank = makeHand(p7Hand, p7Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p7Hand
	}

	p8Hand := []int{cc1[0], pc1[0], pc2[0], cc4[0], cc5[0]}
	p8Suits := []int{cc1[1], pc1[1], pc2[1], cc4[1], cc5[1]}
	nRank = makeHand(p8Hand, p8Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p8Hand
	}

	p9Hand := []int{pc1[0], cc2[0], pc2[0], cc4[0], cc5[0]}
	p9Suits := []int{pc1[1], cc2[1], pc2[1], cc4[1], cc5[1]}
	nRank = makeHand(p9Hand, p9Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p9Hand
	}

	p10Hand := []int{pc1[0], pc2[0], cc3[0], cc4[0], cc5[0]}
	p10Suits := []int{pc1[1], pc2[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(p10Hand, p10Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p10Hand
	}

	/// First Hole card
	p11Hand := []int{cc1[0], cc2[0], cc3[0], cc4[0], pc1[0]}
	p11Suits := []int{cc1[1], cc2[1], cc3[1], cc4[1], pc1[1]}
	nRank = makeHand(p11Hand, p11Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p11Hand
	}

	p12Hand := []int{cc1[0], cc2[0], cc3[0], pc1[0], cc5[0]}
	p12Suits := []int{cc1[1], cc2[1], cc3[1], pc1[1], cc5[1]}
	nRank = makeHand(p12Hand, p12Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p12Hand
	}

	p13Hand := []int{cc1[0], cc2[0], pc1[0], cc4[0], cc5[0]}
	p13Suits := []int{cc1[1], cc2[1], pc1[1], cc4[1], cc5[1]}
	nRank = makeHand(p13Hand, p13Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p13Hand
	}

	p14Hand := []int{cc1[0], pc1[0], cc3[0], cc4[0], cc5[0]}
	p14Suits := []int{cc1[1], pc1[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(p14Hand, p14Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p14Hand
	}

	p15Hand := []int{pc1[0], cc2[0], cc3[0], cc4[0], cc5[0]}
	p15Suits := []int{pc1[1], cc2[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(p15Hand, p15Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p15Hand
	}

	/// Second Hole card
	p16Hand := []int{cc1[0], cc2[0], cc3[0], cc4[0], pc2[0]}
	p16Suits := []int{cc1[1], cc2[1], cc3[1], cc4[1], pc2[1]}
	nRank = makeHand(p16Hand, p16Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p16Hand
	}

	p17Hand := []int{cc1[0], cc2[0], cc3[0], pc2[0], cc5[0]}
	p17Suits := []int{cc1[1], cc2[1], cc3[1], pc2[1], cc5[1]}
	nRank = makeHand(p17Hand, p17Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p17Hand
	}

	p18Hand := []int{cc1[0], cc2[0], pc2[0], cc4[0], cc5[0]}
	p18Suits := []int{cc1[1], cc2[1], pc2[1], cc4[1], cc5[1]}
	nRank = makeHand(p18Hand, p18Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p18Hand
	}

	p19Hand := []int{cc1[0], pc2[0], cc3[0], cc4[0], cc5[0]}
	p19Suits := []int{cc1[1], pc2[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(p19Hand, p19Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p19Hand
	}

	p20Hand := []int{pc2[0], cc2[0], cc3[0], cc4[0], cc5[0]}
	p20Suits := []int{pc2[1], cc2[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(p20Hand, p20Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p20Hand
	}

	/// All community
	p0Hand := []int{cc1[0], cc2[0], cc3[0], cc4[0], cc5[0]}
	p0Suits := []int{cc1[1], cc2[1], cc3[1], cc4[1], cc5[1]}
	nRank = makeHand(p0Hand, p0Suits)
	if nRank < fRank {
		fRank = nRank
		fHighCardArr = p0Hand
	} else if nRank == fRank {
		fRank = nRank
		fHighCardArr = makeCommHighCard(fRank, p0Hand)

	}

	return fRank
}
