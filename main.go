package main

import (
	"fmt"
	"math/rand"
)


func main() {
    choices := []string{"rock", "paper", "scissors"}
	var user string
	fmt.Print("Enter (rock, paper, scissors): ")
	fmt.Scanln(&user)
	computerChoice := choices[rand.Intn(len(choices))]
	
	fmt.Printf("\nYour Choice: %v", user)
	fmt.Printf("\nComputer's Choice: %v", computerChoice +"\n")

	switch user { 

	case "rock":
		switch computerChoice {
		case "rock":
			fmt.Println("Tie!")
			break	
		case "paper":
			fmt.Println("You Lost!")
			break
		case "scissors":
			fmt.Println("You Won!")
			break
		default:
			fmt.Println("There was an error..")
		}

	case "paper":
		switch computerChoice {
		case "rock":
			fmt.Println("You won!")
			break
		case "paper":
			fmt.Println("Tie!")
			break
		case "scissors":
			fmt.Println("You Lost!")
			break
		
		default:
			fmt.Println("There was an error..")
		}
	case "scissors":
		switch computerChoice {
			case "rock":
				fmt.Println("You Lost!")
				break
			case "paper":
				fmt.Println("You Won!")
				break
			case "scissors":
				fmt.Println("Tie!")
				break
			default:
				fmt.Println("There was an error!")
		}
	}
}	
