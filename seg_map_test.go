package noun

import (
	"log"
	"testing"
)

func TestNewSegMap(t *testing.T) {
	s:=&segMap{}
	log.Println(s.hashKeyM("12312312"))
	log.Println(s.hashKeyM("qerqer"))
	log.Println(s.hashKeyM("123qwerqwrqr12312"))
	log.Println(s.hashKeyM("sdfgsafdgs"))

}
