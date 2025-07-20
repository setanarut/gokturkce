package main

import (
	"fmt"

	"github.com/setanarut/gokturkce"
)

func main() {
	s := "Biz, ilhamlarımızı gökten ve gâipten değil, doğrudan doğruya hayattan almış bulunuyoruz"
	c := gokturkce.TR2GTR(s, true)
	fmt.Println(gokturkce.TersÇevir(c))
}
