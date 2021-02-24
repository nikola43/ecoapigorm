package utils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func Message(status int, message string) (map[string]interface{}) {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ReadHttpRequestIntegerParam(w http.ResponseWriter, r *http.Request, key string) (val uint) {
	vars := mux.Vars(r)
	param, err := strconv.Atoi(vars[key])
	if err != nil {
		RespondHttpError(w, http.StatusBadRequest, "Invalid ID")
		return
	} else {
		return uint(param)
	}
}

func ReadHttpRequestStringParam(w http.ResponseWriter, r *http.Request, key string) (val string) {
	vars := mux.Vars(r)
	return vars[key]
}

func DecodeHttpRequestPayload(w http.ResponseWriter, r *http.Request, o interface{}) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&o); err != nil {
		RespondHttpError(w, http.StatusUnprocessableEntity, "Invalid resquest payload")
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()
}

func RespondHttpError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func RespondHttpRequest(w http.ResponseWriter, err error, o interface{}) {
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondHttpError(w, http.StatusNotFound, "Not found")
		default:
			RespondHttpError(w, http.StatusUnauthorized, err.Error())
		}
	} else {
		if o == nil {
			RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
		} else {
			RespondWithJSON(w, http.StatusOK, o)
		}
	}
}

func PostRequest(url string, jsonData string, token string) string {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func GetRequest(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		return string(contents)
	}
	return "error"
}

func PutRequest(url string, jsonData string) string {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
