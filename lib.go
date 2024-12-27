package main

import (
	"time"
)

//Cruud

var 

func KeyValueCreationOrUpdate(key string, value any){
	// expiresAt:=time.Now()
	query,err:=Db.Prepare(`
	INSERT INTO kv (kv_key,kv_value) VALUES (?,?)
	`)
}