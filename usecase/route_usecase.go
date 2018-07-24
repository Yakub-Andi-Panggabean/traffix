package usecase

import (
	"github.com/traffix/repository"
	log "github.com/sirupsen/logrus"
	"github.com/traffix/util"
	"github.com/traffix/entity"
)

type routeUsecase struct {
	routeRepo repository.RouteRepository
}

func NewRouteUsecase(r repository.RouteRepository) *routeUsecase {

	return &routeUsecase{
		routeRepo: r,
	}
}

func (r *routeUsecase) GetRouteCongestionLevel(id string) (util.CongestionLevel, error) {

	log.Debug(id)
	duration,err:=r.routeRepo.FetchDuration(id)
	if(err!=nil){

		return util.UNKNOWN_STATE,err

	}else{

		return util.TranslateCongestionLevel(duration), nil

	}



}

func (r *routeUsecase) RegisterRoute(route entity.Route) (bool, error) {

	log.WithFields(log.Fields{
		"route":route,
	}).Debug("register new route")

	res,err:=r.routeRepo.Save(&route)

	if(err!=nil){

		log.WithFields(log.Fields{
			"error":err.Error(),
		}).Error("registration error")
		return false,err
	}else{

		log.Debug("registration success")
		return res, nil
	}



}

func (r *routeUsecase) EditRoute(route entity.Route)(bool,error){

	log.WithFields(log.Fields{
		"route":route,
	}).Debug("update route")


	res,err:=r.routeRepo.Update(&route)


	if(err!=nil){

		log.WithField("error",err.Error()).Error("failed editing route")

		return false,err

	}else{

		log.Debug("edit route success")
		return res,nil
	}
}

func (r *routeUsecase)  FetchRoute(id string)(entity.Route,error) {

	log.WithField("route id",id).Debug("fetch route data")

	res,err:=r.routeRepo.FetchRoute(id)

	if(err!=nil){
		log.WithFields(log.Fields{
			"error":err.Error(),
			"route id":id,
		}).Error("failed to fetch route")

		return nil,err
	}else{

		log.WithField("route data",res).Debug("fetch route success")

		return res,nil
	}

}









