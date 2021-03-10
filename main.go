package main

import (
	"crud/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/project_gl?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err !=nil {
		fmt.Println("Koneksi ERROR")
	}
	fmt.Println("Koneksi BERHASIL")

	mhs := model.Mahasiswa{
		Nama: "Majdi",
		Nim: "200602048",
		Alamat: "Selong",
	}
	// Dosen
	data := model.Dosen{
		Nama: "Firman",
		Nidn: "0000000001",
		Alamat: "NTB",
	}
	db.Create(&data)
	// Berfungsi utk menyimpan data ke dalam database
	db.Create(&mhs)

	// utk update data pada filed nama dengan id = 1
	// db.Model(&mhs).Where("id =?", 1).Update("nama", "M" )

	// utk delete data dengan id = 2
	// db.Where("id = ?", 2).Delete(&model.Mahasiswa{})

	// query atau mengambil semua data 
	// var mahasiswa []model.Mahasiswa
	// db.Find(&mahasiswa)

	// for _, m := range mahasiswa {
	// 	fmt.Println("Nama = " +m.Nama)
	// 	fmt.Println("Nim = " +m.Nim)
	// 	fmt.Println("Alamat = " +m.Alamat)
	// 	fmt.Println("=======================")
	// }

	route := gin.Default()

	route.GET("/getmahasiswa", func (c *gin.Context)  {
		var mhs []model.Mahasiswa
		db.Find(&mhs)
		c.JSON(http.StatusOK, &mhs)	
	})

	route.Run()

}