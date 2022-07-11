package datasource

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func Stdout(idColumn string, docs []map[string]interface{}) {
	count := len(docs)
	fmt.Printf("文档总数:%d", count)
	if len(docs) > 0 {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(count)
		doc := docs[i]
		fmt.Printf("，其中第%d条(id:%s)为：\n", i+1, doc[idColumn])
		fmt.Println("****************************************")
		text, err := json.MarshalIndent(doc, "", "    ")
		if err == nil {
			fmt.Println(string(text))
			fmt.Println("****************************************")
		}
	}
}
