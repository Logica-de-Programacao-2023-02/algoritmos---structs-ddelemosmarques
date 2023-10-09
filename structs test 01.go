//Crie uma struct chamada Retângulo
//com os campos "largura" e "altura".
//Escreva uma função que receba um
//Retângulo como parâmetro e calcule
//a área do retângulo (área = largura * altura).

package main

import "fmt"

type Retangulo struct {
	Largura int
	Altura  int
}

func main() {

	r := Retangulo{Largura: 5, Altura: 10}
	resultado := areaRetangulo(r)
	fmt.Print(resultado)
}
func areaRetangulo(r Retangulo) int {

	return r.Largura * r.Altura
}
