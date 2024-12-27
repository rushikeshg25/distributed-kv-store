package main

import (
	"log"
	"time"
)


func KeyValueCreationOrUpdate(key string, value any){
	expiresAt:=time.Now().Add(time.Minute*10).Unix()
	query,err:= Db.Prepare(`INSERT INTO kv (kv_key, kv_value, kv_expires_at) 
	VALUES (?, ?, ?) 
	ON DUPLICATE KEY UPDATE 
	kv_value = VALUES(kv_value),
	kv_expires_at = VALUES(kv_expires_at)`)
	if err!=nil{
		log.Fatalf("%s",err)
	}
	_, err = query.Exec("key1", []byte("value"), expiresAt)
	if(err!=nil){
		log.Fatalf("%s",err)
	}
}