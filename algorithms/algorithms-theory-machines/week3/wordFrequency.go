package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	File string `short:"f" long:"file" default:"generatedStrings.txt" description:"a file to read strings from."`
}

func main() {
	//interpret the CL stdin
	flags.Parse(&opts)
	file, err := ioutil.ReadFile(opts.File)
	checkErr(err)
	//create an array of strings that is split by the newline character
	splitByWord := strings.Split(string(file), " ")
	table := createFreqTable(splitByWord)
	fmt.Println("The frequency table is")
	wordCount := 0
	for _, node := range table.keys() {
		wordCount++
		fmt.Printf("%v -> %v \n", node.key, node.val)
	}
	fmt.Println("The number of unique words was", wordCount)
}

func createFreqTable(arr []string) SymbolTable {
	st := &SymbolTable{
		data: &BSTree{},
	}
	r := regexp.MustCompile("[ .,123456789]")
	for _, word := range arr {
		clean := cleanWord(word)
		if r.Match([]byte(clean)) {
			//skip punctuation or spaces
			continue
		}
		found := st.get(word)
		if found != nil {
			st.put(word, found.val+1)
		} else {
			st.put(word, 1)
		}
	}
	return *st
}

func cleanWord(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "'", "", -1)
	s = strings.Replace(s, "\"", "", -1)
	return s
}

type SymbolTable struct {
	data *BSTree
}

func (st *SymbolTable) put(key string, val int) {
	st.data.root = put(st.data.root, key, val)
}

func (st *SymbolTable) get(key string) *Node {
	return get(st.data.root, key)
}

func (st *SymbolTable) contains(key string) bool {
	return get(st.data.root, key) == nil
}

func (st *SymbolTable) keys() []Node {
	return returnKeys(st.data.root)
}

type BSTree struct {
	root *Node
}

func (btree *BSTree) isEmpty() bool {
	return btree.root == nil
}

//search for the key
func get(x *Node, key string) *Node {
	if x == nil {
		return nil
	}
	if x.key > key {
		return get(x.left, key)
	} else if x.key < key {
		return get(x.right, key)
	} else {
		return x
	}
}

//update a key, or insert a new node into tree
func put(x *Node, key string, value int) *Node {
	if x == nil {
		return &Node{
			key: key,
			val: value,
		}
	}
	if x.key > key {
		x.left = put(x.left, key, value)
	} else if x.key < key {
		x.right = put(x.right, key, value)
	} else {
		x.val = value
	}
	return x
}

func returnKeys(x *Node) []Node {
	ordered := []Node{}
	if x == nil {
		return ordered
	}
	//in-order traversal
	ordered = append(ordered, returnKeys(x.left)...)
	ordered = append(ordered, *x)
	ordered = append(ordered, returnKeys(x.right)...)
	return ordered
}

type Node struct {
	key   string
	val   int
	left  *Node
	right *Node
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
