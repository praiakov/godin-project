package entity_test

import (
	"testing"
	"time"

	"github.com/praiakov/godin/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	dueDate := time.Date(2022, time.April, 15, 0, 0, 0, 0, time.Now().Location())
	paidDate := time.Date(2023, time.April, 15, 0, 0, 0, 0, time.Now().Location())

	u := entity.NewUser("Steve Jobs", "sjobs@apple.com", dueDate, paidDate, 12)
	assert.Nil(t, nil)
	assert.Equal(t, u.Name, "Steve Jobs")
	assert.NotNil(t, u.ID)
}
