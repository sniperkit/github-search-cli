package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func Test_Ghs(t *testing.T) {
	assert := func(result interface{}, want interface{}) {
		if !reflect.DeepEqual(result, want) {
			t.Errorf("Returned %+v, want %+v", result, want)
		}
	}

	Setup()
	defer Teardown()

	var handler func(w http.ResponseWriter, r *http.Request)
	// Normal response
	handler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Link", HeaderLink(100, 10))
		var items []string
		for i := 1; i < 100+1; i++ {
			items = append(items, fmt.Sprintf(`{"id":%d, "full_name": "test/search_word%d"}`, i, i))
		}
		fmt.Fprintf(w, `{"total_count": 1000, "items": [%s]}`, strings.Join(items, ","))
	}

	mux.HandleFunc("/search/repositories", handler)
	num, err := ghs(strings.Split(fmt.Sprintf("-e %s -m 1000 SEARCH_WORD", server.URL), " "))
	assert(num, 1000)
	assert(err, nil)

	num, err = ghs(strings.Split(fmt.Sprintf("-e %s -m 110 SEARCH_WORD", server.URL), " "))
	assert(num, 110)
	assert(err, nil)
}
