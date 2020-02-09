package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDummyData(t *testing.T) {
	assert.Equal(t, "{123:22}", GetDummyData())

	//t.Run("Subtest fo r fun", func(t *testing.T) {
	//	assert.Equal(t, "{123:22}", getDummyData())
	//})

}
