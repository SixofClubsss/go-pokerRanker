/*
dReam Tables Poker Hand Ranker - Five Card

Copyright (C) 2022  dReam Tables

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

Always play responsibly.

https://dreamtables.net
*/

package main

import (
	"fmt"
	"sort"
)

var totalHands int
var p1HandRaw [5]int
var p2HandRaw [5]int
var p3HandRaw [5]int
var p4HandRaw [5]int
var p5HandRaw [5]int
var p6HandRaw [5]int
var arrSplit [2]int

var pc1 [2]int
var pc2 [2]int
var pc3 [2]int
var pc4 [2]int
var pc5 [2]int
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
var p1HighCardArr [5]int
var p2HighCardArr [5]int
var p3HighCardArr [5]int
var p4HighCardArr [5]int
var p5HighCardArr [5]int
var p6HighCardArr [5]int

func clearHands() { /// Clears hand arrays before new input
	p1HandRaw = [5]int{0, 0, 0, 0, 0}
	p2HandRaw = [5]int{0, 0, 0, 0, 0}
	p3HandRaw = [5]int{0, 0, 0, 0, 0}
	p4HandRaw = [5]int{0, 0, 0, 0, 0}
	p5HandRaw = [5]int{0, 0, 0, 0, 0}
	p6HandRaw = [5]int{0, 0, 0, 0, 0}
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
		p1Rank = getHand1()
		p2Rank = getHand2()
	case 3:
		p1Rank = getHand1()
		p2Rank = getHand2()
		p3Rank = getHand3()
	case 4:
		p1Rank = getHand1()
		p2Rank = getHand2()
		p3Rank = getHand3()
		getHand4()
	case 5:
		p1Rank = getHand1()
		p2Rank = getHand2()
		p3Rank = getHand3()
		p4Rank = getHand4()
		p5Rank = getHand5()
	case 6:
		p1Rank = getHand1()
		p2Rank = getHand2()
		p3Rank = getHand3()
		p4Rank = getHand4()
		p5Rank = getHand5()
		p6Rank = getHand6()
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

func getHand1() int { /// Splitting card value and suit for making individual hand, highcard value and rank
	for {
		fmt.Println("Input the card values for Player 1: ")

		_, err := fmt.Scanln(&p1HandRaw[0], &p1HandRaw[1], &p1HandRaw[2], &p1HandRaw[3], &p1HandRaw[4])
		if err != nil {
			if p1HandRaw[4] == 0 { /// Did not input 5 cards
				fmt.Println("Incorrect Format, Try Again")
			}

			if p1HandRaw[4] > 0 { /// Input had extra card, removed
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p1HandRaw[0] > 0 && p1HandRaw[0] < 53 && p1HandRaw[1] > 0 && p1HandRaw[1] < 53 && p1HandRaw[2] > 0 &&
			p1HandRaw[2] < 53 && p1HandRaw[3] > 0 && p1HandRaw[3] < 53 && p1HandRaw[4] > 0 && p1HandRaw[4] < 53 { /// Ensure correct input format
			break
		}
	}

	suitSplit(p1HandRaw[0])
	pc1 = [2]int{arrSplit[0], arrSplit[1]}
	p1HighCardArr[0] = arrSplit[0]

	suitSplit(p1HandRaw[1])
	pc2 = [2]int{arrSplit[0], arrSplit[1]}
	p1HighCardArr[1] = arrSplit[0]

	suitSplit(p1HandRaw[2])
	pc3 = [2]int{arrSplit[0], arrSplit[1]}
	p1HighCardArr[2] = arrSplit[0]

	suitSplit(p1HandRaw[3])
	pc4 = [2]int{arrSplit[0], arrSplit[1]}
	p1HighCardArr[3] = arrSplit[0]

	suitSplit(p1HandRaw[4])
	pc5 = [2]int{arrSplit[0], arrSplit[1]}
	p1HighCardArr[4] = arrSplit[0]

	Rank := makeHand()

	p1HighPair = getHighPair(p1HighCardArr)

	return Rank
}

func getHand2() int {
	for {
		fmt.Println("Input the card values for Player 2: ")

		_, err := fmt.Scanln(&p2HandRaw[0], &p2HandRaw[1], &p2HandRaw[2], &p2HandRaw[3], &p2HandRaw[4])
		if err != nil {
			if p2HandRaw[4] == 0 {
				fmt.Println("Incorrect Format, Try Again")
			}

			if p2HandRaw[4] > 0 {
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p2HandRaw[0] > 0 && p2HandRaw[0] < 53 && p2HandRaw[1] > 0 && p2HandRaw[1] < 53 && p2HandRaw[2] > 0 &&
			p2HandRaw[2] < 53 && p2HandRaw[3] > 0 && p2HandRaw[3] < 53 && p2HandRaw[4] > 0 && p2HandRaw[4] < 53 {
			break
		}
	}

	suitSplit(p2HandRaw[0])
	pc1 = [2]int{arrSplit[0], arrSplit[1]}
	p2HighCardArr[0] = arrSplit[0]

	suitSplit(p2HandRaw[1])
	pc2 = [2]int{arrSplit[0], arrSplit[1]}
	p2HighCardArr[1] = arrSplit[0]

	suitSplit(p2HandRaw[2])
	pc3 = [2]int{arrSplit[0], arrSplit[1]}
	p2HighCardArr[2] = arrSplit[0]

	suitSplit(p2HandRaw[3])
	pc4 = [2]int{arrSplit[0], arrSplit[1]}
	p2HighCardArr[3] = arrSplit[0]

	suitSplit(p2HandRaw[4])
	pc5 = [2]int{arrSplit[0], arrSplit[1]}
	p2HighCardArr[4] = arrSplit[0]

	Rank := makeHand()

	p2HighPair = getHighPair(p2HighCardArr)

	return Rank
}

func getHand3() int {
	for {
		fmt.Println("Input the card values for Player 3: ")

		_, err := fmt.Scanln(&p3HandRaw[0], &p3HandRaw[1], &p3HandRaw[2], &p3HandRaw[3], &p3HandRaw[4])
		if err != nil {
			if p3HandRaw[4] == 0 {
				fmt.Println("Incorrect Format, Try Again")
			}

			if p3HandRaw[4] > 0 {
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p3HandRaw[0] > 0 && p3HandRaw[0] < 53 && p3HandRaw[1] > 0 && p3HandRaw[1] < 53 && p3HandRaw[2] > 0 &&
			p3HandRaw[2] < 53 && p3HandRaw[3] > 0 && p3HandRaw[3] < 53 && p3HandRaw[4] > 0 && p3HandRaw[4] < 53 {
			break
		}
	}

	suitSplit(p3HandRaw[0])
	pc1 = [2]int{arrSplit[0], arrSplit[1]}
	p3HighCardArr[0] = arrSplit[0]

	suitSplit(p3HandRaw[1])
	pc2 = [2]int{arrSplit[0], arrSplit[1]}
	p3HighCardArr[1] = arrSplit[0]

	suitSplit(p3HandRaw[2])
	pc3 = [2]int{arrSplit[0], arrSplit[1]}
	p3HighCardArr[2] = arrSplit[0]

	suitSplit(p3HandRaw[3])
	pc4 = [2]int{arrSplit[0], arrSplit[1]}
	p3HighCardArr[3] = arrSplit[0]

	suitSplit(p3HandRaw[4])
	pc5 = [2]int{arrSplit[0], arrSplit[1]}
	p3HighCardArr[4] = arrSplit[0]

	Rank := makeHand()

	p3HighPair = getHighPair(p3HighCardArr)

	return Rank
}

func getHand4() int {
	for {
		fmt.Println("Input the card values for Player 4: ")

		_, err := fmt.Scanln(&p4HandRaw[0], &p4HandRaw[1], &p4HandRaw[2], &p4HandRaw[3], &p4HandRaw[4])
		if err != nil {
			if p4HandRaw[4] == 0 {
				fmt.Println("Incorrect Format, Try Again")
			}

			if p4HandRaw[4] > 0 {
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p4HandRaw[0] > 0 && p4HandRaw[0] < 53 && p4HandRaw[1] > 0 && p4HandRaw[1] < 53 && p4HandRaw[2] > 0 &&
			p4HandRaw[2] < 53 && p4HandRaw[3] > 0 && p4HandRaw[3] < 53 && p4HandRaw[4] > 0 && p4HandRaw[4] < 53 {
			break
		}
	}

	suitSplit(p4HandRaw[0])
	pc1 = [2]int{arrSplit[0], arrSplit[1]}
	p4HighCardArr[0] = arrSplit[0]

	suitSplit(p4HandRaw[1])
	pc2 = [2]int{arrSplit[0], arrSplit[1]}
	p4HighCardArr[1] = arrSplit[0]

	suitSplit(p4HandRaw[2])
	pc3 = [2]int{arrSplit[0], arrSplit[1]}
	p4HighCardArr[2] = arrSplit[0]

	suitSplit(p4HandRaw[3])
	pc4 = [2]int{arrSplit[0], arrSplit[1]}
	p4HighCardArr[3] = arrSplit[0]

	suitSplit(p4HandRaw[4])
	pc5 = [2]int{arrSplit[0], arrSplit[1]}
	p4HighCardArr[4] = arrSplit[0]

	Rank := makeHand()

	p4HighPair = getHighPair(p4HighCardArr)

	return Rank
}

func getHand5() int {
	for {
		fmt.Println("Input the card values for Player 5: ")

		_, err := fmt.Scanln(&p5HandRaw[0], &p5HandRaw[1], &p5HandRaw[2], &p5HandRaw[3], &p5HandRaw[4])
		if err != nil {
			if p5HandRaw[4] == 0 {
				fmt.Println("Incorrect Format, Try Again")
			}

			if p5HandRaw[4] > 0 {
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p5HandRaw[0] > 0 && p5HandRaw[0] < 53 && p5HandRaw[1] > 0 && p5HandRaw[1] < 53 && p5HandRaw[2] > 0 &&
			p5HandRaw[2] < 53 && p5HandRaw[3] > 0 && p5HandRaw[3] < 53 && p5HandRaw[4] > 0 && p5HandRaw[4] < 53 {
			break
		}
	}

	suitSplit(p5HandRaw[0])
	pc1 = [2]int{arrSplit[0], arrSplit[1]}
	p5HighCardArr[0] = arrSplit[0]

	suitSplit(p5HandRaw[1])
	pc2 = [2]int{arrSplit[0], arrSplit[1]}
	p5HighCardArr[1] = arrSplit[0]

	suitSplit(p5HandRaw[2])
	pc3 = [2]int{arrSplit[0], arrSplit[1]}
	p5HighCardArr[2] = arrSplit[0]

	suitSplit(p5HandRaw[3])
	pc4 = [2]int{arrSplit[0], arrSplit[1]}
	p5HighCardArr[3] = arrSplit[0]

	suitSplit(p5HandRaw[4])
	pc5 = [2]int{arrSplit[0], arrSplit[1]}
	p5HighCardArr[4] = arrSplit[0]

	Rank := makeHand()

	p5HighPair = getHighPair(p5HighCardArr)

	return Rank
}

func getHand6() int {
	for {
		fmt.Println("Input the card values for Player 6: ")

		_, err := fmt.Scanln(&p6HandRaw[0], &p6HandRaw[1], &p6HandRaw[2], &p6HandRaw[3], &p6HandRaw[4])
		if err != nil {
			if p6HandRaw[4] == 0 {
				fmt.Println("Incorrect Format, Try Again")
			}

			if p6HandRaw[4] > 0 {
				fmt.Println("Extra input was removed")
				var discard string
				fmt.Scanln(&discard)
			}

		}

		if p6HandRaw[0] > 0 && p6HandRaw[0] < 53 && p6HandRaw[1] > 0 && p6HandRaw[1] < 53 && p6HandRaw[2] > 0 &&
			p6HandRaw[2] < 53 && p6HandRaw[3] > 0 && p6HandRaw[3] < 53 && p6HandRaw[4] > 0 && p6HandRaw[4] < 53 {
			break
		}
	}

	suitSplit(p6HandRaw[0])
	pc1 = [2]int{arrSplit[0], arrSplit[1]}
	p6HighCardArr[0] = arrSplit[0]

	suitSplit(p6HandRaw[1])
	pc2 = [2]int{arrSplit[0], arrSplit[1]}
	p6HighCardArr[1] = arrSplit[0]

	suitSplit(p6HandRaw[2])
	pc3 = [2]int{arrSplit[0], arrSplit[1]}
	p6HighCardArr[2] = arrSplit[0]

	suitSplit(p6HandRaw[3])
	pc4 = [2]int{arrSplit[0], arrSplit[1]}
	p6HighCardArr[3] = arrSplit[0]

	suitSplit(p6HandRaw[4])
	pc5 = [2]int{arrSplit[0], arrSplit[1]}
	p6HighCardArr[4] = arrSplit[0]

	Rank := makeHand()

	p6HighPair = getHighPair(p6HighCardArr)

	return Rank
}

func makeHand() int { /// Determines hand rank after suit slipt

	pHand := []int{pc1[0], pc2[0], pc3[0], pc4[0], pc5[0]}
	pSuits := []int{pc1[1], pc2[1], pc3[1], pc4[1], pc5[1]}

	sort.Ints(pHand)

	/// Royal flush
	if pHand[0] == 10 && pHand[1] == 11 && pHand[2] == 12 && pHand[3] == 13 && pHand[4] == 14 &&
		pSuits[0] == pSuits[1] && pSuits[0] == pSuits[2] && pSuits[0] == pSuits[3] && pSuits[0] == pSuits[4] {

		return 1

	}

	/// Straight flush
	if (pHand[0]+1 == pHand[1] && pHand[1]+1 == pHand[2] && pHand[2]+1 == pHand[3] && pHand[3]+1 == pHand[4] && pHand[0]+4 == pHand[4] &&
		pSuits[0] == pSuits[1] && pSuits[0] == pSuits[2] && pSuits[0] == pSuits[3] && pSuits[0] == pSuits[4]) ||
		(pHand[0] == 2 && pHand[1] == 3 && pHand[2] == 4 && pHand[3] == 5 && pHand[4] == 14 &&
			pSuits[0] == pSuits[1] && pSuits[0] == pSuits[2] && pSuits[0] == pSuits[3] && pSuits[0] == pSuits[4]) {

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
		pHand[0] == 2 && pHand[1] == 3 && pHand[2] == 4 && pHand[3] == 5 && pHand[4] == 14 {
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
