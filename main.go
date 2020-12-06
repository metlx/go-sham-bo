package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func game(){
	rand.Seed(time.Now().UnixNano())
	randomAf := rand.Intn(3)
	scanner := bufio.NewScanner(os.Stdin)
	print("make a move!!!  (rock, paper, scissors)\n")
	scanner.Scan()
	move := scanner.Text()
	compMove := ""
	switch {
	case randomAf == 0:
		compMove = "rock"
	case randomAf == 1:
		compMove = "paper"
	case randomAf == 2:
		compMove = "scissors"
	}
	if move == compMove {
		fmt.Printf("computer chose %s \n", compMove)
		fmt.Println("tie")
	} else if move == "rock" && compMove == "paper" {
		fmt.Printf("computer chose %s \n", compMove)
		fmt.Println("you loose") // rock vs paper
	} else if move == "rock" && compMove == "scissors" {
		fmt.Printf("computer chose %s \n", compMove)
		fmt.Println("you win") // rock vs scissors
	} else if move == "paper" && compMove == "rock" {
		fmt.Printf("computer chose %s \n", compMove)
		fmt.Println("you win") // paper vs rock
	} else if move == "paper" && compMove == "scissors" {
		fmt.Printf("computer chose %s \n", compMove)
		fmt.Println("you loose") // paper vs scissors
	} else if move == "scissors" && compMove == "rock" {
		fmt.Printf("computer chose %s \n", compMove)
		fmt.Println("you loose") // scissors vs rock
	} else if move == "scissors" && compMove == "paper" {
		fmt.Printf("computer chose %s \n", compMove)
		fmt.Println("you win") // scissors vs paper
	}
}

func main(){
	game()
}
