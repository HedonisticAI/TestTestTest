package c_ud

type UserInfo struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Age        int    `json:"age,omitempty"`
	Nation     string `json:"nation,omitempty"`
}

type AgeRequest struct {
	Age int `json:"age"`
}

type Country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}
type NationRequest struct {
	//Count   int    `json:"count"`
	//Name    string `json:"name"`
	Country []Country `json:"country"`
}

type GenderRequst struct {
	Gender string `json:"gender"`
	// Name        string  `json:"name"`
	// Count       int     `json:"count"`
	// Probability float32 `json:"probability"`
}
