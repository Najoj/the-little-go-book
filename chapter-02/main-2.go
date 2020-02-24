package main

import ( "fmt" )

type Saiyan struct {
        Name string
        Power int
}

func Super(s Saiyan) {
        s.Power += 10000
}

func SuperS(s *Saiyan) {
        s = &Saiyan{"Gogogo", 365}
        s.Power += 10000
}

func main() {
        goku := Saiyan{Name: "Goku"}
        goku.Power = 2
        Super(goku)
        fmt.Println(goku.Power)
        
        go2 := &Saiyan{Name: "Go2"}
        go2.Power = 132     // This still works
        SuperS(go2)
        fmt.Println(go2.Power)
        
        go5 := &Saiyan{Name: "Go5"}
        SuperS(go5)
        fmt.Println(go5.Power)
}
