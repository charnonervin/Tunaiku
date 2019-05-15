package domain

import "time"

type Cicilan struct {
	MonthInstallment int       `json:"months_installment"`
	DueDate          time.Time `json:"due_date"`
	Capital          int       `json:"capital"`
	Interest         int       `json:"interest"`
	Total            int       `json:"total"`
}

type Summary struct {
	SummaryCapital  int `json:"summary_capital"`
	SummaryInterest int `json:"summary_interest"`
	SummaryTotal    int `json:"summary_total"`
}

type CicilanResponse struct {
	Cicilan []Cicilan `json:"detail"`
	Summary Summary   `json:"summary"`
}

type CicilanParam struct {
	DateParam string `json:"date"`
	Date      time.Time
	Amount    int `json:"amount"`
	Period    int `json:"period"`
}

func SimulateCicilan(input *CicilanParam) *CicilanResponse {
	result := &CicilanResponse{}
	tableForInterest := make(map[int]float32)
	tableForInterest[12] = 0.0168
	tableForInterest[18] = 0.0168
	tableForInterest[24] = 0.0159
	tableForInterest[30] = 0.0159
	tableForInterest[36] = 0.0159

	for i := 0; i < input.Period; i++ {
		tempCicilan := Cicilan{
			MonthInstallment: i + 1,
			DueDate:          input.Date.AddDate(0, 1, 0),
			Capital:          input.Amount / input.Period,
			Interest:         int(tableForInterest[input.Period] * float32(input.Amount)),
			Total:            (input.Amount / input.Period) + int(tableForInterest[input.Period]*float32(input.Amount)),
		}
		result.Cicilan = append(result.Cicilan, tempCicilan)
		result.Summary.SummaryCapital += tempCicilan.Capital
		result.Summary.SummaryInterest += tempCicilan.Interest
		result.Summary.SummaryTotal += tempCicilan.Total
		input.Date = input.Date.AddDate(0, 1, 0)
	}

	return result
}
