package main

import (
    "bufio"
    "fmt"
    "os"
    "slices"
    "sort"
    "strconv"
    "strings"
)

func clasiffyHand(hand string, usejoker bool) string{
    // count the number of each character in the hand
    count := make(map[string]int)
    for _, c := range hand {
        count[string(c)]++
    }

    reverseCount := make(map[int][]string)

    values := make([]int, 0, len(count))
    for k, v := range count {
        values = append(values, v)
        reverseCount[v] = append(reverseCount[v], k)
    }

    jCount, jExists := count["J"]

    if usejoker && jExists && jCount < 5 {
        maxNotJ := ""
        founded := false

        for i := 5; i > 0; i-- {
            if founded {
                break
            }
            maxCount, exists := reverseCount[i]
            if exists {
                for _, v := range maxCount {
                    if v != "J" {
                        maxNotJ = v
                        founded = true
                        break
                    }
                }
            }
        }
        replacedString := strings.ReplaceAll(hand, "J", maxNotJ)
        return clasiffyHand(replacedString, false)
    }

    // sort the values
    slices.Sort(values)
    slices.Reverse(values)
    //fmt.Println(values)

    handType := ""

    // classify the hand, this is ugly !!
    for _, v := range values {
        if v == 5 {
            handType = "Five of a Kind"
            break
        }
        if v == 4 {
            handType = "Four of a Kind"
            break
        }
        if v == 3 {
            handType = "Three of a Kind"
            continue
        }
        if v == 2 && handType == "Three of a Kind" {
            handType = "Full House"
            break
        }
        if v == 2 && handType == "One Pair" {
            handType = "Two Pair"
            break
        }
        if v == 2 {
            handType = "One Pair"
            continue
        }
        if v == 1 && handType == "" {
            handType = "High Card"
            break
        } else {
            break
        }
    }

    return handType
}

type hand struct {
    cards string
    bid int
}

var handTypes = []string{
    "High Card",
    "One Pair",
    "Two Pair",
    "Three of a Kind",
    "Full House",
    "Four of a Kind",
    "Five of a Kind",
}

// Part 1
type byCards []hand

func (h byCards) Len() int {
    return len(h)
}

func (h byCards) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h byCards) Less(i, j int) bool {
    // Define a custom comparison logic for the cards field
    return customCompare(h[i].cards, h[j].cards, false)
}

// Part 2, could be a better way?
type byCards2 []hand

func (h byCards2) Len() int {
    return len(h)
}

func (h byCards2) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h byCards2) Less(i, j int) bool {
    // Define a custom comparison logic for the cards field
    return customCompare(h[i].cards, h[j].cards, true)
}

// for part 1 and part 2
func customCompare(cardsA, cardsB string, useJoker bool) bool {

    var letterValues = map[rune]int{'2': 2,'3': 3,'4': 4,'5': 5,'6': 6,'7': 7,'8': 8,'9': 9,'T': 10,'J': 11,'Q': 12,'K': 13,'A': 14}

    if useJoker {
        letterValues['J'] = 1
    }

    runesA := []rune(cardsA)
    runesB := []rune(cardsB)

    i, j := 0, 0
    for i < len(runesA) && j < len(runesB) {
        charA := runesA[i]
        charB := runesB[j]

        valueA, _ := letterValues[charA]
        valueB, _ := letterValues[charB]

        if valueA == valueB {
            i++
            j++
            continue
        } else {
            return valueA < valueB
        }

    }

    return cardsA < cardsB
}

func main() {
    args := os.Args
    if len(args) != 2 {
        fmt.Println("Usage: go run main.go <filename>")
        return
    }

    file, err := os.Open(args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    allHands := make(map[string][]hand)
    allHandsPart2 := make(map[string][]hand)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        lineParts := strings.Split(line, " ")
        handString := lineParts[0]
        bidString, _ := strconv.Atoi(lineParts[1])
        // part 1
        handType := clasiffyHand(handString, false)
        allHands[handType] = append(allHands[handType], hand{cards: handString, bid: bidString})
        // part 2
        handTypePart2 := clasiffyHand(handString, true)
        allHandsPart2[handTypePart2] = append(allHandsPart2[handTypePart2], hand{cards: handString, bid: bidString})
    }

    rank := 1
    winnings := 0

    for _, handType := range handTypes {
        hands, exists := allHands[handType]
        if exists {
            sort.Sort(byCards(hands))
            for _, hand := range hands {
                //fmt.Printf("%s %d %d\n", hand.cards, hand.bid, rank)
                winnings += hand.bid*rank
                rank++
            }
        }
    }
    fmt.Println("Part One")
    fmt.Println("Winnings: ", winnings)

    rank = 1
    winnings = 0

    for _, handType := range handTypes {
        hands, exists := allHandsPart2[handType]
        if exists {
            sort.Sort(byCards2(hands))
            for _, hand := range hands {
                //fmt.Printf("%s %d %d\n", hand.cards, hand.bid, rank)
                winnings += hand.bid*rank
                rank++
            }
        }
    }


    fmt.Println("Part Two")
    fmt.Println("Winnings: ", winnings)

}

