package learm_redis

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// 172.23.96.128:6379
// Addr: "localhost:6379",
// var client = redis.NewClient(&redis.Options{
// 	Addr:     "172.23.96.128:6381",
// 	Password: "",
// 	DB:       0,
// })

var client = redis.NewClient(&redis.Options{
	Addr:     "172.23.96.128:6379",
	Password: "",
	DB:       0,
})

// var opt, _ = redis.ParseURL("redis://ugikdev@172.23.96.128:6384/0")
// var client = redis.NewClient(opt)

func TestConnection(t *testing.T) {
	fmt.Println(client)
	assert.NotNil(t, client)
	// err := client.Close()
	// assert.Nil(t, err)
}

var ctx = context.Background()

func TestPing(t *testing.T) {

	result, err := client.Ping(ctx).Result()
	// fmt.Println(result)
	// fmt.Println("err", err)
	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Ugik Developer", time.Second*4)

	result, err := client.Get(ctx, "name").Result()
	assert.Equal(t, "Ugik Developer", result)
	assert.Nil(t, err)
	// jedah 5 second seharusnya tidak ada datanya
	// fmt.Println("Data diredis", result)
	time.Sleep(time.Second * 5)
	result, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
	assert.Equal(t, "", result)
}

/*
*
RPush artinya push Right atau dari kanan
.Result() mengembalikan result dan err
.Val() mengembalikan nilainya saja
.Err Mengembalikan errornya saja
*/
func TestList(t *testing.T) {

	// defer client.Del(ctx, "names")
	client.RPush(ctx, "names", "Ugik")
	client.RPush(ctx, "names", "Fullstack")
	client.RPush(ctx, "names", "Developer")

	result := client.LRange(ctx, "names", 0, -1).Val()
	fmt.Println("re", result)
	// for i := 0; i <= 5; i++ {
	// 	// menghindari racecondition
	// 	// time.Sleep(10 * time.Millisecond)
	// 	fmt.Println(client.LPop(ctx, "names").Val())
	// }
	assert.Equal(t, "Ugik", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Fullstack", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Developer", client.LPop(ctx, "names").Val())
	client.Del(ctx, "names")
	// fmt.Println("byValue", client.LPop(ctx, "names").Val())
	// fmt.Println("byValue", client.LPop(ctx, "names").Val())
	// fmt.Println("byValue", client.LPop(ctx, "names").Val())

	// result, _ := client.LPop(ctx, "names").Result()
	// fmt.Println("byResult", result)
	// result, _ = client.LPop(ctx, "names").Result()
	// fmt.Println("byResult", result)
	// result, _ = client.LPop(ctx, "names").Result()
	// fmt.Println("byResult", result)
	// result, _ = client.LPop(ctx, "names").Result()
	// fmt.Println("byResult", result)
	// fmt.Println(client.LPop(ctx, "names").Result())
	// fmt.Println(client.LPop(ctx, "names").Result())
	// fmt.Println(client.LPop(ctx, "names").Result())
	// fmt.Println("byValue", byValue)

	// result, _ = client.LPop(ctx, "names").Result()
	// fmt.Println(result)

	// result, _ = client.LPop(ctx, "names").Result()
	// fmt.Println(result)

}

func sleepSec() {
	time.Sleep(3 * time.Second)
}
func TestSet(t *testing.T) {

	/**
	Set (SAdd, SMembers): Set menyimpan elemen-elemen unik dan tidak berurutan.
	Ketika kamu menambahkan elemen dengan SAdd, Redis memastikan bahwa elemen tersebut tidak duplikat,
	tetapi tidak menjamin urutan penyimpanannya.

	Masalah Assertion: Kamu menggunakan assert.Equal untuk membandingkan dua slice,
	tetapi urutan elemen dalam slice yang dihasilkan oleh SMembers tidak dijamin sama
	dengan urutan slice yang kamu harapkan.

	Solusi assets.ElementsMatch


	SAdd Menghindari duplikat data
	SCard mengembalikan nilai panjang list
	SMembers mengembalikan nilai array
	*/
	client.SAdd(ctx, "students", "Fazea")
	client.SAdd(ctx, "students", "Fazea")
	client.SAdd(ctx, "students", "Audrilia")
	client.SAdd(ctx, "students", "Audrilia")
	client.SAdd(ctx, "students", "Pramita")
	client.SAdd(ctx, "students", "Pramita")
	sleepSec()
	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.ElementsMatch(t, []string{"Fazea", "Audrilia", "Pramita"}, client.SMembers(ctx, "students").Val())
	client.Del(ctx, "students")

}

func TestSortedSet(t *testing.T) {
	sleepSec()
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Zea"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "Ayah"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "Mama"})
	sleepSec()

	// Descending, terkecil ke besar
	assert.Equal(t, []string{"Ayah", "Mama", "Zea"}, client.ZRange(ctx, "scores", 0, -1).Val())
	// Ascending, terbesar ke yang kecil
	assert.Equal(t, []string{"Zea", "Mama", "Ayah"}, client.ZRevRange(ctx, "scores", 0, -1).Val())

	assert.Equal(t, "Zea", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Mama", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Ayah", client.ZPopMax(ctx, "scores").Val()[0].Member)
	// client.Del(ctx, "scores")

}

