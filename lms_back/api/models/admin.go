package models

type Admin struct {
	Id         string `json:"id"`
	Full_Name  string `json:"full_name"`
	Email      string `json:"email"`
	Age        uint   `json:"age"`
	Status     string `json:"status"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Created_at string `json:"create_at"`
	Updated_at string `json:"updated_at"`
}

type CreateAdmin struct {
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
	Status    string `json:"status"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

type UpdateAdmin struct {
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
	Status    string `json:"status"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

type GetAdmin struct {
	Id         string `json:"id"`
	Full_Name  string `json:"full_name"`
	Email      string `json:"email"`
	Age        uint   `json:"age"`
	Status     string `json:"status"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Created_at string `json:"create_at"`
	Updated_at string `json:"updated_at"`
}

type GetAllAdminsResponse struct {
	Admins []Admin `json:"admins"`
	Count  int16   `json:"count"`
}
type GetAllAdminsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}


type AdminKey struct {
	Id string `json:"id"`
}

type AdminReport struct {
	ID       string `json:"id"`
	FullName string `json:"fullname"`
}

type PaymentReport struct {
	ID        string  `json:"id"`
	Price     float64 `json:"price"`
	StudentID string  `json:"student_id"`
	BranchID  string  `json:"branch_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type AdminPayment struct {
	Admin   Admin   `json:"admin"`
	Payment Payment `json:"payment"`
}
