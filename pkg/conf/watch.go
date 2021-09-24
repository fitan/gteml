package conf

type Watch struct {
	l []chan struct{}
}

func (w *Watch) GetSignal() chan struct{} {
	c := make(chan struct{})
	w.l = append(w.l, c)
	return c
}

func (w *Watch) Send() {
	for _, c := range w.l {
		c <- struct{}{}
	}
}
