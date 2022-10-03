package test

import (
	"testing"

	"github.com/gomodule/redigo/redis"
)

func InitConnection() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", "localhost:6379", redis.DialPassword("eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81"))
	if err != nil {
		panic(err)
	}
	return conn, nil
}

func TestRedisConnection(t *testing.T) {
	conn, err := redis.Dial("tcp", "localhost:6379", redis.DialPassword("eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = conn.Do("HSET", "mahasiswa:1", "nama", "Redha Juanda", "nim", "12345", "ipk", 3.34, "semester", 4)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetData(t *testing.T) {
	conn, err := InitConnection()
	if err != nil {
		t.Fatal(err)
	}

	nama, err := redis.String(conn.Do("HGET", "mahasiswa:1", "nama"))
	if err != nil {
		t.Fatal(err)
	}

	nim, err := redis.String(conn.Do("HGET", "mahasiswa:1", "nim"))
	if err != nil {
		t.Fatal(err)
	}

	ipk, err := redis.Float64(conn.Do("HGET", "mahasiswa:1", "ipk"))
	if err != nil {
		t.Fatal(err)
	}

	semester, err := redis.Int(conn.Do("HGET", "mahasiswa:1", "semester"))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(nama)
	t.Log(nim)
	t.Log(ipk)
	t.Log(semester)
}
