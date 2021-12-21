package main

import (
	"common"
	"common/module"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	http.Handle(common.PathDataJson, metricsMiddle(http.HandlerFunc(jsonData)))
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}

func metricsMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/metrics" {
			next.ServeHTTP(w, r)
		}
		start := time.Now()
		next.ServeHTTP(w, r)
		common.PromSummaryRequestLatency.Observe(float64(time.Since(start)))
	})
}

func jsonData(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var jsonDataReq module.JsonDataReq
	err := decoder.Decode(&jsonDataReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonDataResp := common.ServiceJsonData(&jsonDataReq)
	bytes, err := json.Marshal(jsonDataResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}
