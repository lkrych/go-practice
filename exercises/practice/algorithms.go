package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type UF struct {
	objects    []int
	objectSize []int
}

func (u *UF) Union(p int, q int) {
	pID := u.objects[p]
	qID := u.objects[q]
	if pID == qID {
		return
	}
	for i, val := range u.objects {
		if val == pID {
			u.objects[i] = qID
		}
	}
}

func (u *UF) Connected(p int, q int) bool {
	return u.objects[p] == u.objects[q]
}

func (u *UF) Count() int {
	rootMap := make(map[int]bool)
	count := 0
	for _, root := range u.objects {
		if _, ok := rootMap[root]; ok {
			continue
		}
		count++
		rootMap[root] = true
	}
	return count
}

func NewUF(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	byLine := strings.Split(string(content), "\n")
	size, err := strconv.Atoi(byLine[0])
	if err != nil {
		log.Fatal(err)
	}
	objects := make([]int, size)
	objectSize := make([]int, size)
	for i := 0; i < size; i++ {
		objects[i] = i
		objectSize[i] = 1
	}
	unionFindObject := &UF{
		objects:    objects,
		objectSize: objects,
	}

	for _, line := range byLine[1:] {
		characterInts := strings.Split(line, " ")
		p, err := strconv.Atoi(characterInts[0])
		if err != nil {
			log.Fatal(err)
		}
		q, err := strconv.Atoi(characterInts[1])
		if err != nil {
			log.Fatal(err)
		}
		if unionFindObject.Connected(p, q) {
			fmt.Printf("%v is connected to %v \n", p, q)
			continue
		}
		unionFindObject.Union(p, q)
		fmt.Printf("%v, %v \n", p, q)
	}

	fmt.Printf("There are %v connected objects! \n", unionFindObject.Count())

}

func main() {
	NewUF("./algorithm_input/tinyUF.txt")
}
