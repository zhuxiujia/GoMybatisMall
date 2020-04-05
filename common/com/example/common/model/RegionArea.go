package model

type RegionArea struct {
	Id     string `json:"id"`
	AreaId string `json:"area_id"`
	Area   string `json:"area"`
	CityId string `json:"city_id"`
}
