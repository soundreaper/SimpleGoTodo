package main

import (
	"io"
	"net/http"

<<<<<<< HEAD
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

=======
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var db, _ = gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")

type TodoItemModel struct {
	Id          int `gorm:"primary_key"`
	Description string
	Completed   bool
}

>>>>>>> 18ae470b958cc7208bc866cdb908c3253360ea29
func APIStatus(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
<<<<<<< HEAD
=======
	defer db.Close()
	db.Debug().DropTableIfExists(&TodoItemModel{})
	db.Debug().AutoMigrate(&TodoItemModel{})

>>>>>>> 18ae470b958cc7208bc866cdb908c3253360ea29
	log.Info("Starting Todolist API server")
	router := mux.NewRouter()
	router.HandleFunc("/apistatus", APIStatus).Methods("GET")
	http.ListenAndServe(":8000", router)
}
