package horoshop

import "time"

type ExportRequest struct {
	Token  string `json:"token"`
	Parent int    `json:"parent"`
}

type ResponseData struct {
	Status   string `json:"status"`
	Response struct {
		Pages []Product `json:"pages"`
	} `json:"response"`
	Timestamp time.Time `json:"timestamp"`
}

type Categories struct {
	Pages []Page `json:"pages"`
}

type Page struct {
	ID     int   `json:"id"`
	Parent int   `json:"parent"`
	Title  Title `json:"title"`
}

type Title struct {
	RU string `json:"ru"`
	UA string `json:"ua"`
}
