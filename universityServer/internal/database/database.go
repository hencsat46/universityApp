package database

import (
	"errors"
	"log"
	"strconv"

	"universityServer/internal/migrations"
	"universityServer/internal/models"
	dbhelp "universityServer/internal/pkg/dbHelp"
	"universityServer/internal/tools/handle"

	"gorm.io/gorm"
)

type repostitory struct {
	db *gorm.DB
}

func NewRepostitory() handle.RepostitoryInterfaces {
	db := migrations.InitDB()

	return &repostitory{db: db}
}

func (r *repostitory) ReadUniversity() []models.Universities {
	uni := make([]models.Universities, 0)

	if err := r.db.Select("Uni_name", "Uni_des", "Uni_img").Find(&uni).Error; err != nil {
		log.Println(err)
	}

	return uni

}

func (r *repostitory) GetRemain() (int64, error) {

	var remain int64

	if err := r.db.Model(&models.Universities{}).Count(&remain).Error; err != nil {
		log.Println(err)
		return -1, err
	}

	return remain, nil

}

func (r *repostitory) GetRecords() ([]models.StudentInfo, error) {

	var countQuery int64

	if err := r.db.Model(&models.Students_records{}).Count(&countQuery).Error; err != nil {
		return nil, err
	}

	recordsArr := make([]models.StudentInfo, countQuery)

	if err := r.db.Table("users").Select([]string{"users.student_name", "users.student_surname", "universities.uni_name", "students_records.student_points"}).Joins("RIGHT JOIN students_records ON users.user_id = students_records.student_id").Joins("LEFT JOIN universities ON universities.uni_id = students_records.student_university").Find(&recordsArr).Error; err != nil {
		return nil, err
	}

	return recordsArr, nil

}

func (r *repostitory) AddStudentRecord(studentName string, studentUniversity string, studentPoints string) error {
	pointsInt, err := strconv.Atoi(studentPoints)
	if err != nil {
		return errors.New("cannot convert points")
	}

	if err = addRecord(studentName, studentUniversity, pointsInt, r.db); err != nil {
		return err
	}

	return nil
}

func addRecord(username string, universityName string, points int, database *gorm.DB) error {
	student := models.Users{}
	university := models.Universities{}
	var hasStudent int64
	var recordId int

	recordName, err := dbhelp.GetRecordId()
	if err != nil {
		return err
	}

	if err := database.Model(&models.Users{}).Where(&models.Users{Username: username}).Find(&student).Error; err != nil {
		return err
	}

	if err := database.Where(&models.Universities{Uni_name: universityName}).Find(&university).Error; err != nil {
		return err
	}

	if err := database.Model(&models.Students_records{Student_id: int(student.User_id), Student_university: int(university.Uni_id)}).Count(&hasStudent).Error; err != nil {
		return err
	}

	if hasStudent > 0 { // student exist

		if err := database.Model(&models.Students_records{}).Where(&models.Students_records{Student_id: int(student.User_id)}).Select(recordName).Find(&recordId).Error; err != nil {
			return err
		}

		if err := database.Save(&models.Students_records{Record_id: uint(recordId), Student_id: int(student.User_id), Student_university: int(university.Uni_id), Student_points: points}).Error; err != nil {
			return err
		}
		return nil
	}

	if err := database.Create(&models.Students_records{Student_id: int(student.User_id), Student_university: int(university.Uni_id), Student_points: points}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repostitory) GetStatus() (bool, error) {

	var result bool

	if err := r.db.Model(&models.Records_status{}).Select("status").Find(&result).Error; err != nil {
		return false, err
	}
	return result, nil

}

func (r *repostitory) GetId(username string) (uint, error) {

	var userModel models.Users

	if err := r.db.Model(&models.Users{}).Where("username = ?", username).Find(&userModel).Error; err != nil {
		return 0, err
	}

	return userModel.User_id, nil

	// conn := ConnectDB()
	// defer conn.Close(context.Background())

	// var userId string

	// response, err := conn.Query(context.Background(), fmt.Sprintf("SELECT user_id FROM users WHERE username = '%s';", username))
	// if err != nil {
	// 	return "", err
	// }
	// response.Next()
	// err = response.Scan(&userId)
	// if err != nil {
	// 	return "", err
	// }
	// return userId, nil
}

func (r *repostitory) SetUser(username string, password string, studentName string, studentSurname string) error {

	user := models.Users{Username: username, Passwd: password, Student_name: studentName, Student_surname: studentSurname}
	log.Println(user)

	if status, err := checkUser(r.db, &user); err != nil {
		return err
	} else {
		if status == 1 {
			return errors.New("this username already exists")
		}
	}

	if err := r.db.Omit("User_id").Create(&user).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func checkUser(database *gorm.DB, model *models.Users) (int64, error) {

	var result int64

	if err := database.Model(&model).Where("username = ?", model.Username).Count(&result).Error; err != nil {
		return -1, err
	}

	return result, nil

}

func (r *repostitory) SignIn(username, password string) error {
	var result int64

	if err := r.db.Model(&models.Users{}).Where(&models.Users{Username: username, Passwd: password}).Count(&result).Error; err != nil {
		return err
	}

	if result != 1 {
		return errors.New("wrong username or password")
	}
	return nil

}

func (r *repostitory) GetInfoDb(username string) (models.StudentInfo, error) {

	var studentData models.StudentInfo

	if err := r.db.Model(&models.Users{}).Select([]string{"users.username", "users.student_name", "users.student_surname", "universities.uni_name"}).Where("username = ?", username).Joins("LEFT JOIN students_records ON students_records.student_id = users.user_id").Joins("LEFT JOIN universities ON universities.uni_id = students_records.student_university").Find(&studentData).Error; err != nil {
		return models.StudentInfo{}, nil
	}

	return studentData, nil

}

func (r *repostitory) GetResult() ([]models.ResultRecord, int, error) {
	var uniCount int64
	var universityName []string
	var universityId []uint

	if err := r.db.Model(&models.Universities{}).Count(&uniCount).Error; err != nil {
		return nil, 0, err
	}

	result := make([]models.ResultRecord, uniCount)

	if err := r.db.Model(&models.Universities{}).Select("uni_name").Find(&universityName).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Model(&models.Universities{}).Select("uni_id").Find(&universityId).Error; err != nil {
		return nil, 0, err
	}

	var emptyCount int

	for i := 0; i < int(uniCount); i++ {
		var studentCount int64
		var seatsCount int

		if err := r.db.Model(models.Universities{}).Select("seats_count").Where("uni_id = ?", universityId[i]).Find(&seatsCount).Error; err != nil {
			return nil, 0, err
		}

		if err := r.db.Model(models.Students_records{}).Where("student_university = ?", universityId[i]).Count(&studentCount).Error; err != nil {
			return nil, 0, err
		}

		students := make([]models.ResultStudent, studentCount)

		if err := r.db.Table("users").Limit(seatsCount).Where("student_university = ?", universityId[i]).Joins("LEFT JOIN students_records ON users.user_id = students_records.student_id").Select([]string{"users.student_name", "users.student_surname", "students_records.student_points"}).Order("students_records.student_points DESC").Find(&students).Error; err != nil {
			return nil, 0, err
		}

		if len(students) == 0 {
			emptyCount++
		}
		//log.Println(students)

		result[i].Student_university = universityName[i]
		result[i].Student_information = students
		//result[i] = models.ResultRecord{Student_university: universityName[i], }

	}

	//log.Println(result)

	// for i := 0; i < 4; i++ {
	// 	log.Println(universityName[i])
	// }
	return result, emptyCount, nil
}
