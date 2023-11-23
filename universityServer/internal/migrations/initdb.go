package migrations

import (
	"log"
	"os"

	"universityServer/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var err error
	log.Println(os.Getenv("DB_URL"))
	DB, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Println(err)
		return nil
	}

	err = DB.Migrator().AutoMigrate(&models.Users{}, &models.Records_status{}, &models.Students_records{}, &models.Universities{})

	if err != nil {
		log.Println(err)
		return nil
	}

	createUniversities(DB)

	return DB

}

func createUniversities(db *gorm.DB) {
	db.Create(&models.Universities{Uni_name: "Московский государственный университет", Uni_des: "Московский государственный университет имени М. В. Ломоносова (с 1755 по 1917 год — Императорский Московский университет) — один из старейших и крупнейших классических университетов России, один из центров российской науки и культуры, расположенный в Москве. C 1940 года носит имя Михаила Васильевича Ломоносова. Полное название — федеральное государственное бюджетное образовательное учреждение высшего образования «Московский государственный университет имени М. В. Ломоносова».", Uni_img: "img/mgu.jpg", Min_point: 280, Seats_count: 90})

	db.Create(&models.Universities{Uni_name: "Московский политехнический университет", Uni_des: "Московский политехнический университет — высшее учебное заведение в Москве. В соответствии с приказом Министерства образования и науки Российской Федерации от 21 марта 2016 года создан путём реорганизации в форме слияния двух российских вузов — МГУП им. Ивана Фёдорова и Университета машиностроения (МАМИ). Организационно-правовая форма — федеральное государственное автономное образовательное учреждение высшего образования.", Uni_img: "img/mpu.jpg", Min_point: 240, Seats_count: 100})
	db.Create(&models.Universities{Uni_name: "НИУ ВШЭ", Uni_des: "Национальный исследовательский университет «Высшая школа экономики» — автономное учреждение, федеральное государственное высшее учебное заведение. ВШЭ создана в 1992 году, нынешний статус носит с 2009 года. Основной кампус находится в Москве, ещё три — в Санкт-Петербурге, Нижнем Новгороде и Перми.", Uni_img: "img/hse.jpg", Min_point: 290, Seats_count: 40})
	db.Create(&models.Universities{Uni_name: "Московский государственный технический университет им. Н. Э. Баумана", Uni_des: "Московский государственный технический университет им. Н. Э. Баумана — российский национальный исследовательский университет, научный центр, особо ценный объект культурного наследия народов России", Uni_img: "img/mgtu.jpg", Min_point: 290, Seats_count: 40})
	// DB.Create(&models.Universities{Uni_name: "", Uni_des: "", Uni_img: "", Min_point: , Seats_count: })
}
