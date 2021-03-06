package clock_test

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	clock "github.com/omniboost/go-clock"
)

func TestChargeLogs(t *testing.T) {
	userName := os.Getenv("CLOCK_USER_NAME")
	apiKey := os.Getenv("CLOCK_API_KEY")
	subscriptionID, err := strconv.Atoi(os.Getenv("CLOCK_SUBSCRIPTION_ID"))
	if err != nil {
		t.Error(err)
	}
	accountID, err := strconv.Atoi(os.Getenv("CLOCK_ACCOUNT_ID"))
	if err != nil {
		t.Error(err)
	}

	client := clock.NewClient(nil, userName, apiKey)
	client.SetDebug(true)
	client.SetSubscriptionID(subscriptionID)
	client.SetAccountID(accountID)
	client.SetDisallowUnknownFields(true)

	req := client.NewChargeLogsRequest()

	req.PathParams().Date = time.Date(2018, 12, 5, 0, 0, 0, 0, time.UTC)
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
