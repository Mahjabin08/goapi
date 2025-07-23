package middleware

import (
	"errors"
	"net/http"
	"github.com/Mahjabin08/goapi/api"
	"github.com/Mahjabin08/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")
//it needs to take in and return a http handler interface
func Authorization(next http.Handler) http.Handler {
	
}