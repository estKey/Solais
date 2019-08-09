package user

import (
	"Solais/utils/log"
	"encoding/json"
)

// User Entity 
type User struct {
	id string `json:"id"`			// id => timestamp + ip address
	name string `json:"name"`
	role string `json:"role"`		// admin, holder, guest
}

// Json to Struct
func JtoS(jsonStr string) User {
	var user User
	json.Unmarshal([]byte(jsonStr), &user)
	return user
}

// Struct to Json
func StoJ(user User) []byte {
	jsonBytes, err := json.Marshal(user)
	if err != nil {
					log.Error.Fatal("redis set failed:", err)
				}
	return jsonBytes
}

func new(id string, name string, role string) *User {
	&User {
		id: id
		name: name
		role: role
	}
}

// The User will only be registered to the browser and wait until it join or create a room
func Register(id string, name string, role string) {
	user := new(id, name, role)
}

