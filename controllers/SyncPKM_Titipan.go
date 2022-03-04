package controllers

import (
	"beego/conf"
	"encoding/json"
	"fmt"
	"strconv"
)

// "PKM_mekaar/conf"
// "strings"

type Titipan_GetResponse struct {
	ResponseCode        string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
}

type Titipan_Hasil1 struct {
	Hasil string
}
type Titipan_Hasil2 struct {
	Hasil string
	Code  string
}

type SyncPKM_Titipan struct {
	OurBranchID      string  `json:"OurBranchID"`
	AccountID        string  `json:"AccountID"`
	DebitInstallment float64 `json:"DebitInstallment"`
	TrxDate          string  `json:"TrxDate"`
}

type Titipan_ParamBody struct {
	OurBranchID string `json:"OurBranchID"`
	GroupID     string `json:"GroupID"`
	TrxDate     string `json:"TrxDate"`
	Kode        string `json:"Kode"`
}

func (c *MainController) SyncPKM_Titipan() Titipan_GetResponse {
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	var dt Titipan_ParamBody
	err := decoder.Decode(&dt)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dt)

	//ADD DATA MENTAH DARI 226
	data_sync := conf.DbPKM.QueryRow(`EXEC PPP_SYNC_Titipan @OurBranchID=?,@GroupID=?,@TrxDate=?`, dt.OurBranchID, dt.GroupID, dt.TrxDate)

	var x Titipan_Hasil1
	errx := data_sync.Scan(&x.Hasil)

	fmt.Println(x.Hasil)
	if errx != nil {
		fmt.Println(errx)
		c.Ctx.ResponseWriter.WriteHeader(400)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		var response = Titipan_GetResponse{
			ResponseCode:        "0" + statusCode,
			ResponseDescription: "Bad Request",
		}
		c.Data["json"] = response

		c.ServeJSON()

		return response
	}

	//fmt.Println(x.Hasil)
	//SYNC ke PNM_LIVE

	rows := conf.DbBRNET.QueryRow(`EXEC syncPKM_Titipan @JSON='` + string(x.Hasil) + `'`)

	var y Titipan_Hasil2
	erry := rows.Scan(&y.Hasil, &y.Code)

	if erry != nil {
		fmt.Println(erry)
		c.Ctx.ResponseWriter.WriteHeader(400)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		var response = Titipan_GetResponse{
			ResponseCode:        "0" + statusCode,
			ResponseDescription: "Bad Request",
		}
		c.Data["json"] = response

		c.ServeJSON()

		return response
	}
	//BALIK ke PKMMobile
	fmt.Println(y.Code)
	fmt.Println(y.Hasil)
	balik := conf.DbPKM.QueryRow(`EXEC PPP_BalikanSync_Titipan @JSON='`+string(y.Hasil)+`',@Code='`+string(y.Code)+`',@OurBranchID=?,@GroupID=?,@TrxDate=?`, dt.OurBranchID, dt.GroupID, dt.TrxDate)

	var a string
	errBalik := balik.Scan(&a)

	if errBalik != nil {
		fmt.Println(errBalik)
		c.Ctx.ResponseWriter.WriteHeader(400)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		var response = Titipan_GetResponse{
			ResponseCode:        "0" + statusCode,
			ResponseDescription: "Bad Request",
		}
		c.Data["json"] = response

		c.ServeJSON()

		return response
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	statcod := c.Ctx.ResponseWriter.Status
	statusCode := strconv.Itoa(statcod)

	var response = Titipan_GetResponse{
		ResponseCode:        "0" + statusCode,
		ResponseDescription: a,
	}
	c.Data["json"] = response

	c.ServeJSON()

	return response
}
