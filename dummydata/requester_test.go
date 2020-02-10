package dummydata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetResponse(t *testing.T) {
	assert.Equal(t, `{"data":123, "data1": 22}`, GetResponse())
}

func TestZendeskRequest(t *testing.T) {
	//assert.Equal(t, "{data:123}", GetResponse())
}

func MockRequester(_ string) []byte {
	return []byte(`{"origin":"123"}`)
}

func TestSimpleGet(t *testing.T) {
	resp := SimpleGet(MockRequester, "https://httpbin.org/ip")
	assert.IsType(t, resp, &httpBinIpResponse{})
	assert.NotEmpty(t, resp.Origin)
	assert.Equal(t, "123", resp.Origin)
}
