package types

type ServicesI interface {
	User() User
}

type User interface {
	Create()
	Update()
	Delete()
	Read() string
}
