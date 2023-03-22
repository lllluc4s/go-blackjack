package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Define as cartas disponíveis
	cards := []string{
		"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A",
	}

	// Define o valor de cada carta
	cardValues := map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
		"J": 10, "Q": 10, "K": 10, "A": 11,
	}

	// Inicializa o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Inicializa o jogo
	var playerCards []string
	var dealerCards []string
	var playerScore int
	var dealerScore int

	// Dá duas cartas para o jogador
	playerCards = append(playerCards, drawCard(cards))
	playerCards = append(playerCards, drawCard(cards))
	playerScore = calculateScore(playerCards, cardValues)

	// Dá duas cartas para o dealer
	dealerCards = append(dealerCards, drawCard(cards))
	dealerCards = append(dealerCards, drawCard(cards))
	dealerScore = calculateScore(dealerCards, cardValues)

	fmt.Println("Suas cartas:", playerCards)
	fmt.Println("Sua pontuação:", playerScore)

	// Verifica se o jogador tem blackjack
	if playerScore == 21 {
		fmt.Println("Blackjack! Você venceu!")
		return
	}

	// Verifica se o dealer tem blackjack
	if dealerScore == 21 {
		fmt.Println("Dealer tem blackjack! Você perdeu.")
		return
	}

	// Pede ao jogador para bater ou ficar
	for playerScore < 21 {
		var input string
		fmt.Print("Bater (hit) ou ficar (stand)? ")
		fmt.Scanln(&input)

		if input == "hit" {
			card := drawCard(cards)
			playerCards = append(playerCards, card)
			playerScore = calculateScore(playerCards, cardValues)

			fmt.Println("Suas cartas:", playerCards)
			fmt.Println("Sua pontuação:", playerScore)

			if playerScore > 21 {
				fmt.Println("Você estourou! Você perdeu.")
				return
			}
		} else if input == "stand" {
			break
		}
	}

	// Dealer bate até que tenha pelo menos 17
	for dealerScore < 17 {
		card := drawCard(cards)
		dealerCards = append(dealerCards, card)
		dealerScore = calculateScore(dealerCards, cardValues)

		fmt.Println("Carta do dealer:", card)
		fmt.Println("Pontuação do dealer:", dealerScore)

		if dealerScore > 21 {
			fmt.Println("Dealer estourou! Você venceu!")
			return
		}
	}

	// Verifica quem tem a pontuação mais alta
	if playerScore > dealerScore {
		fmt.Println("Você venceu!")
	} else if dealerScore > playerScore {
		fmt.Println("Você perdeu.")
	} else {
		fmt.Println("Empate.")
	}
}

func drawCard(cards []string) string {
	// Seleciona aleatoriamente uma carta do baralho
	cardIndex := rand.Intn(len(cards))
	card := cards[cardIndex]

	// Remove a carta do baralho
	cards = append(cards[:cardIndex], cards[cardIndex+1:]...)

	return card
}

func calculateScore(cards []string, cardValues map[string]int) int {
	score := 0
	numAces := 0

	// Calcula a pontuação das cartas
	for _, card := range cards {
		value := cardValues[card]
		score += value

		// Conta o número de Ases
		if card == "A" {
			numAces++
		}
	}

	// Converte os Ases de 11 para 1, se necessário
	for numAces > 0 && score > 21 {
		score -= 10
		numAces--
	}

	return score
}
