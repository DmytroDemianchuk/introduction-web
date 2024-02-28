package models

type Data struct {
	Sheets []Sheet `json:"sheets"`
}

type Sheet struct {
	Title   string   `json:"title"`
	Columns []Column `json:"columns"`
	Set     Set      `json:"set"`
}

type Column struct {
	Title  string   `json:"title"`
	Values []string `json:"values,omitempty"`
	Index  Index    `json:"index"`
}

type Index struct {
	Column int `json:"column"`
	Row    int `json:"row"`
}
