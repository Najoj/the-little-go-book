package main

import ( "fmt" )


func main() {
        scores := [4] int {51,12,63,57}
        fmt.Println(scores)

        makescores := make([]int, 2, 12)
        fmt.Println(makescores)
        fmt.Println(len(makescores))
        makescores = append(makescores, 1337)
        fmt.Println(makescores)
        fmt.Println(len(makescores))
        fmt.Println(cap(makescores))

        // Will not work
        // var varscores = [9]int
        // fmt.Println(varscores)
}
