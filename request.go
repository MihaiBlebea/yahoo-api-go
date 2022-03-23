package yahooapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Request interface {
	GetResponseBody(url string) (string, error)
}

type HttpRequest struct {
	logger     *logrus.Logger
	silent     bool
	cache      map[string]string
	lastUpdate time.Time
	sync.RWMutex
}

func NewHttpRequestClient(silent bool) *HttpRequest {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	return &HttpRequest{
		logger: logger,
		silent: silent,
		cache:  map[string]string{},
	}
}

func (r *HttpRequest) GetResponseBody(url string) (string, error) {
	val, exists := r.cache[url]
	if exists && !r.shouldRefreshCache() {
		if !r.silent {
			r.logger.Info(fmt.Sprintf("Getting %s from cache", url))
		}
		return val, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	res := string(body)

	r.Lock()
	defer r.Unlock()

	r.cache[url] = res

	r.setCacheUpdatedTime()

	if !r.silent {
		r.logger.Info(fmt.Sprintf("Getting %s from api call", url))
	}

	return res, nil
}

func (r *HttpRequest) GetCache() map[string]string {
	return r.cache
}

func (r *HttpRequest) shouldRefreshCache() bool {
	if r.lastUpdate.IsZero() {
		return true
	}

	now := time.Now()
	duration := time.Minute * 60

	return r.lastUpdate.Add(duration).Before(now)
}

func (r *HttpRequest) setCacheUpdatedTime() {
	r.lastUpdate = time.Now()
}
