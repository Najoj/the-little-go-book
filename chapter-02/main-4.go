package main

import ( "fmt" )

type Saiyan struct {
        Name string
        Power int
}


func NewSaiyan(name string, power int) *Saiyan {
        return &Saiyan {
                Name: name,
                Power: power,
        }
}

func main() {
        goku := NewSaiyan("Hello", 20)
        fmt.Println(goku.Power)
}
