package utils

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestPing(t *testing.T) {
	db := Db()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	err := db.Client().Ping(ctx, readpref.Primary())
	assert.Equal(t, err, nil)
	cancel()
}
