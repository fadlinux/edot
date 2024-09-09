package order

type (
	SearchParam struct {
		ID   string `json:"id"`
		Q    string `json:"q"`
		Page int64
		Size int64
	}

	Order struct {
		UserID   int            `json:"user_id"`
		Products []ProductOrder `json:"products"`
	}

	ProductOrder struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
		Price     int `json:"price"`
	}

	// Response for fetching order
	Response struct {
		Header Header `json:"header"`
		Data   []Data `json:"data"`
	}

	// Data for fetching order
	Data struct {
		ID          int64   `json:"id"`
		Name        string  `json:"name"`
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
