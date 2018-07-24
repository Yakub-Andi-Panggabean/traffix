package entity

import "time"

type Route struct{

	Id string
	Name string
	Origin string
	Destination string
	CreatedTime time.Time
	UpdatedTime time.Time

}



