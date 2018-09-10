package mtproto

type UpdateListener interface {
	Notify(update TL_updates)
}

type FuncUpdateListener func(update TL_updates)

func NewFuncUpdateListener(f func(update TL_updates)) FuncUpdateListener {
	return FuncUpdateListener(f)
}

func (f FuncUpdateListener) Notify(update TL_updates) {
	f(update)
}
