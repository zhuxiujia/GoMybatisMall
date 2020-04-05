package model

type RegionCity struct {
	Id         string `json:"id"`
	CityId     string `json:"city_id"`
	City       string `json:"city"`
	ProvinceId string `json:"province_id"`
}
