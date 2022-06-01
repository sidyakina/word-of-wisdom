package challenge

import "math/rand"

type Challenge struct {
	Challenge string
}

func NewChallenge(str string) Challenge {
	return Challenge{Challenge: str}
}

func GenerateNewChallenge() Challenge {
	return Challenge{Challenge: generateRandomString(lenChallengeString)}
}

func generateRandomString(lenStr int) string {
	str := ""
	for i := 0; i < lenStr; i++ {
		k := rand.Intn(len(symbols))
		str += string(symbols[k])
	}

	return str
}
