package router

import (
	"net/http"
	"strconv"

	"github.com/astawan/go-api-sore/core"
	"github.com/astawan/go-api-sore/service/entities"
	"github.com/gin-gonic/gin"
)

type Router struct {
	gin *gin.Engine
}

type RouterContract interface {
	NewRouter() http.Handler
}

func RouterConstructor(gin *gin.Engine) RouterContract {
	return &Router{
		gin: gin,
	}
}

func (r *Router) NewRouter() http.Handler {

	r.gin.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": true,
			"msg":    "Hello world",
		})
	})

	// mengambil semua data buku
	// GET http://localhost:3001/bukus
	r.gin.GET("/bukus", func(c *gin.Context) {
		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().
			Joins("LEFT JOIN Penulis ON Buku.penulisId = Penulis.id").
			Select(
				"Buku.id",
				"Buku.name",
				"Buku.penulisId",
				"Penulis.name AS penulisName",
			)

		var data []entities.Buku

		q.Find(&data)

		// mengembalikan respon berhasil
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	// mengambil satu data buku berdasarkan ID buku
	// GET http://localhost:3001/buku?id=1
	r.gin.GET("/buku", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id")) // get ID dan convert ke integer

		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().
			Joins("LEFT JOIN Penulis ON Buku.penulisId = Penulis.id").
			Where("Buku.id = ?", &id).
			Select(
				"Buku.id",
				"Buku.name",
				"Buku.penulisId",
				"Penulis.name AS penulisName",
			)

		var data []entities.Buku

		q.Find(&data)

		// mengembalikan respon berhasil
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	// menambah satu data buku berdasarkan input
	// contoh akses API untuk menambah data baru
	//
	// POST http://localhost:3001/buku
	// Content-Type: application/json
	//
	// {
	// 	"name": "nama buku",
	// 	"penulisId": 1
	// }
	r.gin.POST("/buku", func(c *gin.Context) {
		// menerima input dari payload API
		input := &entities.BukuInsert{}
		if err := c.ShouldBindJSON(&input); err != nil {
			// mengembalikan respon gagal
			e := err.Error()
			c.JSON(500, gin.H{
				"status":     false,
				"errMessage": &e,
			})
			return
		}

		app := core.NewApp()
		db := app.Mysql

		db.Debug().Create(&input)

		// mengembalikan respon berhasil
		c.JSON(200, gin.H{
			"status": true,
			"result": &input,
		})
	})

	// menambah banyak data buku berdasarkan input
	// contoh akses API untuk menambah data baru
	//
	// POST http://localhost:3001/bukus
	// Content-Type: application/json
	//
	// {
	// 	"data": [
	// 		{
	// 			"name": "buku 1",
	// 			"penulisId": 1
	// 		},
	// 		{
	// 			"name": "buku 2",
	// 			"penulisId": 1
	// 		},
	// 		{
	// 			"name": "buku 3",
	// 			"penulisId": 1
	// 		}
	// 	]
	// }
	// payload disesuaikan dengan atribut entitas,
	// di sini menggunakan entits BukuInsertMany dengan atribut Data dengan tipe array
	// maka payload juga harus menyediakan "data" dengan value array
	r.gin.POST("/bukus", func(c *gin.Context) {
		// menerima input dari payload API
		input := &entities.BukuInsertMany{}
		if err := c.ShouldBindJSON(&input); err != nil {
			// mengembalikan respon gagal
			e := err.Error()
			c.JSON(500, gin.H{
				"status":     false,
				"errMessage": &e,
			})
			return
		}

		app := core.NewApp()
		db := app.Mysql

		db.Debug().Create(&input.Data)

		// mengembalikan respon berhasil
		c.JSON(200, gin.H{
			"status": true,
			"result": &input.Data,
		})
	})

	// mengubah data buku berdasarkan ID buku
	// contoh akses API untuk mengubah data
	//
	// PUT http://localhost:3001/buku/1
	// Content-Type: application/json
	//
	// {
	// 	"name": "nama buku baru",
	// }
	r.gin.PUT("/buku/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) // get ID dan convert ke integer

		// menerima input dari payload API
		input := &entities.BukuInsert{}
		if err := c.ShouldBindJSON(&input); err != nil {
			// mengembalikan respon gagal
			e := err.Error()
			c.JSON(500, gin.H{
				"status":     false,
				"errMessage": &e,
			})
			return
		}

		app := core.NewApp()
		db := app.Mysql

		// mengambil entitas  yang akan diupdate berdasarkan ID
		var entity entities.Buku
		if err := db.Where("id = ?", id).First(&entity).Error; err != nil {
			e := err.Error()
			c.JSON(500, gin.H{
				"status":     false,
				"errMessage": &e,
			})
			return
		}

		// membuat variabel baru untuk menyimpan data baru dari payload yang tidak null
		// sesuaikan dengan entitas dan field database
		updateFields := make(map[string]interface{})
		if input.Name != nil {
			updateFields["name"] = input.Name
		}
		if input.PenulisID != nil {
			updateFields["penulisId"] = input.PenulisID
		}

		// mengubah entitas yang telah diambil sebelumnya dengan data baru
		if err := db.Model(&entity).Updates(updateFields).Error; err != nil {
			e := err.Error()
			// mengembalikan respon gagal
			c.JSON(500, gin.H{
				"status":     false,
				"errMessage": &e,
			})
			return
		}

		// mengembalikan respon berhasil
		c.JSON(200, gin.H{
			"status": true,
			"result": &entity,
		})
	})

	// menghapus data buku berdasarkan ID buku
	// DELETE http://localhost:3001/buku/1
	r.gin.DELETE("/buku/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) // get ID dan convert ke integer

		app := core.NewApp()
		db := app.Mysql

		// mengambil entitas  yang akan diupdate berdasarkan ID
		db.Delete(&entities.Buku{}, id)

		// mengembalikan respon berhasil
		c.JSON(200, gin.H{
			"status": true,
		})
	})

	return r.gin
}
