package dto

import "github.com/adeemgoogle/gowork/src/model/enum"

type UserDto struct {
	Id        int64          `json:"id"`
	DeviceId  string         `json:"deviceId"`
	Gender    enum.Gender    `json:"gender"`
	Locations []*LocationDto `json:"location"`
}
