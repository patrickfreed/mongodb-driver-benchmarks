package main

import (
    "context"
    "time"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"encoding/json"

    "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func loadSourceDocument(pathParts ...string) SmallDoc {
	data, err := ioutil.ReadFile(filepath.Join(pathParts...))
	if err != nil {
		panic(err)
	}
	doc := SmallDoc{}
	err = json.Unmarshal(data, &doc)
	if err != nil {
		panic(err)
	}
	// fmt.Println(doc.OggaoR4O);
	// fmt.Println(doc.Vxri7mmI)

	return doc
}

type SmallDoc struct {
	OggaoR4O string `json:oggaoR4O`
	Vxri7mmI string `json:vxri7mmI`
	IQ8K4ZMG string `json:IQ8K4ZMG`
	WzI8s1W0 string `json:wzI8s1W0`
	E5Aj2zB3 string `json:e5Aj2zB3`
	TMXe8Wi7 string
	PHPSSV51 string
	CxSCo4jD int64
	FC8GSDC5 int64
	FC63DsLR int64
	L6e0U4bR int64
	TLRpkltp int64
	Ph9CZN5L int64
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println("%s", err)
	}

	coll := client.Database("bench").Collection("go")

	coll.Drop(context.Background())

	doc := loadSourceDocument("../small_doc.json")
	fmt.Println(doc);

	payload := make([]interface{}, 10000)
	for idx := range payload {
		payload[idx] = doc;
	}

	start := time.Now()
	_, err = coll.InsertMany(ctx, payload)
	if err != nil {
		fmt.Println(err)
	}
	duration := time.Since(start)

	fmt.Println("insert: ", duration.Milliseconds(), "ms")

	c, _ := coll.EstimatedDocumentCount(context.Background());
	fmt.Println("docs: ", c)
}
