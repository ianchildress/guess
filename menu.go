package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type menu struct {
	scanner *bufio.Scanner
	answer  answer
	player  string
}

func newMenu(a answer) menu {
	return menu{answer: a, scanner: bufio.NewScanner(os.Stdin)}
}

func (m menu) run() {
	fmt.Println("Hello, what is your name?")

	m.scanner.Scan()
	if err := m.scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	m.player = m.scanner.Text()
	fmt.Println("Hello", m.player, "let's play a game!")
	fmt.Println("I will think of an animal and you will try to guess what it is.")
	fmt.Printf("The possible answers are: ")
	for i := range m.answer.all {
		fmt.Printf("%s ", m.answer.all[i].name)
	}
	fmt.Printf("\n")
	m.top()
}

func (m menu) top() {
	for {
		fmt.Println("Would what you like to do?")
		fmt.Println("1) Ask a question about the animal")
		fmt.Println("2) Guess the animal")
		fmt.Println("3) Quit the program")
		m.scanner.Scan()
		if err := m.scanner.Err(); err != nil {
			log.Println(err)
			os.Exit(1)
		}

		switch m.scanner.Text() {
		case "1":
			m.question()
			continue
		case "2":
			m.guess()
			continue
		case "3":
			fmt.Printf("Good bye %s, thanks for playing!\n", m.player)
			os.Exit(1)
		default:
			fmt.Println("I'm sorry, that isn't one of your choices. Please try again.")
		}
	}
}

func (m menu) question() {
	for {
		fmt.Println("What would you like to ask?")
		fmt.Println("1) Does the animal have scales?")
		fmt.Println("2) Does the animal have a tail?")
		fmt.Println("3) Does the animal have legs?")
		fmt.Println("4) Does the animal live in the water?")
		m.scanner.Scan()
		if err := m.scanner.Err(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		switch m.scanner.Text() {
		case "1":
			fmt.Println(m.answer.hasScales())
			return
		case "2":
			fmt.Println(m.answer.hasTail())
			return
		case "3":
			fmt.Println(m.answer.hasLegs())
			return
		case "4":
			fmt.Println(m.answer.livesInWater())
			return
		default:
			fmt.Println("I'm sorry, I didn't understand that answer. " +
				"Please enter a number that corresponds with your question.")
		}
	}
}
func (m menu) guess() {
	for {
		fmt.Println("What animal am I thinking of?")
		for i := range m.answer.all {
			fmt.Printf("%v) %s\n", i+1, m.answer.all[i].name)
		}
		m.scanner.Scan()
		if err := m.scanner.Err(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		i, err := strconv.Atoi(m.scanner.Text())
		if err != nil || i > len(m.answer.all) || i < 1 {
			fmt.Println("I'm sorry, I didn't understand that answer. Please enter an answer by typing in a number.")
			continue
		}

		if m.answer.all[i-1].name == m.answer.chosen.name {
			fmt.Printf("That's right %s! I was thinking of %s\n", m.player, m.answer.chosen.name)
			return
		}
		fmt.Printf("Sorry %s, that wasn't the animal I was thinking of.\n", m.player)
		return
	}
}
