// Copyright (c) 2018-2019 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/cybavo/VAULTX_MOCK_SERVER/api"
)

type OuterController struct {
	beego.Controller
}

func getQueryString(ctx *context.Context) []string {
	var qs []string
	tokens := strings.Split(ctx.Request.URL.RawQuery, "&")
	for _, token := range tokens {
		qs = append(qs, token)
	}
	return qs
}

var debugPrint = func(ctx *context.Context) {
	var params string
	qs := getQueryString(ctx)
	if qs != nil {
		params = strings.Join(qs, "&")
	}
	logs.Debug(fmt.Sprintf("Requst => %s, params: %s, body: %s", ctx.Input.URL(), params, ctx.Input.RequestBody))
}

func init() {
	beego.InsertFilter("/v1/mock/*", beego.BeforeExec, debugPrint)
}

func (c *OuterController) AbortWithError(status int, err error) {
	resp := api.ErrorCodeResponse{
		ErrMsg:  err.Error(),
		ErrCode: status,
	}
	c.Data["json"] = resp
	c.Abort(strconv.Itoa(status))
}

// @router /users [post]
func (c *OuterController) RegisterUser() {
	defer c.ServeJSON()

	var request api.RegisterUserRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.RegisterUser(&request, qs)
	if err != nil {
		logs.Error("RegisterUser failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /devices [post]
func (c *OuterController) PairDevice() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.PairDevice(qs)
	if err != nil {
		logs.Error("PairDevice failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /users/pin [post]
func (c *OuterController) SetupPIN() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.SetupPIN(qs)
	if err != nil {
		logs.Error("SetupPIN failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /devices/repair [post]
func (c *OuterController) RepairDevice() {
	defer c.ServeJSON()

	request := &api.RepairDeviceRequest{}
	if len(c.Ctx.Input.RequestBody) > 0 {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, request)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.RepairDevice(request, qs)
	if err != nil {
		logs.Error("RepairDevice failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /loginverify [post]
func (c *OuterController) LoginVerify() {
	defer c.ServeJSON()

	var request api.LoginVerifyRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.LoginVerify(&request, qs)
	if err != nil {
		logs.Error("LoginVerify failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /users/me [get]
func (c *OuterController) GetUser() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.GetUser(qs)
	if err != nil {
		logs.Error("GetUser failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /wallets/backup [post]
func (c *OuterController) BackupWallet() {
	defer c.ServeJSON()

	var request api.BackupWalletRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.BackupWallet(&request, qs)
	if err != nil {
		logs.Error("BackupWallet failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /wallets/backup/:token [get]
func (c *OuterController) DownloadWalletBackupFile() {
	defer c.ServeJSON()

	token := c.Ctx.Input.Param(":token")
	if token == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	fileBinary, err := api.DownloadWalletBackupFile(token, qs)
	if err != nil {
		logs.Error("DownloadWalletBackupFile failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = ioutil.WriteFile("./backup.dat", fileBinary, 0644)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = struct {
		Result int
	}{
		Result: 1,
	}
}

// @router /wallets/upload [post]
func (c *OuterController) UploadWalletBackupFile() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	fileContent, err := ioutil.ReadFile("./backup.dat")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	resp, err := api.UploadWalletBackupFile(fileContent, qs)
	if err != nil {
		logs.Error("UploadWalletBackupFile failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /wallets/restore [post]
func (c *OuterController) RestoreWallets() {
	defer c.ServeJSON()

	var request api.RestoreWalletsRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.RestoreWallets(&request, qs)
	if err != nil {
		logs.Error("RestoreWallets failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /wallets/signature [post]
func (c *OuterController) SignMessage() {
	defer c.ServeJSON()

	var request api.SignMessageRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.SignMessage(&request, qs)
	if err != nil {
		logs.Error("SignMessage failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /wallets/rawtx [post]
func (c *OuterController) SignTransaction() {
	defer c.ServeJSON()

	var request api.SignTransactionRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.SignTransaction(&request, qs)
	if err != nil {
		logs.Error("SignTransaction failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @router /order/status [post]
func (c *OuterController) GetOrderStatus() {
	defer c.ServeJSON()

	var request api.GetOrderStatusRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.GetOrderStatus(&request, qs)
	if err != nil {
		logs.Error("GetOrderStatus failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}