func TestHash(t *testing.T) {
	client.HSet(ctx, "user:1", "id", "1")
	client.HSet(ctx, "user:1", "name", "Ugik Dev")
	client.HSet(ctx, "user:1", "email", "ugik.dev@gmail.com")

	user := client.HGetAll(ctx, "user:1").Val()

	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "Ugik Dev", user["name"])
	assert.Equal(t, "ugik.dev@gmail.com", user["email"])

	client.Del(ctx, "user:1")
}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko A",
		Longitude: 106.818489,
		Latitude:  -6.178966,
	})
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko B",
		Longitude: 106.821568,
		Latitude:  -6.180662,
	})

	distance := client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val()
	assert.Equal(t, 0.3892, distance)

	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude:  106.819143,
		Latitude:   -6.180182,
		Radius:     5, // 1 meter radius
		RadiusUnit: "km",
		Sort:       "ASC", // Or "DESC" depending on needs
		Count:      10,    // Limit the number of results
	}).Val()
	// fmt.Println("distance", distance)
	// fmt.Println(sellers)

	// pos, _ := client.GeoPos(ctx, "sellers", "Toko A").Result()
	// fmt.Println(pos[0].Longitude, pos[0].Latitude)
	// pos, _ = client.GeoPos(ctx, "Sicily", "Palermo").Result()
	// fmt.Println(pos[0].Longitude, pos[0].Latitude)
	assert.Equal(t, []string{"Toko A", "Toko B"}, sellers)
}

func TestHyperLogLog(t *testing.T) {
	// Menyimpan data unique dengan size memory lebih kecil
	client.PFAdd(ctx, "visitors", "eko", "kurniawan", "khannedy")
	client.PFAdd(ctx, "visitors", "eko", "budi", "joko")
	client.PFAdd(ctx, "visitors", "rully", "budi", "joko")

	total := client.PFCount(ctx, "visitors").Val()
	assert.Equal(t, int64(6), total)
}

func TestPipeline(t *testing.T) {
	// digunakan untuk menjalankan perintah sekaligus tanpa menunggu satu persatu
	// sehingga lebih optimal dan cepat
	_, err := client.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "Ugik", 5*time.Second)
		pipeliner.SetEx(ctx, "address", "Pringsewu", 5*time.Second)
		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Ugik", client.Get(ctx, "name").Val())
	assert.Equal(t, "Pringsewu", client.Get(ctx, "address").Val())
}

func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "Joko", 5*time.Second)
		pipeliner.SetEx(ctx, "address", "Cirebon", 5*time.Second)
		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Joko", client.Get(ctx, "name").Val())
	assert.Equal(t, "Cirebon", client.Get(ctx, "address").Val())
}

/*
*
Data Strem di simpan di redis walaupun belum ada subscriber
Data disimpan dalam bentuk map
*/
func TestPublishStream(t *testing.T) {

	for i := 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "datastream",
			Values: map[string]interface{}{
				"name":    "Ugik",
				"address": "Indonesia",
				"desc":    "Data " + strconv.Itoa(i),
			},
		}).Err()
		fmt.Println("Create Data " + strconv.Itoa(i))
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	// client.Del(ctx, "datastream")
	// client.XGroupDestroy(ctx, "datastream", "group-1")
	client.XGroupCreate(ctx, "datastream", "group-1", "0")
	client.XGroupCreateConsumer(ctx, "datastream", "group-1", "consumer-1")
	client.XGroupCreateConsumer(ctx, "datastream", "group-1", "consumer-2")
	client.XGroupCreateConsumer(ctx, "datastream", "group-1", "consumer-3")
}

func TestConsumeStream(t *testing.T) {
	// tanda > menandakan membaca data yang belum dibaca
	// Block = saat tidak ada data, maka akan menunggu selama 5 detik untuk menunggu data masuk
	streams := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-2",
		Streams:  []string{"datastream", ">"},
		// Streams: []string{"datastream >", ">"},
		Count: 2,
		Block: 5 * time.Second,
		NoAck: true,
	}).Val()
	fmt.Println(streams)
	for _, stream := range streams {
		// fmt.Println("strean", stream)
		for _, message := range stream.Messages {
			fmt.Println("line 290", message.ID)
			fmt.Println("line 291", message.Values)
			client.XAck(ctx, "datastream", "group-1", message.ID).Err()
			// if ackErr != nil {
			// 	assert.Fail(t, "Failed to acknowledge message", ackErr)
			// } else {
			// 	fmt.Printf("Acknowledged Message ID: %s\n", message.ID)
			// }
		}
	}
}

/*
*
PUBSUB
lebih simple dari  stream
1. tidak ada consumer group
2. jika data diaksess akan hilang
3. butuh subscriber sebagai penganti cunsumer , dan jika tidak ada subscriber data akan hilann

jalankan dlu testsubscriber dan akan menunggu sampai testpublish ketika di kirim data maka baru dirender di subscriber
*/
func TestSubscribePubSub(t *testing.T) {
	subscriber := client.Subscribe(ctx, "channel-1")
	defer subscriber.Close()
	for i := 0; i < 10; i++ {
		// message, err := subscriber.ReceiveTimeout(ctx, 1*time.Minute)
		message, err := subscriber.ReceiveMessage(ctx)
		assert.Nil(t, err)
		fmt.Println(message.Payload)
	}
	fmt.Println("ended")
}

func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.Publish(ctx, "channel-1", "Hello "+strconv.Itoa(i)).Err()
		assert.Nil(t, err)
	}
}
