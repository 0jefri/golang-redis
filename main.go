package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Buat koneksi ke Redis
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Alamat Redis (default)
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Redis")

	// Simpan data ke Redis
	err := rdb.Set(ctx, "mykey", "hello redis", 0).Err()
	if err != nil {
		fmt.Println("Error saat menyimpan data:", err)
		return
	}

	// Ambil data dari Redis
	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		fmt.Println("Error saat mengambil data:", err)
		return
	}
	fmt.Println("Nilai dari 'mykey':", val)

	// Hapus data dari Redis
	// err = rdb.Del(ctx, "mykey").Err()
	// if err != nil {
	// 	fmt.Println("Error saat menghapus data:", err)
	// 	return
	// }
	// fmt.Println("Data berhasil dihapus!")
}
