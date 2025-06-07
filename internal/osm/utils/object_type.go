package utils

import (
	"github.com/samber/lo"

	"github.com/yasonofriychuk/real-estate-insight/internal/osm"
)

var tagsTypeMap = map[[2]string]osm.ObjType{
	{"amenity", "hospital"}: osm.Hospital,
	{"amenity", "clinic"}:   osm.Hospital,

	{"leisure", "sports_centre"}:  osm.Sport,
	{"leisure", "stadium"}:        osm.Sport,
	{"leisure", "pitch"}:          osm.Sport,
	{"leisure", "fitness_centre"}: osm.Sport,

	{"amenity", "kindergarten"}: osm.Kindergarten,

	{"amenity", "school"}:     osm.School,
	{"amenity", "university"}: osm.School,

	{"amenity", "bus_stop"}:               osm.BusStop,
	{"public_transport", "stop_position"}: osm.BusStop,
	{"amenity", "station"}:                osm.BusStop,
	{"railway", "subway_entrance"}:        osm.BusStop,
}

func ObjectTypeByTags(tags map[string]string) []osm.ObjType {
	var result []osm.ObjType

	if _, ok := tags["shop"]; ok {
		result = append(result, osm.Shops)
	}

	for k, v := range tagsTypeMap {
		if tags[k[0]] == k[1] {
			result = append(result, v)
		}
	}

	return lo.Uniq(result)
}
