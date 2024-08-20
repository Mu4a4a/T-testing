package service

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{prod: prod, pres: pres}
}

type Service struct {
	prod Producer
	pres Presenter
}

func (s Service) Run() {

	in, _ := s.prod.Produce()

	s.pres.Present(in)
}

func spammyMasker(input []byte) string {
	link := "http://"
	nlink := len(link)
	outputSlice := make([]byte, len(input))
	copy(outputSlice, input)

	for i := 0; i <= len(input)-nlink; i++ {
		if string(input[i:i+nlink]) == link {
			j := i + nlink
			for j < len(input) && (libray(input[j]) || input[j] == '_' || input[j] == '.' || input[j] == '~' || input[j] == '-') {
				outputSlice[j] = '*'
				j++
			}
			i = j - 1
		}
	}
	return string(outputSlice)
}

func libray(a byte) bool {
	return (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}
