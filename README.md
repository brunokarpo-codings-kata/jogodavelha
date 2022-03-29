# jogodavelha - coding kata

## O que é um Coding Kata
Coding kata são exercícios de programação que ajudam os programadores desenvolverem suas habilidades através da prática e repetição, este termo Coding kata foi usado por Dave Thomas em 1999, ele foi co-autor do livro The Pragmatic Programmer, em uma referência ao conceito japonês de kata nas artes marciais. O conceito foi implementado por Laurent Bossavit e Emmanuel Gaillot, que falaram sobre isso no XP2005 em Sheffield. Após esta conferência, Robert C. Martin descreveu o conceito e os usos iniciais em seu artigo “The Programming Dojo” [Source](http://ninjadolinux.com.br/coding-kata/)

## Jogo da velha
Esse Kata envolve modelagem e codificação. Consiste em implementar uma lógica simples de um jogo da velha.
O programa deve iniciar solicitando o nome de dois jogadores. Ao receber os nomes, o sistema deve apresentar o campo de um jogo da velha onde os jogadores devem, cada um em seu turno, escolher um campo para marcar, no intuíto de enfileirar sua marcação, ou na vertical, horizontal ou diagonal em três campos consecutivos. O jogo pode terminar com um vencedor ou empatado, caso ninguém consiga realizar as três marcações sequenciais no preenchimento do tabuleiro.

## Objetivos
Tente executar o Kata nos seguintes passos:
- Use 10 minutos para pensar na modelagem do sistema. Como o código será divido? Quais serão as entidades que existiram? O que elas fazem?
- Tente implementar o código em 1 hora e meia.
- Não há necessidade de se preocupar com a UI, o foco deve ser na implementação da lógica do jogo.
- Tente escrever testes automatizados

## Regras do repositório
A branch main existe apenas para o arquivo README.md que explica o Kata. Crie uma branch e implemente apenas nessa branch. Não haverá merge. A ideia é que o Kata seja genérico o suficiente para ser implementado em qualquer linguagem.


### Modelagem
O jogo vai ter 3 entidades.
- Board: É a entidade que representa a matriz 3x3 do jogo da velha. Ela conterá as regras de negócio relacionadas ao tabuleiro, a saber: 
  - iniciar os campos com valores zerados
  - permitir alguém marcar um campo do tabuleiro
  - garantir que ninguém marque um campo já marcado
  - verificar se as ordem das marcações configuram um ganhador ou um empate
- Game: É a entidade que guarda o estado do jogo e suas interações. Guarda as instâncias dos jogadores, um tabuleiro relacionado àquele jogo, qual o turno e se o jogo está em execução ou já terminou
- Main: orienta a lógica do jogo para sua execução. Interage com os usuários, recebendo os inputs e apresentando o estado do jogo à medida que ele acontece.