package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)



func main() {
	url := "http://localhost:9200/hotel/_search"

	// 並列実行するリクエストボディの定義
	bodies := make([][]byte, 12)
	bodies[0] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"JANUARY","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[1] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"FEBRUARY","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[2] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"MARCH","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[3] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"APRIL","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[4] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"MAY","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[5] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"JUNE","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[6] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"JULY","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[7] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"AUGUST","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[8] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"SEPTEMBER","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[9] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"OCTOBER","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[10] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"NOVEMBER","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)
	bodies[11] = []byte(`{"query":{"bool":{"must":[{"multi_match":{"query":"DECEMBER","fields":["city.keyword","country.keyword","hotel_name^1.0","busy_month.keyword"],"type":"best_fields"}}],"must_not":[{"match":{"travelerOrphanFlag":{"query":true,"operator":"OR","prefix_length":0,"max_expansions":50,"fuzzy_transpositions":true,"lenient":false,"zero_terms_query":"NONE","auto_generate_synonyms_phrase_query":true,"boost":1}}}],"adjust_pure_negative":true,"boost":1}},"aggs":{"test":{"terms":{"field":"hotel_name.keyword","size":1000}}}}`)

	wg := sync.WaitGroup{}

	builder := strings.Builder{}
	// 時間計測開始
	start := time.Now()
	//var allTime []float64
	for _, body := range bodies {
		wg.Add(1)
		b := body
		go func() {

			req, err := http.NewRequest("GET", url, bytes.NewBuffer(b))
			if err != nil {
				panic(err)
			}

			// Headerの設定
			encoder := base64.StdEncoding
			//auth := encoder.EncodeToString([]byte(os.Getenv("USER")+":"+ os.Getenv("PASS")))
			auth := encoder.EncodeToString([]byte("elastic"+":"+ "changeme"))
			req.Header.Set("Authorization", "Basic "+auth)

			req.Header.Set("Content-Type", "application/json")

			// search リクエスト実行
			client := http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			body, err = ioutil.ReadAll(resp.Body)

			var f interface{}

			err2 := json.Unmarshal(body, &f)
			if err2 != nil {
				panic(err2)
			}

			m := f.(map[string]interface{})
			eachTaken := m["took"].(float64)
			s := fmt.Sprint(eachTaken)
			builder.WriteString(s + ",")

			if err != nil {
				panic(err)
			}

			wg.Done()
		}()
	}
	wg.Wait()

	end := time.Now()
	fmt.Println("")

	// compile error
	//fmt.Printf("%d,", (end.Sub(start)).Milliseconds())
	fmt.Printf("%d,",(end.Sub(start)))

	final := strings.TrimRight(builder.String(), ",")
	fmt.Printf("%s", final)
}
