package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var playerName string
	fmt.Print("Enter Player Name : ")
	fmt.Scanf("%s", &playerName)
	replayFlag := true
	for replayFlag {
		replayStr := ""
		playHangman()
		fmt.Print("\n\n Do you want to replay ? (Y/N) : ")
		fmt.Scanf("\n %s", &replayStr)
		if !strings.EqualFold(replayStr, "Y") {
			replayFlag = false
		}
	}
}

func playHangman() {
	fmt.Println()
	fmt.Println("******************************* Find the Word *******************************************")
	listOfWords := []string{"abruptly", "absurd", "abyss", "affix", "avenue", "awkward", "axiom", "azure", "bagpipes", "bandwagon", "banjo", "bayou", "beekeeper", "blitz", "blizzard", "boggle", "bookworm", "boxcar", "boxful", "buckaroo", "buffalo", "buffoon", "buzzard", "buzzing", "buzzwords", "caliph", "cobweb", "cockiness", "croquet", "crypt", "curacao", "cycle", "daiquiri", "dirndl", "disavow", "dizzying", "duplex", "dwarves", "embezzle", "equip", "espionage", "exodus", "faking", "fishhook", "fixable", "fjord", "flapjack", "flopping", "fluffiness", "flyby", "foxglove", "frazzled", "frizzled", "fuchsia", "funny", "gabby", "galaxy", "galvanize", "gazebo", "giaour", "gizmo", "glowworm", "glyph", "gnarly", "gnostic", "gossip", "grogginess", "haiku", "haphazard", "hyphen", "iatrogenic", "icebox", "injury", "ivory", "ivy", "jackpot", "jaundice", "jawbreaker", "jaywalk", "jazziest", "jazzy", "jelly", "jigsaw", "jinx", "jiujitsu", "jockey", "jogging", "joking", "jovial", "joyful", "juicy", "jukebox", "jumbo", "kayak", "kazoo", "keyhole", "khaki", "kilobyte", "kiosk", "kitsch", "kiwifruit", "klutz", "knapsack", "larynx", "lengths", "lucky", "luxury", "lymph", "marquis", "matrix", "megahertz", "microwave", "mnemonic", "mystify", "naphtha", "nightclub", "nowadays", "numbskull", "nymph", "onyx", "ovary", "oxidize", "oxygen", "pajama", "peekaboo", "phlegm", "pixel", "pizazz", "pneumonia", "polka", "pshaw", "psyche", "puppy", "puzzling", "quartz", "queue", "quips", "quixotic", "quiz", "quizzes", "quorum", "razzmatazz", "rhubarb", "rhythm", "rickshaw", "schnapps", "scratch", "shiv", "snazzy", "sphinx", "spritz", "squawk", "staff", "strength", "strengths", "stretch", "stronghold", "stymied", "subway", "swivel", "syndrome", "thriftless", "thumbscrew", "topaz", "transcript", "transgress", "transplant", "triphthong", "twelfth", "twelfths", "unknown", "unworthy", "unzip", "uptown", "vaporize", "vixen", "vodka", "voodoo", "vortex", "voyeurism", "walkway", "waltz", "wave", "wavy", "waxy", "wellspring", "wheezy", "whiskey", "whizzing", "whomever", "wimpy", "witchcraft", "wizard", "woozy", "wristwatch", "wyvern", "xylophone", "yachtsman", "yippee", "yoked", "youthful", "yummy", "zephyr", "zigzag", "zigzagging", "zilch", "zipper", "zodiac", "zombie"}
	randomChosenWord := getRandomWord(listOfWords)
	listForEnteringLetters := make([]string, 0)
	for i := 0; i < len(randomChosenWord); i++ {
		listForEnteringLetters = append(listForEnteringLetters, "_")
	}
	guessedLetters := []string{}
	noOfTries := 0
	for noOfTries < 4 {
		var letterOrWordGuessed string
		fmt.Printf("\n%v \n\n", listForEnteringLetters)
		fmt.Print("Guess the letter or full word correctly : ")
		fmt.Scanf("\n %s", &letterOrWordGuessed)
		if len(letterOrWordGuessed) == 1 {
			if strings.Contains(randomChosenWord, letterOrWordGuessed) {
				guessedLetters = append(guessedLetters, letterOrWordGuessed)
				replacePlaceHolders(listForEnteringLetters, letterOrWordGuessed, randomChosenWord)
				if !strings.Contains(randomChosenWord, "_") && strings.Join(listForEnteringLetters, "") == randomChosenWord {
					fmt.Printf("**** You Won, U have guessed the word correct, it is %s  ****", randomChosenWord)
					break
				}
			} else if !strings.Contains(randomChosenWord, "_") && strings.Join(listForEnteringLetters, "") == randomChosenWord {
				fmt.Printf("\n**** You Won, U have guessed the word correct, it is %s  ****", randomChosenWord)
				break
			} else {
				noOfTries += 1
			}
		} else {
			if strings.EqualFold(randomChosenWord, letterOrWordGuessed) {
				fmt.Printf("\n**** You Won, U have guessed the word correct, it is %s  ****", randomChosenWord)
			} else {
				fmt.Printf("\n**** Game Over, %s is correct answer ****", randomChosenWord)
			}
			break
		}
		if noOfTries == 4 {
			fmt.Printf("\n**** Game Over, %s is correct answer ****", randomChosenWord)
			break
		}
		fmt.Printf("No of tries left are %v \n", 3-noOfTries)
	}
}

func replacePlaceHolders(listForEnteringLetters []string, letterOrWordGuessed, randomChosenWord string) {
	tempWord := randomChosenWord
	for {
		indActual := strings.Index(randomChosenWord, letterOrWordGuessed)
		indTemp := strings.Index(tempWord, letterOrWordGuessed)
		if indTemp == -1 {
			break
		}
		tempWord = tempWord[indTemp+1:]
		if listForEnteringLetters[indActual] != letterOrWordGuessed {
			listForEnteringLetters[indActual] = letterOrWordGuessed
		} else {
			listForEnteringLetters[len(randomChosenWord)-len(tempWord)-1] = letterOrWordGuessed
		}
	}
}

func getRandomWord(listOfWords []string) string {
	rand.Seed(time.Now().Unix())
	return listOfWords[rand.Intn(len(listOfWords))]
}
