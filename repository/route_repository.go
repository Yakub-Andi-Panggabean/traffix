package repository

import(
	
	"github.com/traffix/entity"

)

type RouteRepository interface{

	FetchDuration(id string)(int,error)

	Save(route *entity.Route)(bool,error)

	Update(route *entity.Route)(bool,error)

	FetchRoute(id string)(*entity.Route,error)

}