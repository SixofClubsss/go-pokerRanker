/*
dReam Tables Poker Hand Ranker - Holdem

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

	"github.com/fatih/color"
)

func compareAll() { /// Main compare to determine winner

	winningRank := []int{p1Rank, p2Rank, p3Rank, p4Rank, p5Rank, p6Rank}
	sort.Ints(winningRank)

	comText := color.New(color.Bold, color.FgGreen).PrintlnFunc()

	comText("Community:", cardEquiv(communityCard[0]), cardEquiv(communityCard[1]), cardEquiv(communityCard[2]), cardEquiv(communityCard[3]), cardEquiv(communityCard[4]), communityCard)
	if totalHands >= 2 { /// Print results
		fmt.Println("Player 1 has:", cardEquiv(p1HandRaw[0]), cardEquiv(p1HandRaw[1]), p1HandRaw, handToText(p1Rank), p1HighCardArr, p1HighPair, p1LowPair)
		fmt.Println("Player 2 has:", cardEquiv(p2HandRaw[0]), cardEquiv(p2HandRaw[1]), p2HandRaw, handToText(p2Rank), p2HighCardArr, p2HighPair, p2LowPair)
	}
	if totalHands >= 3 {
		fmt.Println("Player 3 has:", cardEquiv(p3HandRaw[0]), cardEquiv(p3HandRaw[1]), p3HandRaw, handToText(p3Rank), p3HighCardArr)
	}
	if totalHands >= 4 {
		fmt.Println("Player 4 has:", cardEquiv(p4HandRaw[0]), cardEquiv(p4HandRaw[1]), p4HandRaw, handToText(p4Rank), p4HighCardArr)
	}
	if totalHands >= 5 {
		fmt.Println("Player 5 has:", cardEquiv(p5HandRaw[0]), cardEquiv(p5HandRaw[1]), p5HandRaw, handToText(p5Rank), p5HighCardArr)
	}
	if totalHands >= 6 {
		fmt.Println("Player 6 has:", cardEquiv(p6HandRaw[0]), cardEquiv(p6HandRaw[1]), p6HandRaw, handToText(p6Rank), p6HighCardArr)
	}
	fmt.Println()

	if p1Rank < p2Rank && p1Rank < p3Rank && p1Rank < p4Rank && p1Rank < p5Rank && p1Rank < p6Rank { /// Outright win, player has highest rank
		fmt.Println("Player 1 Wins Outright")
	} else if p2Rank < p1Rank && p2Rank < p3Rank && p2Rank < p4Rank && p2Rank < p5Rank && p2Rank < p6Rank {
		fmt.Println("Player 2 Wins Outright")
	} else if p3Rank < p1Rank && p3Rank < p2Rank && p3Rank < p4Rank && p3Rank < p5Rank && p3Rank < p6Rank {
		fmt.Println("Player 3 Wins Outright")
	} else if p4Rank < p1Rank && p4Rank < p2Rank && p4Rank < p3Rank && p4Rank < p5Rank && p4Rank < p6Rank {
		fmt.Println("Player 4 Wins Outright")
	} else if p5Rank < p1Rank && p5Rank < p2Rank && p5Rank < p3Rank && p5Rank < p4Rank && p5Rank < p6Rank {
		fmt.Println("Player 5 Wins Outright")
	} else if p6Rank < p1Rank && p6Rank < p2Rank && p6Rank < p3Rank && p6Rank < p4Rank && p6Rank < p5Rank {
		fmt.Println("Player 6 Wins Outright")
	} else {

		highestPair := []int{p1HighPair, p2HighPair, p3HighPair, p4HighPair, p5HighPair, p6HighPair}
		sort.Ints(highestPair)

		if p1Rank != winningRank[0] || (winningRank[0] == 9 && p1HighPair != highestPair[5]) { /// If player hand is not the higest rank or if doesn't have high pair stip cards
			less1()
		}

		if p2Rank != winningRank[0] || (winningRank[0] == 9 && p2HighPair != highestPair[5]) {
			less2()
		}

		if p3Rank != winningRank[0] || (winningRank[0] == 9 && p3HighPair != highestPair[5]) {
			less3()
		}

		if (p4Rank != winningRank[0]) || (winningRank[0] == 9 && p4HighPair != highestPair[5]) {
			less4()
		}

		if p5Rank != winningRank[0] || (winningRank[0] == 9 && p5HighPair != highestPair[5]) {
			less5()
		}

		if p6Rank != winningRank[0] || (winningRank[0] == 9 && p6HighPair != highestPair[5]) {
			less6()
		}

		if winningRank[0] == 10 { /// Compares and strips loosing hands in high card situations
			compare1_2()
			compare2_1()
			if p1HighCardArr[4] > p2HighCardArr[4] {
				compare3_1()
				compare1_3()
			} else {
				compare3_2()
				compare2_3()
			}

			if p1HighCardArr[4] > p3HighCardArr[4] {
				compare1_4()
				compare4_1()
			} else if p2HighCardArr[4] > p3HighCardArr[4] {
				compare2_4()
				compare4_2()
			} else {
				compare3_4()
				compare4_3()
			}

			if p1HighCardArr[4] > p4HighCardArr[4] {
				compare1_5()
				compare5_1()
			} else if p2HighCardArr[4] > p4HighCardArr[4] {
				compare2_5()
				compare5_2()
			} else if p3HighCardArr[4] > p4HighCardArr[4] {
				compare3_5()
				compare5_3()
			} else {
				compare4_5()
				compare5_4()
			}

			if p1HighCardArr[4] > p5HighCardArr[4] {
				compare1_6()
				compare6_1()
			} else if p2HighCardArr[4] > p5HighCardArr[4] {
				compare2_6()
				compare6_2()
			} else if p3HighCardArr[4] > p5HighCardArr[4] {
				compare3_6()
				compare6_3()
			} else if p4HighCardArr[4] > p5HighCardArr[4] {
				compare4_6()
				compare6_4()
			} else {
				compare5_6()
				compare6_5()
			}
		}
		/// No outright win, highest pairing first used to compare two hands of same rank
		if p1HighPair > p2HighPair && p1HighPair > p3HighPair && p1HighPair > p4HighPair && p1HighPair > p5HighPair && p1HighPair > p6HighPair {
			if p1Rank == winningRank[0] {
				fmt.Println("Player 1 Wins (HighPair Compare)")
			}
		} else if p2HighPair > p1HighPair && p2HighPair > p3HighPair && p2HighPair > p4HighPair && p2HighPair > p5HighPair && p2HighPair > p6HighPair {
			if p2Rank == winningRank[0] {
				fmt.Println("Player 2 Wins (HighPair Compare)")
			}
		} else if p3HighPair > p1HighPair && p3HighPair > p2HighPair && p3HighPair > p4HighPair && p3HighPair > p5HighPair && p3HighPair > p6HighPair {
			if p3Rank == winningRank[0] {
				fmt.Println("Player 3 Wins (HighPair Compare)")
			}
		} else if p4HighPair > p1HighPair && p4HighPair > p2HighPair && p4HighPair > p3HighPair && p4HighPair > p5HighPair && p4HighPair > p6HighPair {
			if p4Rank == winningRank[0] {
				fmt.Println("Player 4 Wins (HighPair Compare)")
			}
		} else if p5HighPair > p1HighPair && p5HighPair > p2HighPair && p5HighPair > p3HighPair && p5HighPair > p4HighPair && p5HighPair > p6HighPair {
			if p5Rank == winningRank[0] {
				fmt.Println("Player 5 Wins (HighPair Compare)")
			}
		} else if p6HighPair > p1HighPair && p6HighPair > p2HighPair && p6HighPair > p3HighPair && p6HighPair > p4HighPair && p6HighPair > p5HighPair {
			if p6Rank == winningRank[0] {
				fmt.Println("Player 6 Wins (HighPair Compare)")
			}
		} else if winningRank[0] == 8 && p1LowPair > p2LowPair && p1LowPair > p3LowPair && p1LowPair > p4LowPair && p1LowPair > p5LowPair && p1LowPair > p6LowPair {
			if p1Rank == winningRank[0] {
				fmt.Println("Player 1 Wins (LowPair Compare)")
			}
		} else if winningRank[0] == 8 && p2LowPair > p1LowPair && p2LowPair > p3LowPair && p2LowPair > p4LowPair && p2LowPair > p5LowPair && p2LowPair > p6LowPair {
			if p2Rank == winningRank[0] {
				fmt.Println("Player 2 Wins (LowPair Compare)")
			}
		} else if winningRank[0] == 8 && p3LowPair > p1LowPair && p3LowPair > p2LowPair && p3LowPair > p4LowPair && p3LowPair > p5LowPair && p3LowPair > p6LowPair {
			if p3Rank == winningRank[0] {
				fmt.Println("Player 3 Wins (LowPair Compare)")
			}
		} else if winningRank[0] == 8 && p4LowPair > p1LowPair && p4LowPair > p2LowPair && p4LowPair > p3LowPair && p4LowPair > p5LowPair && p4LowPair > p6LowPair {
			if p4Rank == winningRank[0] {
				fmt.Println("Player 4 Wins (LowPair Compare)")
			}
		} else if winningRank[0] == 8 && p5LowPair > p1LowPair && p5LowPair > p2LowPair && p5LowPair > p3LowPair && p5LowPair > p4LowPair && p5LowPair > p6LowPair {
			if p5Rank == winningRank[0] {
				fmt.Println("Player 5 Wins (LowPair Compare)")
			}
		} else if winningRank[0] == 8 && p6LowPair > p1LowPair && p6LowPair > p2LowPair && p6LowPair > p3LowPair && p6LowPair > p4LowPair && p6LowPair > p5LowPair {
			if p6Rank == winningRank[0] {
				fmt.Println("Player 6 Wins (LowPair Compare)")
			}

			/// No outright or HighPair win so we comapre all left over hands to determine HighCard winner
		} else if (p1HighCardArr[4] > p2HighCardArr[4] && p1HighCardArr[4] > p3HighCardArr[4] && p1HighCardArr[4] > p4HighCardArr[4] && p1HighCardArr[4] > p5HighCardArr[4] && p1HighCardArr[4] > p6HighCardArr[4]) ||

			(p1HighCardArr[4] >= p2HighCardArr[4] && p1HighCardArr[4] >= p3HighCardArr[4] && p1HighCardArr[4] >= p4HighCardArr[4] && p1HighCardArr[4] >= p5HighCardArr[4] && p1HighCardArr[4] >= p6HighCardArr[4] &&
				p1HighCardArr[3] > p2HighCardArr[3] && p1HighCardArr[3] > p3HighCardArr[3] && p1HighCardArr[3] > p4HighCardArr[3] && p1HighCardArr[3] > p5HighCardArr[3] && p1HighCardArr[3] > p6HighCardArr[3]) ||

			(p1HighCardArr[4] >= p2HighCardArr[4] && p1HighCardArr[4] >= p3HighCardArr[4] && p1HighCardArr[4] >= p4HighCardArr[4] && p1HighCardArr[4] >= p5HighCardArr[4] && p1HighCardArr[4] >= p6HighCardArr[4] &&
				p1HighCardArr[3] >= p2HighCardArr[3] && p1HighCardArr[3] >= p3HighCardArr[3] && p1HighCardArr[3] >= p4HighCardArr[3] && p1HighCardArr[3] >= p5HighCardArr[3] && p1HighCardArr[3] >= p6HighCardArr[3] &&
				p1HighCardArr[2] > p2HighCardArr[2] && p1HighCardArr[2] > p3HighCardArr[2] && p1HighCardArr[2] > p4HighCardArr[2] && p1HighCardArr[2] > p5HighCardArr[2] && p1HighCardArr[2] > p6HighCardArr[2]) ||

			(p1HighCardArr[4] >= p2HighCardArr[4] && p1HighCardArr[4] >= p3HighCardArr[4] && p1HighCardArr[4] >= p4HighCardArr[4] && p1HighCardArr[4] >= p5HighCardArr[4] && p1HighCardArr[4] >= p6HighCardArr[4] &&
				p1HighCardArr[3] >= p2HighCardArr[3] && p1HighCardArr[3] >= p3HighCardArr[3] && p1HighCardArr[3] >= p4HighCardArr[3] && p1HighCardArr[3] >= p5HighCardArr[3] && p1HighCardArr[3] >= p6HighCardArr[3] &&
				p1HighCardArr[2] >= p2HighCardArr[2] && p1HighCardArr[2] >= p3HighCardArr[2] && p1HighCardArr[2] >= p4HighCardArr[2] && p1HighCardArr[2] >= p5HighCardArr[2] && p1HighCardArr[2] >= p6HighCardArr[2] &&
				p1HighCardArr[1] > p2HighCardArr[1] && p1HighCardArr[1] > p3HighCardArr[1] && p1HighCardArr[1] > p4HighCardArr[1] && p1HighCardArr[1] > p5HighCardArr[1] && p1HighCardArr[1] > p6HighCardArr[1]) ||

			(p1HighCardArr[4] >= p2HighCardArr[4] && p1HighCardArr[4] >= p3HighCardArr[4] && p1HighCardArr[4] >= p4HighCardArr[4] && p1HighCardArr[4] >= p5HighCardArr[4] && p1HighCardArr[4] >= p6HighCardArr[4] &&
				p1HighCardArr[3] >= p2HighCardArr[3] && p1HighCardArr[3] >= p3HighCardArr[3] && p1HighCardArr[3] >= p4HighCardArr[3] && p1HighCardArr[3] >= p5HighCardArr[3] && p1HighCardArr[3] >= p6HighCardArr[3] &&
				p1HighCardArr[2] >= p2HighCardArr[2] && p1HighCardArr[2] >= p3HighCardArr[2] && p1HighCardArr[2] >= p4HighCardArr[2] && p1HighCardArr[2] >= p5HighCardArr[2] && p1HighCardArr[2] >= p6HighCardArr[2] &&
				p1HighCardArr[1] >= p2HighCardArr[1] && p1HighCardArr[1] >= p3HighCardArr[1] && p1HighCardArr[1] >= p4HighCardArr[1] && p1HighCardArr[1] >= p5HighCardArr[1] && p1HighCardArr[1] >= p6HighCardArr[1] &&
				p1HighCardArr[0] > p2HighCardArr[0] && p1HighCardArr[0] > p3HighCardArr[0] && p1HighCardArr[0] > p4HighCardArr[0] && p1HighCardArr[0] > p5HighCardArr[0] && p1HighCardArr[0] > p6HighCardArr[0]) {

			if p1Rank == winningRank[0] {
				fmt.Println("Player 1 Wins (HighCard Compare)")
			}

		} else if (p2HighCardArr[4] > p1HighCardArr[4] && p2HighCardArr[4] > p3HighCardArr[4] && p2HighCardArr[4] > p4HighCardArr[4] && p2HighCardArr[4] > p5HighCardArr[4] && p2HighCardArr[4] > p6HighCardArr[4]) ||

			(p2HighCardArr[4] >= p1HighCardArr[4] && p2HighCardArr[4] >= p3HighCardArr[4] && p2HighCardArr[4] >= p4HighCardArr[4] && p2HighCardArr[4] >= p5HighCardArr[4] && p2HighCardArr[4] >= p6HighCardArr[4] &&
				p2HighCardArr[3] > p1HighCardArr[3] && p2HighCardArr[3] > p3HighCardArr[3] && p2HighCardArr[3] > p4HighCardArr[3] && p2HighCardArr[3] > p5HighCardArr[3] && p2HighCardArr[3] > p6HighCardArr[3]) ||

			(p2HighCardArr[4] >= p1HighCardArr[4] && p2HighCardArr[4] >= p3HighCardArr[4] && p2HighCardArr[4] >= p4HighCardArr[4] && p2HighCardArr[4] >= p5HighCardArr[4] && p2HighCardArr[4] >= p6HighCardArr[4] &&
				p2HighCardArr[3] >= p1HighCardArr[3] && p2HighCardArr[3] >= p3HighCardArr[3] && p2HighCardArr[3] >= p4HighCardArr[3] && p2HighCardArr[3] >= p5HighCardArr[3] && p2HighCardArr[3] >= p6HighCardArr[3] &&
				p2HighCardArr[2] > p1HighCardArr[2] && p2HighCardArr[2] > p3HighCardArr[2] && p2HighCardArr[2] > p4HighCardArr[2] && p2HighCardArr[2] > p5HighCardArr[2] && p2HighCardArr[2] > p6HighCardArr[2]) ||

			(p2HighCardArr[4] >= p1HighCardArr[4] && p2HighCardArr[4] >= p3HighCardArr[4] && p2HighCardArr[4] >= p4HighCardArr[4] && p2HighCardArr[4] >= p5HighCardArr[4] && p2HighCardArr[4] >= p6HighCardArr[4] &&
				p2HighCardArr[3] >= p1HighCardArr[3] && p2HighCardArr[3] >= p3HighCardArr[3] && p2HighCardArr[3] >= p4HighCardArr[3] && p2HighCardArr[3] >= p5HighCardArr[3] && p2HighCardArr[3] >= p6HighCardArr[3] &&
				p2HighCardArr[2] >= p1HighCardArr[2] && p2HighCardArr[2] >= p3HighCardArr[2] && p2HighCardArr[2] >= p4HighCardArr[2] && p2HighCardArr[2] >= p5HighCardArr[2] && p2HighCardArr[2] >= p6HighCardArr[2] &&
				p2HighCardArr[1] > p1HighCardArr[1] && p2HighCardArr[1] > p3HighCardArr[1] && p2HighCardArr[1] > p4HighCardArr[1] && p2HighCardArr[1] > p5HighCardArr[1] && p2HighCardArr[1] > p6HighCardArr[1]) ||

			(p2HighCardArr[4] >= p1HighCardArr[4] && p2HighCardArr[4] >= p3HighCardArr[4] && p2HighCardArr[4] >= p4HighCardArr[4] && p2HighCardArr[4] >= p5HighCardArr[4] && p2HighCardArr[4] >= p6HighCardArr[4] &&
				p2HighCardArr[3] >= p1HighCardArr[3] && p2HighCardArr[3] >= p3HighCardArr[3] && p2HighCardArr[3] >= p4HighCardArr[3] && p2HighCardArr[3] >= p5HighCardArr[3] && p2HighCardArr[3] >= p6HighCardArr[3] &&
				p2HighCardArr[2] >= p1HighCardArr[2] && p2HighCardArr[2] >= p3HighCardArr[2] && p2HighCardArr[2] >= p4HighCardArr[2] && p2HighCardArr[2] >= p5HighCardArr[2] && p2HighCardArr[2] >= p6HighCardArr[2] &&
				p2HighCardArr[1] >= p1HighCardArr[1] && p2HighCardArr[1] >= p3HighCardArr[1] && p2HighCardArr[1] >= p4HighCardArr[1] && p2HighCardArr[1] >= p5HighCardArr[1] && p2HighCardArr[1] >= p6HighCardArr[1] &&
				p2HighCardArr[0] > p1HighCardArr[0] && p2HighCardArr[0] > p3HighCardArr[0] && p2HighCardArr[0] > p4HighCardArr[0] && p2HighCardArr[0] > p5HighCardArr[0] && p2HighCardArr[0] > p6HighCardArr[0]) {

			if p2Rank == winningRank[0] {
				fmt.Println("Player 2 Wins (HighCard Compare)")

			}

		} else if (p3HighCardArr[4] > p1HighCardArr[4] && p3HighCardArr[4] > p2HighCardArr[4] && p3HighCardArr[4] > p4HighCardArr[4] && p3HighCardArr[4] > p5HighCardArr[4] && p3HighCardArr[4] > p6HighCardArr[4]) ||

			(p3HighCardArr[4] >= p1HighCardArr[4] && p3HighCardArr[4] >= p2HighCardArr[4] && p3HighCardArr[4] >= p4HighCardArr[4] && p3HighCardArr[4] >= p5HighCardArr[4] && p3HighCardArr[4] >= p6HighCardArr[4] &&
				p3HighCardArr[3] > p1HighCardArr[3] && p3HighCardArr[3] > p2HighCardArr[3] && p3HighCardArr[3] > p4HighCardArr[3] && p3HighCardArr[3] > p5HighCardArr[3] && p3HighCardArr[3] > p6HighCardArr[3]) ||

			(p3HighCardArr[4] >= p1HighCardArr[4] && p3HighCardArr[4] >= p2HighCardArr[4] && p3HighCardArr[4] >= p4HighCardArr[4] && p3HighCardArr[4] >= p5HighCardArr[4] && p3HighCardArr[4] >= p6HighCardArr[4] &&
				p3HighCardArr[3] >= p1HighCardArr[3] && p3HighCardArr[3] >= p2HighCardArr[3] && p3HighCardArr[3] >= p4HighCardArr[3] && p3HighCardArr[3] >= p5HighCardArr[3] && p3HighCardArr[3] >= p6HighCardArr[3] &&
				p3HighCardArr[2] > p1HighCardArr[2] && p3HighCardArr[2] > p2HighCardArr[2] && p3HighCardArr[2] > p4HighCardArr[2] && p3HighCardArr[2] > p5HighCardArr[2] && p3HighCardArr[2] > p6HighCardArr[2]) ||

			(p3HighCardArr[4] >= p1HighCardArr[4] && p3HighCardArr[4] >= p2HighCardArr[4] && p3HighCardArr[4] >= p4HighCardArr[4] && p3HighCardArr[4] >= p5HighCardArr[4] && p3HighCardArr[4] >= p6HighCardArr[4] &&
				p3HighCardArr[3] >= p1HighCardArr[3] && p3HighCardArr[3] >= p2HighCardArr[3] && p3HighCardArr[3] >= p4HighCardArr[3] && p3HighCardArr[3] >= p5HighCardArr[3] && p3HighCardArr[3] >= p6HighCardArr[3] &&
				p3HighCardArr[2] >= p1HighCardArr[2] && p3HighCardArr[2] >= p2HighCardArr[2] && p3HighCardArr[2] >= p4HighCardArr[2] && p3HighCardArr[2] >= p5HighCardArr[2] && p3HighCardArr[2] >= p6HighCardArr[2] &&
				p3HighCardArr[1] > p1HighCardArr[1] && p3HighCardArr[1] > p2HighCardArr[1] && p3HighCardArr[1] > p4HighCardArr[1] && p3HighCardArr[1] > p5HighCardArr[1] && p3HighCardArr[1] > p6HighCardArr[1]) ||

			(p3HighCardArr[4] >= p1HighCardArr[4] && p3HighCardArr[4] >= p2HighCardArr[4] && p3HighCardArr[4] >= p4HighCardArr[4] && p3HighCardArr[4] >= p5HighCardArr[4] && p3HighCardArr[4] >= p6HighCardArr[4] &&
				p3HighCardArr[3] >= p1HighCardArr[3] && p3HighCardArr[3] >= p2HighCardArr[3] && p3HighCardArr[3] >= p4HighCardArr[3] && p3HighCardArr[3] >= p5HighCardArr[3] && p3HighCardArr[3] >= p6HighCardArr[3] &&
				p3HighCardArr[2] >= p1HighCardArr[2] && p3HighCardArr[2] >= p2HighCardArr[2] && p3HighCardArr[2] >= p4HighCardArr[2] && p3HighCardArr[2] >= p5HighCardArr[2] && p3HighCardArr[2] >= p6HighCardArr[2] &&
				p3HighCardArr[1] >= p1HighCardArr[1] && p3HighCardArr[1] >= p2HighCardArr[1] && p3HighCardArr[1] >= p4HighCardArr[1] && p3HighCardArr[1] >= p5HighCardArr[1] && p3HighCardArr[1] >= p6HighCardArr[1] &&
				p3HighCardArr[0] > p1HighCardArr[0] && p3HighCardArr[0] > p2HighCardArr[0] && p3HighCardArr[0] > p4HighCardArr[0] && p3HighCardArr[0] > p5HighCardArr[0] && p3HighCardArr[0] > p6HighCardArr[0]) {

			if p3Rank == winningRank[0] {
				fmt.Println("Player 3 Wins (HighCard Compare)")

			}

		} else if (p4HighCardArr[4] > p1HighCardArr[4] && p4HighCardArr[4] > p2HighCardArr[4] && p4HighCardArr[4] > p3HighCardArr[4] && p4HighCardArr[4] > p5HighCardArr[4] && p4HighCardArr[4] > p6HighCardArr[4]) ||

			(p4HighCardArr[4] >= p1HighCardArr[4] && p4HighCardArr[4] >= p2HighCardArr[4] && p4HighCardArr[4] >= p3HighCardArr[4] && p4HighCardArr[4] >= p5HighCardArr[4] && p4HighCardArr[4] >= p6HighCardArr[4] &&
				p4HighCardArr[3] > p1HighCardArr[3] && p4HighCardArr[3] > p2HighCardArr[3] && p4HighCardArr[3] > p3HighCardArr[3] && p4HighCardArr[3] > p5HighCardArr[3] && p4HighCardArr[3] > p6HighCardArr[3]) ||

			(p4HighCardArr[4] >= p1HighCardArr[4] && p4HighCardArr[4] >= p2HighCardArr[4] && p4HighCardArr[4] >= p3HighCardArr[4] && p4HighCardArr[4] >= p5HighCardArr[4] && p4HighCardArr[4] >= p6HighCardArr[4] &&
				p4HighCardArr[3] >= p1HighCardArr[3] && p4HighCardArr[3] >= p2HighCardArr[3] && p4HighCardArr[3] >= p3HighCardArr[3] && p4HighCardArr[3] >= p5HighCardArr[3] && p4HighCardArr[3] >= p6HighCardArr[3] &&
				p4HighCardArr[2] > p1HighCardArr[2] && p4HighCardArr[2] > p2HighCardArr[2] && p4HighCardArr[2] > p3HighCardArr[2] && p4HighCardArr[2] > p5HighCardArr[2] && p4HighCardArr[2] > p6HighCardArr[2]) ||

			(p4HighCardArr[4] >= p1HighCardArr[4] && p4HighCardArr[4] >= p2HighCardArr[4] && p4HighCardArr[4] >= p3HighCardArr[4] && p4HighCardArr[4] >= p5HighCardArr[4] && p4HighCardArr[4] >= p6HighCardArr[4] &&
				p4HighCardArr[3] >= p1HighCardArr[3] && p4HighCardArr[3] >= p2HighCardArr[3] && p4HighCardArr[3] >= p3HighCardArr[3] && p4HighCardArr[3] >= p5HighCardArr[3] && p4HighCardArr[3] >= p6HighCardArr[3] &&
				p4HighCardArr[2] >= p1HighCardArr[2] && p4HighCardArr[2] >= p2HighCardArr[2] && p4HighCardArr[2] >= p3HighCardArr[2] && p4HighCardArr[2] >= p5HighCardArr[2] && p4HighCardArr[2] >= p6HighCardArr[2] &&
				p4HighCardArr[1] > p1HighCardArr[1] && p4HighCardArr[1] > p2HighCardArr[1] && p4HighCardArr[1] > p3HighCardArr[1] && p4HighCardArr[1] > p5HighCardArr[1] && p4HighCardArr[1] > p6HighCardArr[1]) ||

			(p4HighCardArr[4] >= p1HighCardArr[4] && p4HighCardArr[4] >= p2HighCardArr[4] && p4HighCardArr[4] >= p3HighCardArr[4] && p4HighCardArr[4] >= p5HighCardArr[4] && p4HighCardArr[4] >= p6HighCardArr[4] &&
				p4HighCardArr[3] >= p1HighCardArr[3] && p4HighCardArr[3] >= p2HighCardArr[3] && p4HighCardArr[3] >= p3HighCardArr[3] && p4HighCardArr[3] >= p5HighCardArr[3] && p4HighCardArr[3] >= p6HighCardArr[3] &&
				p4HighCardArr[2] >= p1HighCardArr[2] && p4HighCardArr[2] >= p2HighCardArr[2] && p4HighCardArr[2] >= p3HighCardArr[2] && p4HighCardArr[2] >= p5HighCardArr[2] && p4HighCardArr[2] >= p6HighCardArr[2] &&
				p4HighCardArr[1] >= p1HighCardArr[1] && p4HighCardArr[1] >= p2HighCardArr[1] && p4HighCardArr[1] >= p3HighCardArr[1] && p4HighCardArr[1] >= p5HighCardArr[1] && p4HighCardArr[1] >= p6HighCardArr[1] &&
				p4HighCardArr[0] > p1HighCardArr[0] && p4HighCardArr[0] > p2HighCardArr[0] && p4HighCardArr[0] > p3HighCardArr[0] && p4HighCardArr[0] > p5HighCardArr[0] && p4HighCardArr[0] > p6HighCardArr[0]) {

			if p4Rank == winningRank[0] {
				fmt.Println("Player 4 Wins (HighCard Compare)")

			}

		} else if (p5HighCardArr[4] > p1HighCardArr[4] && p5HighCardArr[4] > p2HighCardArr[4] && p5HighCardArr[4] > p3HighCardArr[4] && p5HighCardArr[4] > p4HighCardArr[4] && p5HighCardArr[4] > p6HighCardArr[4]) ||

			(p5HighCardArr[4] >= p1HighCardArr[4] && p5HighCardArr[4] >= p2HighCardArr[4] && p5HighCardArr[4] >= p3HighCardArr[4] && p5HighCardArr[4] >= p4HighCardArr[4] && p5HighCardArr[4] >= p6HighCardArr[4] &&
				p5HighCardArr[3] > p1HighCardArr[3] && p5HighCardArr[3] > p2HighCardArr[3] && p5HighCardArr[3] > p3HighCardArr[3] && p5HighCardArr[3] > p4HighCardArr[3] && p5HighCardArr[3] > p6HighCardArr[3]) ||

			(p5HighCardArr[4] >= p1HighCardArr[4] && p5HighCardArr[4] >= p2HighCardArr[4] && p5HighCardArr[4] >= p3HighCardArr[4] && p5HighCardArr[4] >= p4HighCardArr[4] && p5HighCardArr[4] >= p6HighCardArr[4] &&
				p5HighCardArr[3] >= p1HighCardArr[3] && p5HighCardArr[3] >= p2HighCardArr[3] && p5HighCardArr[3] >= p3HighCardArr[3] && p5HighCardArr[3] >= p4HighCardArr[3] && p5HighCardArr[3] >= p6HighCardArr[3] &&
				p5HighCardArr[2] > p1HighCardArr[2] && p5HighCardArr[2] > p2HighCardArr[2] && p5HighCardArr[2] > p3HighCardArr[2] && p5HighCardArr[2] > p4HighCardArr[2] && p5HighCardArr[2] > p6HighCardArr[2]) ||

			(p5HighCardArr[4] >= p1HighCardArr[4] && p5HighCardArr[4] >= p2HighCardArr[4] && p5HighCardArr[4] >= p3HighCardArr[4] && p5HighCardArr[4] >= p4HighCardArr[4] && p5HighCardArr[4] >= p6HighCardArr[4] &&
				p5HighCardArr[3] >= p1HighCardArr[3] && p5HighCardArr[3] >= p2HighCardArr[3] && p5HighCardArr[3] >= p3HighCardArr[3] && p5HighCardArr[3] >= p4HighCardArr[3] && p5HighCardArr[3] >= p6HighCardArr[3] &&
				p5HighCardArr[2] >= p1HighCardArr[2] && p5HighCardArr[2] >= p2HighCardArr[2] && p5HighCardArr[2] >= p3HighCardArr[2] && p5HighCardArr[2] >= p4HighCardArr[2] && p5HighCardArr[2] >= p6HighCardArr[2] &&
				p5HighCardArr[1] > p1HighCardArr[1] && p5HighCardArr[1] > p2HighCardArr[1] && p5HighCardArr[1] > p3HighCardArr[1] && p5HighCardArr[1] > p4HighCardArr[1] && p5HighCardArr[1] > p6HighCardArr[1]) ||

			(p5HighCardArr[4] >= p1HighCardArr[4] && p5HighCardArr[4] >= p2HighCardArr[4] && p5HighCardArr[4] >= p3HighCardArr[4] && p5HighCardArr[4] >= p4HighCardArr[4] && p5HighCardArr[4] >= p6HighCardArr[4] &&
				p5HighCardArr[3] >= p1HighCardArr[3] && p5HighCardArr[3] >= p2HighCardArr[3] && p5HighCardArr[3] >= p3HighCardArr[3] && p5HighCardArr[3] >= p4HighCardArr[3] && p5HighCardArr[3] >= p6HighCardArr[3] &&
				p5HighCardArr[2] >= p1HighCardArr[2] && p5HighCardArr[2] >= p2HighCardArr[2] && p5HighCardArr[2] >= p3HighCardArr[2] && p5HighCardArr[2] >= p4HighCardArr[2] && p5HighCardArr[2] >= p6HighCardArr[2] &&
				p5HighCardArr[1] >= p1HighCardArr[1] && p5HighCardArr[1] >= p2HighCardArr[1] && p5HighCardArr[1] >= p3HighCardArr[1] && p5HighCardArr[1] >= p4HighCardArr[1] && p5HighCardArr[1] >= p6HighCardArr[1] &&
				p5HighCardArr[0] > p1HighCardArr[0] && p5HighCardArr[0] > p2HighCardArr[0] && p5HighCardArr[0] > p3HighCardArr[0] && p5HighCardArr[0] > p4HighCardArr[0] && p5HighCardArr[0] > p6HighCardArr[0]) {

			if p5Rank == winningRank[0] {
				fmt.Println("Player 5 Wins (HighCard Compare)")

			}

		} else if (p6HighCardArr[4] > p1HighCardArr[4] && p6HighCardArr[4] > p2HighCardArr[4] && p6HighCardArr[4] > p3HighCardArr[4] && p6HighCardArr[4] > p4HighCardArr[4] && p6HighCardArr[4] > p5HighCardArr[4]) ||

			(p6HighCardArr[4] >= p1HighCardArr[4] && p6HighCardArr[4] >= p2HighCardArr[4] && p6HighCardArr[4] >= p3HighCardArr[4] && p6HighCardArr[4] >= p4HighCardArr[4] && p6HighCardArr[4] >= p5HighCardArr[4] &&
				p6HighCardArr[3] > p1HighCardArr[3] && p6HighCardArr[3] > p2HighCardArr[3] && p6HighCardArr[3] > p3HighCardArr[3] && p6HighCardArr[3] > p4HighCardArr[3] && p6HighCardArr[3] > p5HighCardArr[3]) ||

			(p6HighCardArr[4] >= p1HighCardArr[4] && p6HighCardArr[4] >= p2HighCardArr[4] && p6HighCardArr[4] >= p3HighCardArr[4] && p6HighCardArr[4] >= p4HighCardArr[4] && p6HighCardArr[4] >= p5HighCardArr[4] &&
				p6HighCardArr[3] >= p1HighCardArr[3] && p6HighCardArr[3] >= p2HighCardArr[3] && p6HighCardArr[3] >= p3HighCardArr[3] && p6HighCardArr[3] >= p4HighCardArr[3] && p6HighCardArr[3] >= p5HighCardArr[3] &&
				p6HighCardArr[2] > p1HighCardArr[2] && p6HighCardArr[2] > p2HighCardArr[2] && p6HighCardArr[2] > p3HighCardArr[2] && p6HighCardArr[2] > p4HighCardArr[2] && p6HighCardArr[2] > p5HighCardArr[2]) ||

			(p6HighCardArr[4] >= p1HighCardArr[4] && p6HighCardArr[4] >= p2HighCardArr[4] && p6HighCardArr[4] >= p3HighCardArr[4] && p6HighCardArr[4] >= p4HighCardArr[4] && p6HighCardArr[4] >= p5HighCardArr[4] &&
				p6HighCardArr[3] >= p1HighCardArr[3] && p6HighCardArr[3] >= p2HighCardArr[3] && p6HighCardArr[3] >= p3HighCardArr[3] && p6HighCardArr[3] >= p4HighCardArr[3] && p6HighCardArr[3] >= p5HighCardArr[3] &&
				p6HighCardArr[2] >= p1HighCardArr[2] && p6HighCardArr[2] >= p2HighCardArr[2] && p6HighCardArr[2] >= p3HighCardArr[2] && p6HighCardArr[2] >= p4HighCardArr[2] && p6HighCardArr[2] >= p5HighCardArr[2] &&
				p6HighCardArr[1] > p1HighCardArr[1] && p6HighCardArr[1] > p2HighCardArr[1] && p6HighCardArr[1] > p3HighCardArr[1] && p6HighCardArr[1] > p4HighCardArr[1] && p6HighCardArr[1] > p5HighCardArr[1]) ||

			(p6HighCardArr[4] >= p1HighCardArr[4] && p6HighCardArr[4] >= p2HighCardArr[4] && p6HighCardArr[4] >= p3HighCardArr[4] && p6HighCardArr[4] >= p4HighCardArr[4] && p6HighCardArr[4] >= p5HighCardArr[4] &&
				p6HighCardArr[3] >= p1HighCardArr[3] && p6HighCardArr[3] >= p2HighCardArr[3] && p6HighCardArr[3] >= p3HighCardArr[3] && p6HighCardArr[3] >= p4HighCardArr[3] && p6HighCardArr[3] >= p5HighCardArr[3] &&
				p6HighCardArr[2] >= p1HighCardArr[2] && p6HighCardArr[2] >= p2HighCardArr[2] && p6HighCardArr[2] >= p3HighCardArr[2] && p6HighCardArr[2] >= p4HighCardArr[2] && p6HighCardArr[2] >= p5HighCardArr[2] &&
				p6HighCardArr[1] >= p1HighCardArr[1] && p6HighCardArr[1] >= p2HighCardArr[1] && p6HighCardArr[1] >= p3HighCardArr[1] && p6HighCardArr[1] >= p4HighCardArr[1] && p6HighCardArr[1] >= p5HighCardArr[1] &&
				p6HighCardArr[0] > p1HighCardArr[0] && p6HighCardArr[0] > p2HighCardArr[0] && p6HighCardArr[0] > p3HighCardArr[0] && p6HighCardArr[0] > p4HighCardArr[0] && p6HighCardArr[0] > p5HighCardArr[0]) {

			if p6Rank == winningRank[0] {
				fmt.Println("Player 6 Wins (HighCard Compare)")

			}

		} else {
			fmt.Println(p1HighCardArr)
			fmt.Println(p2HighCardArr)
			fmt.Println("Push")
		}
	}
}

func suitSplit(card int) { /// Splits cards inside getHand#() funcs
	switch card {
	////// Spades
	case 1:
		arrSplit[0] = 14
		arrSplit[1] = 0

	case 2:
		arrSplit[0] = 2
		arrSplit[1] = 0

	case 3:
		arrSplit[0] = 3
		arrSplit[1] = 0

	case 4:
		arrSplit[0] = 4
		arrSplit[1] = 0

	case 5:
		arrSplit[0] = 5
		arrSplit[1] = 0

	case 6:
		arrSplit[0] = 6
		arrSplit[1] = 0

	case 7:
		arrSplit[0] = 7
		arrSplit[1] = 0

	case 8:
		arrSplit[0] = 8
		arrSplit[1] = 0

	case 9:
		arrSplit[0] = 9
		arrSplit[1] = 0

	case 10:
		arrSplit[0] = 10
		arrSplit[1] = 0

	case 11:
		arrSplit[0] = 11
		arrSplit[1] = 0

	case 12:
		arrSplit[0] = 12
		arrSplit[1] = 0

	case 13:
		arrSplit[0] = 13
		arrSplit[1] = 0

		////// Hearts
	case 14:
		arrSplit[0] = 14
		arrSplit[1] = 13

	case 15:
		arrSplit[0] = 2
		arrSplit[1] = 13

	case 16:
		arrSplit[0] = 3
		arrSplit[1] = 13

	case 17:
		arrSplit[0] = 4
		arrSplit[1] = 13

	case 18:
		arrSplit[0] = 5
		arrSplit[1] = 13

	case 19:
		arrSplit[0] = 6
		arrSplit[1] = 13

	case 20:
		arrSplit[0] = 7
		arrSplit[1] = 13

	case 21:
		arrSplit[0] = 8
		arrSplit[1] = 13

	case 22:
		arrSplit[0] = 9
		arrSplit[1] = 13

	case 23:
		arrSplit[0] = 10
		arrSplit[1] = 13

	case 24:
		arrSplit[0] = 11
		arrSplit[1] = 13

	case 25:
		arrSplit[0] = 12
		arrSplit[1] = 13

	case 26:
		arrSplit[0] = 13
		arrSplit[1] = 13

		////// Clubs
	case 27:
		arrSplit[0] = 14
		arrSplit[1] = 26

	case 28:
		arrSplit[0] = 2
		arrSplit[1] = 26

	case 29:
		arrSplit[0] = 3
		arrSplit[1] = 26

	case 30:
		arrSplit[0] = 4
		arrSplit[1] = 26

	case 31:
		arrSplit[0] = 5
		arrSplit[1] = 26

	case 32:
		arrSplit[0] = 6
		arrSplit[1] = 26

	case 33:
		arrSplit[0] = 7
		arrSplit[1] = 26

	case 34:
		arrSplit[0] = 8
		arrSplit[1] = 26

	case 35:
		arrSplit[0] = 9
		arrSplit[1] = 26

	case 36:
		arrSplit[0] = 10
		arrSplit[1] = 26

	case 37:
		arrSplit[0] = 11
		arrSplit[1] = 26

	case 38:
		arrSplit[0] = 12
		arrSplit[1] = 26

	case 39:
		arrSplit[0] = 13
		arrSplit[1] = 26

		////// Diamonds
	case 40:
		arrSplit[0] = 14
		arrSplit[1] = 39

	case 41:
		arrSplit[0] = 2
		arrSplit[1] = 39

	case 42:
		arrSplit[0] = 3
		arrSplit[1] = 39

	case 43:
		arrSplit[0] = 4
		arrSplit[1] = 39

	case 44:
		arrSplit[0] = 5
		arrSplit[1] = 39

	case 45:
		arrSplit[0] = 6
		arrSplit[1] = 39

	case 46:
		arrSplit[0] = 7
		arrSplit[1] = 39

	case 47:
		arrSplit[0] = 8
		arrSplit[1] = 39

	case 48:
		arrSplit[0] = 9
		arrSplit[1] = 39

	case 49:
		arrSplit[0] = 10
		arrSplit[1] = 39

	case 50:
		arrSplit[0] = 11
		arrSplit[1] = 39

	case 51:
		arrSplit[0] = 12
		arrSplit[1] = 39

	case 52:
		arrSplit[0] = 13
		arrSplit[1] = 39

	}
}

func compare1_2() { /// Compare two individual hands and strip loosing hand
	if (p1HighCardArr[4] > p2HighCardArr[4]) ||
		(p1HighCardArr[4] == p2HighCardArr[4] && p1HighCardArr[3] > p2HighCardArr[3]) ||
		(p1HighCardArr[4] == p2HighCardArr[4] && p1HighCardArr[3] == p2HighCardArr[3] && p1HighCardArr[2] > p2HighCardArr[2]) ||
		(p1HighCardArr[4] == p2HighCardArr[4] && p1HighCardArr[3] == p2HighCardArr[3] && p1HighCardArr[2] == p2HighCardArr[2] && p1HighCardArr[1] > p2HighCardArr[1]) ||
		(p1HighCardArr[4] == p2HighCardArr[4] && p1HighCardArr[3] == p2HighCardArr[3] && p1HighCardArr[2] == p2HighCardArr[2] && p1HighCardArr[1] == p2HighCardArr[1] && p1HighCardArr[0] > p2HighCardArr[0]) {

		less2()
	}

}

func compare1_3() {
	if (p1HighCardArr[4] > p3HighCardArr[4]) ||
		(p1HighCardArr[4] == p3HighCardArr[4] && p1HighCardArr[3] > p3HighCardArr[3]) ||
		(p1HighCardArr[4] == p3HighCardArr[4] && p1HighCardArr[3] == p3HighCardArr[3] && p1HighCardArr[2] > p3HighCardArr[2]) ||
		(p1HighCardArr[4] == p3HighCardArr[4] && p1HighCardArr[3] == p3HighCardArr[3] && p1HighCardArr[2] == p3HighCardArr[2] && p1HighCardArr[1] > p3HighCardArr[1]) ||
		(p1HighCardArr[4] == p3HighCardArr[4] && p1HighCardArr[3] == p3HighCardArr[3] && p1HighCardArr[2] == p3HighCardArr[2] && p1HighCardArr[1] == p3HighCardArr[1] && p1HighCardArr[0] > p3HighCardArr[0]) {

		less3()
	}

}

func compare1_4() {
	if (p1HighCardArr[4] > p4HighCardArr[4]) ||
		(p1HighCardArr[4] == p4HighCardArr[4] && p1HighCardArr[3] > p4HighCardArr[3]) ||
		(p1HighCardArr[4] == p4HighCardArr[4] && p1HighCardArr[3] == p4HighCardArr[3] && p1HighCardArr[2] > p4HighCardArr[2]) ||
		(p1HighCardArr[4] == p4HighCardArr[4] && p1HighCardArr[3] == p4HighCardArr[3] && p1HighCardArr[2] == p4HighCardArr[2] && p1HighCardArr[1] > p4HighCardArr[1]) ||
		(p1HighCardArr[4] == p4HighCardArr[4] && p1HighCardArr[3] == p4HighCardArr[3] && p1HighCardArr[2] == p4HighCardArr[2] && p1HighCardArr[1] == p4HighCardArr[1] && p1HighCardArr[0] > p4HighCardArr[0]) {

		less4()
	}

}

func compare1_5() {
	if (p1HighCardArr[4] > p5HighCardArr[4]) ||
		(p1HighCardArr[4] == p5HighCardArr[4] && p1HighCardArr[3] > p5HighCardArr[3]) ||
		(p1HighCardArr[4] == p5HighCardArr[4] && p1HighCardArr[3] == p5HighCardArr[3] && p1HighCardArr[2] > p5HighCardArr[2]) ||
		(p1HighCardArr[4] == p5HighCardArr[4] && p1HighCardArr[3] == p5HighCardArr[3] && p1HighCardArr[2] == p5HighCardArr[2] && p1HighCardArr[1] > p5HighCardArr[1]) ||
		(p1HighCardArr[4] == p5HighCardArr[4] && p1HighCardArr[3] == p5HighCardArr[3] && p1HighCardArr[2] == p5HighCardArr[2] && p1HighCardArr[1] == p5HighCardArr[1] && p1HighCardArr[0] > p5HighCardArr[0]) {

		less5()
	}

}

func compare1_6() {
	if (p1HighCardArr[4] > p6HighCardArr[4]) ||
		(p1HighCardArr[4] == p6HighCardArr[4] && p1HighCardArr[3] > p6HighCardArr[3]) ||
		(p1HighCardArr[4] == p6HighCardArr[4] && p1HighCardArr[3] == p6HighCardArr[3] && p1HighCardArr[2] > p6HighCardArr[2]) ||
		(p1HighCardArr[4] == p6HighCardArr[4] && p1HighCardArr[3] == p6HighCardArr[3] && p1HighCardArr[2] == p6HighCardArr[2] && p1HighCardArr[1] > p6HighCardArr[1]) ||
		(p1HighCardArr[4] == p6HighCardArr[4] && p1HighCardArr[3] == p6HighCardArr[3] && p1HighCardArr[2] == p6HighCardArr[2] && p1HighCardArr[1] == p6HighCardArr[1] && p1HighCardArr[0] > p6HighCardArr[0]) {

		less6()
	}

}

func compare2_1() {
	if (p2HighCardArr[4] > p1HighCardArr[4]) ||
		(p2HighCardArr[4] == p1HighCardArr[4] && p2HighCardArr[3] > p1HighCardArr[3]) ||
		(p2HighCardArr[4] == p1HighCardArr[4] && p2HighCardArr[3] == p1HighCardArr[3] && p2HighCardArr[2] > p1HighCardArr[2]) ||
		(p2HighCardArr[4] == p1HighCardArr[4] && p2HighCardArr[3] == p1HighCardArr[3] && p2HighCardArr[2] == p1HighCardArr[2] && p2HighCardArr[1] > p1HighCardArr[1]) ||
		(p2HighCardArr[4] == p1HighCardArr[4] && p2HighCardArr[3] == p1HighCardArr[3] && p2HighCardArr[2] == p1HighCardArr[2] && p2HighCardArr[1] == p1HighCardArr[1] && p2HighCardArr[0] > p1HighCardArr[0]) {

		less1()
	}

}

func compare2_3() {
	if (p2HighCardArr[4] > p3HighCardArr[4]) ||
		(p2HighCardArr[4] == p3HighCardArr[4] && p2HighCardArr[3] > p3HighCardArr[3]) ||
		(p2HighCardArr[4] == p3HighCardArr[4] && p2HighCardArr[3] == p3HighCardArr[3] && p2HighCardArr[2] > p3HighCardArr[2]) ||
		(p2HighCardArr[4] == p3HighCardArr[4] && p2HighCardArr[3] == p3HighCardArr[3] && p2HighCardArr[2] == p3HighCardArr[2] && p2HighCardArr[1] > p3HighCardArr[1]) ||
		(p2HighCardArr[4] == p3HighCardArr[4] && p2HighCardArr[3] == p3HighCardArr[3] && p2HighCardArr[2] == p3HighCardArr[2] && p2HighCardArr[1] == p3HighCardArr[1] && p2HighCardArr[0] > p3HighCardArr[0]) {

		less3()
	}

}

func compare2_4() {
	if (p2HighCardArr[4] > p4HighCardArr[4]) ||
		(p2HighCardArr[4] == p4HighCardArr[4] && p2HighCardArr[3] > p4HighCardArr[3]) ||
		(p2HighCardArr[4] == p4HighCardArr[4] && p2HighCardArr[3] == p4HighCardArr[3] && p2HighCardArr[2] > p4HighCardArr[2]) ||
		(p2HighCardArr[4] == p4HighCardArr[4] && p2HighCardArr[3] == p4HighCardArr[3] && p2HighCardArr[2] == p4HighCardArr[2] && p2HighCardArr[1] > p4HighCardArr[1]) ||
		(p2HighCardArr[4] == p4HighCardArr[4] && p2HighCardArr[3] == p4HighCardArr[3] && p2HighCardArr[2] == p4HighCardArr[2] && p2HighCardArr[1] == p4HighCardArr[1] && p2HighCardArr[0] > p4HighCardArr[0]) {

		less4()
	}

}

func compare2_5() {
	if (p2HighCardArr[4] > p5HighCardArr[4]) ||
		(p2HighCardArr[4] == p5HighCardArr[4] && p2HighCardArr[3] > p5HighCardArr[3]) ||
		(p2HighCardArr[4] == p5HighCardArr[4] && p2HighCardArr[3] == p5HighCardArr[3] && p2HighCardArr[2] > p5HighCardArr[2]) ||
		(p2HighCardArr[4] == p5HighCardArr[4] && p2HighCardArr[3] == p5HighCardArr[3] && p2HighCardArr[2] == p5HighCardArr[2] && p2HighCardArr[1] > p5HighCardArr[1]) ||
		(p2HighCardArr[4] == p5HighCardArr[4] && p2HighCardArr[3] == p5HighCardArr[3] && p2HighCardArr[2] == p5HighCardArr[2] && p2HighCardArr[1] == p5HighCardArr[1] && p2HighCardArr[0] > p5HighCardArr[0]) {

		less5()
	}

}

func compare2_6() {
	if (p2HighCardArr[4] > p6HighCardArr[4]) ||
		(p2HighCardArr[4] == p6HighCardArr[4] && p2HighCardArr[3] > p6HighCardArr[3]) ||
		(p2HighCardArr[4] == p6HighCardArr[4] && p2HighCardArr[3] == p6HighCardArr[3] && p2HighCardArr[2] > p6HighCardArr[2]) ||
		(p2HighCardArr[4] == p6HighCardArr[4] && p2HighCardArr[3] == p6HighCardArr[3] && p2HighCardArr[2] == p6HighCardArr[2] && p2HighCardArr[1] > p6HighCardArr[1]) ||
		(p2HighCardArr[4] == p6HighCardArr[4] && p2HighCardArr[3] == p6HighCardArr[3] && p2HighCardArr[2] == p6HighCardArr[2] && p2HighCardArr[1] == p6HighCardArr[1] && p2HighCardArr[0] > p6HighCardArr[0]) {

		less6()
	}

}

func compare3_1() {
	if (p3HighCardArr[4] > p1HighCardArr[4]) ||
		(p3HighCardArr[4] == p1HighCardArr[4] && p3HighCardArr[3] > p1HighCardArr[3]) ||
		(p3HighCardArr[4] == p1HighCardArr[4] && p3HighCardArr[3] == p1HighCardArr[3] && p3HighCardArr[2] > p1HighCardArr[2]) ||
		(p3HighCardArr[4] == p1HighCardArr[4] && p3HighCardArr[3] == p1HighCardArr[3] && p3HighCardArr[2] == p1HighCardArr[2] && p3HighCardArr[1] > p1HighCardArr[1]) ||
		(p3HighCardArr[4] == p1HighCardArr[4] && p3HighCardArr[3] == p1HighCardArr[3] && p3HighCardArr[2] == p1HighCardArr[2] && p3HighCardArr[1] == p1HighCardArr[1] && p3HighCardArr[0] > p1HighCardArr[0]) {

		less1()
	}

}

func compare3_2() {
	if (p3HighCardArr[4] > p2HighCardArr[4]) ||
		(p3HighCardArr[4] == p2HighCardArr[4] && p3HighCardArr[3] > p2HighCardArr[3]) ||
		(p3HighCardArr[4] == p2HighCardArr[4] && p3HighCardArr[3] == p2HighCardArr[3] && p3HighCardArr[2] > p2HighCardArr[2]) ||
		(p3HighCardArr[4] == p2HighCardArr[4] && p3HighCardArr[3] == p2HighCardArr[3] && p3HighCardArr[2] == p2HighCardArr[2] && p3HighCardArr[1] > p2HighCardArr[1]) ||
		(p3HighCardArr[4] == p2HighCardArr[4] && p3HighCardArr[3] == p2HighCardArr[3] && p3HighCardArr[2] == p2HighCardArr[2] && p3HighCardArr[1] == p2HighCardArr[1] && p3HighCardArr[0] > p2HighCardArr[0]) {

		less2()
	}

}

func compare3_4() {
	if (p3HighCardArr[4] > p4HighCardArr[4]) ||
		(p3HighCardArr[4] == p4HighCardArr[4] && p3HighCardArr[3] > p4HighCardArr[3]) ||
		(p3HighCardArr[4] == p4HighCardArr[4] && p3HighCardArr[3] == p4HighCardArr[3] && p3HighCardArr[2] > p4HighCardArr[2]) ||
		(p3HighCardArr[4] == p4HighCardArr[4] && p3HighCardArr[3] == p4HighCardArr[3] && p3HighCardArr[2] == p4HighCardArr[2] && p3HighCardArr[1] > p4HighCardArr[1]) ||
		(p3HighCardArr[4] == p4HighCardArr[4] && p3HighCardArr[3] == p4HighCardArr[3] && p3HighCardArr[2] == p4HighCardArr[2] && p3HighCardArr[1] == p4HighCardArr[1] && p3HighCardArr[0] > p4HighCardArr[0]) {

		less4()
	}

}

func compare3_5() {
	if (p3HighCardArr[4] > p5HighCardArr[4]) ||
		(p3HighCardArr[4] == p5HighCardArr[4] && p3HighCardArr[3] > p5HighCardArr[3]) ||
		(p3HighCardArr[4] == p5HighCardArr[4] && p3HighCardArr[3] == p5HighCardArr[3] && p3HighCardArr[2] > p5HighCardArr[2]) ||
		(p3HighCardArr[4] == p5HighCardArr[4] && p3HighCardArr[3] == p5HighCardArr[3] && p3HighCardArr[2] == p5HighCardArr[2] && p3HighCardArr[1] > p5HighCardArr[1]) ||
		(p3HighCardArr[4] == p5HighCardArr[4] && p3HighCardArr[3] == p5HighCardArr[3] && p3HighCardArr[2] == p5HighCardArr[2] && p3HighCardArr[1] == p5HighCardArr[1] && p3HighCardArr[0] > p5HighCardArr[0]) {

		less5()
	}

}

func compare3_6() {
	if (p3HighCardArr[4] > p6HighCardArr[4]) ||
		(p3HighCardArr[4] == p6HighCardArr[4] && p3HighCardArr[3] > p6HighCardArr[3]) ||
		(p3HighCardArr[4] == p6HighCardArr[4] && p3HighCardArr[3] == p6HighCardArr[3] && p3HighCardArr[2] > p6HighCardArr[2]) ||
		(p3HighCardArr[4] == p6HighCardArr[4] && p3HighCardArr[3] == p6HighCardArr[3] && p3HighCardArr[2] == p6HighCardArr[2] && p3HighCardArr[1] > p6HighCardArr[1]) ||
		(p3HighCardArr[4] == p6HighCardArr[4] && p3HighCardArr[3] == p6HighCardArr[3] && p3HighCardArr[2] == p6HighCardArr[2] && p3HighCardArr[1] == p6HighCardArr[1] && p3HighCardArr[0] > p6HighCardArr[0]) {

		less6()
	}

}

func compare4_1() {
	if (p4HighCardArr[4] > p1HighCardArr[4]) ||
		(p4HighCardArr[4] == p1HighCardArr[4] && p4HighCardArr[3] > p1HighCardArr[3]) ||
		(p4HighCardArr[4] == p1HighCardArr[4] && p4HighCardArr[3] == p1HighCardArr[3] && p4HighCardArr[2] > p1HighCardArr[2]) ||
		(p4HighCardArr[4] == p1HighCardArr[4] && p4HighCardArr[3] == p1HighCardArr[3] && p4HighCardArr[2] == p1HighCardArr[2] && p4HighCardArr[1] > p1HighCardArr[1]) ||
		(p4HighCardArr[4] == p1HighCardArr[4] && p4HighCardArr[3] == p1HighCardArr[3] && p4HighCardArr[2] == p1HighCardArr[2] && p4HighCardArr[1] == p1HighCardArr[1] && p4HighCardArr[0] > p1HighCardArr[0]) {

		less1()
	}

}

func compare4_2() {
	if (p4HighCardArr[4] > p2HighCardArr[4]) ||
		(p4HighCardArr[4] == p2HighCardArr[4] && p4HighCardArr[3] > p2HighCardArr[3]) ||
		(p4HighCardArr[4] == p2HighCardArr[4] && p4HighCardArr[3] == p2HighCardArr[3] && p4HighCardArr[2] > p2HighCardArr[2]) ||
		(p4HighCardArr[4] == p2HighCardArr[4] && p4HighCardArr[3] == p2HighCardArr[3] && p4HighCardArr[2] == p2HighCardArr[2] && p4HighCardArr[1] > p2HighCardArr[1]) ||
		(p4HighCardArr[4] == p2HighCardArr[4] && p4HighCardArr[3] == p2HighCardArr[3] && p4HighCardArr[2] == p2HighCardArr[2] && p4HighCardArr[1] == p2HighCardArr[1] && p4HighCardArr[0] > p2HighCardArr[0]) {

		less2()
	}

}

func compare4_3() {
	if (p4HighCardArr[4] > p3HighCardArr[4]) ||
		(p4HighCardArr[4] == p3HighCardArr[4] && p4HighCardArr[3] > p3HighCardArr[3]) ||
		(p4HighCardArr[4] == p3HighCardArr[4] && p4HighCardArr[3] == p3HighCardArr[3] && p4HighCardArr[2] > p3HighCardArr[2]) ||
		(p4HighCardArr[4] == p3HighCardArr[4] && p4HighCardArr[3] == p3HighCardArr[3] && p4HighCardArr[2] == p3HighCardArr[2] && p4HighCardArr[1] > p3HighCardArr[1]) ||
		(p4HighCardArr[4] == p3HighCardArr[4] && p4HighCardArr[3] == p3HighCardArr[3] && p4HighCardArr[2] == p3HighCardArr[2] && p4HighCardArr[1] == p3HighCardArr[1] && p4HighCardArr[0] > p3HighCardArr[0]) {

		less3()
	}

}

func compare4_5() {
	if (p4HighCardArr[4] > p5HighCardArr[4]) ||
		(p4HighCardArr[4] == p5HighCardArr[4] && p4HighCardArr[3] > p5HighCardArr[3]) ||
		(p4HighCardArr[4] == p5HighCardArr[4] && p4HighCardArr[3] == p5HighCardArr[3] && p4HighCardArr[2] > p5HighCardArr[2]) ||
		(p4HighCardArr[4] == p5HighCardArr[4] && p4HighCardArr[3] == p5HighCardArr[3] && p4HighCardArr[2] == p5HighCardArr[2] && p4HighCardArr[1] > p5HighCardArr[1]) ||
		(p4HighCardArr[4] == p5HighCardArr[4] && p4HighCardArr[3] == p5HighCardArr[3] && p4HighCardArr[2] == p5HighCardArr[2] && p4HighCardArr[1] == p5HighCardArr[1] && p4HighCardArr[0] > p5HighCardArr[0]) {

		less5()
	}

}

func compare4_6() {
	if (p4HighCardArr[4] > p6HighCardArr[4]) ||
		(p4HighCardArr[4] == p6HighCardArr[4] && p4HighCardArr[3] > p6HighCardArr[3]) ||
		(p4HighCardArr[4] == p6HighCardArr[4] && p4HighCardArr[3] == p6HighCardArr[3] && p4HighCardArr[2] > p6HighCardArr[2]) ||
		(p4HighCardArr[4] == p6HighCardArr[4] && p4HighCardArr[3] == p6HighCardArr[3] && p4HighCardArr[2] == p6HighCardArr[2] && p4HighCardArr[1] > p6HighCardArr[1]) ||
		(p4HighCardArr[4] == p6HighCardArr[4] && p4HighCardArr[3] == p6HighCardArr[3] && p4HighCardArr[2] == p6HighCardArr[2] && p4HighCardArr[1] == p6HighCardArr[1] && p4HighCardArr[0] > p6HighCardArr[0]) {

		less6()
	}

}

func compare5_1() {
	if (p5HighCardArr[4] > p1HighCardArr[4]) ||
		(p5HighCardArr[4] == p1HighCardArr[4] && p5HighCardArr[3] > p1HighCardArr[3]) ||
		(p5HighCardArr[4] == p1HighCardArr[4] && p5HighCardArr[3] == p1HighCardArr[3] && p5HighCardArr[2] > p1HighCardArr[2]) ||
		(p5HighCardArr[4] == p1HighCardArr[4] && p5HighCardArr[3] == p1HighCardArr[3] && p5HighCardArr[2] == p1HighCardArr[2] && p5HighCardArr[1] > p1HighCardArr[1]) ||
		(p5HighCardArr[4] == p1HighCardArr[4] && p5HighCardArr[3] == p1HighCardArr[3] && p5HighCardArr[2] == p1HighCardArr[2] && p5HighCardArr[1] == p1HighCardArr[1] && p5HighCardArr[0] > p1HighCardArr[0]) {

		less1()
	}

}

func compare5_2() {
	if (p5HighCardArr[4] > p2HighCardArr[4]) ||
		(p5HighCardArr[4] == p2HighCardArr[4] && p5HighCardArr[3] > p2HighCardArr[3]) ||
		(p5HighCardArr[4] == p2HighCardArr[4] && p5HighCardArr[3] == p2HighCardArr[3] && p5HighCardArr[2] > p2HighCardArr[2]) ||
		(p5HighCardArr[4] == p2HighCardArr[4] && p5HighCardArr[3] == p2HighCardArr[3] && p5HighCardArr[2] == p2HighCardArr[2] && p5HighCardArr[1] > p2HighCardArr[1]) ||
		(p5HighCardArr[4] == p2HighCardArr[4] && p5HighCardArr[3] == p2HighCardArr[3] && p5HighCardArr[2] == p2HighCardArr[2] && p5HighCardArr[1] == p2HighCardArr[1] && p5HighCardArr[0] > p2HighCardArr[0]) {

		less2()
	}

}

func compare5_3() {
	if (p5HighCardArr[4] > p3HighCardArr[4]) ||
		(p5HighCardArr[4] == p3HighCardArr[4] && p5HighCardArr[3] > p3HighCardArr[3]) ||
		(p5HighCardArr[4] == p3HighCardArr[4] && p5HighCardArr[3] == p3HighCardArr[3] && p5HighCardArr[2] > p3HighCardArr[2]) ||
		(p5HighCardArr[4] == p3HighCardArr[4] && p5HighCardArr[3] == p3HighCardArr[3] && p5HighCardArr[2] == p3HighCardArr[2] && p5HighCardArr[1] > p3HighCardArr[1]) ||
		(p5HighCardArr[4] == p3HighCardArr[4] && p5HighCardArr[3] == p3HighCardArr[3] && p5HighCardArr[2] == p3HighCardArr[2] && p5HighCardArr[1] == p3HighCardArr[1] && p5HighCardArr[0] > p3HighCardArr[0]) {

		less3()
	}

}

func compare5_4() {
	if (p5HighCardArr[4] > p4HighCardArr[4]) ||
		(p5HighCardArr[4] == p4HighCardArr[4] && p5HighCardArr[3] > p4HighCardArr[3]) ||
		(p5HighCardArr[4] == p4HighCardArr[4] && p5HighCardArr[3] == p4HighCardArr[3] && p5HighCardArr[2] > p4HighCardArr[2]) ||
		(p5HighCardArr[4] == p4HighCardArr[4] && p5HighCardArr[3] == p4HighCardArr[3] && p5HighCardArr[2] == p4HighCardArr[2] && p5HighCardArr[1] > p4HighCardArr[1]) ||
		(p5HighCardArr[4] == p4HighCardArr[4] && p5HighCardArr[3] == p4HighCardArr[3] && p5HighCardArr[2] == p4HighCardArr[2] && p5HighCardArr[1] == p4HighCardArr[1] && p5HighCardArr[0] > p4HighCardArr[0]) {

		less4()
	}

}

func compare5_6() {
	if (p5HighCardArr[4] > p6HighCardArr[4]) ||
		(p5HighCardArr[4] == p6HighCardArr[4] && p5HighCardArr[3] > p6HighCardArr[3]) ||
		(p5HighCardArr[4] == p6HighCardArr[4] && p5HighCardArr[3] == p6HighCardArr[3] && p5HighCardArr[2] > p6HighCardArr[2]) ||
		(p5HighCardArr[4] == p6HighCardArr[4] && p5HighCardArr[3] == p6HighCardArr[3] && p5HighCardArr[2] == p6HighCardArr[2] && p5HighCardArr[1] > p6HighCardArr[1]) ||
		(p5HighCardArr[4] == p6HighCardArr[4] && p5HighCardArr[3] == p6HighCardArr[3] && p5HighCardArr[2] == p6HighCardArr[2] && p5HighCardArr[1] == p6HighCardArr[1] && p5HighCardArr[0] > p6HighCardArr[0]) {

		less6()
	}

}

func compare6_1() {
	if (p6HighCardArr[4] > p1HighCardArr[4]) ||
		(p6HighCardArr[4] == p1HighCardArr[4] && p6HighCardArr[3] > p1HighCardArr[3]) ||
		(p6HighCardArr[4] == p1HighCardArr[4] && p6HighCardArr[3] == p1HighCardArr[3] && p6HighCardArr[2] > p1HighCardArr[2]) ||
		(p6HighCardArr[4] == p1HighCardArr[4] && p6HighCardArr[3] == p1HighCardArr[3] && p6HighCardArr[2] == p1HighCardArr[2] && p6HighCardArr[1] > p1HighCardArr[1]) ||
		(p6HighCardArr[4] == p1HighCardArr[4] && p6HighCardArr[3] == p1HighCardArr[3] && p6HighCardArr[2] == p1HighCardArr[2] && p6HighCardArr[1] == p1HighCardArr[1] && p6HighCardArr[0] > p1HighCardArr[0]) {

		less1()
	}

}

func compare6_2() {
	if (p6HighCardArr[4] > p2HighCardArr[4]) ||
		(p6HighCardArr[4] == p2HighCardArr[4] && p6HighCardArr[3] > p2HighCardArr[3]) ||
		(p6HighCardArr[4] == p2HighCardArr[4] && p6HighCardArr[3] == p2HighCardArr[3] && p6HighCardArr[2] > p2HighCardArr[2]) ||
		(p6HighCardArr[4] == p2HighCardArr[4] && p6HighCardArr[3] == p2HighCardArr[3] && p6HighCardArr[2] == p2HighCardArr[2] && p6HighCardArr[1] > p2HighCardArr[1]) ||
		(p6HighCardArr[4] == p2HighCardArr[4] && p6HighCardArr[3] == p2HighCardArr[3] && p6HighCardArr[2] == p2HighCardArr[2] && p6HighCardArr[1] == p2HighCardArr[1] && p6HighCardArr[0] > p2HighCardArr[0]) {

		less2()
	}

}

func compare6_3() {
	if (p6HighCardArr[4] > p3HighCardArr[4]) ||
		(p6HighCardArr[4] == p3HighCardArr[4] && p6HighCardArr[3] > p3HighCardArr[3]) ||
		(p6HighCardArr[4] == p3HighCardArr[4] && p6HighCardArr[3] == p3HighCardArr[3] && p6HighCardArr[2] > p3HighCardArr[2]) ||
		(p6HighCardArr[4] == p3HighCardArr[4] && p6HighCardArr[3] == p3HighCardArr[3] && p6HighCardArr[2] == p3HighCardArr[2] && p6HighCardArr[1] > p3HighCardArr[1]) ||
		(p6HighCardArr[4] == p3HighCardArr[4] && p6HighCardArr[3] == p3HighCardArr[3] && p6HighCardArr[2] == p3HighCardArr[2] && p6HighCardArr[1] == p3HighCardArr[1] && p6HighCardArr[0] > p3HighCardArr[0]) {

		less3()
	}

}

func compare6_4() {
	if (p6HighCardArr[4] > p4HighCardArr[4]) ||
		(p6HighCardArr[4] == p4HighCardArr[4] && p6HighCardArr[3] > p4HighCardArr[3]) ||
		(p6HighCardArr[4] == p4HighCardArr[4] && p6HighCardArr[3] == p4HighCardArr[3] && p6HighCardArr[2] > p4HighCardArr[2]) ||
		(p6HighCardArr[4] == p4HighCardArr[4] && p6HighCardArr[3] == p4HighCardArr[3] && p6HighCardArr[2] == p4HighCardArr[2] && p6HighCardArr[1] > p4HighCardArr[1]) ||
		(p6HighCardArr[4] == p4HighCardArr[4] && p6HighCardArr[3] == p4HighCardArr[3] && p6HighCardArr[2] == p4HighCardArr[2] && p6HighCardArr[1] == p4HighCardArr[1] && p6HighCardArr[0] > p4HighCardArr[0]) {

		less4()
	}

}

func compare6_5() {
	if (p6HighCardArr[4] > p5HighCardArr[4]) ||
		(p6HighCardArr[4] == p5HighCardArr[4] && p6HighCardArr[3] > p5HighCardArr[3]) ||
		(p6HighCardArr[4] == p5HighCardArr[4] && p6HighCardArr[3] == p5HighCardArr[3] && p6HighCardArr[2] > p5HighCardArr[2]) ||
		(p6HighCardArr[4] == p5HighCardArr[4] && p6HighCardArr[3] == p5HighCardArr[3] && p6HighCardArr[2] == p5HighCardArr[2] && p6HighCardArr[1] > p5HighCardArr[1]) ||
		(p6HighCardArr[4] == p5HighCardArr[4] && p6HighCardArr[3] == p5HighCardArr[3] && p6HighCardArr[2] == p5HighCardArr[2] && p6HighCardArr[1] == p5HighCardArr[1] && p6HighCardArr[0] > p5HighCardArr[0]) {

		less5()
	}

}

func less6() { /// Strip func
	p6HighCardArr[0] = 0
	p6HighCardArr[1] = 0
	p6HighCardArr[2] = 0
	p6HighCardArr[3] = 0
	p6HighCardArr[4] = 0
	p6HighPair = 0
}

func less5() {
	p5HighCardArr[0] = 0
	p5HighCardArr[1] = 0
	p5HighCardArr[2] = 0
	p5HighCardArr[3] = 0
	p5HighCardArr[4] = 0
	p5HighPair = 0
}

func less4() {
	p4HighCardArr[0] = 0
	p4HighCardArr[1] = 0
	p4HighCardArr[2] = 0
	p4HighCardArr[3] = 0
	p4HighCardArr[4] = 0
	p4HighPair = 0
}

func less3() {
	p3HighCardArr[0] = 0
	p3HighCardArr[1] = 0
	p3HighCardArr[2] = 0
	p3HighCardArr[3] = 0
	p3HighCardArr[4] = 0
	p3HighPair = 0
}

func less2() {
	p2HighCardArr[0] = 0
	p2HighCardArr[1] = 0
	p2HighCardArr[2] = 0
	p2HighCardArr[3] = 0
	p2HighCardArr[4] = 0
	p2HighPair = 0
}

func less1() {
	p1HighCardArr[0] = 0
	p1HighCardArr[1] = 0
	p1HighCardArr[2] = 0
	p1HighCardArr[3] = 0
	p1HighCardArr[4] = 0
	p1HighPair = 0
}
