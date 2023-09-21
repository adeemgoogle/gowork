package request

import "github.com/adeemgoogle/gowork/src/model/enum"

type RqUser struct {
	Gender      enum.Gender `json:"gender"`
	LocationIds []int64     `json:"locationIds"`
}
