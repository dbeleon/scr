package main

import (
	"encoding/base64"
	"log"
	"math"
	"math/rand"
	"strings"

	"github.com/dbeleon/scr"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	s := strings.Join([]string{gofakeit.Street(), gofakeit.Street(), gofakeit.Street(), gofakeit.Street(), gofakeit.Street(), gofakeit.Street()}, " ")
	log.Println(s)
	for t := 0; t < 1000; t++ {
		l := int(math.Max(5, float64(rand.Intn(65))))
		feedbacks := make([]int, int(math.Max(2, float64(rand.Intn(int(float32(l)/3.0*2.0))))))
		idxs := make([]int, l)
		for i := 0; i < l; i++ {
			idxs[i] = i
		}
		gofakeit.ShuffleInts(idxs)
		for i := 0; i < len(feedbacks); i++ {
			feedbacks[i] = idxs[i]
		}
		poly := gofakeit.Uint64()
		poly = poly & (uint64(0xFFFF_FFFF_FFFF_FFFF) >> (64 - l))
		log.Printf("Length: %d, Feedbacks: %v, Polynomial: %X\n", l, feedbacks, poly)
		scram := scr.New(l, feedbacks, poly)
		data := []byte(s)
		scram.ScrambleAdditive(data)
		log.Println(base64.StdEncoding.EncodeToString(data))
		scram.DescrambleAdditive(data)
		log.Println(string(data))
		res := s == string(data)
		if !res {
			panic("descramble error")
		}
		log.Printf("Result: %v\n", res)
	}
}
