package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

var (
	db     *bolt.DB
	bucket []byte
)

const dbname = "ljl.db"

func init() {
	//创建bolt数据库本地文件
	dbc, err := bolt.Open(dbname, 0600, nil)

	//初始化bucket
	bucket = []byte("demoBucket")
	if err != nil {
		fmt.Println("open err:", err)
		return
	} else {
		db = dbc
	}
	//创建bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket(bucket)
		return err
	})
}

//把数据插入到bolt数据库中，相当于redis中的set命令
func insert(key, value string) {

	k := []byte(key)
	v := []byte(value)
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)

		err := b.Put(k, v)
		return err
	})
}

//删除一个指定的key中的数据
func rm(key string) {
	k := []byte(key)
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Delete(k)
		return err
	})
}

//读取一条数据
func read(key string) string {
	k := []byte(key)
	var val []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		val = b.Get(k)
		return nil
	})
	return string(val)
}

//遍历指定的bucket中的数据
func fetchAll(buk []byte) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(buk)
		cur := b.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			fmt.Printf("key is %s,value is %s\n", k, v)
		}
		return nil
	})
}

func main() {
	defer db.Close()

	insert("hello", "Hello World,This is Bolt Database11..")
	insert("hello2", "Hello World,This is Bolt Database22..")
	fmt.Println(read("hello"))
	fmt.Println(read("hello2"))
	rm("hello")
	insert("news1", "this is a title.")

	fetchAll(bucket)

}
