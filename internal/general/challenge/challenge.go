package challenge

type Challenge struct {
}

func NewChallenge() Challenge {
	return Challenge{}
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

func (c *Challenge) FindAnswer() (*Answer, error) {
	return &Answer{}, nil
}
