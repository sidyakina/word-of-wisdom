package test

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func Test_Algorithms(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	n := 20

	for lenServerString := 10; lenServerString <= 40; lenServerString = lenServerString + 10 {
		t1, t2 := time.Duration(0), time.Duration(0)
		p1, p2 := 0, 0

		for k := 0; k < n; k++ {
			serverString := generateRandomString(lenServerString)

			// log.Printf("%v = %s", k, serverString)

			time1 := time.Now()
			for i := 0; ; i++ {
				b := generateRandomString(10)

				data := []byte(serverString + b)

				d := fmt.Sprintf("%x", sha1.Sum(data))

				if d[0] == '0' && d[1] == '0' && d[2] == '0' && d[3] == '0' && d[4] == '0' {
					p1 += i
					break
				}
			}
			t1 += time.Since(time1)

			time2 := time.Now()
			for i := 0; ; i++ {
				b := generateRandomString(10)

				data := []byte(serverString + b)

				d := fmt.Sprintf("%x", sha256.Sum256(data))

				if d[0] == '0' && d[1] == '0' && d[2] == '0' && d[3] == '0' && d[4] == '0' {
					p2 += i
					break
				}
			}
			t2 += time.Since(time2)
		}

		log.Printf("len(server string) = %v", lenServerString)
		log.Printf("sha1: total time %v, avg time: %v, avg attempts: %v", t1, t1/time.Duration(n), p1/n)
		log.Printf("sha2: total time %v, avg time: %v, avg attempts: %v\n", t2, t2/time.Duration(n), p2/n)
	}
}

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(lenStr int) string {
	str := ""
	for i := 0; i < lenStr; i++ {
		k := rand.Intn(len(symbols))
		str += string(symbols[k])
	}

	return str
}

//2022/06/02 21:04:12 len(server string) = 10
//2022/06/02 21:04:12 sha1: total time 13.114112559s, avg time: 655.705627ms, avg attempts: 825302
//2022/06/02 21:04:12 sha2: total time 20.147067551s, avg time: 1.007353377s, avg attempts: 1126127

//2022/06/02 21:04:59 len(server string) = 20
//2022/06/02 21:04:59 sha1: total time 23.480415041s, avg time: 1.174020752s, avg attempts: 1487351
//2022/06/02 21:04:59 sha2: total time 24.268615528s, avg time: 1.213430776s, avg attempts: 1364958

//2022/06/02 21:05:32 len(server string) = 30
//2022/06/02 21:05:32 sha1: total time 14.642746618s, avg time: 732.13733ms, avg attempts: 837336
//2022/06/02 21:05:32 sha2: total time 17.441058231s, avg time: 872.052911ms, avg attempts: 900725

//2022/06/02 21:06:20 len(server string) = 40
//2022/06/02 21:06:20 sha1: total time 24.163660117s, avg time: 1.208183005s, avg attempts: 1358091
//2022/06/02 21:06:20 sha2: total time 23.917443648s, avg time: 1.195872182s, avg attempts: 1208527
