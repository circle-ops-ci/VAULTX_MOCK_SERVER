// Copyright (c) 2018-2019 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package api

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

type ErrorCodeResponse struct {
	ErrMsg  string `json:"error,omitempty"`
	ErrCode int    `json:"error_code,omitempty"`
}

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Language string `json:"locale"`
}

type RegisterUserResponse struct {
	Email string `json:"email"`
}

type PairDeviceResponse struct {
	OrderID int64  `json:"order_id"`
	URL     string `json:"url"`
}

type SetupPINResponse struct {
	OrderID int64 `json:"order_id"`
}

type RepairDeviceRequest struct {
	Token     string `json:"token,omitempty"`
	VerifyNum int64  `json:"verify_num,omitempty"`
}

type RepairDeviceResponse struct {
	OrderID int64  `json:"order_id,omitempty"`
	Token   string `json:"token,omitempty"`
	URL     string `json:"url,omitempty"`
}

type LoginVerifyRequest struct {
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	ExpiresAt int64  `json:"expires_at"`
}

type LoginVerifyResponse struct {
	OrderID   int64 `json:"order_id"`
	ExpiresAt int64 `json:"expires_at"`
}

type GetUserResponse struct {
	UserEmail    string `json:"user_email"`
	CompanyID    int64  `json:"company_id"`
	IsPairDevice bool   `json:"is_pair_device"`
	IsSetupPin   bool   `json:"is_setup_pin"`
	IsDoBackup   bool   `json:"is_do_backup"`
	Wallets      []struct {
		Type    string `json:"type"`
		Address string `json:"address"`
	} `json:"wallets"`
}

type BackupWalletRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type BackupWalletResponse struct {
	OrderID int64  `json:"order_id"`
	Token   string `json:"token"`
}

type UploadWalletBackupFileResponse struct {
	Token    string `json:"token"`
	Question string `json:"question"`
}

