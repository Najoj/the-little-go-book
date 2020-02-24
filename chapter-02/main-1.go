package main

type Saiyan struct {
        Name string
        Power int
}

func main() {
        goku := Saiyan{Name: "Goku"}
        goku.Power = 2
        
        /* Does not work
        hello := Test{A: "tSt",
                      B: 13,
              }
        hello.B = 20
        */

        go2 := Saiyan{"Go2", 1337}
        go2.Power = 2
}
