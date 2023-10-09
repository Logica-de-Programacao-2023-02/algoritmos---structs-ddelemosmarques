//Crie uma struct chamada Livro com os campos
//"título" ,"autor" e "ano de publicação".
//Escreva uma função que receba um Livro
//como parâmetro e imprima suas informações.

package main

import "fmt"

type Livro struct {
	título string
	autor  string
	ano    int
}

func imprimeLivro(l Livro) {

	fmt.Println("título: ", l.título)
	fmt.Println("autor: ", l.autor)
	fmt.Println("ano de publicação: ", l.ano)
}

func main() {
	l := Livro{
		título: "48 leis do poder",
		autor:  "Robert Greene",
		ano:    2007,
	}

	imprimeLivro(l)
}
