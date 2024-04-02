package models

type Man struct {
	Age         int
	Height      float32
	Weight      float32
	IsBreathing bool
	IsThinking  bool
	IsEating    bool
	IsAlive     bool
}

func (m *Man) Breath() {
	m.IsBreathing = true
}

func (m *Man) Eat() {
	m.IsEating = true
}

func (m *Man) Think() {
	m.IsThinking = true
}

func (m *Man) GetGender() string {
	return "Man"
}
