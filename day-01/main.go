package main

import (
	"bufio"
	"fmt"
	"os"
	//"regexp"
	"strconv"
	"strings"
	"unicode"
)

func find_numbers(s string) (rune, rune){
    
    var found_first bool = false
    var found_last bool = false

    var first rune;
    var last rune;

    l := len(s)

    for i := 0; i < l; i++ {
        head :=[]rune(string(s[i]))[0]
        if unicode.IsDigit(head) && found_first == false {
            first = head
            found_first = true
            //fmt.Println("first ", first)
        }

        tail := []rune(string(s[l-i-1]))[0]
        if unicode.IsDigit(tail) && found_last == false {
            last = tail
            found_last = true
            //fmt.Println("last ", last)
        }

        if found_first && found_last {
            break
        }
    }
    return first, last
}


func part1() {
    f, err := os.Open("test_input")
    if err != nil {
        panic(err)
    }

    var total int = 0;

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        var s = scanner.Text()
        f, l := find_numbers(s)
        //fmt.Printf("First digit: %c, Last digit: %c\n", f, l)
        concatenatedString := string(f) + string(l)
        concatenatedInt, _ := strconv.Atoi(concatenatedString)
        //fmt.Println(concatenatedInt)
        total += concatenatedInt
    }
    fmt.Println(total)  
}

// This convert the words to numbers, but just the first word, not all, this is my lazy take on the cases where the last letter of a number is the first letter of the other.
// so xxoneeight123 will be converted to xx1ne8ight123
// this is unelegant but sifficient for these specific scenario
func convert_text_to_int(s string) string {

    numberWords := map[string]rune{
        "one": '1', "two":   '2', "three": '3',
        "four":  '4', "five": '5', "six":   '6', "seven": '7',
        "eight": '8', "nine": '9',
    }

    result := []rune{}
    i := 0

    for i < len(s) {
        matched := false
        for word, digit := range numberWords {
            if strings.HasPrefix(s[i:], string(word)) {
                result = append(result, digit)
                matched = true
                break
            }
        }
        if !matched {
            result = append(result, rune(s[i]))
        }
        i++
    }

    return string(result)
}

func part2() {
     f, err := os.Open("test_input2")
    if err != nil {
        panic(err)
    }

    var total int = 0;

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        var s = scanner.Text()
        s = convert_text_to_int(s)
        //fmt.Println(s)
        f, l := find_numbers(s)
        //fmt.Printf("First digit: %c, Last digit: %c\n", f, l)
        concatenatedString := string(f) + string(l)
        concatenatedInt, _ := strconv.Atoi(concatenatedString)
        //fmt.Println(concatenatedInt)
        total += concatenatedInt
    }
    fmt.Println(total)  
}

func main() {
    fmt.Println("Advent of Code 2023 - Day 1")
    fmt.Println("Part 1")
    part1()

    fmt.Println("Part 2")
    part2()

}
