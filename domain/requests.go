package domain


type UserRegister struct {
	Name string `json:"name"`
	Email	string `json:"email"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address string `json:"address"`
	PlateNumber string `json:"plate_number"`
	LicenseNumber string `json:"license_number"`
	TruckCapacity int `json:"truck_capacity"`
	
}




type UserLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`

}

type AddProduct struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	Description string `json:"description"`
	Status string `json:"status"`
	Shipping bool `json:"shipping"`
}

type UpdateProduct struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	Description string `json:"description"`
	Status string `json:"status"`
	Shipping bool `json:"shipping"`
}


