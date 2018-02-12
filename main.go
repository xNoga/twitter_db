package main

import (
	"log"

	"github.com/gin-gonic/gin"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checkNet(err error, c *gin.Context) {
	if err != nil {
		log.Print(err)
		c.JSON(500, "error!")
	}
}

var db *mgo.Session

func main() {
	var err error
	db, err = mgo.Dial("mongodb://tester:test123@138.68.69.47:27017/social_net")
	check(err)
	defer db.Close()

	done := make(chan bool, 1)
	go setIndexes(done) // setting indexes on "text" and "user" propeties.
	<-done

	router := gin.Default()
	router.GET("/users", getUsersCount)
	router.GET("/mentions", getTopMentioners)
	router.GET("/active", getMostActive)
	router.GET("/negative", getMostNegative)
	router.GET("/positive", getMostPositive)
	router.Run()
}

func getUsersCount(c *gin.Context) {
	var res []interface{}
	err := getColl().Find(bson.M{}).Distinct("user", &res)
	checkNet(err, c)
	c.JSON(200, gin.H{"numberOfUsers": len(res)})
}

func getTopMentioners(c *gin.Context) {
	var res []interface{}
	pipe := getColl().Pipe([]bson.M{
		{
			"$unwind": "$text",
		},
		{
			"$match": bson.M{
				"text": bson.M{"$regex": `@\w+`},
			},
		},
		{
			"$group": bson.M{
				"_id":      bson.M{"user": "$user"},
				"mentions": bson.M{"$sum": 1},
			},
		},
		{
			"$sort": bson.M{
				"mentions": -1,
			},
		},
		{
			"$limit": 10,
		},
	})
	err := pipe.All(&res)
	checkNet(err, c)
	c.JSON(200, res)
}

func getMostActive(c *gin.Context) {
	var res []interface{}
	pipe := getColl().Pipe([]bson.M{
		{
			"$group": bson.M{
				"_id":      bson.M{"user": "$user"},
				"mentions": bson.M{"$sum": 1},
			},
		},
		{
			"$sort": bson.M{
				"mentions": -1,
			},
		},
		{
			"$limit": 10,
		},
	})
	err := pipe.All(&res)
	checkNet(err, c)
	c.JSON(200, res)
}

func getMostNegative(c *gin.Context) {
	var res []interface{}
	pipe := getColl().Pipe([]bson.M{
		{
			"$match": bson.M{
				"text": bson.M{"$regex": `shit`},
			},
		},
		{
			"$group": bson.M{
				"_id":      bson.M{"user": "$user"},
				"mentions": bson.M{"$sum": 1},
				"text":     bson.M{"$push": "$text"},
			},
		},
		{
			"$sort": bson.M{
				"mentions": -1,
			},
		},
		{
			"$limit": 10,
		},
	})
	err := pipe.All(&res)
	checkNet(err, c)
	c.JSON(200, res)
}

func getMostPositive(c *gin.Context) {
	var res []interface{}
	pipe := getColl().Pipe([]bson.M{
		{
			"$match": bson.M{
				"text": bson.M{"$regex": `awesome`},
			},
		},
		{
			"$group": bson.M{
				"_id":      bson.M{"user": "$user"},
				"mentions": bson.M{"$sum": 1},
				"text":     bson.M{"$push": "$text"},
			},
		},
		{
			"$sort": bson.M{
				"mentions": -1,
			},
		},
		{
			"$limit": 10,
		},
	})
	err := pipe.All(&res)
	checkNet(err, c)
	c.JSON(200, res)
}

func setIndexes(done chan bool) {
	var coll = getColl()
	indexes, err := coll.Indexes()
	check(err)

	if len(indexes) < 2 {
		coll.EnsureIndexKey("user")
		coll.EnsureIndexKey("text")
	}
	done <- true
}

func getColl() *mgo.Collection {
	return db.DB("social_net").C("tweets")
}
