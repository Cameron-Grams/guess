package main

import (
	"bufio"
	"fmt"
	"github.com/Cameron-Grams/randnum"
	"log"
	"os"
	"strconv"
	"strings"
)

func makeguess(guess_ int, target int) string {
	if guess_ > target {
		return "Too High"
	} else if guess_ < target {
		return "Too Low"
	}
	return "Success"
}

func main() {
	var user_guess int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter range for target: ")

	target_range, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	target_range = strings.TrimSpace(target_range)
	tgt_range, err := strconv.Atoi(target_range)
	if err != nil {
		log.Fatal(err)
	}

	target, err := randnum.RandomInteger(tgt_range)
	if err != nil {
		log.Fatal(err)
	}

	message := fmt.Sprintf("Target is: %v", target)
	fmt.Println(message)

	for user_guess != target {
		user_pointer := &user_guess

		fmt.Println("Make a guess at the number")
		guess_, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess_ = strings.TrimSpace(guess_)
		user_value, err := strconv.Atoi(guess_)
		if err != nil {
			log.Fatal(err)
		}
		feedback := makeguess(user_value, target)
		fmt.Println(feedback)

		*user_pointer = user_value
	}

	fmt.Println("Congratulations!")
}
