package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text : ")
	text, err := reader.ReadString('\n') // Accept input text
	if err != nil {
		fmt.Println(err)
	}
	result := getMostUsedWords(text)
	if resultJson, err := json.Marshal(result); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(resultJson))
	}

}
func getMostUsedWords(text string) map[string]int {
	text = strings.ReplaceAll(text, ",", "") // remove comma
	text = strings.ReplaceAll(text, ".", "") // remove fullstop
	wordsArray := strings.Fields(text)       // split text to words array
	wordCount := make(map[string]int)        // make map  key as word and value as number of count that word occurs
	for _, word := range wordsArray {
		if _, found := wordCount[word]; found {
			wordCount[word]++
		} else {
			wordCount[word] = 1
		}
	}
	wordsArray = nil // clear the wordsArray and store words in wordCount map
	for word := range wordCount {
		wordsArray = append(wordsArray, word)
	}
	sort.Slice(wordsArray, func(i, j int) bool { //sort the wordsArray according to the value from wordCount map
		return wordCount[wordsArray[i]] > wordCount[wordsArray[j]]
	})
	result := make(map[string]int)
	if len(wordCount) >= 10 { // check length of word count to overcome runtime error
		for i := 0; i < 10; i++ {
			//fmt.Println(wordsArray[i], wordCount[wordsArray[i]])
			result[wordsArray[i]] = wordCount[wordsArray[i]]
		}
	} else {
		for _, word := range wordsArray {
			result[word] = wordCount[word]
		}
	}
	return result

}