type RestoreWalletsRequest struct {
	Token    string `json:"token"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type RestoreWalletsResponse struct {
	Token string `json:"token"`
}

type SignMessageRequest struct {
	Message string `json:"message"`
}

type CommonResponse struct {
	OrderID int64 `json:"order_id"`
}

type SignTransactionRequest struct {
	To       *string `json:"to"`
	GasLimit uint64  `json:"gas_limit"`
	GasPrice int64   `json:"gas_price"`
	Value    int64   `json:"value"`
	Input    string  `json:"input"`
	Private  bool    `json:"private"`
}

type GetOrderStatusRequest struct {
	OrderIDs []int64 `json:"order_ids"`
}

type OrderStatus struct {
	IsExist        bool                   `json:"is_exist"`
	OrderID        int64                  `json:"order_id"`
	BehaviorType   int                    `json:"behavior_type"`
	BehaviorResult int                    `json:"behavior_result"`
	Addon          map[string]interface{} `json:"addon"`
}

type GetOrderStatusResponse struct {
	OrderStatus []OrderStatus `json:"order_status"`
}

const (
	BehaviorTypeLogin         = 1
	BehaviorTypeSignRawTx     = 2
	BehaviorTypeSignSignature = 3
	BehaviorTypePairedDevice  = 4
	BehaviorTypeSetupPIN      = 5
	BehaviorTypeBackup        = 6
)

const (
	BehaviorResultPending = 0
	BehaviorResultReject  = 1
	BehaviorResultAccept  = 2
	BehaviorResultExpired = 3
	BehaviorResultFailed  = 4
)

type CallbackStruct struct {
	OrderID        int64  `json:"order_id"`
	CompanyID      int64  `json:"company_id"`
	BehaviorType   int    `json:"behavior_type"`
	BehaviorResult int    `json:"behavior_result"`
	Input          string `json:"input"`
	Output         string `json:"output"`
}

type GetNonceRequest struct {
	Address   string `json:"address"`
	CompanyID int64  `json:"company_id"`
}

type GetNonceResponse struct {
	Nonce int64 `json:"nonce"`
}

func RegisterUser(request *RegisterUserRequest, qs []string) (response *RegisterUserResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequest("POST", "/v1/vaultx/users", qs, jsonRequest, false)
	if err != nil {
		return nil, err
	}

	response = &RegisterUserResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("RegisterUser() => ", response)
	return
}

func PairDevice(qs []string) (response *PairDeviceResponse, err error) {
	resp, err := makeRequest("POST", "/v1/vaultx/devices", qs, nil, false)
	if err != nil {
		return nil, err
	}

	response = &PairDeviceResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("PairDevice() => ", response)
	return
}

func RepairDevice(request *RepairDeviceRequest, qs []string) (response *RepairDeviceResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequest("POST", "/v1/vaultx/devices/repair", qs, jsonRequest, false)
	if err != nil {
		return nil, err
	}

	response = &RepairDeviceResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("RepairDevice() => ", response)
	return
}

func SetupPIN(qs []string) (response *SetupPINResponse, err error) {
	resp, err := makeRequest("POST", "/v1/vaultx/users/pin", qs, nil, false)
	if err != nil {
		return nil, err
	}

	response = &SetupPINResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("SetupPIN() => ", response)
	return
}

func LoginVerify(request *LoginVerifyRequest, qs []string) (response *LoginVerifyResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequest("POST", "/v1/vaultx/loginverify", qs, jsonRequest, false)
	if err != nil {
		return nil, err
	}

	response = &LoginVerifyResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("LoginVerify() => ", response)
	return
}

func GetUser(qs []string) (response *GetUserResponse, err error) {
	resp, err := makeRequest("GET", "/v1/vaultx/users/me", qs, nil, false)
	if err != nil {
		return nil, err
	}

	response = &GetUserResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("GetUser() => ", response)
	return
}

func BackupWallet(request *BackupWalletRequest, qs []string) (response *BackupWalletResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequest("POST", "/v1/vaultx/wallets/backup", qs, jsonRequest, false)
	if err != nil {
		return nil, err
	}

	response = &BackupWalletResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("BackupWallet() => ", response)
	return
}

func DownloadWalletBackupFile(token string, qs []string) (fileBinary []byte, err error) {
	resp, err := makeRequest("GET", fmt.Sprintf("/v1/vaultx/wallets/backup/%s", token), qs, nil, false)
	if err != nil {
		return nil, err
	}

	logs.Debug("DownloadWalletBackupFile() receive %d bytes file", len(resp))
	return resp, nil
}

func UploadWalletBackupFile(fileBinary []byte, qs []string) (response *UploadWalletBackupFileResponse, err error) {
	resp, err := makeRequest("POST", "/v1/vaultx/wallets/upload", qs, fileBinary, true)
	if err != nil {
		return nil, err
	}

	response = &UploadWalletBackupFileResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("UploadWalletBackupFile() => ", response)
	return
}

func RestoreWallets(request *RestoreWalletsRequest, qs []string) (response *RestoreWalletsResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := makeRequest("POST", "/v1/vaultx/wallets/restore", qs, jsonRequest, false)
	if err != nil {
		return nil, err
	}

	response = &RestoreWalletsResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("RestoreWallets() => ", response)
	return
}

func SignMessage(request *SignMessageRequest, qs []string) (response *CommonResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := makeRequest("POST", "/v1/vaultx/wallets/signature", qs, jsonRequest, false)
	if err != nil {
		return nil, err
	}

	response = &CommonResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("SignMessage() => ", response)
	return
}

func SignTransaction(request *SignTransactionRequest, qs []string) (response *CommonResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := makeRequest("POST", "/v1/vaultx/wallets/rawtx", qs, jsonRequest, false)
	if err != nil {
		return nil, err
	}

	response = &CommonResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("SignTransaction() => ", response)
	return
}

func GetOrderStatus(request *GetOrderStatusRequest, qs []string) (response *GetOrderStatusResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := makeRequest("POST", "/v1/vaultx/order/status", qs, jsonRequest, false)
	if err != nil {
		return nil, err
	}

	response = &GetOrderStatusResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	logs.Debug("GetOrderStatus() => ", response)
	return
}
