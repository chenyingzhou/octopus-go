package datasource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func ElasticSearchOut(targetSource string, targetSet string, docs []map[string]interface{}, idColumn string) {
	client, err := GetESClient(targetSource)
	if err != nil {
		log.Println("获取ES连接失败，将转为标准输出...")
		Stdout(targetSource, targetSet, docs, idColumn)
		return
	}
	_, err = client.Bulk(bulkBody(idColumn, docs), client.Bulk.WithIndex(targetSet))
	if err != nil {
		log.Println("写入ES失败，将转为标准输出...")
		Stdout(targetSource, targetSet, docs, idColumn)
	}
}

func bulkBody(idColumn string, docs []map[string]interface{}) io.Reader {
	body := &bytes.Buffer{}
	for _, doc := range docs {
		meta := []byte(fmt.Sprintf(`{"index":{"_id":"%s"}}%s`, doc[idColumn], "\n"))
		data, _ := json.Marshal(doc)
		data = append(data, "\n"...)
		body.Grow(len(meta) + len(data))
		body.Write(meta)
		body.Write(data)
	}
	return body
}
