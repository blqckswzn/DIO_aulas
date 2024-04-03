package main // declaração do pacote
import "fmt"// importação de biblioteca nativa
const ebulicaoF float64 = 212.0
func mains () {
	var temperaturaF = ebulicaoF
	var temperaturaC = (temperaturaF - 32) * 5/9
	fmt.Println("resultado é :",temperaturaC)
}