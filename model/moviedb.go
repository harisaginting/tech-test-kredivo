type MDBListMovie struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []MDBMovie `json:"results"`
}


type MDBMovie struct {
	Iso6391       string  `json:"iso_639_1"`
	ID            int     `json:"id"`
	Featured      int     `json:"featured"`
	Description   string  `json:"description"`
	Revenue       string  `json:"revenue"`
	Public        int     `json:"public"`
	Name          string  `json:"name"`
	UpdatedAt     string  `json:"updated_at"`
	CreatedAt     string  `json:"created_at"`
	SortBy        int     `json:"sort_by"`
	BackdropPath  string  `json:"backdrop_path,omitempty"`
	Runtime       int     `json:"runtime"`
	AverageRating float64 `json:"average_rating"`
	Iso31661      string  `json:"iso_3166_1"`
	Adult         int     `json:"adult"`
	NumberOfItems int     `json:"number_of_items"`
	PosterPath    string  `json:"poster_path,omitempty"`
}

type MDBCreateRequestTokenResponse struct {
	StatusMessage string `json:"status_message"`
	RequestToken  string `json:"request_token"`
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
}

type MDBCreateAccessTokenResponse struct {
	AccountID     string `json:"account_id"`
	AccessToken   string `json:"access_token"`
	Success       bool   `json:"success"`
	StatusMessage string `json:"status_message"`
	StatusCode    int    `json:"status_code"`
}