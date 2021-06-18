package mypkg_test

import (
	"encoding/json"
	"mypkg"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {
	cases := map[string]struct {
		n        int
		hasError bool
	}{
		"100": {n: 100},
		"200": {n: 200},
	}
	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			var requested bool
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requested = true
				if r.FormValue("n") != strconv.Itoa(tc.n) {
					t.Errorf("param n want %s got %d", r.FormValue("n"), tc.n)
				}
				resp := &mypkg.ExportGetResponse{
					Value: "hoge",
				}
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					t.Fatal("unexpected error:", err)
				}
			}))
			defer s.Close()
			defer mypkg.SetBaseURL(s.URL)()
			cli := mypkg.Client{HTTPClient: s.Client()}
			_, err := cli.Get(tc.n)
			switch {
			case err != nil && !tc.hasError:
				t.Error("unexpected error:", err)
			case err == nil && tc.hasError:
				t.Error("expected error has not occurred")
			}
			if !requested {
				t.Error("no request")
			}
		})
	}
}
