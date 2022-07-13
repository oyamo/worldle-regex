package dict

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

type WordDB struct {
	words *[]string
}

func LoadDict(path string) (*WordDB, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("file closing error ", err)
		}
	}(file)

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &WordDB{
		words: &words,
	}, nil
}

func (db *WordDB) SearchByRegex(regex string) (*[]string, error) {
	var result []string
	re, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	for _, word := range *db.words {
		if re.MatchString(word) {
			result = append(result, word)
		}
	}
	return &result, nil
}
