package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day3Part1() {

	content, err := os.ReadFile("input.txt");

	if err != nil {
		log.Fatal(err);
	}

	total, batteries := 0, strings.Fields(string(content));

	for _, banks := range batteries {

		first, second := 0, 0;

		for  i, digits := range banks {

			num, err := strconv.Atoi(string(digits));

			if err != nil {
				log.Fatal(err);
			}

			if num > first && i < len(banks) - 1{
				first = num;
				second = 0;
				continue;
			}

			if num > second {
				second = num;
			}
		}
		result := (first*10) + second;
		total += result;
	}
	fmt.Printf("Total: %d\n", total);	
}