package contract

const IDKey = "start:id"

type IDService interface {
	NewID() string
}
