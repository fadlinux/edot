package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	cHttp "github/fadlinux/edot/common/http"

	util "github/fadlinux/edot/common/util"
	mUser "github/fadlinux/edot/internal/model/user"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
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

	if totalData > 0 {
		result.Header.Message = "Email and phone data already exist, please input another data"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return
	} else {

		//password convert to bcrypt
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dataReq.PasswordHash), 12)
		_, _ = d.userUC.AddUser(context.Background(), mUser.User{
			Name:         dataReq.Name,
			Phone:        dataReq.Phone,
			Email:        dataReq.Email,
			PasswordHash: string(hashedPassword),
		})

		result.Header.Message = "Success, Create user!"
		result.Header.StatusCode = 200

		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return
	}

}

// HandleUserLogin : handler for user login
func (d Delivery) HandleUserLogin(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	startTime := time.Now()
	var err error
	var result mUser.Response
	var dataReq mUser.User
	//var totalData int64

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

	//password convert to bcrypt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dataReq.PasswordHash), 12)
	result, _ = d.userUC.FetchUser(context.Background(), mUser.User{
		Phone: dataReq.Phone,
		Email: dataReq.Email,
	})
	abc := len(result.Data)
	if abc > 0 {
		fmt.Println("hashedPassword:L ", string(hashedPassword), result.Data)
		if err = bcrypt.CompareHashAndPassword([]byte(result.Data[0].Password), []byte(hashedPassword)); err != nil {
			// If the two passwords don't match, return a 401 status
			fmt.Print("wo")
		}
	}

	// if result.Data[0].Password == string(hashedPassword) {
	// 	result.Header.Message = "user logged"
	// 	result.Header.StatusCode = 200
	// } else {
	result.Header.Message = "email/phone cannot exist"
	result.Header.StatusCode = 400
	//}

	result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
	cHttp.Render(w, result, 0, req.FormValue("callback"))

}
