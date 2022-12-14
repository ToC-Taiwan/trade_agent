// Package routers package routers
package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"

	"trade_agent/pkg/dbagent"
	"trade_agent/pkg/log"

	"github.com/gin-gonic/gin"
)

// AddBalanceHandlersV1 AddBalanceHandlersV1
func AddBalanceHandlersV1(group *gin.RouterGroup) {
	group.GET("/balance", GetAllBalance)
	group.POST("/balance", ImportBalance)
	group.DELETE("/balance", DeletaAllBalance)
}

// GetAllBalance GetAllBalance
// @Summary GetAllBalance V1
// @tags Balance
// @accept json
// @produce json
// @success 200 {object} []dbagent.Balance
// @failure 500 {object} ErrorResponse
// @Router /v1/balance [get]
func GetAllBalance(c *gin.Context) {
	var res ErrorResponse
	allBalance, err := dbagent.Get().GetAllBalance()
	if err != nil {
		log.Get().Error(err)
		res.Response = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if len(allBalance) > 1 {
		sort.Slice(allBalance, func(i, j int) bool {
			return allBalance[i].TradeDay.Before(allBalance[j].TradeDay)
		})
	}
	c.JSON(http.StatusOK, allBalance)
}

// ImportBalance ImportBalance
// @Summary ImportBalance V1
// @tags Balance
// @accept json
// @produce json
// @param body body []dbagent.Balance{} true "Body"
// @success 200
// @failure 500 {object} ErrorResponse
// @Router /v1/balance [post]
func ImportBalance(c *gin.Context) {
	var res ErrorResponse
	body := []dbagent.Balance{}
	if byteArr, err := ioutil.ReadAll(c.Request.Body); err != nil {
		log.Get().Error(err)
		res.Response = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	} else if err := json.Unmarshal(byteArr, &body); err != nil {
		log.Get().Error(err)
		res.Response = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if len(body) > 1 {
		sort.Slice(body, func(i, j int) bool {
			return body[i].TradeDay.Before(body[j].TradeDay)
		})
	}
	for _, v := range body {
		tmp := v
		err := dbagent.Get().InsertOrUpdateBalance(&tmp)
		if err != nil {
			log.Get().Error(err)
			res.Response = err.Error()
			c.JSON(http.StatusInternalServerError, res)
			return
		}
	}
	c.JSON(http.StatusOK, nil)
}

// DeletaAllBalance DeletaAllBalance
// @Summary DeletaAllBalance V1
// @tags Balance
// @accept json
// @produce json
// @success 200
// @failure 500 {object} ErrorResponse
// @Router /v1/balance [delete]
func DeletaAllBalance(c *gin.Context) {
	var res ErrorResponse
	if err := dbagent.Get().DeleteAllBalance(); err != nil {
		log.Get().Error(err)
		res.Response = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusOK, nil)
}
