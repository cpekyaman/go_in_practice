package gotest

type Messager interface {
	Send(email string, subject string, body string) error
}

func Alert(m Messager, err string) error {
	return m.Send("support@nowhere.com", "Problem", err)
}
