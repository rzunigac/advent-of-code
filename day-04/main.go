package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

func checkCard(card string) (int, int, int){
    //fmt.Println("Checking: ", card)
    words := strings.Fields(card)
    cleanCard := strings.Join(words, " ")
    
    f := func(c rune) bool {
        return c == ':' || c == '|'
    }

    fields := strings.FieldsFunc(cleanCard, f)
    id, _ :=  strconv.Atoi(strings.Split(fields[0], " ")[1])
    winningNumbers := strings.Split(strings.Trim(fields[1], " "), " ")
    myNumbers := strings.Split(strings.Trim(fields[2], " "), " ")

    //fmt.Println(id, winningNumbers, myNumbers)
    // we can do a nested loop on each slice to checkCard the numbers taht match, or we can create a map
    myNumberMap := make(map[string]bool)
    for _, v := range myNumbers {
        myNumberMap[v] = true
    }
    
    totalMatches := 0
    for _, number := range winningNumbers {
        if myNumberMap[number] {
            totalMatches++
        }
    }

    cardPoints := math.Pow(2.0, float64(totalMatches-1))
    //fmt.Println("Total matches: ", totalMatches)
    //fmt.Println("Card Points: ", cardPoints)

    // when there are no matches we have Pow(2, 0-1) = 0.5
    // we will abuse the fact that we return an integer so 0.5 turns into 0
    return id, totalMatches, int(cardPoints)

}


func main () {
    f, err := os.Open("test_input")
    if err != nil {
        panic(err)
    }

    totalPoints := 0
    totalCards := 0
    scratchCards := make(map[int]int)

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        var s = scanner.Text()
        cardId, cardMatches, cardPoints := checkCard(s)
        totalPoints += cardPoints

        // Update current card count
        if val, exists := scratchCards[cardId]; exists {
        // Key exists, increment its value
            scratchCards[cardId] = val + 1
        } else {
        // Key does not exist, set it to 1
            scratchCards[cardId] = 1
        }

        // Update next card count
        currentCount := scratchCards[cardId]
        if cardMatches >= 1 {
            for i := 1; i <= cardMatches; i++{
                if val, exists := scratchCards[cardId + i]; exists {
                    scratchCards[cardId + i] = val + currentCount
                } else {
                    scratchCards[cardId + i] = currentCount
                }
            }
        }
    }

    for _, v := range scratchCards {
        totalCards += v
    }

    fmt.Println("Advent of Code 2023 - Day 4")
    fmt.Println("Part 1")
    fmt.Println("Total points: ", totalPoints)

    fmt.Println("Part 2")
    fmt.Println("total cards: ", totalCards)
}



