package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func processSeedList(line string, slice *[]int) error {
    stringNumbers := strings.TrimSpace(strings.Split(line, ":")[1])
    
    for _, s := range strings.Split(stringNumbers, " ") {
        n, err := strconv.Atoi(s)

        if err != nil {
            return err
        }

        *slice = append(*slice, n)
    }

    return nil
}

func addLineToMap(line string, m *map[int]int) error {

    parts := strings.Split(line, " ")

    mvalue, err2 := strconv.Atoi(parts[0])
    mkey, err1 := strconv.Atoi(parts[1])
    mlen, err3 := strconv.Atoi(parts[2])

    if err1 != nil || err2 != nil || err3 != nil {
        return fmt.Errorf("could not convert string to int")
    }

    if len(parts) != 3 {
        return fmt.Errorf("input string does not contain two elements")
    }

    for i := 0; i < mlen; i++ {
        (*m)[mkey+i] = mvalue+i
    } 

    return nil
}

func addLineToMap2(line string, m *map[int][]int) error {
    
    parts := strings.Split(line, " ")

    mvalue, err2 := strconv.Atoi(parts[0])
    mkey, err1 := strconv.Atoi(parts[1])
    mlen, err3 := strconv.Atoi(parts[2])

    if err1 != nil || err2 != nil || err3 != nil {
        return fmt.Errorf("could not convert string to int")
    }

    if len(parts) != 3 {
        return fmt.Errorf("input string does not contain two elements")
    }

    (*m)[mkey] = append((*m)[mkey], []int{mvalue, mlen}...)

    return nil
}

func getValueOrKey(key int, convert *map[int]int) int {
    value, exists := (*convert)[key]
    if exists {
        return value
    }
    return key
}

func getValueOrKey2(number int, myMap *map[int][]int) int {
    //fmt.Println("-----------------")
    closestKey := math.MinInt64
    found := false
    //fmt.Println("Number: ", number)

    for key := range (*myMap) {
        if key <= number && key > closestKey {
            closestKey = key
            found = true
        }
    }

    //fmt.Println("Closestkey: ", closestKey)
    //fmt.Println("found: ", found)

    if found {
        upper := closestKey + (*myMap)[closestKey][1]
        //fmt.Println("Upper: ",upper)
        if number < upper {
            // return the difference between the number and the closest key
            //fmt.Println((*myMap)[closestKey][0])
            //fmt.Println((*myMap)[closestKey][1])
            a := number - closestKey + (*myMap)[closestKey][0]
            //fmt.Println("a", a)
            return a
        }            
    }

    return number
}

func seedToLocation(seed int, 
    seed_to_soil *map[int][]int, 
    soil_to_fertilizer *map[int][]int, 
    fertilizer_to_water *map[int][]int, 
    water_to_light *map[int][]int, 
    light_to_temperature *map[int][]int, 
    temperature_to_humidity *map[int][]int, 
    humidity_to_location *map[int][]int) int {

    soil := getValueOrKey2(seed, seed_to_soil)
    fertilizer := getValueOrKey2(soil, soil_to_fertilizer)
    water := getValueOrKey2(fertilizer, fertilizer_to_water)
    light := getValueOrKey2(water, water_to_light)
    temperature := getValueOrKey2(light, light_to_temperature)
    humidity := getValueOrKey2(temperature, temperature_to_humidity)
    location := getValueOrKey2(humidity, humidity_to_location)

    return location
    
}

func main () {
    
    seedList := make([]int, 0)
    //seed_to_soil := make(map[int]int)
    //soil_to_fertilizer := make(map[int]int)
    //fertilizer_to_water := make(map[int]int)
    //water_to_light := make(map[int]int)
    //light_to_temperature := make(map[int]int)
    //temperature_to_humidity := make(map[int]int)
    //humidity_to_location := make(map[int]int)

    simple_seed_to_soil := make(map[int][]int)
    simple_soil_to_fertilizer := make(map[int][]int)
    simple_fertilizer_to_water := make(map[int][]int)
    simple_water_to_light := make(map[int][]int)
    simple_light_to_temperature := make(map[int][]int)
    simple_temperature_to_humidity := make(map[int][]int)
    simple_humidity_to_location := make(map[int][]int)

    f, err := os.Open("input")
    if err != nil {
        panic(err)
    }

    defer f.Close()

    firstLineRead := false
    nextIsMapLine := false
    inMap := false
    currentMap := ""

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {

        var line = scanner.Text()
        //fmt.Println(line)

        if !firstLineRead {
            firstLineRead = true
            // process seed list
            processSeedList(line, &seedList)
            //fmt.Println("First Line")
            //fmt.Println(seedList)
            continue
        }

        if strings.TrimSpace(line) == "" {
            nextIsMapLine = true
            inMap = false
            //fmt.Println("Empty line found")
        }

        if nextIsMapLine && strings.TrimSpace(line) != ""{
            nextIsMapLine = false
            inMap = true
            // process map
            currentMap = strings.Split(line, " ")[0]
            fmt.Println("Map Line", currentMap)
            continue
        }

        if(inMap){
            if currentMap == "seed-to-soil" {
                _ = addLineToMap2(line, &simple_seed_to_soil)
            } else if currentMap == "soil-to-fertilizer" {
                _ = addLineToMap2(line, &simple_soil_to_fertilizer)
            } else if currentMap == "fertilizer-to-water" {
                _ = addLineToMap2(line, &simple_fertilizer_to_water)
            } else if currentMap == "water-to-light" {
                _ = addLineToMap2(line, &simple_water_to_light)
            } else if currentMap == "light-to-temperature" {
                _ = addLineToMap2(line, &simple_light_to_temperature)
            } else if currentMap == "temperature-to-humidity" {
                _ = addLineToMap2(line, &simple_temperature_to_humidity)
            } else if currentMap == "humidity-to-location" {
                _ = addLineToMap2(line, &simple_humidity_to_location)
            }

        }
    }
    
    //fmt.Println(seed_to_soil)
    //fmt.Println(simple_seed_to_soil)
    //fmt.Println(soil_to_fertilizer)
    //fmt.Println(fertilizer_to_water)
    //fmt.Println(water_to_light)
    //fmt.Println(light_to_temperature)
    //fmt.Println(temperature_to_humidity)
    //fmt.Println(humidity_to_location)

    minLocation := math.MaxInt
    for _, seed := range seedList {
        location := seedToLocation(seed, 
            &simple_seed_to_soil, 
            &simple_soil_to_fertilizer, 
            &simple_fertilizer_to_water, 
            &simple_water_to_light, 
            &simple_light_to_temperature, 
            &simple_temperature_to_humidity, 
            &simple_humidity_to_location)
        if location < minLocation {
            minLocation = location
        }
        //fmt.Println(seed, location)
    }

    fmt.Println("Advent of Code 2023 - Day 5")
    fmt.Println("Part 1")
    fmt.Println("Closest Location: ", minLocation)

    //fmt.Println("Part 2")
    //fmt.Println("Ans1: ")


}



