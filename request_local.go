package yahooapi

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type JsonRequest struct {
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
