package park

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"taoyuan_carpark/db"
	"taoyuan_carpark/model"
)

const taoyuan_url = "http://data.tycg.gov.tw/api/v1/rest/datastore/0daad6e6-0632-44f5-bd25-5e1de1e9146f?format=json"

func UpdateCarPark() {
	resp, err := http.Get(taoyuan_url)
	if err != nil {
		fmt.Println(err)
	}
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
	log.Info("Success")

}

func Find(c *gin.Context) []model.CarPark {
	query := c.Request.URL.Query()
	place := query["place"]
	return db.FindPostgres(place[0])
}
