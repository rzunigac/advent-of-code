package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
)

func checkEmptyspace(image *[][]rune) ([]bool, []bool) {

    rowIsDot := make([]bool, len(*image))
    colIsDot := make([]bool, len((*image)[0]))
    for i := range rowIsDot {
        rowIsDot[i] = true
    }
    for i := range colIsDot {
        colIsDot[i] = true
    }
    for i:=0; i < len(*image); i++ {
        for j:=0; j < len((*image)[i]); j++ {
            if (*image)[i][j] != '.' {
                rowIsDot[i] = false
                colIsDot[j] = false
            }
        }
    }

    return rowIsDot, colIsDot
}

func createEmptyRow(length int) []rune {
    result := make([]rune, length)
    for i:=0; i < length; i++ {
        result[i] = '.'
    }
    return result
}

func expandImage(image *[][]rune, rowEmptys *[]bool, colEmptys *[]bool) [][]rune {
    result := make([][]rune, 0)

    for i := 0; i < len(*rowEmptys); i++ {
        tmp_row := make([]rune, 0)
        for j := 0; j < len(*colEmptys); j++ {
            tmp_row = append(tmp_row, (*image)[i][j])
            if (*colEmptys)[j] {
                tmp_row = append(tmp_row, '.')
            }
        }
        if (*rowEmptys)[i] {
            result = append(result, createEmptyRow(len(tmp_row)))
        }
        result = append(result, tmp_row)
    }
    return result
}

func detectGalaxy(image *[][]rune) [][]int {
    galaxyIndex := make([][]int, 0)
    for i:=0; i < len(*image); i++ {
        for j:=0; j < len((*image)[i]); j++ {
            if (*image)[i][j] == '#' {
                galaxyIndex = append(galaxyIndex, []int{i, j})
            }
        }
    }
    return galaxyIndex
}

func distance(x [] int, y []int) int {
    return int(math.Abs(float64(x[0] - y[0])) + math.Abs(float64(x[1] - y[1])))
}

func crossesExpansion(a []int, b []int, rowEmptys *[]bool, colEmptys *[]bool) int {
    horizontalExpansions := 0
    verticalExpansions := 0

    rAxis := []int{a[0], b[0]}
    cAxis := []int{a[1], b[1]}

    sort.Ints(rAxis)
    sort.Ints(cAxis)

    for i:= rAxis[0]; i < rAxis[1]; i++ {
        if (*rowEmptys)[i] {
            horizontalExpansions++
        }
    }

    for i:= cAxis[0]; i < cAxis[1]; i++ {
        if (*colEmptys)[i] {
            verticalExpansions++
        }
    }

    return horizontalExpansions + verticalExpansions
}

func distances(galaxyIndex [][]int, expansionFactor int, rowEmptys *[]bool, colEmptys *[]bool) []int {
    // calculate the distance between all the points in galaxyIndex
    distances := make([]int, 0)
    for i:=0; i < len(galaxyIndex); i++ {
        for j:=i+1; j < len(galaxyIndex); j++ {
            // we need to convert the int values to float before given to the abs function
            distance := distance(galaxyIndex[i], galaxyIndex[j]) + crossesExpansion(galaxyIndex[i], galaxyIndex[j], rowEmptys, colEmptys)*(expansionFactor-1)
            distances = append(distances, int(distance))
        }        
    }
    return distances
}

func printImage(image *[][]rune) {
    for i:=0; i < len(*image); i++ {
        for j:=0; j < len((*image)[i]); j++ {
            fmt.Print(string((*image)[i][j]))
        }
        fmt.Println()
    }
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

    image := make([][]rune, 0)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        image = append(image, []rune(line))
    }

    rowEmptys, colEmptys := checkEmptyspace(&image)

    // For part one we do the expanded image function, but we replaced it for the part 2 for the expansion part on the distances function, now we dont expand
    // expandedImage := expandImage(&image, &rowEmptys, &colEmptys)
    expandedImage := image

    galaxyIndex := detectGalaxy(&expandedImage)

    galaxyDistances := distances(galaxyIndex, 2, &rowEmptys, &colEmptys)

    total := 0
    for dist := range galaxyDistances {
        total += galaxyDistances[dist]
    }

    fmt.Println("Part One")
    fmt.Println("Sum of distances: ", total )

    galaxyDistances2 := distances(galaxyIndex, 1000000, &rowEmptys, &colEmptys)

    total = 0
    for dist := range galaxyDistances2 {
        total += galaxyDistances2[dist]
    }

    fmt.Println("Part Two")
    fmt.Println("Sum of distances: ", total)
    

}

