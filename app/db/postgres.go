package db

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"strconv"
	"taoyuan_carpark/model"
	"time"

	_ "github.com/lib/pq"
)

func UpdatePostgres(parkInfo model.CarPark) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.Get("postgres.host"), viper.Get("postgres.port"), viper.Get("postgres.user"), viper.Get("postgres.password"), viper.Get("postgres.db"))

	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)

	// insert new or update old data
	stmt, err := db.Prepare("" +
		"INSERT INTO carparks (id, parkid, areaid, areaname, parkname, totalspace, surplusspace, payguide, introduction, address, wgsx, wgsy) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) " +
		"ON CONFLICT (id) DO UPDATE " +
		"SET parkid = excluded.parkid, areaid = excluded.areaid, areaname = excluded.areaname, parkname = excluded.parkname, totalspace = excluded.totalspace, surplusspace = excluded.surplusspace, payguide = excluded.payguide, introduction = excluded.introduction, address = excluded.address, wgsx = excluded.wgsx, wgsy = excluded.wgsy;")

	checkErr(err)

	res, err := stmt.Exec(parkInfo.Id, parkInfo.ParkId, parkInfo.AreaId, parkInfo.AreaName, parkInfo.ParkName, parkInfo.TotalSpace, parkInfo.SurplusSpace, parkInfo.PayGuide, parkInfo.Introduction, parkInfo.Address, parkInfo.WgsX, parkInfo.WgsY)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)

	db.Close()
}

func UpdateHistoryPostgres(parkInfo model.CarPark) {
	weekday := time.Now().Weekday()
	hour := time.Now().Hour()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.Get("postgres.host"), viper.Get("postgres.port"), viper.Get("postgres.user"), viper.Get("postgres.password"), viper.Get("postgres.db"))

	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)

	// insert new or update old data
	stmt, err := db.Prepare("" +
		"INSERT INTO history (hashkey, parkid, week, hour, totalspace, surplusspace) " +
		"VALUES ($1, $2, $3, $4, $5, $6) " +
		"ON CONFLICT (hashkey) DO UPDATE " +
		"SET hashkey = excluded.hashkey, parkid = excluded.parkid, week = excluded.week, hour = excluded.hour, totalspace = excluded.totalspace, surplusspace = excluded.surplusspace;")

	checkErr(err)

	hashKey := parkInfo.ParkId + "-" + strconv.Itoa(int(weekday)) + "-" + strconv.Itoa(int(hour))

	res, err := stmt.Exec(hashKey, parkInfo.ParkId, weekday, hour, parkInfo.TotalSpace, parkInfo.SurplusSpace)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)

	db.Close()
}

func FindPostgres(place string) []model.CarPark {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.Get("postgres.host"), viper.Get("postgres.port"), viper.Get("postgres.user"), viper.Get("postgres.password"), viper.Get("postgres.db"))

	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)

	res, err := db.Query("SELECT * FROM carparks WHERE areaname=$1", place)
	checkErr(err)

	carparks := []model.CarPark{}

	defer res.Close()
	for res.Next() {
		var carpark model.CarPark
		if err := res.Scan(&carpark.Id, &carpark.ParkId, &carpark.AreaId, &carpark.AreaName,
			&carpark.ParkName, &carpark.TotalSpace, &carpark.SurplusSpace, &carpark.PayGuide,
			&carpark.Introduction, &carpark.Address, &carpark.WgsX, &carpark.WgsY); err != nil {
			log.Error().Msgf("%s", err)
		}
		carparks = append(carparks, carpark)
	}
	return carparks
}

func FindHistoryPostgres(parkId string, hour int) []model.CarParkHistory {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.Get("postgres.host"), viper.Get("postgres.port"), viper.Get("postgres.user"), viper.Get("postgres.password"), viper.Get("postgres.db"))

	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)

	res, err := db.Query("SELECT * FROM history WHERE parkid=$1 and hour=$2", parkId, hour)
	checkErr(err)

	carparks := []model.CarParkHistory{}

	defer res.Close()
	for res.Next() {
		var carpark model.CarParkHistory
		if err := res.Scan(&carpark.HashKey, &carpark.Week, &carpark.Hour,
			&carpark.TotalSpace, &carpark.SurplusSpace, &carpark.ParkId); err != nil {
			log.Error().Msgf("%s", err)
		}
		carparks = append(carparks, carpark)
	}
	return carparks
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
