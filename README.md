# Tutorial Golang Gin Gonic

> Tutorial ini membutuhkan
>
> 1. Installed Go Languange
> 2. Code Editor (Visual Studio Code Preferred)
> 3. Postman
> 4. MongoDB Compass
>

## Konten

1. Menginstall Golang
2. Write your first Golang code
3. Golang and beyond
4. RESTAPI with Gin-Gonic
5. Folder structure for Go Backend Gin
6. Backend for Internet of Things using Gin

---

### Menginstall Golang

Menginstall golang kadang agak tricky, kita coba untuk menyederhanakan.

1. Download Go language di link [https://go.dev/doc/install](https://go.dev/doc/install)
2. Klik Install
3. Lanjutkan seperti installasi pada umumnya.

---

### Menulis kode Golang pertamamu

Seperti biasa, kita akan coba menulis kode "Hello World"

1. Buka Code Editor
2. Open folder
3. Pada terminal, ketikkan command `go mod init <nama folder>`
   Akan terbuat file `go.mod`, untuk mengidentifikasi module yang baru saja kita buat.
4. Buat file baru dengan nama `app.go`
5. Tuliskan kode berikut

   ```go
    package main

    func main() {
      print("Hello World")
    }
   ```

6. Pada terminal, ketikkan command `go run app.go`

---

### Golang dan melampauinya

#### Modules

Konsep module pada golang tidak jauh dengan konsep module pada bahasa pemrograman lain.
Langsung saja ke implementasinya.
Buat struktur folder sebagai berikut

```bash
Folder
|-function
  |-add.go
  |-multiple
    |-app.go
  |-substract
    |-app.go
|-app.go
|-go.mod
```

* `function/add.go`

  ```go
  package function

  func Add(num1 int, num2 int) int {
    return num1 + num2
  }
  ```

* `function/substract/index.go`

  ```go
  package substract

  func Substract(num1 int, num2 int) int {
    return num1 - num2
  }
  ```

* `function/multiple/index.go`

  ```go
  package function

  func Multiplie(num1 int, num2 int) int {
    return num1 * num2
  }
  ```

* `app.go`
  
  ```go
  package main

  import (
    "tutorial/function"
    multiple "tutorial/function/multiple"
    "tutorial/function/substract"
  )

  func main() {
    println(function.Add(2, 3))
    println(substract.Substract(4, 1))
    println(multiple.Multiplier(3, 10))
  }

  ```

Ada banyak cara pemanggilan module yang kita buat sendiri.

#### Menginstall module package

Di golang, penginstallan module agak lain dengan bahasa pemrograman seperti JavaScript ataupun Python.

Yaitu dengan memanggil repository dari package module dengan cara:

1. Membuka terminal
2. Tuliskan `go get -u <repository>`
   contoh: `go get -u github.com/gin-gonic/gin`

Lalu kita dapat memanggil package module tersebut dengan keyword `import`
Contoh

```go
  import "github.com/gin-gonic/gin"
```

---

### RESTAPI dengan Gin-Gonic

Okay, seperti gelar yang akan kita dapat "Sarjana Terapan", kita akan menerapkan sedikit ilmu yang kita capai baru saja.

> Kenapa Gin?
>
> 1. Well documented
> 2. Broad community
> 3. Easy to use

Langsung saja

* `app.go`
  
  ```go
  package main

  import (
    // "tutorial/function"
    // multiple "tutorial/function/multiple"
    // "tutorial/function/substract"

    "github.com/gin-gonic/gin"
  )

  func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
      c.JSON(200, "Hello World")
    })

    router.Run("localhostlocalhost:8080")
  }

  ```

Buka postman, lalu inputkan `http://localhostlocalhost:8080` pada Input Request URL Field.
Lalu tekan **Send**

Jika pada Response tertulis "Hello World", maka RESTAPI pertama kita berhasil dibuat.

---

### Folder Structure

Dalam mengembangkan aplikasi backend yang scalable, dibutuhkan folder structure yang tepat.
Tidak ada pakem untuk folder structure, tapi ada convention untuk folder structure yang dipakai pengembang kebanyakan.

Contoh yang sering digunakan:

```bash
Folder
|-models
|-views
|-controllers
|-routes
|-configs
|-middlewares
```

1. **model**: digunakan untuk menyimpan dan menentukan struktur data yang digunakan pada routine dan subroutine.
2. **views**: digunakan untuk mempresentasikan struktur data, baik masih dalam proses ataupun final.
3. **controllers**: digunakan untuk menhandle fungsi yang dijalankan ketika ada request API.
4. **routes**: digunakan untuk menentukan route API yang akan digunakan.
5. **configs**: digunakan untuk menyimpan fungsi atau konstanta konfigurasi sistem.
6. **middlewares**: digunakan untuk menyimpan fungsi ataupun routine-subroutine pipeline.

Pada project yang telah kita buat, terdapat folder `function`. kita dapat mengabaikannya terlebih dahulu, kita akan pindahkan di tempat yang seharusnya.

Kita akan membuat struktur folder seperti yang dicontohkan.

Let's go to the real problems.

---

### Backend for Internet of Things using Gin

Kita akan membuat sebuah backend untuk Internet of Things dengan:

1. Go Language
2. Gin-Gonic Framework
3. MongoDB

Dengan fitur antara lain:

1. Register Things
2. Log
3. Read
4. Remove

Kita akan menentukan struktur data dari sistem Internet of Things yang akan kita buat.

Project ini akan kita beri judul "Vehicle Tracker". Kita akan menyimpan data kendaraan. Data yang dapat kita simpan antara lain

1. Latitude
2. Longitude
3. Kecepatan

#### Database Connection

* `configs/setup.go`
  
  ```go
  package configs

  import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
  )

  func ConnectDB() *mongo.Client {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
      log.Fatal(err)
    }

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
      log.Fatal(err)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")
    return client
  }

  // getting database collections
  func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("VehicleTracker").Collection(collectionName)
    return collection
  }

  ```

* `app.go`
  
  ```go
  package main

  import (
    "tutorial/configs"

    "github.com/gin-gonic/gin"
  )

  func main() {
    router := gin.Default()
    configs.ConnectDB()

    ...

    router.Run("localhost:8080")
  }
  ```

#### Things (Authentication/Register)

Untuk mengidentifikasi Things yang mengirimkan request ke server, kita perlu memberikan identifier.
Kita akan menyimpan data Things pada collection yang akan kita beri nama "things".

* `models/things.go`
  
  ```go
  package models

  import (
    "time"
  )

  type Things struct {
    Thingname string    `json:"things_name,omitempty" bson:"things_name,omitempty" validate:"required,min=0"`
    Password  string    `json:"password,omitempty" validate:"required,min=3,max=255"`
    CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
  }
  ```

* `routes/things.go`
  
  ```go
  package routes

  import (
    controller_things "tutorial/controllers/things"

    "github.com/gin-gonic/gin"
  )

  func ThingsRoute(router *gin.RouterGroup) {
    router.POST("/things/register", controller_things.RegisterThings())
  }
  ```

* `views/things.go`
  
  ```go
  package views

  type ThingsView struct {
    ThingsId  interface{} `json:"things_id,omitempty" validate:"required"`
    Thingname string      `json:"things_name,omitempty" bson:"things_name,omitempty" validate:"required,min=0"`
  }
  ```

* `controllers\things\register_things.go`
  
  ```go
  package controller_things

  import (
    "context"
    "net/http"
    "time"
    "tutorial/configs"
    "tutorial/models"
    "tutorial/views"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
  )

  var validate = validator.New()

  var thingsCollection *mongo.Collection = configs.GetCollection(configs.DB, "things")

  func RegisterThings() gin.HandlerFunc {
    return func(c *gin.Context) {
      ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
      var things models.Things
      defer cancel()
      c.BindJSON(&things)

      if validationErr := validate.Struct(&things); validationErr != nil {
        c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
        return
      }

      count, err_ := thingsCollection.CountDocuments(ctx, bson.M{"things_name": things.Thingname})

      if err_ != nil {
        c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
        return
      }

      if count >= 1 {
        c.JSON(http.StatusBadRequest, bson.M{"data": "Things name has been taken"})
        return
      }

      bytes, errors := bcrypt.GenerateFromPassword([]byte(things.Password), 14)
      if errors != nil {
        c.JSON(http.StatusBadRequest, bson.M{"data": "Password tidak valid"})
      }

      newThings := models.Things{
        Thingname: things.Thingname,
        Password:  string(bytes),
        CreatedAt: time.Now(),
      }

      result, err := thingsCollection.InsertOne(ctx, newThings)
      if err != nil {
        c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
        return
      }

      finalView := views.ThingsView{
        ThingsId:  result.InsertedID,
        Thingname: things.Thingname,
      }

      c.JSON(http.StatusCreated, bson.M{
        "status":  http.StatusCreated,
        "message": "success",
        "data":    finalView,
      })
    }
  }

  ```

* `app.go`

  ```go
  package main

  import (
    "tutorial/configs"
    "tutorial/routes"

    "github.com/gin-gonic/gin"
  )

  func main() {
    router := gin.Default()
    configs.ConnectDB()

    v1 := router.Group("/v1")

    routes.ThingsRoute(v1)

    router.Run("localhost:8080")
  }
  ```

***Postman test***
![TutorialImg TutorialImg](/tutorial_images/1.PNG "Postman test")
![TutorialImg TutorialImg](/tutorial_images/2.PNG "Postman test")

#### Things (Authentication/RetriveId)

* `routes/things.go`
  
  ```go
  package routes

  import (
    controller_things "tutorial/controllers/things"

    "github.com/gin-gonic/gin"
  )

  func ThingsRoute(router *gin.RouterGroup) {
    router.POST("/things/register", controller_things.RegisterThings())
    router.POST("/things/retriveid", controller_things.GetIdThings())
  }
  ```

* `views/things.go`
  
  ```go
  
  ...

  type PayloadRetriveId struct {
    Thingname string `json:"things_name,omitempty" bson:"things_name,omitempty" validate:"required,min=0"`
    Password  string `json:"password,omitempty" validate:"required,min=3,max=255"`
  }

  type FinalRetriveId struct {
    ThingsId  interface{} `json:"things_id,omitempty" bson:"_id,omitempty" validate:"required"`
    Thingname string      `json:"things_name,omitempty" bson:"things_name,omitempty" validate:"required,min=0"`
  }
  ```

* `controllers\things\get_id_things.go`
  
  ```go
  package controller_things

  import (
    "context"
    "net/http"
    "time"
    "tutorial/models"
    "tutorial/views"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
  )

  func GetIdThings() gin.HandlerFunc {
    return func(c *gin.Context) {
      ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
      var thingPayload views.PayloadRetriveId
      defer cancel()
      c.BindJSON(&thingPayload)

      if validationErr := validate.Struct(&thingPayload); validationErr != nil {
        c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
        return
      }

      var resultThings models.Things
      var finalPayload views.FinalRetriveId
      result := thingsCollection.FindOne(ctx, bson.M{"things_name": thingPayload.Thingname})
      result.Decode(&resultThings)
      result.Decode(&finalPayload)
      err := bcrypt.CompareHashAndPassword([]byte(resultThings.Password), []byte(thingPayload.Password))
      if err != nil {
        c.JSON(http.StatusBadRequest, bson.M{
          "status":  http.StatusBadRequest,
          "message": "Bad request",
          "data":    "Things Name or Password is not valid",
        })
        return
      }

      c.JSON(http.StatusOK,
        bson.M{
          "status":  http.StatusOK,
          "message": "Success",
          "data":    finalPayload,
        },
      )
    }
  }

  ```

***Postman test***
![TutorialImg TutorialImg](/tutorial_images/3.PNG "Postman test")
![TutorialImg TutorialImg](/tutorial_images/4.PNG "Postman test")

#### Point (Logging)

* `models/point.go`

  ```go
  package models

  import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
  )

  type Point struct {
    ThingsId  primitive.ObjectID `json:"things_id,omitempty" bson:"things_id,omitempty" validate:"required,min=0"`
    Latitude  string             `json:"latitude,omitempty" validate:"required"`
    Longitude string             `json:"longitude,omitempty" validate:"required"`
    Velocity  int64              `json:"velocity,omitempty" validate:"required"`
    CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
  }
  ```

* `routes/point.go`

  ```go
  package routes

  import (
    controller_point "tutorial/controllers/point"

    "github.com/gin-gonic/gin"
  )

  func PointRoute(router *gin.RouterGroup) {
    router.POST("/point/:things_id", controller_point.PointLogging())
  }
  ```

* `views/point.go`
  
  ```go
  package views

  type PayloadPoint struct {
    Latitude  string `json:"latitude,omitempty" validate:"required"`
    Longitude string `json:"longitude,omitempty" validate:"required"`
    Velocity  int64  `json:"velocity,omitempty" validate:"required"`
  }

  type FinalPoint struct {
    Id        interface{} `json:"_id,omitempty" bson:"_id,omitempty" validate:"required"`
    Latitude  string      `json:"latitude,omitempty" validate:"required"`
    Longitude string      `json:"longitude,omitempty" validate:"required"`
    Velocity  int64       `json:"velocity,omitempty" validate:"required"`
  }
  ```

* `controllers/point/post_logging.go`

  ```go
  package controller_point

  import (
    "context"
    "net/http"
    "time"
    "tutorial/configs"
    "tutorial/models"
    "tutorial/views"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
  )

  var validate = validator.New()

  var pointsCollection = configs.GetCollection(configs.DB, "points")
  var thingsCollection = configs.GetCollection(configs.DB, "things")

  func PointLogging() gin.HandlerFunc {
    return func(c *gin.Context) {
      ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
      var pointPayload views.PayloadPoint
      defer cancel()
      var thingsId = c.Param("things_id")
      c.BindJSON(&pointPayload)

      if validationErr := validate.Struct(&pointPayload); validationErr != nil {
        c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
        return
      }

      thingsIdObj, err := primitive.ObjectIDFromHex(thingsId)
      if err != nil {
        c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
        return
      }

      count, err_ := thingsCollection.CountDocuments(ctx, bson.M{"_id": thingsIdObj})

      if err_ != nil {
        c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
        return
      }

      if count == 0 {
        c.JSON(http.StatusBadRequest, bson.M{"data": "Things by given Id is not found"})
        return
      }

      newPoint := models.Point{
        ThingsId:  thingsIdObj,
        Latitude:  pointPayload.Latitude,
        Longitude: pointPayload.Longitude,
        Velocity:  pointPayload.Velocity,
      }

      result, err := pointsCollection.InsertOne(ctx, newPoint)
      if err != nil {
        c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
        return
      }

      finalView := views.FinalPoint{
        Id:        result.InsertedID,
        Latitude:  pointPayload.Latitude,
        Longitude: pointPayload.Longitude,
        Velocity:  pointPayload.Velocity,
      }

      c.JSON(http.StatusCreated, bson.M{
        "status":  http.StatusCreated,
        "message": "success",
        "data":    finalView,
      })

    }
  }

  ```

* `app.go`

  ```go
  package main

  import (
    "tutorial/configs"
    "tutorial/routes"

    "github.com/gin-gonic/gin"
  )

  func main() {
    router := gin.Default()
    configs.ConnectDB()

    v1 := router.Group("/v1")

    routes.ThingsRoute(v1)
    routes.PointRoute(v1)

    router.Run("localhost:8080")
  }
  ```
