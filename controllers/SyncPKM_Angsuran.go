package controllers

import (
	"beego/conf"
	"encoding/json"
	"fmt"
	"strconv"
)

// "PKM_mekaar/conf"
// "strings"

type GetResponse struct {
	ResponseCode        string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
}

type Hasil1 struct {
	Hasil string
}
type Hasil2 struct {
	Hasil string
	Code  string
}

type SyncPKM_Angsuran struct {
	OurBranchID      string  `json:"OurBranchID"`
	AccountID        string  `json:"AccountID"`
	DebitInstallment float64 `json:"DebitInstallment"`
	TrxDate          string  `json:"TrxDate"`
}

type ParamBody struct {
	OurBranchID string `json:"OurBranchID"`
	GroupID     string `json:"GroupID"`
	TrxDate     string `json:"TrxDate"`
	Kode        string `json:"Kode"`
}

func (c *MainController) SyncPKM_Angsuran() GetResponse {
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	var dt ParamBody
	err := decoder.Decode(&dt)

	if err != nil {
		conf.ErrorLog("error decode : " + err.Error())
	}
	//fmt.Println(dt)

	//ADD DATA MENTAH DARI 226
	data_sync := conf.DbPKM.QueryRow(`EXEC PPP_SYNC_ANGSURAN @OurBranchID=?,@GroupID=?,@TrxDate=?,@Kode=?`, dt.OurBranchID, dt.GroupID, dt.TrxDate, dt.Kode)

	var x Hasil1
	errx := data_sync.Scan(&x.Hasil)

	fmt.Println(x.Hasil)
	if errx != nil {
		fmt.Println(errx)
		c.Ctx.ResponseWriter.WriteHeader(400)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		var response = GetResponse{
			ResponseCode:        "0" + statusCode,
			ResponseDescription: "Bad Request",
		}
		c.Data["json"] = response

		c.ServeJSON()

		conf.ErrorLog("error query : " + err.Error())

		return response
	}

	//fmt.Println(x.Hasil)
	//SYNC ke PNM_LIVE

	rows := conf.DbBRNET.QueryRow(`EXEC syncPKM_Angsuran @JSON='` + string(x.Hasil) + `'`)

	var y Hasil2
	erry := rows.Scan(&y.Hasil, &y.Code)

	if erry != nil {
		fmt.Println(erry)
		c.Ctx.ResponseWriter.WriteHeader(400)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		var response = GetResponse{
			ResponseCode:        "0" + statusCode,
			ResponseDescription: "Bad Request",
		}
		c.Data["json"] = response

		c.ServeJSON()

		conf.ErrorLog("error exec : " + err.Error())

		return response
	}
	//BALIK ke PKMMobile
	fmt.Println(y.Code)
	//fmt.Println(y.Hasil)
	balik := conf.DbPKM.QueryRow(`EXEC PPP_BalikanSync_Angsuran  @JSON='`+string(y.Hasil)+`',@Code='`+string(y.Code)+`',@OurBranchID=?,@GroupID=?,@TrxDate=?,@Kode=?`, dt.OurBranchID, dt.GroupID, dt.TrxDate, dt.Kode)

	var a string
	errBalik := balik.Scan(&a)

	if errBalik != nil {
		fmt.Println(errBalik)
		c.Ctx.ResponseWriter.WriteHeader(400)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		var response = GetResponse{
			ResponseCode:        "0" + statusCode,
			ResponseDescription: "Bad Request",
		}
		c.Data["json"] = response

		c.ServeJSON()

		conf.ErrorLog("error exec : " + err.Error())

		return response
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	statcod := c.Ctx.ResponseWriter.Status
	statusCode := strconv.Itoa(statcod)

	var response = GetResponse{
		ResponseCode:        "0" + statusCode,
		ResponseDescription: a,
	}
	c.Data["json"] = response

	c.ServeJSON()

	return response
}
