# Blackjack

Implementação simples do jogo Blackjack em Go.

## Regras

- O jogador tenta vencer o dealer alcançando uma pontuação total maior do que o dealer sem exceder 21.
- Cada jogador e o dealer recebem duas cartas do baralho.
- As cartas têm os seguintes valores:
  - Ases podem valer 1 ou 11 (o programa escolhe automaticamente a melhor opção).
  - As outras cartas valem seu valor facial (2 a 10).
  - J, Q e K valem 10 pontos cada.
- Se o jogador receber um Ás e uma carta com valor 10 em sua mão inicial, ele tem um "blackjack" e vence automaticamente, a menos que o dealer também tenha um "blackjack".
- O jogador decide se quer "bater" (receber outra carta) ou "ficar" (não receber mais cartas).
- Se o jogador ultrapassar 21 pontos, ele perde automaticamente.
- Quando o jogador decide ficar, é a vez do dealer jogar.
- O dealer bate (recebe mais cartas) até que sua pontuação total seja pelo menos 17.
- Se o dealer ultrapassar 21 pontos, o jogador vence automaticamente.
- Se o jogador e o dealer tiverem a mesma pontuação, é considerado um empate.
- O jogador pode jogar novamente.

## Como Jogar

1. Baixe e instale o Go: https://golang.org/
2. Abra o terminal e navegue até o diretório onde está o código.
3. Execute o comando `go run main.go`.
4. Siga as instruções na tela para jogar.

