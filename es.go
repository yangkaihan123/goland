package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)

//传入参数
type listParam struct {
	page      int
	size      int
	ip        string
	startTime int64
	endTime   int64
	level     string
	hostname  string
}

//返回查询结果
type listRtn struct {
	total int64
	logs  []log
}
type log struct {
	id        string
	Ip        string `json:"ip"`
	Timestamp int64  `json:"timestamp`
	Level     string `json:"level"`
	Hostname  string `json:"hostname"`
}

func main() {
	//传入参数
	var listPostParam = listParam{
		page:      0,
		size:      10,
		ip:        "192.*",
		startTime: 1605139523095,
		endTime:   1605337534872,
		level:     "INFO",
		hostname:  "HF-DC",
	}

	//连接es
	client, err := elastic.NewClient(elastic.SetURL("http://172.17.6.16:9200/"), elastic.SetBasicAuth("elastic", "Wanzhu@sby!"))
	if err != nil {
		panic(err)
	}

	//es查询
	var res *elastic.SearchResult
	indices := "log"
	boolQ := elastic.NewBoolQuery()
	if len(strings.TrimSpace(listPostParam.level)) > 0 {
		boolQ.Filter(elastic.NewTermQuery("level", listPostParam.level)) //level需准确匹配
	}

	if listPostParam.endTime > 0 && listPostParam.startTime <= 0 { //日期区间查询
		boolQ.Filter(elastic.NewRangeQuery("timestamp").Lte(listPostParam.endTime))
	} else if listPostParam.endTime <= 0 && listPostParam.startTime > 0 {
		boolQ.Filter(elastic.NewRangeQuery("timestamp").Gte(listPostParam.startTime))
	} else if listPostParam.endTime > 0 && listPostParam.startTime > 0 {
		boolQ.Filter(elastic.NewRangeQuery("timestamp").Lte(listPostParam.endTime).Gte(listPostParam.startTime))
	}
	//注意：ip和hostname都需要实现模糊查询，ip通过wildcard实现，hostname通过match实现
	// match查询不一定完全符合，这里只是提供一种思路，比如用户填入内容为HF-DC-OP，同样会查出记录HF-DC-FR-001，这个就需要考虑分词了
	/*
		GET log/_analyze
		{
		  "field": "hostname",
		  "text":  "HF-DC-OP"
		}
		结果：
		{
		  "tokens" : [
		    {
		      "token" : "hf",
		      "start_offset" : 0,
		      "end_offset" : 2,
		      "type" : "<ALPHANUM>",
		      "position" : 0
		    },
		    {
		      "token" : "dc",
		      "start_offset" : 3,
		      "end_offset" : 5,
		      "type" : "<ALPHANUM>",
		      "position" : 1
		    },
		    {
		      "token" : "op",
		      "start_offset" : 6,
		      "end_offset" : 8,
		      "type" : "<ALPHANUM>",
		      "position" : 2
		    }
		  ]
		}
	*/
	if len(strings.TrimSpace(listPostParam.ip)) > 0 {
		boolQ.Filter(elastic.NewWildcardQuery("ip.keyword", fmt.Sprintf("*%s*", listPostParam.ip))) //ip模糊查询
	}
	if len(strings.TrimSpace(listPostParam.hostname)) > 0 {
		boolQ.Filter(elastic.NewMatchQuery("hostname", listPostParam.hostname).Operator("or")) //hostname模糊查询，试着将or调整为and看看结果
	}
	res, err = client.Search(indices).
		Query(boolQ).
		Sort("timestamp", false).
		From(listPostParam.page * listPostParam.size).
		Size(listPostParam.size).
		Do(context.Background()) //分页查询
	if err != nil {
		panic(err)
	}

	//es查询结果处理：不清楚怎么处理结果，可以参考第三步的查询结果，是相互对应的
	var listRtn listRtn
	if res != nil && res.Hits.TotalHits.Value > 0 {
		listRtn.total = res.Hits.TotalHits.Value
		for _, hit := range res.Hits.Hits {
			var log log
			err = json.Unmarshal(hit.Source, &log)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			log.id = hit.Id
			listRtn.logs = append(listRtn.logs, log)
		}
	} else {
		listRtn.total = 0
	}
	fmt.Printf("%v", listRtn)
}
