package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertKey(t *testing.T) {
	assert.Equal(t, "khach-san", convertKey("Khách sạn"))
}
