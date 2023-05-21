package models

type ComplainFilters struct {
	*GenericFilters
	ComplainID   string
	StudentID    string
	HostelID     string
	ComplainType string
}

type GenericFilters struct {
	Limit int
	Page  int
}
