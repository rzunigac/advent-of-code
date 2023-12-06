package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
)

func lineToArray(s string) []int {
    valuesString := strings.Split(s, ":")[1]
    valuesString = strings.Join(strings.Fields(valuesString), " ")

    stringsSlice := strings.Split(valuesString, " ")

    var intSlice []int

    for _, s := range stringsSlice {
        i, _ := strconv.Atoi(s)
        intSlice = append(intSlice, i)
    }

    return intSlice
}

func bruteForceCalulation(times []int, distances []int) int{
    marginToWin := 1
     for i := 0; i < len(times); i++ {
        //fmt.Println("Race: ", i)
        maxTime := times[i]
        minDistance := distances[i]
        waysToWin := 0

        for j := 1; j < maxTime; j++ {
            runDistance := (maxTime - j) * j
            if runDistance > minDistance {
                waysToWin++
            }
        }
        marginToWin = marginToWin * waysToWin
    }   
    return marginToWin
}

func main () {
    fmt.Println("Advent of Code 2023 - Day 4")

    startTime := time.Now()

    f, err := os.Open("test_input")
    if err != nil {
        fmt.Println("Error opening the file:", err)
        return
    }
    scanner := bufio.NewScanner(f)

    scanner.Scan()
    rawLine1 := scanner.Text()
    times := lineToArray(rawLine1)

    scanner.Scan()
    rawLine2 := scanner.Text()
    distances := lineToArray(rawLine2)

    f.Close()

    fmt.Println("Read File execution time:", time.Now().Sub(startTime))
    
    startTime = time.Now()

    marginToWin := bruteForceCalulation(times, distances)  

    fmt.Println("Part 1")
    fmt.Println("Margin to win: ", marginToWin)
    fmt.Println("Part One execution time:", time.Now().Sub(startTime))

    fmt.Println("Part 2")

    startTime = time.Now()

    var timesString []string
    var distancesString []string

    for i := 0; i < len(times); i++ {
        timesString = append(timesString, strconv.Itoa(times[i]))
        distancesString = append(distancesString, strconv.Itoa(distances[i]))
    }
    
    newTime, _ := strconv.Atoi(strings.Join(timesString, ""))
    newDistance, _ := strconv.Atoi(strings.Join(distancesString, ""))
    
    //fmt.Println("New time: ", newTime)
    //fmt.Println("New distance: ", newDistance)

    // we still use bruteforce since the total time is in the order of tens of millions and the calculation should take little time
    // the distance t values are bigger but they are used only for comparison
    marginToWin2 := bruteForceCalulation([]int{newTime}, []int{newDistance}) 

    fmt.Println("Margin to win: ", marginToWin2)
    fmt.Println("Part Two execution time:", time.Now().Sub(startTime))
}



