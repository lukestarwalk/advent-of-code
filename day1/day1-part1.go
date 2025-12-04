package day1

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
);
 
func day1Part1(){
	//ler o ficheiro
	content, err := os.ReadFile("input.txt");
	
	//verifica se houve erro na leitura
	if err != nil {
		log.Fatal(err);
	}
	
	//converte os dados em string
	lines := string(content);
	
	//separa a string dos dados em strings de cada movimento
	splitted := strings.Fields(lines);
	
	//variáveis position e password
	position, password := 50, 0;

	for _, dial_move := range splitted {

		//variável direction que pode ser L ou R (Esquerda ou Direita)
		direction := string(dial_move[0]);

		//string de cada direção (ex.: L32) convertida em número (32)
		number, err := strconv.Atoi(dial_move[1:]);

		//verifica se a conversão foi feita sem erros
		if err != nil {
			log.Fatal(err);
		}

		//verifica se a direção é para à Esquerda
		if direction == "L" {

			//decrementa o número da posição actual
			position -= number;

			//como o intervalo é de 0 à 99, se for negativo, incrementa 100
			for position < 0 {
				position += 100;
			}
		
		//se a direção é para a direita (R)...
		} else {

			//incrementa a posição actual
			position += number;

			//como o intervalo é de 0 à 99, se for superior à 100, decrementa 100
			for position >= 100 {
				position -= 100;
			}
			fmt.Println();
		}

		//sempre que estiver na posição 0, a password incrementa
		if position == 0 {
			password++;
		}
	}
	fmt.Println(password);
}