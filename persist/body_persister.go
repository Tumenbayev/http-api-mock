package persist

import (
	"github.com/jmartin82/mmock/definition"
)

//BodyPersister contains the functions to persist and read body 
type BodyPersister interface {
	//Persists Response body if needed
	Persist(*definition.Persist, *definition.Request, *definition.Response)
}