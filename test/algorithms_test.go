package test

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_Algorithms(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	n := 20

	for lenServerString := 10; lenServerString <= 40; lenServerString = lenServerString + 10 {
		var t1, t2, t1max, t2max time.Duration
		var p1, p2, p1max, p2max int

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
					if i > p1max {
						p1max = i
					}

					break
				}
			}
			t := time.Since(time1)
			t1 += t
			if t > t1max {
				t1max = t
			}

			time2 := time.Now()
			for i := 0; ; i++ {
				b := generateRandomString(10)

				data := []byte(serverString + b)

				d := fmt.Sprintf("%x", sha256.Sum256(data))

				if d[0] == '0' && d[1] == '0' && d[2] == '0' && d[3] == '0' && d[4] == '0' {
					p2 += i
					if i > p2max {
						p2max = i
					}

					break
				}
			}
			t = time.Since(time2)
			t2 += t
			if t > t2max {
				t2max = t
			}
		}

		fmt.Printf("\nlen(server string) = %v", lenServerString)
		fmt.Printf("\nsha1: \ntotal time %v \navg time: %v, avg attempts: %v", t1, t1/time.Duration(n), p1/n)
		fmt.Printf("\nmax time %v, max attempts: %v", t1max, p1max)
		fmt.Printf("\nsha2: \ntotal time %v, \navg time: %v, avg attempts: %v", t2, t2/time.Duration(n), p2/n)
		fmt.Printf("\nmax time %v, max attempts: %v\n", t2max, p2max)
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

/*
len(server string) = 10
sha1:
total time 13.285327532s,
avg time: 664.266376ms, avg attempts: 851535
max time 1.733282266s, max attempts: 2221810
sha2:
total time 20.339648831s,
avg time: 1.016982441s, avg attempts: 1144626
max time 3.351335874s, max attempts: 3751754

len(server string) = 20
sha1:
total time 16.205209778s,
avg time: 810.260488ms, avg attempts: 1043763
max time 4.598956506s, max attempts: 5943154
sha2:
total time 15.566848979s,
avg time: 778.342448ms, avg attempts: 879487
max time 3.450856599s, max attempts: 3878805

len(server string) = 30
sha1:
total time 16.468408262s,
avg time: 823.420413ms, avg attempts: 971582
max time 2.948801529s, max attempts: 3425847
sha2:
total time 23.025300398s,
avg time: 1.151265019s, avg attempts: 1211975
max time 4.54900777s, max attempts: 4839195

len(server string) = 40
sha1:
total time 21.759263977s,
avg time: 1.087963198s, avg attempts: 1285492
max time 3.349470774s, max attempts: 3883787
sha2:
total time 23.058677067s,
avg time: 1.152933853s, avg attempts: 1201996
max time 3.576479828s, max attempts: 3801497
--- PASS: Test_Algorithms (149.71s)
PASS

Process finished with the exit code 0

*/
