package main

import (
	"fmt"
	"os"
)

func main() {
	const (
		emojiZombie      = "👻"
		emojiEnforcement = "👮"
		emojiJudgement   = "💀"
	)

	fmt.Printf(" Gothonic: The Gothic Go Linter %s\n", emojiZombie)
	fmt.Println("Interrogating your Go code for precise 'gothonic' idiomatic Go style...")

	// Placeholdet for business logic...
	fmt.Printf("Detected: %s shadowed variable 'err' in function main()\n", emojiEnforcement)
	fmt.Printf("Judgement: %s Not gothonic enough.\n", emojiJudgement)

	os.Exit(1)
}
