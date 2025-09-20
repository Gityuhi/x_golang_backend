package model

import "time"

type User struct {
    UserID         int32
    Email          string
    HashedPassword string
    CreatedAt      time.Time
}