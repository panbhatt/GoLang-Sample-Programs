package main

import (
	"fmt"
	"bufio"
	"os"
	//"strings"
	"sort"
)

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	mpOfWords := make(map[string]int)

	scanner.Split(bufio.ScanWords)


	for scanner.Scan() {
		mpOfWords[scanner.Text()]++
	}  // Press CTRL + D to end the input

	fmt.Println("Total NO of Unique Words = ",len(mpOfWords) )

	// Print top 5 words
	type kv struct {
		key string
		occurence int
	}

	var slOfAllWords []kv


	for k,v := range mpOfWords {
		slOfAllWords = append(slOfAllWords, kv{k,v})
	}

	

	sort.Slice(slOfAllWords, func(i , j int) bool {
		return slOfAllWords[i].occurence > slOfAllWords[j].occurence	
	})

	fmt.Printf("%v", slOfAllWords)

	for _,val := range slOfAllWords[:3] {
		fmt.Println(val.key , " -> " , val.occurence)
	}


}