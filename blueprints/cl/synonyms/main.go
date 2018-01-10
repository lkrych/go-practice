package main

import (
	"bufio"
	"cl/thesaurus"
	"fmt"
	"log"
	"os"
)

func main() {
	//parse secrets
	var thesaurus = &thesaurus.BigHuge{}
	thesaurus = thesaurus.GetConfig()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for %v: %v", word, err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find synonyms for %v", word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}

}
