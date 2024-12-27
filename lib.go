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

func GetKeyValue(key string) any{
	var value any
	var expiresAt int64
	err:=Db.QueryRow(`SELECT (kv_value,kv_expires_at) FROM kv WHERE kv_key=key`).Scan(&value,&expiresAt)
	if(err!=nil){
		log.Fatalf("%s",err)
	}
	if(time.Now().Unix()>expiresAt){
		_,err:=Db.Exec(`DELETE FROM kv WHERE kv_key=key`,key)
		if(err!=nil){
			log.Fatalf("%s",err)
		}
		return nil
	}
	return value
}

//Incase of Inme DB use the <20% rule to delete keys on cron job 
func DeleteKeyValue(key string){}