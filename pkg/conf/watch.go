package conf

type Watch struct {
	l chan struct{}
}

func (w *Watch) GetSignal() chan struct{} {
	return w.l
}

func (w *Watch) Send() {
	w.l <- struct{}{}
}
