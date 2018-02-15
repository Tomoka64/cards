package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for _, v := range d {
		fmt.Println(v)
	}
}

func (d deck) ToString() string {
	return strings.Join(d, ",")
}

func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func NewDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

//
// func removeDuplicates(d deck) deck {
// 	count := 0
// 	ret := []deck{}
// 	trash := []deck{}
// 	for i, v := range d {
// 		if i == 0 {
// 			count++
// 			ret = append(ret, v)
// 		} else if d[i-1] != v {
// 			count++
// 			ret = append(ret, v)
// 		} else if d[i-1] == v {
// 			trash = append(trash, d[i])
// 		}
// 	}
// 	for i, v := range ret {
// 		d[i] = v
// 	}
// 	result := []deck{}
// 	for i := 0; i < len(trash)-1; i++ {
// 		for j := 0; j < len(ret)-1; j++ {
// 			if trash[i] != ret[j] {
// 				result = append(result, ret[j])
// 			}
// 		}
// 	}
// 	return result
// }

// func (d deck) deleteTheSame() {
// 	// var m []string
// 	a := d.split()
// 	fmt.Println(a)
// }

func split(d []string) [][]string {
	var result [][]string
	// var res []string
	for _, val := range d {
		result = append(result, strings.Split(val, ","))
	}

	return result

}

func removeDuplicates(cards [][]string) {
	// var result [][]string
	// var trash [][]string
	// var ret [][]string
	// var count = 0
	for _, v := range cards {
		fmt.Println(v)
	}
}

// func remove(arr []string) {
// 	m := make(map[string]struct{})
// 	for _, ele := range arr {
// 		m[ele] = struct{}{} // m["a"] = struct{}{} が二度目は同じものとみなされて重複が消える。
// 	}
//
// 	uniq := []string{}
// 	for i := range m {
// 		uniq = append(uniq, i)
// 	}
//
// 	fmt.Printf("%v\n", uniq) // ["a", "b", "c"]
//
// }
