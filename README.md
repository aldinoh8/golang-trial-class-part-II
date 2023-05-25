# Golang-Trial-Class-Part-II
# Trial Class: Mini Ecommerce

## Setup Database
- Buatlah sebuah database dengan nama `trial-class-go`
- Jalankan sql query pada file `migrations/create_tables.sql` lalu `migrations/seed_products.sql`
- Ubah dsn database pada config/db.go line 14, sesuaikan dengan confgi database dan database yang telah dibuat
```go
// contoh untuk db dengan nama trial-class-go
host=localhost user=postgres password=postgres dbname=trial-class-go port=5432
```

## How to Setup Server API
## Initialization
### **init go mod project**
```bash
go mod init [project-name]
```
### **install packages**
```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/driver/postgres
go get -u gorm.io/gorm
```
### **init gin server**, di func main di file main.go
```go
func main() {
  r := gin.Default()
  
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```
untuk memastikan server sudah berjalan, bisa coba akses `localhost:8080/ping`, jika sudah muncul "pong", berarti server sudah berjalan dengan semestinya

### **Membuat endpoint**
```go
r.GET("/example", func(c *gin.Context) {
	// logic di dalam endpoint.
	c.JSON(http.StatusOK, gin.H{
      "message": "OK",
    })
})
```


### **Setup aplikasi dan database**
untuk menghubungkan aplikasi dengan database jangan lupa isi dsn untuk db postgresql, di bagian `config/db.go`. dan menjalankan function `DBConnect` ketika aplikasi dijalankan, dengan 
cara menambahkan perintah `config.DBConnect()` di func main di `main.go`

```go
func main() {
  r := gin.Default()
  config.DBConnect()

  // setup endpoint etc
  
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

memberi response/output pada endpoint, dapat menggunakan function dari gin.Context yaitu JSON
```go
// parameter pertama adalah status kode
// parameter kedua adalah response/outputnya
ctx.JSON(200, "hello world")
```

menerima input dari request body menggunakan method ShouldBindJson. Kita perlu mendefinisikan struct dari request body yang diinginkan
```go
// buat struct body data yang diinginkan
type Body struct {
	Name string `json:"name"`
	Age string `json:"age"`
}

// di dalam handler endpoint
var body Body
err := ctx.ShouldBindJson(&body)
```
dengan menggunakan method ShouldBindJson diatas, value dari request body akan di-bind(dimasukkan) ke dalam variable body yang sudah di-define. Dan selanjutnya body tadi bisa digunakan untuk flow logic berikutnya, seperti insert data ke database

### **Contoh untuk GET products**
```go
r.GET("/example", func(c *gin.Context) {
	// entity.Product sesuai dengan definisi yang sudah dibuat di class part I
	var products []entity.Product

	// config & find product sesuai dengan yang sudah dibuat di class part I
	if err := config.DB.Find(&products).Error; err != nil {
		// response error ketika terdapat error
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response ok dan data products
	ctx.JSON(http.StatusOK, products)
})
```


### **Dokumentasi & Swagger**
untuk membuat dokumentasi dengan swagger. kita perlu meng-install package swagger
```bash
go get -u github.com/swaggo/swag/cmd/swag  
go get -u github.com/swaggo/gin-swagger  
go get -u github.com/swaggo/files  
go get -u github.com/alecthomas/template
```
selanjutnya kita perlu menambahkan comment pada file main.go kita berupa informasi-informasi terkait aplikasi kita yang akan ditampilkan di dokumentasi swagger
```go
// @title           Trial Class Mini Ecommerce
// @version         1.0
// @description     Dokomentasi REST API project Mini Ecommerce II

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:8000
```

Untuk dapat meng-akses dokumentasi yang dibuat oleh swagger kita perlu menambahkan route
```go
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```
dan menjalankan perintah `swag ini` pada terminal

Untuk membuat dokumentasi per-endpoint(bagaimana request response dll) kita perlu menambahkan comment untuk setiap endpoint handler kita
```go
// @Summary Get Product
// @Schemes Product
// @Description Get list of all available Products
// @Tags Product
// @Produce json
// @Success 200 {array} entity.Product
// @Router /products [get]
func GetProductHandler(ctx *gin.Context) {
	var products []entity.Product
	if err := config.DB.Find(&products).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
```

