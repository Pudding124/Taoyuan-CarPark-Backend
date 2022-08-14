package model

import "container/list"

type TaoyuanData struct {
	success bool
	Result  Result
}

type Result struct {
	include_total  bool
	resource_id    string
	fields         []list.List
	records_format string
	Records        []CarPark
	offset         int
	total          int
	limit          int
}

type CarPark struct {
	Id           int `json:"_id"`
	ParkId       string
	AreaId       string
	AreaName     string
	ParkName     string
	TotalSpace   int
	SurplusSpace string
	PayGuide     string
	Introduction string
	Address      string
	WgsX         float64
	WgsY         float64
}

type CarParkHistory struct {
	HashKey      string
	ParkId       string
	Week         int
	Hour         int
	TotalSpace   int
	SurplusSpace string
}
