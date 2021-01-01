package sign_up

type EventData struct {
	accountID int64
	name      string
}

func (e EventData) AccountID() int64 {
	return e.accountID
}

func (e EventData) Name() string {
	return e.name
}
