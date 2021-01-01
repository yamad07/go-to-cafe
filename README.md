# go-to-cafe

This is a sample repository of modular monolith architecture.


The main features are below.
- High Module Independency
    - Each domain-specific module is independent.
    - There are strong constraints for crossing domain boundaries.
- High Deployability
    - Low deployment costs due to a single server.

# Architecture
```
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── api
│       └── server # This is a main entry point.
├── db
│   ├── migrations
│   └── schema.sql
├── docker-compose.yaml
├── domain # This is the directory with the independent domain modules.
│   ├── auth
│   │   ├── adapter
│   │   ├── pkg
│   │   └── usecase
│   ├── cafe
│   │   ├── adapter
│   │   ├── pkg
│   │   └── usecase
│   ├── search
│   │   ├── adapter
│   │   ├── pkg
│   │   └── usecase
│   └── user
│       ├── adapter
│       ├── pkg
│       └── usecase
├── go.mod
├── go.sum
├── pkg
│   ├── apperror
│   ├── boundary # These are the interfaces to use when crossing the domain boundary. (by using `singleton`, `observer`, `go interface`)
│   │   ├── create_cafe
│   │   │   ├── event.go
│   │   │   └── notifer.go
│   │   └── sign_up
│   │       ├── event.go
│   │       └── notifer.go
│   ├── config
│   ├── elasticsearch
│   ├── mysql
│   ├── presenter
│   └── registry
└── server
```

# Usage

```
$ make migrate
$ make run
```

# Crossing the Baundary
**Before modular monolith**
```
func (u Usecase) SignUp(input Input) (output Output, aerr apperror.Error) {
	p := value.NewCreateAccountParam(input.Email, input.Password)
	aerr = u.accountRepository.Create(p)
	if aerr != nil {
		return output, aerr
	}
	act, aerr := u.accountRepository.GetByUID(p.UID)
	if aerr != nil {
		return output, aerr
	}
	aerr := u.userRepository.Create(act.ID)
	if aerr != nil {
		return output, aerr
	}

	output = Output{
		Account: act,
	}

	return output, nil
}
```
In this example, the usecase of `auth` dependent on the `user`.

**In modular monolith**
```
func (u Usecase) SignUp(input Input) (output Output, aerr apperror.Error) {
	p := value.NewCreateAccountParam(input.Email, input.Password)
	aerr = u.accountRepository.Create(p)
	if aerr != nil {
		return output, aerr
	}
	act, aerr := u.accountRepository.GetByUID(p.UID)
	if aerr != nil {
		return output, aerr
	}

	evt := EventData{
		accountID: act.ID,
		name:      input.Name,
	}
    // notification to subscribers
	sign_up.NewNotifier().Notify(evt)

	output = Output{
		Account: act,
	}

	return output, nil
}
```
In this example, `auth` does not dependent on `user`.

To achieve this, I used the singleton pattern, observer pattern, and the go interface.
For more details, please read [this code](https://github.com/yamad07/go-to-cafe/tree/main/pkg/boundary).
