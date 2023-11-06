package models

type Users struct {
	User_id         uint `gorm:"primaryKey"`
	Username        string
	Passwd          string
	Student_name    string
	Student_surname string
}

type Universities struct {
	Uni_id      uint   `gorm:"primaryKey"`
	Uni_name    string `gorm:"unique"`
	Uni_des     string
	Uni_img     string
	Min_point   int
	Seats_count int
}

type Students_records struct {
	Record_id          uint `gorm:"primaryKey"`
	Student_id         int
	Student_university int
	Student_points     int
}

type Records_status struct {
	Status bool
}

type Response struct {
	Status  int
	Payload interface{}
}
