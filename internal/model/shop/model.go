package shop

// Data and Model Related
type (

	// Shop : represent data model of shpp
	Shop struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)

// Usecase Related Data Type
type (
	// Response for fetching shop
	Response struct {
		Header Header `json:"header"`
		Data   []Data `json:"data"`
	}

	// Data for fetching shop
	Data struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	Header struct {
		ProcessTime float64 `json:"process_time"`
		StatusCode  int     `json:"status_code"`
		Message     string  `json:"message"`
	}
)
