package model

type Check struct {
	Id     int
	Url_id int
	Count  int
}

type CheckResponse struct {
	Url   string
	Count int
}

func NewCheck(url_id int) Check {
	return Check{0, url_id, 0}
}
