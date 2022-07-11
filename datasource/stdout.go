package datasource

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func Stdout(targetSource string, targetSet string, docs []map[string]interface{}, idColumn string) {
	count := len(docs)
	fmt.Printf("连接: %s, 数据集: %s, 文档总数:%d", targetSource, targetSet, count)
	if len(docs) > 0 {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(count)
		doc := docs[i]
		fmt.Printf(", 随机选择第%d条(id=%s):\n", i+1, doc[idColumn])
		fmt.Println("****************************************")
		text, err := json.MarshalIndent(doc, "", "    ")
		if err == nil {
			fmt.Println(string(text))
			fmt.Println("****************************************")
		}
	}
}
