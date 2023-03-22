README - Blackjack Game
Este é um código simples em Go que implementa o jogo de Blackjack. O jogo é jogado pelo usuário contra o dealer, que é representado pelo computador.


Requisitos
Go 1.16 ou superior


Como jogar
Execute o arquivo main.go usando o comando "go run main.go" no terminal.
Você receberá duas cartas, assim como o dealer. A pontuação de suas cartas é exibida.
Escolha entre "hit" ou "stand". Se escolher "hit", você receberá uma nova carta e sua pontuação será atualizada. Se escolher "stand", você manterá sua pontuação atual.
O dealer irá bater até que sua pontuação seja igual ou maior que 17.
Se sua pontuação for maior que a do dealer e menor ou igual a 21, você ganha. Se a pontuação do dealer for maior que a sua e menor ou igual a 21, você perde. Se ambas as pontuações forem iguais, o jogo é um empate.


Entendendo o código
O código é dividido em três funções principais:

- drawCard(cards []string) string: seleciona aleatoriamente uma carta do baralho e a remove do mesmo.

- calculateScore(cards []string, cardValues map[string]int) int: calcula a pontuação das cartas de acordo com seus valores e retorna o resultado.

- main(): gerencia todo o fluxo do jogo.

O jogo usa duas variáveis para acompanhar as cartas do jogador e do dealer, bem como as pontuações de cada um. Os valores das cartas são definidos no mapa cardValues e as cartas disponíveis são armazenadas em uma fatia cards.

As cartas são selecionadas aleatoriamente usando o gerador de números aleatórios fornecido pela biblioteca "math/rand". O jogo usa um loop para pedir ao jogador para bater ou ficar, e outro loop para o dealer bater até que sua pontuação seja igual ou maior que 17.
