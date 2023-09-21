package enum

type Gender string

func (s Gender) String() string {
	return string(s)
}

const (
	MALE   Gender = "male"
	FEMALE Gender = "female"
)
