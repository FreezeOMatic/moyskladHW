package employers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BaseApiURL = "https://online.moysklad.ru/api/remap/1.2/entity/employee"

func CreateEmployee(employee Employee, token string) error {

	url := BaseApiURL

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(employee); err != nil {
		return fmt.Errorf("body encode error: %s", err.Error())
	}
	fmt.Println("Connecting to API ... ")
	fmt.Println("Sending Request ...")

	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return fmt.Errorf("request prepare error: %s", err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do request error: %s", err.Error())
	}
	fmt.Println("Almost done ...")
	fmt.Println(res.StatusCode)
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return fmt.Errorf("wrong response status code: %d", res.StatusCode)
	}
	body2, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("response body read error: %s", err.Error())
	}

	fmt.Println("Completed, response is on the way ...")

	fmt.Println(string(body2))

	return nil
}

func GetEmpList(token string) (result GetEmployers, err error) {
	url := BaseApiURL

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GetEmployers{}, fmt.Errorf("request prepare error: %s", err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+token)

	fmt.Println("Loading employee list ...")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return GetEmployers{}, fmt.Errorf("do request error: %s", err.Error())
	}

	fmt.Println(res.StatusCode)
	if res.StatusCode >= 300 {
		return GetEmployers{}, fmt.Errorf("wrong response status code: %d", res.StatusCode)
	}

	Result := GetEmployers{}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return GetEmployers{}, fmt.Errorf("response body read error: %s", err.Error())
	}
	err = json.Unmarshal(body, &Result)
	if err != nil {
		return GetEmployers{}, fmt.Errorf("response body unmarshal error: %s", err.Error())
	}
	return Result, nil
}

func ChangeEmployee(token, employeeID, firstName, lastName string) error {
	url := BaseApiURL
	newone := make([]UpdateEmployee, 0)
	var app UpdateEmployee

	app.Meta.Href = "https://online.moysklad.ru/api/remap/1.2/entity/employee/" + employeeID
	app.Meta.MetadataHref = "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata"
	app.Meta.MediaType = "application/json"
	app.Meta.Type = "employee"
	app.FirstName = firstName
	app.LastName = lastName

	newone = append(newone, app)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(newone); err != nil {
		return fmt.Errorf("body encode error: %s", err.Error())
	}

	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return fmt.Errorf("request prepare error: %s", err.Error())
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do request error: %s", err.Error())
	}

	if res.StatusCode >= 300 {
		return fmt.Errorf("wrong response status code: %d", res.StatusCode)
	}
	defer res.Body.Close()
	body2, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("response body read error: %s", err.Error())
	}

	fmt.Println(string(body2))
	return nil
}

func DeleteEmployee(token, employeeID string) error {
	url := BaseApiURL + "/delete"

	newone := make([]Delete, 0)
	var app Delete

	app.Meta.Href = "https://online.moysklad.ru/api/remap/1.2/entity/employee/" + employeeID
	app.Meta.MetadataHref = "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata"
	app.Meta.MediaType = "application/json"
	app.Meta.Type = "employee"

	newone = append(newone, app)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(newone); err != nil {
		return fmt.Errorf("body encode error: %s", err.Error())
	}

	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return fmt.Errorf("request prepare error: %s", err.Error())
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do request error: %s", err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return fmt.Errorf("wrong response status code: %d", res.StatusCode)
	}
	body2, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("response body read error: %s", err.Error())
	}

	fmt.Println(string(body2))
	return nil
}

func GetToken(login, password string) (token string, err error) {

	url := "https://online.moysklad.ru/api/remap/1.2/security/token"

	fmt.Println("Connecting to API ...")

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", fmt.Errorf("request prepare error: %s", err.Error())
	}
	fmt.Println("Sending Request ...")
	fmt.Println("Adding Headers ...")

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(login+":"+password)))

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request error: %s", err.Error())
	}

	resp := tokenResponse{}

	if res.StatusCode == 201 {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", fmt.Errorf("response body read error: %s", err.Error())
		}

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return "", fmt.Errorf("response body unmarshal error: %s", err.Error())
		}
		return resp.Token, nil
	}
	return "", fmt.Errorf("wrong response status code: %d", res.StatusCode)

}
