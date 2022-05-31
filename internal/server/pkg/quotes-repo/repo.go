package quotesrepo

type Repo struct {
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) GetQuote() (string, error) {
	return "test", nil
}
