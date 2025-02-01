package main // kind of a required first line in go for now

import "fmt"

func main() {
	fmt.Println("Welcome All To My Quiz Game!")
	var name string

	fmt.Println("What might your name be?")
	fmt.Scan(&name)
	fmt.Printf("Ah yes, %v, how lovely!\n", name)

	fmt.Printf("Please enter your age: ") // asks user for age
	var age uint
	fmt.Scan(&age)
	fmt.Println(age != 18) // checks users age

	if age >= 18 { // a conditional block checks if user is old enough
		fmt.Println("Yay! you're old enough to play my game!")
	} else {
		fmt.Println("Ooo, so sorry. Seems you're too young to play.")
		return
	}

	fmt.Println("Whats better?, RTX 3090 or RTX 4090 ")
	var answer string
	var answer2 string
	score := 0
	total_points := 10

	fmt.Scan(&answer, &answer2)
	fmt.Println(answer, answer2)

	if answer+" "+answer2 == "RTX 4090" {
		fmt.Println("That is correct!")
		score += 5
		fmt.Println("You now have", score, "points. Congratulations!", name)
	} else if answer+" "+answer2 == "rtx 4090" {
		fmt.Println("That is correct!")
		score += 5
		fmt.Println("You now have", score, "points. Congratulations", name)
	} else {
		fmt.Println("Incorrect!")
		score += 0
		fmt.Println("Oh no, unfortunate. You did not gain any points this time.")
		fmt.Println("Current amount of points:", score)
	}

	fmt.Printf("How many cores does a Ryzen 7 9700 have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!, wow you're smart!")
		score += 5
		fmt.Println("You're on a roll! You now have", score, "points.")
	} else {
		fmt.Println("Incorrect!")
		score += 0
		fmt.Println("Oh jeez.. You gained no points this time.")
	}

	fmt.Printf("You scored %v out of %v.\n", score, total_points)
}
