package omisego_test

import (
	omg "github.com/Alainy/OmiseGo-Go-SDK"
	"net/url"
	"testing"
)

var (
	accessKey = "dqG1xQec9uWRCSUyg70zWZ_MFYKXc5n1vBVNzajejBw"
	secretKey = "4yd0m3eFR9yWO3niEnTUodM0ZfSg0KsKkRSbwK1U4_J_SY3IdCEbDcUsADm2-X1VVulFyVPfcli7wMX5NJsgY0mzg6VTd8ul9isAbVoEomGl0TwENPRCUnC444CSEpyyCthHf7CFaMitCbc2f0KK0lQxarAkkLRlfL-EAmLYiAg"

	adminURL = &url.URL{
		Scheme: "http",
		Host:   "localhost:4000",
		Path:   "/api/admin",
	}
)

func TestAdminAll(t *testing.T) {
	c, _ := omg.NewClient(accessKey, secretKey, adminURL)
	adminClient := omg.AdminAPI{
		Client: c,
	}

	body := omg.ListParams{}
	res, err := adminClient.AdminAll(body)

	if err != nil {
		t.Fatal(err)
	}
}
