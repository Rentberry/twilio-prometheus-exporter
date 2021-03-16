package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c collector) getUsageRecords() (*usageRecordsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Usage/Records/Today.json", c.accountId), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.sid, c.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	result := usageRecordsResponse{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != 0 || result.Status != 0 {
		return nil, fmt.Errorf("twilio request failed with code %d and status %d", result.Code, result.Status)
	}

	return &result, nil
}

type usageRecord struct {
	AccountSid string `json:"account_sid"`
	Category   string
	Count      string
	Usage      string
	Price      string
}
type usageRecordsResponse struct {
	Code         int
	Status       int
	UsageRecords []usageRecord `json:"usage_records"`
}
