package sign_up

type Event interface {
	Name() string
	AccountID() int64
}
