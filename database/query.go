package database

type ParamList struct {
	Page       int    `json:"page" valid:"Required"`
	PerPage    int    `json:"per_page" valid:"Required"`
	Search     string `json:"search,omitempty"`
	InitSearch string `json:"init_search,omitempty"`
	SortField  string `json:"sort_field,omitempty"`
}

// ResponseModelList :
type ResultList struct {
	Page         int         `json:"page"`
	Total        int         `json:"total"`
	LastPage     int         `json:"last_page"`
	DefineSize   string      `json:"define_size"`
	DefineColumn string      `json:"define_column"`
	AllColumn    string      `json:"all_column"`
	Data         interface{} `json:"data"`
	Msg          string      `json:"message"`
}
