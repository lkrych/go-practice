package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type UF interface {
	Union(p int, q int)
	Connected(p int, q int) bool
	Find(p int) int
	New(objects []int, objectSize []int) UF
	getObjects() []int
}

func Count(objects []int) int {
	rootMap := make(map[int]bool)
	count := 0
	for _, root := range objects {
		if _, ok := rootMap[root]; ok {
			continue
		}
		count++
		rootMap[root] = true
	}
	return count
}

func NewUF(filename string, u UF) {
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
	unionFindObject := u.New(objects, objectSize)

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

	fmt.Printf("There are %v connected objects! \n", Count(unionFindObject.getObjects()))
}

//NaiveUF has O(n) union method,
type NaiveUF struct {
	objects    []int
	objectSize []int
}

func (u *NaiveUF) Union(p int, q int) {
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

func (u *NaiveUF) Find(p int) int {
	return u.objects[p]
}

func (u *NaiveUF) Connected(p int, q int) bool {
	return u.objects[p] == u.objects[q]
}

func (u *NaiveUF) getObjects() []int {
	return u.objects
}

func (u *NaiveUF) New(objects []int, objectSize []int) UF {
	return &NaiveUF{
		objects:    objects,
		objectSize: objectSize,
	}
}

//QuickUnionUF has a O(1) union method, but the connected method now takes longer, possibly O(n), or the depth of the tree
type QuickUnionUF struct {
	objects    []int
	objectSize []int
}

func (u *QuickUnionUF) Union(p int, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)

	if rootP == rootQ {
		return
	}
	u.objects[rootP] = rootQ
}

func (u *QuickUnionUF) Find(p int) int {
	for p != u.objects[p] {
		p = u.objects[p]
	}
	return p
}

func (u *QuickUnionUF) Connected(p int, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *QuickUnionUF) getObjects() []int {
	return u.objects
}

func (u *QuickUnionUF) New(objects []int, objectSize []int) UF {
	return &QuickUnionUF{
		objects:    objects,
		objectSize: objectSize,
	}
}

func main() {
	NewUF("./algorithm_input/tinyUF.txt", &NaiveUF{})
}
