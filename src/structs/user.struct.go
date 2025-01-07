package structs

import "time"

type UserStruct struct {
	id        int
	name      string
	birthday  string
	telephone string
	password  string

	created_at time.Time
	updated_at time.Time
	deleted_at time.Time
}
