package objects_find_nearest_infrastructure

import (
	"github.com/paulmach/orb"

	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/osm"
)

func paramsToPoint(params api.ObjectsFindNearestInfrastructureParams) orb.Point {
	return orb.Point{
		params.Lon.Value,
		params.Lat.Value,
	}
}

func paramsToObjTypes(params api.ObjectsFindNearestInfrastructureParams) []osm.ObjType {
	objTypes := make([]osm.ObjType, 0, len(params.ObjectTypes))
	for _, ot := range params.ObjectTypes {
		switch ot {
		case api.ObjectsFindNearestInfrastructureObjectTypesItemBusStop:
			objTypes = append(objTypes, osm.BusStop)
		case api.ObjectsFindNearestInfrastructureObjectTypesItemHospital:
			objTypes = append(objTypes, osm.Hospital)
		case api.ObjectsFindNearestInfrastructureObjectTypesItemKindergarten:
			objTypes = append(objTypes, osm.Kindergarten)
		case api.ObjectsFindNearestInfrastructureObjectTypesItemShops:
			objTypes = append(objTypes, osm.Shops)
		case api.ObjectsFindNearestInfrastructureObjectTypesItemSport:
			objTypes = append(objTypes, osm.Sport)
		}
	}
	return objTypes
}
