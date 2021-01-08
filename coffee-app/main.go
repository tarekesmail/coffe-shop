package main
import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"gopkg.in/gorp.v1"
	_ "github.com/lib/pq"
	"log"
	"os"

	//"fmt"
	//"reflect"
)
////

var initdb = initDb()
var db_host = os.Getenv("DB_HOST")
var db_user = os.Getenv("DB_USER")
var db_password = os.Getenv("DB_PASSWORD")
var db_name = os.Getenv("DB_NAME")
var server_port = os.Getenv("SERVER_PORT")

func initDb() *gorp.DbMap {
	db, err := sql.Open("postgres", "host="+db_host+" user="+db_user+" password="+db_password+" dbname="+db_name+" sslmode=disable")
	checkErr(err, "sql.Open failed")
	initdb := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	return initdb
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Get_Invalid_Deliveries(c *gin.Context) {
	var deliveries []Deliveries
	_, err := initdb.Select(&deliveries, "SELECT delivery.*" +
		"FROM delivery LEFT JOIN (SELECT delivery.id," +
		"delivery.supplier_id," +
		"delivery.driver_id," +
		"supplier_bean_type.bean_type_id," +
		"carrier_bean_type.bean_type_id " +
		"FROM delivery LEFT JOIN supplier_bean_type " +
		"on delivery.supplier_id = supplier_bean_type.supplier_id " +
		"JOIN driver on driver.id = delivery.driver_id " +
		"JOIN carrier_bean_type on carrier_bean_type.carrier_id = driver.id) " +
		"valid on (delivery.supplier_id = valid.supplier_id and delivery.driver_id = valid.driver_id) " +
		"where valid.id is null")
	if err == nil {
		c.JSON(200, deliveries)
	} else {
		c.JSON(404, gin.H{"error": "no data into the table"})
	}
}

func index (c *gin.Context) {
	content := gin.H{"This is": "API Index"}
	c.JSON(200, content)
}
////
type Deliveries struct {
	Id int64 `db:"id" json:"id"`
	Supplier_ID string `db:"supplier_id" json:"supplier_id"`
	Driver_ID string `db:"driver_id" json:"driver_id"`
	Updated_At string `db:"updated_at" json:"updated_at"`
	Created_At string `db:"created_at" json:"created_at"`
}
func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/invalid-deliveries", Get_Invalid_Deliveries)
	}
	r.GET("/", index)
	r.Run(":"+server_port)
}


