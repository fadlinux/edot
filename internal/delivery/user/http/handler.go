package http

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	cHttp "github/fadlinux/edot/common/http"

	util "github/fadlinux/edot/common/util"
	mUser "github/fadlinux/edot/internal/model/user"

	"github.com/julienschmidt/httprouter"
)

// HandleUserRegister : handler for user register
func (d Delivery) HandleUserRegister(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	startTime := time.Now()
	var err error
	var result mUser.Response
	var dataReq mUser.User
	var totalData int64

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&dataReq)

	result.Data = make([]mUser.Data, 0)
	email := strings.TrimSpace(dataReq.Email)
	password := strings.TrimSpace(dataReq.PasswordHash)

	if err != nil {
		result.Header.Message = err.Error()
		result.Header.StatusCode = 400
	}

	if strings.TrimSpace(dataReq.Name) == "" {
		result.Header.Message = "Name must be filled out"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return

	} else if strings.TrimSpace(dataReq.Phone) == "" {
		result.Header.Message = "Phone must be filled out"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return

	} else if strings.TrimSpace(dataReq.Email) == "" {
		result.Header.Message = "Email must be filled out"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return

	} else if strings.TrimSpace(password) == "" {
		result.Header.Message = "Password must be filled out"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return

	} else if !util.ValidEmail(email) {
		result.Header.Message = "Format email must be valid!"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return
	}

	//check exist user
	totalData, _ = d.userUC.GetUserLogin(context.Background(), mUser.User{
		Phone: dataReq.Phone,
		Email: dataReq.Email,
	})

	result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
	if totalData > 0 {
		result.Header.Message = "Email and phone data already exist, please input another data"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return
	}

	_, _ = d.userUC.AddUser(context.Background(), mUser.User{
		Name:         dataReq.Name,
		Phone:        dataReq.Phone,
		Email:        dataReq.Email,
		PasswordHash: getMD5Hash(password),
	})

	result.Header.Message = "Success, Create user!"
	result.Header.StatusCode = 200

	cHttp.Render(w, result, 0, req.FormValue("callback"))

}

// HandleUserLogin : handler for user login
func (d Delivery) HandleUserLogin(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	startTime := time.Now()
	var err error
	var result mUser.Response
	var dataReq mUser.User

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&dataReq)

	result.Data = make([]mUser.Data, 0)
	email := strings.TrimSpace(dataReq.Email)

	if err != nil {
		result.Header.Message = err.Error()
		result.Header.StatusCode = 400
	}

	if strings.TrimSpace(dataReq.Phone) == "" {
		result.Header.Message = "Phone must be filled out"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return

	} else if strings.TrimSpace(dataReq.Email) == "" {
		result.Header.Message = "Email must be filled out"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return

	} else if !util.ValidEmail(email) {
		result.Header.Message = "Format email must be valid!"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return
	}

	hashedPassword := getMD5Hash(dataReq.PasswordHash)
	result, _ = d.userUC.FetchUser(context.Background(), mUser.User{
		Phone: dataReq.Phone,
		Email: dataReq.Email,
	})

	checkData := len(result.Data)
	if checkData > 0 {
		pass := result.Data[0].Password
		if pass == hashedPassword {
			result.Header.Message = "success login"
			result.Header.StatusCode = 200
			cHttp.Render(w, result, 0, req.FormValue("callback"))
			return
		} else {
			result.Data = []mUser.Data{}
			result.Header.Message = "email/phone cannot exist"
			result.Header.StatusCode = 400
			result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
			cHttp.Render(w, result, 0, req.FormValue("callback"))

			return
		}
	} else {
		result.Data = []mUser.Data{}
		result.Header.Message = "email/phone cannot exist"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))

		return
	}

}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
