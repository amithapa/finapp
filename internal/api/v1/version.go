package v1

import (
	"encoding/json"
)

//API for returning version
//When server starts, we set version and than use it if necessary.

// ServerVersion represents the server version
type ServerVersion struct {
	Version string `json:"version"`
}

//Marsheled JSON
var versionJSON []byte

func init() {
	var err error
	versionJSON, err = json.Marshal(ServerVersion{
		Version: config.Version,
	})
	if err != nil {
		panic(err)
	}
}