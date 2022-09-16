package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	//"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func TestRegisterValidUsernamePassword(t *testing.T) {

	r := getRouter(true)
	r.POST("/user/register", Register)

	w := httptest.NewRecorder()
	regPayload := payloadValid()
	req := httptest.NewRequest("POST", "/user/register", strings.NewReader(regPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(regPayload)))

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	require.Contains(t, w.Body.String(), "Logout")
}

func TestRegisterUsernameAlreadyTaken(t *testing.T) {

	r := getRouter(true)
	r.POST("/user/register", Register)

	w := httptest.NewRecorder()
	regPayload := payloadValid()
	req := httptest.NewRequest("POST", "/user/register", strings.NewReader(regPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(regPayload)))
	log.Info(regPayload)

	r.ServeHTTP(w, req)

	assert.NotEqual(t, w.Code, http.StatusOK)
	require.NotContains(t, w.Body.String(), "Logout")

}

func TestRegisterUsernameInvalid(t *testing.T) {

	r := getRouter(true)
	r.POST("/user/register", Register)

	w := httptest.NewRecorder()
	regPayload := payloadInvalid()
	req := httptest.NewRequest("POST", "/user/register", strings.NewReader(regPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(regPayload)))

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
	require.NotContains(t, w.Body.String(), "Logout")

}

func TestLoginValidCredentials(t *testing.T) {

	r := getRouter(true)
	r.POST("/user/login", Login)

	w := httptest.NewRecorder()
	payloadValid := payloadValid()
	req := httptest.NewRequest("POST", "/user/login", strings.NewReader(payloadValid))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payloadValid)))

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	require.Contains(t, w.Body.String(), "Logout")

}

func TestLoginInvalidCredentials(t *testing.T) {

	r := getRouter(true)
	r.POST("/user/login", Login)

	w := httptest.NewRecorder()
	payloadInvalid := payloadInvalid()
	req := httptest.NewRequest("POST", "/user/login", strings.NewReader(payloadInvalid))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payloadInvalid)))

	r.ServeHTTP(w, req)

	assert.NotEqual(t, w.Code, http.StatusOK)
	require.NotContains(t, w.Body.String(), "Logout")
}

func TestLogout(t *testing.T) {

	TestLoginValidCredentials(t)
	r := getRouter(true)
	r.GET("/user/logout", Logout)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/user/logout", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNotFound)
	require.NotEmpty(t, w.Body.String())

}

func TestGetUsers(t *testing.T) {

	TestLoginValidCredentials(t)
	r := getRouter(true)
	r.GET("/user/users", GetUsers)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/user/users", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusFound)
	require.NotEmpty(t, w.Body.String())
	require.Contains(t, w.Body.String(), "testing123@testing.com")
}

func TestHomepage(t *testing.T) {

	r := getRouter(true)
	r.GET("/user", Homepage)

	req := httptest.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	require.NotEmpty(t, w.Body.String())
	require.Contains(t, w.Body.String(), "Register")
}
