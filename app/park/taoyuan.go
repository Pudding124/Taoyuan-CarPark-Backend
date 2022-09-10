package park

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"strconv"
	"taoyuan_carpark/db"
	"taoyuan_carpark/logging"
	"taoyuan_carpark/model"
)

const taoyuan_url = "http://data.tycg.gov.tw/api/v1/rest/datastore/0daad6e6-0632-44f5-bd25-5e1de1e9146f?format=json"

func UpdateCarPark() {
	resp, err := http.Get(taoyuan_url)

	if err != nil {
		fmt.Println(err)
	}
	if resp != nil {
		defer resp.Body.Close()

		data := model.TaoyuanData{}
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
		}
		error2 := json.Unmarshal([]byte(bodyBytes), &data)
		if error2 != nil {
			fmt.Println(err)
		}

		for _, record := range data.Result.Records {
			db.UpdatePostgres(record)
		}
		logging.Print(zerolog.InfoLevel, "", "Update Real Time Success")
	} else {
		logging.Print(zerolog.InfoLevel, "", "Update Real Time Fail")
	}
}

func UpdateHistoryCarPark() {
	resp, err := http.Get(taoyuan_url)
	if err != nil {
		fmt.Println(err)
	}
	if resp != nil {
		defer resp.Body.Close()

		data := model.TaoyuanData{}
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
		}
		error2 := json.Unmarshal([]byte(bodyBytes), &data)
		if error2 != nil {
			fmt.Println(err)
		}

		for _, record := range data.Result.Records {
			db.UpdateHistoryPostgres(record)
		}
		logging.Print(zerolog.InfoLevel, "", "Update History Success")
	} else {
		logging.Print(zerolog.InfoLevel, "", "Update History Fail")
	}

}

func Find(c *gin.Context) []model.CarPark {
	query := c.Request.URL.Query()
	place := query["place"]
	return db.FindPostgres(place[0])
}

func FindHistory(c *gin.Context) []model.CarParkHistory {
	query := c.Request.URL.Query()
	parkId := query["parkId"]
	hour, _ := strconv.Atoi(query["hour"][0])
	return db.FindHistoryPostgres(parkId[0], hour)
}
