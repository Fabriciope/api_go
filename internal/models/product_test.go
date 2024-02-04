package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPRoduct(t *testing.T) {
	p, err := NewProduct("Phone", 1500)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Phone", p.Name)
	assert.Equal(t, 1500, p.Price)
	assert.Nil(t, p.Validate())
}

func TestProductWhenNameIsEmpty(t *testing.T) {
	p, err := NewProduct("", 1500)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsEmpty(t *testing.T) {
	p, err := NewProduct("Phone", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Phone", -150)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}