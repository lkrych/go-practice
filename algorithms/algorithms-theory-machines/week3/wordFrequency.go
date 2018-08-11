package wordFrequency

import (
	"fmt"
	"io/ioutil"
	"log"
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
	fmt.Println(table)
}

func createFreqTable(arr []string) SymbolTable {

}

type SymbolTable struct {
	data *BSTree
}

func (st *SymbolTable) put(key string, val int) {

}

func (st *SymbolTable) get(key string) {

}

func (st *SymbolTable) contains(key string) {

}

func (st *SymbolTable) keys() []string {

}

type BSTree struct {
	root *Node
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

func put(x *Node, key string, value int) *Node {
	if x == nil {
		return &Node{
			key: key,
			val: value,
		}
	}
	if x.key > key {
		put(x.left, key, value)
	} else if x.key < key {
		put(x.right, key, value)
	} else {
		x.val = value
	}
	return x
}

func (btree *BSTree) returnKeys() []string {

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
