package models

type Woman struct {
	Man
}

func (w *Woman) Breath() {
	w.IsBreathing = true
}

func (w *Woman) Eat() {
	w.IsEating = true
}

func (w *Woman) Think() {
	w.IsThinking = true
}

func (w *Woman) GetGender() string {
	return "Woman"
}
