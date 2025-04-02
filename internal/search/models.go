package search

type UserInfo struct {
	ID         int    `json:"ID"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Age        int    `json:"age,omitempty"`
	Nation     string `json:"nation,omitempty"`
}

type SearchResponse struct {
	UserInfo []UserInfo `json:"users"`
}
