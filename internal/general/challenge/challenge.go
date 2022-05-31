package challenge

type Challenge struct {
}

func GenerateNewChallenge() Challenge {
	return Challenge{}
}

type Answer struct {
}

func NewAnswer() Answer {
	return Answer{}
}

func (c *Challenge) CheckAnswer(answer Answer) bool {
	return true
}
