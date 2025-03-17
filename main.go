package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

func (g *GameState) ProcessCSV() {
	f, err := os.Open("quiz-go.csv")
	if err != nil {
		panic("Erro ao abrir o arquivo CSV")
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler o CSV")
	}

	
	for index, record := range records {
		if index > 0 { 
			correctAnswer, err := toInt(record[5])
			if err != nil {
				fmt.Println("Erro ao converter resposta correta:", err)
				continue
			}

			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return 0, errors.New("Entrada inválida. Digite um número válido")
	}
	return i, nil
}

func (g *GameState) Init() {
	fmt.Println("Seja bem-vindo(a) ao quiz!")
	fmt.Print("Escreva o seu nome: ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		panic("Erro ao ler a entrada do usuário")
	}

	g.Name = strings.TrimSpace(name)
	fmt.Printf("Vamos ao jogo, %s!\n\n", g.Name)
}

func (g *GameState) Run() {
	for index, question := range g.Questions {
		fmt.Printf("\033[33m%d. %s\033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Print("Digite uma alternativa: ")

		var answer int
		var err error

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')
			answer, err = toInt(read)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}

		
		if answer == question.Answer {
			fmt.Println("\033[32mResposta correta!\033[0m\n")
			g.Points++
		} else {
			fmt.Println("\033[31mResposta errada.\033[0m\n")
		}
	}

	fmt.Printf("\nFim do jogo! Pontuação final: %d\n", g.Points)
}

func main() {
	game := &GameState{}

	game.ProcessCSV()
	game.Init()
	game.Run()
}
