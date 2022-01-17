package language

import (
	"log"
	"testing"
)

func TestKVLoop(t *testing.T) {
	log.Println("hi")
	a := map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
	}
	b := []int{1, 2, 3, 4, 5}
	for k, v := range a {
		log.Printf("the ptr to k,v is %x %x\n", &k, &v)
		log.Printf("the ptr to map[%s] is %x", a[k], &a)
	}
	for k, v := range b { // k,v loop can not modify the value in slice
		log.Printf("the ptr to k=%v,v=%v is %x %x\n", k, v, &k, &v)
		v = 111111
		// log.Printf("the ptr to map[%s] is %x", a[k], &a)
	}
	log.Println(b)

}
