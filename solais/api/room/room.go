package room

import (

)

// Room Entity
// When inserting to redis, the key will be the room id
type Room struct {
	name: string `json:"name"`
	address: string `json:"address"`
	population: string `json:"population"`
}

// Json to Struct
func JtoS(jsonStr string) Room {
	var room Room
	json.Unmarshal([]byte(jsonStr), &user)
	return room
}

// Struct to Json
func StoJ(room Room) []byte {
	jsonBytes, err := json.Marshal(room)
	if err != nil {
					log.Error.Fatal("redis set failed:", err)
				}
	return jsonBytes
}

func new() *Room {
	&Room {

	}
}

// The lifespan of the room will be set based on the token, right now it's 7hrs
func Create() {

}