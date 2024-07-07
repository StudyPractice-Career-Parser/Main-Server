package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	parser "main-server"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summury SearchVacancies
// @Tags Search
// @ID parse-v
// @Param name query string false "filter1"
// @Param company query string false "filter2"
// @Param salary query int false "filter3"
// @Success 200 {object} parser.VacancyAnalitics
// @Router /search/vacancies [get]
func (h *Handler) getVacancies(c *gin.Context) {

	querys := c.Request.URL.Query()

	var salary int
	var name, company string
	salaries, ok := querys["salary"]
	if ok && len(salaries) != 0 {
		salary, _ = strconv.Atoi(salaries[0])
	}
	names, ok := querys["name"]
	if ok && len(names) != 0 {
		name = names[0]
	}
	companies, ok := querys["company"]
	if ok && len(companies) != 0 {
		company = companies[0]
	}

	fmt.Println("Input:", name, company, salary)

	searchBody := parser.SearchVacancies{
		Name:     name,
		Company:  company,
		Salary:   salary,
		IsChosen: false,
	}
	postBody, _ := json.Marshal(searchBody)

	req, _ := http.NewRequest("POST", "http://habr-parser:8001/vacancies", bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	req.Close = true

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)

	if resp == nil {
		fmt.Println("response is nil")
	}
	//resp, err := http.Post("http://127.0.0.1:8001/vacancies", "application/json", responseBody)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error occured while sending request to service with port: 8001")
		return
	}
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error occured while getting response on service with port:%d\n%s\n", 8001, err.Error())
	}

	var result parser.VacancyAnalitics
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		fmt.Printf("Error unmarshaling data from request.\n%s\n", err.Error())
	}
	fmt.Printf("Output: %#v\n", result)
	resp.Body.Close()
	// jsonResult, err := json.Marshal(result)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// }
	c.IndentedJSON(http.StatusOK, result)
}

func (h *Handler) getResume(c *gin.Context) {

}
