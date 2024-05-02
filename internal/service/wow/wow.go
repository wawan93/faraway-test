package wow

import "math/rand"

// Service provides a random quote
type Service struct {
	quotes []string
}

// New creates a new Service
func New() *Service {
	return &Service{
		quotes: []string{
			"Be yourself; everyone else is already taken.",
			"Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
			"So many books, so little time.",
			"A room without books is like a body without a soul.",
			"You only live once, but if you do it right, once is enough.",
			"Be the change that you wish to see in the world.",
			"In three words I can sum up everything I've learned about life: it goes on.",
			"If you tell the truth, you don't have to remember anything.",
		},
	}
}

// Quote returns a random quote
func (s Service) Quote() string {
	return s.quotes[rand.Int31n(int32(len(s.quotes)))]
}
