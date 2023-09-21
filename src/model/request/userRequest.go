package request

import "github.com/adeemgoogle/gowork/src/model/enum"

type RqUser struct {
	Gender   enum.Gender `json:"gender"`
	Location string      `json:"location"`
}
