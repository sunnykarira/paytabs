package app

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewContextForApp(t *testing.T) {
	app := App{
		externals: map[string]interface{}{
			"rand": rand.Int(),
		},
	}
	ctx := NewContext(app)
	newApp := ctx.App
	if app.externals["rand"].(int) != newApp.externals["rand"].(int) {
		t.Error("Invalid app found in context")
	}
}

func TestNewContextForValues(t *testing.T) {
	app := App{}
	ctx := NewContext(app)

	value := rand.Int()
	ctx.AddValue("rand", value)
	newValue := ctx.GetValue("rand").(int)

	if value != newValue {
		t.Errorf("Context values mismached Expeced %d; Got %d", value, newValue)
	}

	// Test for invalid values should return nil
	invalidKeyValue := ctx.GetValue("invalid_key")

	if invalidKeyValue != nil {
		t.Error("Invalid key does not return nil value")
	}

}

func TestAddRespHeader(t *testing.T) {
	app := App{}
	ctx := NewContext(app)

	key := "content-type"
	value := "html"
	ctx.AddRespHeader(key, value)

	newValue := ctx.ResponseHeaders[key]

	if value != newValue {
		t.Error("Response header not set correctly")
	}
}

func TestDeadline(t *testing.T) {
	app := NewApp()

	t.Run("without dealine", func(t *testing.T) {
		ctx := NewContext(app)
		_, ok := ctx.Deadline()
		assert.Equal(t, false, ok)
	})

	t.Run("with dealine", func(t *testing.T) {
		expectedDeadline := time.Now().Add(50 * time.Millisecond)
		ctx, cancel := context.WithDeadline(NewContext(app), expectedDeadline)
		defer cancel()

		actualDeadline, ok := ctx.Deadline()
		assert.Equal(t, expectedDeadline, actualDeadline)
		assert.Equal(t, true, ok)
	})
}

func TestDone(t *testing.T) {
	app := NewApp()
	ctx, cancel := context.WithCancel(NewContext(app))

	done := false

	cancel()

	select {
	case <-ctx.Done():
		done = true
		return
	}

	assert.Equal(t, true, done)
}

func TestErr(t *testing.T) {
	app := NewApp()
	crx := NewContext(app)
	ctx, cancel := context.WithCancel(crx)
	cancel()
	assert.Nil(t, crx.Err())
	assert.NotNil(t, ctx.Err())
	assert.Equal(t, "context canceled", ctx.Err().Error())
}

func TestValue(t *testing.T) {
	app := NewApp()
	key := "new"
	value := "test"
	crx := NewContext(app)
	ctx := context.WithValue(crx, key, value)
	actualValue := ctx.Value(key).(string)
	crxValue := crx.Value(key)
	assert.Equal(t, value, actualValue)
	assert.Nil(t, crxValue)
}
