package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type Transaction struct {
	Date   time.Time `json:"date"`
	Type   string    `json:"type"`
	Amount float64   `json:"amount"`
}

type GraphData struct {
	Dates         []string  `json:"dates"`
	Revenues      []float64 `json:"revenues"`
	Expenses      []float64 `json:"expenses"`
	NetDifference []float64 `json:"netDifference"`
}

func processExcelFile(filePath string) (GraphData, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return GraphData{}, err
	}
	defer f.Close()

	var transactions []Transaction
	rows, err := f.GetRows("Transactions")
	if err != nil {
		return GraphData{}, err
	}

	for _, row := range rows[1:] { // Skip header row
		date, _ := time.Parse("2006-01-02", row[0])
		ht, _ := strconv.ParseFloat(row[2], 64)
		ttc, _ := strconv.ParseFloat(row[3], 64)
		tva, _ := strconv.ParseFloat(row[4], 64)
		tvaRate := parseTvaRate(row[5])

		// Calculate missing values
		if ht == 0 {
			ht = ttc / (1 + tvaRate/100)
		} else if ttc == 0 {
			ttc = ht * (1 + tvaRate/100)
		} else if tva == 0 {
			tva = ht * tvaRate / 100
		}

		amount := ht
		if row[1] == "expense" {
			amount = -amount
		}

		transactions = append(transactions, Transaction{
			Date:   date,
			Type:   row[1],
			Amount: amount,
		})
	}

	data := aggregateData(transactions)
	return data, nil
}

func parseTvaRate(rate string) float64 {
	switch strings.TrimSpace(rate) {
	case "0":
		return 0
	case "5.5":
		return 5.5
	case "10":
		return 10
	case "20":
		return 20
	default:
		return 0 // Default case, should be handled more robustly
	}
}

func aggregateData(transactions []Transaction) GraphData {
	dateMap := make(map[string]map[string]float64)

	for _, t := range transactions {
		dateStr := t.Date.Format("2006-01-02")
		if _, ok := dateMap[dateStr]; !ok {
			dateMap[dateStr] = map[string]float64{"revenue": 0, "expense": 0}
		}
		if t.Type == "revenue" {
			dateMap[dateStr]["revenue"] += t.Amount
		} else if t.Type == "expense" {
			dateMap[dateStr]["expense"] += t.Amount
		}
	}

	var dates []string
	var revenues, expenses, netDifference []float64

	for dateStr, amounts := range dateMap {
		dates = append(dates, dateStr)
		revenues = append(revenues, amounts["revenue"])
		expenses = append(expenses, amounts["expense"])
		netDifference = append(netDifference, amounts["revenue"]-amounts["expense"])
	}

	return GraphData{
		Dates:         dates,
		Revenues:      revenues,
		Expenses:      expenses,
		NetDifference: netDifference,
	}
}
