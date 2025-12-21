package model

type Session struct {
	Count int
}

type Timer struct {
	Type int
	Time int
	Data []byte
}

func (s *Session) AddCount() {
	if s.Count == 8 {
		s.Count = 0
	}
	s.Count++
}
