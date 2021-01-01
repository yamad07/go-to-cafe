package create_cafe

import "sync"

var (
	once     sync.Once
	instance Notifier
)

type Observer interface {
	OnNotify(Event)
}

type Notifier struct {
	observers map[Observer]struct{}
}

func NewNotifier() Notifier {
	once.Do(func() {
		instance = Notifier{
			observers: map[Observer]struct{}{},
		}
	})
	return instance
}

func (n Notifier) Register(o Observer) {
	n.observers[o] = struct{}{}
}

func (n Notifier) Deregister(o Observer) {
	delete(n.observers, o)
}

func (n Notifier) Notify(e Event) {
	for o := range n.observers {
		o.OnNotify(e)
	}
}
