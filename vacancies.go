package parser

type Vacancy struct {
	Id            int    `json:"id"`
	Url           string `json:"url"`
	Name          string `json:"name"`
	Min_payment   int    `json:"min_payment"`
	Max_payment   int    `json:"max_payment"`
	Description   string `json:"description"`
	Is_active     bool   `json:"is_active"`
	Employer_name string `json:"employer_name"`
}

type GetVacancies struct {
	Name    string `json:"name" validate:"optional"`
	Company string `json:"company" validate:"optional"`
	Salary  int    `json:"salary" validate:"optional"`
}

type SearchVacancies struct {
	Name     string `json:"name" validate:"optional"`
	Company  string `json:"company" validate:"optional"`
	Salary   int    `json:"salary" validate:"optional"`
	IsChosen bool   `json:"is_chosen" binding:"required"`
}

type VacancyAnalitics struct {
	Count              int    `json:"count"`
	MostCompany        string `json:"most_company"`
	CountOfMostCompany int    `json:"count_of_most_company"`
	AvaragePayment     int    `json:"avarage_payment"`
}
