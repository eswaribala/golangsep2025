package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitorsalgado/mocha/v3"
	"github.com/vitorsalgado/mocha/v3/expect"
	"github.com/vitorsalgado/mocha/v3/reply"
)

func TestFetchUser(t *testing.T) {
	m := mocha.New(t).CloseOnCleanup(t)
	m.Start()

	user := User{ID: 1, Name: "nice-name", Email: "sample@gmail.com"}

	m.AddMocks(mocha.
		Get(expect.URLPath("/users/1")).
		Reply(reply.OK().Status(
			http.StatusOK,
		).BodyJSON(user)))
	user, _ = FetchUser(m.URL())

	//assert.NoError(t, err)
	assert.Equal(t, "nice-name", user.Name)
	assert.Equal(t, "sample@gmail.com", user.Email)
}
