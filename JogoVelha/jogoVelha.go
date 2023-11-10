package main

import "fmt"

//Criando variaveis e estruturas de dados para o tabuleiro e jogadores
type jogador struct {
	nome    string
	simbolo string
}
type jogo struct {
	tabuleiro     [3][3]string
	jogadorAtual jogador
	vencedor     jogador
}

//Função para inicialização do jogo
func inicializarJogo() jogo {
	novoJogo := jogo{
		tabuleiro:     [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}},
		jogadorAtual: jogador{"Jogador1", "X"},
		vencedor:     jogador{"", "X"},
	}
	return novoJogo
}

//Função para exibir o tabuleiro
func exibirTabuleiro(jogadorAtual jogo){
	fmt.Println("0 1 2")
	for i, row := range jogadorAtual.tabuleiro{
		fmt.Println(i)
		for _, cell := range row{
			fmt.Printf("%s", cell)
		}
		fmt.Println()
	}
}

//Função para realização da jogada
func realizarJogada(jogoAtual jogo, linha, coluna int)(jogo, bool){
	if linha < 0 || linha > 2 || coluna < 0 || coluna > 2 || jogoAtual.tabuleiro[linha][coluna] != ""  {
		return jogoAtual, false
	}

	jogoAtual.tabuleiro[linha][coluna] =jogoAtual.jogadorAtual.simbolo
	return jogoAtual, true
}


func main() {
	fmt.Println("#### Jogo Da Velha ####")
	jogoAtual := inicializarJogo()
	/* turno := 0 */

	for{
		exibirTabuleiro(jogoAtual)
		fmt.Printf("%s(%s) Faça sua jogada (Linha e Coluna): ", jogoAtual.jogadorAtual.nome, jogoAtual.jogadorAtual.simbolo)
	
		var linha, coluna int
		fmt.Scan(&linha, &coluna)

		jogoAtual, jogadaValida := realizarJogada(jogoAtual, linha, coluna)

		if !jogadaValida{
			fmt.Println("Jogada Inválida tente novamen")
		}
	
	}
}
