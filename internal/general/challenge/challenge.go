package challenge

import (
	cryptorand "crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Challenge struct {
	Challenge string
}

func NewChallenge(str string) Challenge {
	return Challenge{Challenge: str}
}

func GenerateNewChallenge() Challenge {
	return Challenge{Challenge: generateCryptoRandomString(lenChallengeString)}
}

func (c *Challenge) FindAnswer() (string, error) {
	timer := time.NewTimer(generationTimeout)
	defer timer.Stop()

	var (
		i      int
		answer string
	)
	for {
		select {
		case <-timer.C:
			return "", fmt.Errorf("generation time more than max: %v", generationTimeout)
		default:
			i++
			answer = generateRandomString(lenAnswerString)
		}

		if valid(c.Challenge + answer) {
			log.Printf("answer %v (attempts %v)", answer, i)

			break
		}
	}

	return answer, nil
}

func (c *Challenge) CheckAnswer(answer string) bool {
	if len(answer) != lenAnswerString {
		log.Printf("wrong len(answer): %v, want: %v", len(answer), lenAnswerString)

		return false
	}

	for _, v := range answer {
		if !strings.ContainsRune(symbols, v) {
			log.Printf("there is wrong symbol %v in answer", v)

			return false
		}
	}

	return valid(c.Challenge + answer)
}

func valid(data string) bool {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
	zeros := calculateLeadingZeros(hash)

	return zeros >= numberZeroBits
}

func calculateLeadingZeros(hash string) int {
	var zeros int
	for _, v := range hash {
		switch v {
		case '0':
			zeros += 4 // 0 -> 0000

			continue
		case '1':
			zeros += 3 // 1 -> 0001
		case '2', '3':
			zeros += 2
		case '4', '5', '6', '7':
			zeros += 1
		}

		break
	}

	return zeros
}

func generateRandomString(lenStr int) string {
	str := ""
	for i := 0; i < lenStr; i++ {
		k := rand.Intn(len(symbols))
		str += string(symbols[k])
	}

	return str
}

func generateCryptoRandomString(lenStr int) string {
	numberSymbols := len(symbols)

	b := make([]byte, lenStr)
	_, err := cryptorand.Read(b)
	if err != nil {
		log.Printf("failed to generate crypto random string, use math random string")

		return generateRandomString(lenStr)
	}

	for i := 0; i < lenStr; i++ {
		b[i] = symbols[int(b[i])%numberSymbols]
	}

	return string(b)
}
