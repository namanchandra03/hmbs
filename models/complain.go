package models

type ComplainDetails struct {
	StudentID    string `json:"studentId" db:"student_id"`
	HostelID     string `json:"hostelId" db:"hostel_id"`
	ComplainType string `json:"complainType" db:"complain_type"`
	Description  string `json:"description" db:"description"`
}

type Complain struct {
	ComplainID string `json:"complainId"`
}

type ComplainInfo struct {
	StudentName  string `json:"studentName" db:"student_name"`
	Hostel       string `json:"hostel" db:"hostel_name"`
	Floor        string `json:"floor" db:"floor"`
	Room         string `json:"room" db:"room"`
	ComplainType string `json:"complainType" db:"complain_type"`
	Description  string `json:"description" db:"description"`
}
