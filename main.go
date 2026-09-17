package main

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
	"github.com/jemaimedamine22-png/school-api/api"
    "github.com/jemaimedamine22-png/school-api/db"
    "github.com/jemaimedamine22-png/school-api/util"
)

func main() {
	// 1. قراءة الإعدادات من ملف app.env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// 2. الاتصال بقاعدة البيانات باستخدام Driver وفحصه
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// 3. إنشاء الـ Store الخاص بـ sqlc
	store := db.NewStore(conn)
	
	log.Println("connected to database successfully!", store)

	// هنا سنقوم لاحقاً بتشغيل سيرفر Gin وإطلاق الـ API
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}