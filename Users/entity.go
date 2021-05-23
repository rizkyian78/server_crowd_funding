package Users

import "time"

type User struct {
	ID             string
	Name           string
	Occupation     string
	Email          string
	HashedPassword string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
