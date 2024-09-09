package user

// Data and Model Related
type (

	// User : represent data model of user
	User struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Email        string `json:"Email"`
		Phone        string `json:"phone"`
		PasswordHash string `json:"password"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}
)

// Usecase Related Data Type
type (
	// Response for fetching user
	Response struct {
		Header Header `json:"header"`
		Data   []Data `json:"data"`
	}

	// Data for fetching user
	Data struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"Email"`
		Phone     string `json:"phone"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Password  string `json:"password"`
	}

	Header struct {
		ProcessTime float64 `json:"process_time"`
		StatusCode  int     `json:"status_code"`
		Message     string  `json:"message"`
	}
)
