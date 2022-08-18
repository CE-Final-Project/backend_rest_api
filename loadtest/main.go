package main

import (
	"log"
	"net/http"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func SimpleLoadTest() {
	rate := vegeta.Rate{
		Freq: 1,
		Per:  1,
	}

	duration := 1 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: http.MethodGet,
		URL:    "http://localhost:6767/accounts/1",
	})
	attacker := vegeta.NewAttacker()
	vegeta.Connections(0)(attacker)
	vegeta.KeepAlive(false)(attacker)
	vegeta.Timeout(20 * time.Second)(attacker)

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "test") {
		if res.Code != 200 {
			log.Println(res.Error)
		}
		metrics.Add(res)
	}
	metrics.Close()

	report := vegeta.NewTextReporter(&metrics)
	err := report(os.Stdout)
	if err != nil {
		log.Fatalf("Generate report error ===> %v", err)
	}
}

// LoadTestGetRefreshToken load tests refresh token route
func LoadTestGetRefreshToken() {
	rate := vegeta.Rate{
		Freq: 1,
		Per:  1,
	}
	duration := 1 * time.Second
	headers := http.Header{}
	headers.Add("Authorization", "Basic Zm9vYmFyc2FkZmFkZnNmc2Rm")
	headers.Add("Content-Type", "application/x-www-form-urlencoded")
	headers.Add("X-Secret", "mysecret")
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "http://localhost:8080/auth",
		Header: headers,
		Body:   []byte("grant_type=refresh_token&refresh_token=J1lrRYzJR1lY%2BkG4XsQaA4cxNDS2K0cioJAa7ortvxc%3D"),
	})
	attacker := vegeta.NewAttacker()
	vegeta.Connections(0)(attacker)
	vegeta.KeepAlive(false)(attacker)
	vegeta.Timeout(20 * time.Second)(attacker)

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "test") {
		if res.Code != 200 {
			log.Println(res.Error)
		}
		metrics.Add(res)
	}
	metrics.Close()

	report := vegeta.NewTextReporter(&metrics)
	report(os.Stdout)
}

func main() {
	//LoadTestGetRefreshToken()
	SimpleLoadTest()
}
