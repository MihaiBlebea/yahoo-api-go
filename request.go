package yahooapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Request interface {
	GetResponseBody(url string) (string, error)
}

type HttpRequest struct {
}

type JsonRequest struct {
}

func (r *HttpRequest) GetResponseBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (r *JsonRequest) GetResponseBody(url string) (string, error) {
	if strings.Contains(url, "earningsHistory") {
		return readFile("./example/earnings_history.json")
	}

	if strings.Contains(url, "incomeStatementHistoryQuarterly") {
		return readFile("./example/income_statement_history_quarterly.json")
	}

	if strings.Contains(url, "balanceSheetHistoryQuarterly") {
		return readFile("./example/balance_sheet_history_quarterly.json")
	}

	return "", errors.New("could not read file")
}

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
