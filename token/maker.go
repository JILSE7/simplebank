package token

import "time"

//Maker is an interface for maniging tokens

type Maker interface {
	//CreateToken a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
