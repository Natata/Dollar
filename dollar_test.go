package dollar

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDollar(t *testing.T) {
	type Good struct {
		Name   string `json:"name"`
		Amount Dollar `json:"amount"`
	}

	raw := []byte(`{"name":"smart phone","amount":399.99}`)
	g := &Good{}
	err := json.Unmarshal(raw, g)
	assert.Nil(t, err)
	assert.Equal(t, "399.99", g.Amount.String())

	raw01, err := json.Marshal(g)
	assert.Nil(t, err)
	assert.EqualValues(t, raw, raw01)
}
