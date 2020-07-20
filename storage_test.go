package noun

import (
	"log"
	"strconv"
	"testing"
)

type testNoun struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func TestNewNoun(t *testing.T) {
	newNoun := NewNoun(100)
	for i := 0; i < 1000; i++ {
		name := strconv.Itoa(i)
		newNoun.Set(name, &testNoun{i, name})
	}
	t1, ok1 := newNoun.Get("999").(*testNoun)
	if ok1 {
		log.Println(t1.Age, t1.Name)
	}
	t2, ok2 := newNoun.Get("2").(*testNoun)
	if ok2 {
		log.Println(t2.Age, t2.Name)
	}else{
		if newNoun.Get("2")==nil{
			log.Println("not in lru")
		}
	}
}
