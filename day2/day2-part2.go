package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2Part2() {
	content, err := os.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}
	
	var total int64 = 0;
	intervales := strings.Split(string(content), ",");

	for _, interval := range intervales {

		numbers := strings.Split(interval, "-");

		a, err := strconv.Atoi(numbers[0]);
		if err != nil {
			log.Fatal(err);
		}

		b, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err);
		}

		for j := a; j <= b; j++ {

			numStr := fmt.Sprint(j);

			lenght := len(numStr);

			if strings.Compare(numStr[:(lenght/2)], numStr[(lenght/2):]) == 0 {

				total += int64(j);

			} else {

				for i := 0; i < len(numStr)/2; i++{

					subNumber := string(numStr[:i+1]);

					times := strings.Count(numStr, subNumber);

					if times * len(subNumber) == len(numStr) {

						total += int64(j);

					}
				}
			}
		}
	}
	fmt.Println(total);
}