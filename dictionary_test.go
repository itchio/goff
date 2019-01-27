package goff_test

import (
	"testing"

	"github.com/itchio/goff"
	"github.com/stretchr/testify/assert"
)

func TestDictionary(t *testing.T) {
	assert := assert.New(t)
	var err error

	var d *goff.Dictionary
	defer d.Free()

	var s string
	var ok bool

	s, ok = d.Get("a")
	assert.Empty(s)
	assert.False(ok)

	d, err = d.Set("a", "algo")
	assert.NoError(err)
	s, ok = d.Get("a")
	assert.True(ok)
	assert.EqualValues("algo", s)

	d, err = d.Set("a", "amos")
	assert.NoError(err)
	s, ok = d.Get("a")
	assert.True(ok)
	assert.EqualValues("amos", s)

	d, err = d.Set("b", "bob")
	assert.NoError(err)

	d, err = d.Set("c", "chris")
	assert.NoError(err)

	assert.EqualValues([]string{"a", "b", "c"}, d.Keys())

	assert.EqualValues(map[string]string{
		"a": "amos",
		"b": "bob",
		"c": "chris",
	}, d.AsMap())
}
