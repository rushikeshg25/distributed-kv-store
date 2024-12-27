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

func GetKeyValue(key string) ([]byte,error){
	var value []byte
	timeNow:=time.Now().Unix()
	err:=Db.QueryRow(`SELECT (kv_value,kv_expires_at) FROM kv WHERE kv_key=? AND kv_expires_at>?`,key,timeNow).Scan(&value)
	if err!=nil{
		return nil,err
	}
	return value,nil
}

//Incase of Inme DB use the <20% rule to delete keys on cron job 
func DeleteKeyValue(key string){}