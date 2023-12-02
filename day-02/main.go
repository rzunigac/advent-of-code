package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_posible_game(s string) (bool, int){

    game_slice := strings.Split(s, ":")

    id, _ := strconv.Atoi(strings.Split(game_slice[0], " ")[1])
    sets :=  strings.Split(game_slice[1], ";")

    is_posible := true

    for _, v := range sets {
        set := convert_set_string_to_map(v)
        if !is_posible_set(set) {
            is_posible = false
        }
    }

    return is_posible, id
}

func convert_set_string_to_map(set string) map[string]int {
    //set := " 1 red, 2 green, 6 blue"
    set_list := strings.Split(set, ",")
    set_map := map[string]int{"red": 0, "green": 0, "blue": 0}

    for _, v := range set_list {
        //fmt.Println(i, v)
        numcol := strings.Split(strings.TrimSpace(v), " ")
        number, _ := strconv.Atoi(numcol[0])
        color := numcol[1]
        set_map[color] = number
        //fmt.Println(number, color)
    }
    //fmt.Println(set_map)
    return set_map
}

func is_posible_set(set map[string]int) bool {

    //only 12 red cubes, 13 green cubes, and 14 blue cubes
    boxes := map[string]int{"red": 12, "green": 13, "blue": 14}
    //set := map[string]int{"red": 20, "green": 8, "blue": 6}
    is_posible := true

    for key, value1 := range boxes {
        value2, ok := set[key]
            if !ok || value1 < value2 {
                is_posible = false
                break
            }
    }
    return is_posible
}

func calculate_minimum_power_for_game(s string) int {
    game_slice := strings.Split(s, ":")

    //id, _ := strconv.Atoi(strings.Split(game_slice[0], " ")[1])
    sets :=  strings.Split(game_slice[1], ";")
    minimum := map[string]int{"red": 0, "green": 0, "blue": 0}

    for _, v := range sets {
        set := convert_set_string_to_map(v)

        for key, seen := range set {
            needed, _ := minimum[key]
                if needed < seen {
                    minimum[key] = seen
                }
        }

    }
    
    power := 1
    for _, v := range minimum {
        power = power*v
    }

    return power
   
}


func main () {
    f, err := os.Open("test_input")
    if err != nil {
        panic(err)
    }

    sum_posible_games_id := 0
    sum_of_powers := 0

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        var s = scanner.Text()
        posible, id := is_posible_game(s)
        if posible {
            sum_posible_games_id += id
        }

        power := calculate_minimum_power_for_game(s)
        sum_of_powers += power
    }

    fmt.Println("Advent of Code 2023 - Day 2")
    fmt.Println("Part 1")
    fmt.Println("Sum of posible games IDs: ", sum_posible_games_id)

    fmt.Println("Part 2")
    fmt.Println("Sum of min powers: ", sum_of_powers)
}



