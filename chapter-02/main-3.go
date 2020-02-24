package main

import ( "fmt" )

type Saiyan struct {
        Name string
        Power int
}


func (s *Saiyan) Super() {
        s.Power += 10000
}

func main() {
        goku := &Saiyan{Name: "Goku"}
        goku.Power = 9000
        goku.Super()
        fmt.Println(goku.Power)
        
}
