package mtproto

import "sync"

type syncError struct {
	error error
	sync  sync.Mutex
}

func (e *syncError) clear() {
	e.set(nil)
}

func (e *syncError) set(err error) {
	e.sync.Lock()
	e.error = err
	e.sync.Unlock()
}

func (e *syncError) setIfEmpty(err error) {
	e.sync.Lock()
	if e.error == nil {
		e.error = err
	}
	e.sync.Unlock()
}

func (e *syncError) get() error {
	e.sync.Lock()
	err := e.error
	e.sync.Unlock()
	return err
}
