package Product

// Data and Model Related
type (

	// Product : represent data model of product
	Product struct {
		ID          int64   `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Stock       int     `json:"stock"`
		Price       float64 `json:"price"`
		WarehouseId string  `json:"warehouse_id"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
	}

	SearchParam struct {
		ID   string `json:"id"`
		Q    string `json:"q"`
		Pmin string `json:"pmin"`
		Pmax string `json:"pmax"`
		Page int64
		Size int64
	}
)

// Usecase Related Data Type
type (
	// Response for fetching product
	Response struct {
		Header Header `json:"header"`
		Data   []Data `json:"data"`
	}

	// Data for fetching product
	Data struct {
		ID          int64   `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Stock       int     `json:"stock"`
		Price       float64 `json:"price"`
		WarehouseId string  `json:"warehouse_id"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
	}

	Header struct {
		ProcessTime float64 `json:"process_time"`
		StatusCode  int     `json:"status_code"`
		Message     string  `json:"message"`
	}
)
