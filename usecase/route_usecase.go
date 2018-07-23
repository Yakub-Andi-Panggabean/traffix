package usecase

import "github.com/traffix/repository"

type routeUsecase struct {
	routeRepo repository.RouteRepository
}


func NewRouteUsecase(r repository.RouteRepository)routeUsecase{

	return &routeUsecase{
		routeRepo:r,
	}
}

func (r *routeUsecase) getRouteDuration(id string)  {



}