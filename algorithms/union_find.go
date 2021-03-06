package main

import (
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
	Count() int
}

func NewUF(filename string, u UF) *UF {
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
			// fmt.Printf("%v is connected to %v \n", p, q)
			continue
		}
		unionFindObject.Union(p, q)
		// fmt.Printf("%v, %v \n", p, q)
	}
	return &unionFindObject
}

//NaiveUF has O(n) Union method and an O(1) Find method
type NaiveUF struct {
	objects    []int
	objectSize []int
	count      int
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
	u.count--
}

func (u *NaiveUF) Find(p int) int {
	return u.objects[p]
}

func (u *NaiveUF) Connected(p int, q int) bool {
	return u.objects[p] == u.objects[q]
}

func (u *NaiveUF) Count() int {
	return u.count
}

func (u *NaiveUF) New(objects []int, objectSize []int) UF {
	return &NaiveUF{
		objects:    objects,
		objectSize: objectSize,
		count:      len(objects),
	}
}

//QuickUnionUF has a O(1) union method, but the connected method now takes longer,
// possibly O(n), because objects are now stored as trees and the Big-O complexity of
// Find is now dependent upon the depth of the tree
type QuickUnionUF struct {
	objects    []int
	objectSize []int
	count      int
}

func (u *QuickUnionUF) Union(p int, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)

	if rootP == rootQ {
		return
	}
	u.objects[rootP] = rootQ
	u.count--
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

func (u *QuickUnionUF) Count() int {
	return u.count
}

func (u *QuickUnionUF) New(objects []int, objectSize []int) UF {
	return &QuickUnionUF{
		objects:    objects,
		objectSize: objectSize,
		count:      len(objects),
	}
}

//WeightedQuickUnionUF makes up for the inefficiencies in the Find method of
// QuickUnionUF by making sure the smaller tree is always assigned the larger tree.
// This makes sure the trees are wider rather than taller. This makes Union O(logn).
type WeightedQuickUnionUF struct {
	objects    []int
	objectSize []int
	count      int
}

func (u *WeightedQuickUnionUF) Union(p int, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)

	if rootP == rootQ {
		return
	}

	if u.objectSize[rootP] < u.objectSize[rootQ] {
		u.objects[rootP] = rootQ
		u.objectSize[rootQ] += u.objectSize[rootP]
	} else {
		u.objects[rootQ] = rootP
		u.objectSize[rootP] += u.objectSize[rootQ]
	}
	u.count--
}

func (u *WeightedQuickUnionUF) Find(p int) int {
	for p != u.objects[p] {
		p = u.objects[p]
	}
	return p
}

func (u *WeightedQuickUnionUF) Connected(p int, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *WeightedQuickUnionUF) Count() int {
	return u.count
}

func (u *WeightedQuickUnionUF) New(objects []int, objectSize []int) UF {
	return &WeightedQuickUnionUF{
		objects:    objects,
		objectSize: objectSize,
		count:      len(objects),
	}
}

//WeightedQuickUnionWithCompressionUF improves performance by adding path compression
// to weighted quick union
// adding a loop in the find block to make sure that every node we inspect
// has its path re-assigned.
type WeightedQuickUnionWithCompressionUF struct {
	objects    []int
	objectSize []int
	count      int
}

func (u *WeightedQuickUnionWithCompressionUF) Union(p int, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)

	if rootP == rootQ {
		return
	}

	if u.objectSize[rootP] < u.objectSize[rootQ] {
		u.objects[rootP] = rootQ
		u.objectSize[rootQ] += u.objectSize[rootP]
	} else {
		u.objects[rootQ] = rootP
		u.objectSize[rootP] += u.objectSize[rootQ]
	}
	u.count--
}

func (u *WeightedQuickUnionWithCompressionUF) Find(p int) int {
	root := p
	for root != u.objects[root] {
		root = u.objects[root]
	}
	//assign sub-trees to root
	//constant O(logn) sub-procedure amortizes Find calls
	for p != root {
		newp := u.objects[p]
		u.objects[p] = root
		p = newp
	}
	return root
}

func (u *WeightedQuickUnionWithCompressionUF) Connected(p int, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *WeightedQuickUnionWithCompressionUF) Count() int {
	return u.count
}

func (u *WeightedQuickUnionWithCompressionUF) New(objects []int, objectSize []int) UF {
	return &WeightedQuickUnionWithCompressionUF{
		objects:    objects,
		objectSize: objectSize,
		count:      len(objects),
	}
}

func main() {
	NewUF("./algorithm_input/tinyUF.txt", &WeightedQuickUnionWithCompressionUF{})
}
