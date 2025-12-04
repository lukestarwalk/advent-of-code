package day1

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
);

func dayPart2(){

	content, err := os.ReadFile("input.txt");
	
	if err != nil {
		log.Fatal(err);
	}
	
	lines := string(content);
	
	splitted := strings.Fields(lines);
	
	position, password := 50, 0;

	for _, dial_move := range splitted {

		direction := string(dial_move[0]);

		number, err := strconv.Atoi(dial_move[1:]);

		if err != nil {
			log.Fatal(err);
		}

		if direction == "L" {
			for i := 1; i <= number; i++ {
				position--;
				if position < 0 {
					position = 99;
				}
				if position == 0 {
					password++;
				}
			}
		
		} else {
			for i := 1; i <= number; i++ {
				position++;
				if position > 99 {
					position = 0;
				}
				if position == 0 {
					password++;
				}
			}
		}
	}
	fmt.Println(password);
}