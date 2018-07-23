package repository

import(
	
	"github.com/traffix/entity"

)

type RouteRepository interface{

	FetchDuration(id string)(string,error)

	Save(route *entity.Route)(bool,error)

	Delete(id string)(bool,error)

	Update(route *entity.Route)(bool,error)

	FetchRoute(id string)(route *entity.Route,error)

}