package user

import "time"

const ctLayout = "2006-01-02"

type DateOfBirth struct {
	time.Time
}

func (ct *DateOfBirth) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	if s == `""` {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(`"`+ctLayout+`"`, s)
	return
}

type UpdateUserDTO struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Dob   DateOfBirth `json:"dob"`
}
