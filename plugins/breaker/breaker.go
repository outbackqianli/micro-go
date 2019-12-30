package breaker

import (
	"errors"
	"fmt"
	"net/http"
	statusCode "outback/micro-go/plugins/http"

	"github.com/afex/hystrix-go/hystrix"
)

//BreakerWrapper hystrix breaker
func BreakerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		name := r.Method + "-" + r.RequestURI

		runFunc := func() error {
			sct := &statusCode.StatusCodeTracker{ResponseWriter: w, Status: http.StatusOK}
			h.ServeHTTP(sct.WrappedResponseWriter(), r)
			fmt.Printf("降级时 status is %d", sct.Status)
			if sct.Status >= http.StatusInternalServerError {
				str := fmt.Sprintf("status code %d", sct.Status)
				return errors.New(str)
			}
			return nil
		}
		fallbackFunc := func(e error) error {
			//if e == hystrix.ErrCircuitOpen {
			if e != nil {
				fmt.Printf("触发了降级：errr is %s", e.Error())
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("请稍后重试"))
			}
			return e
		}

		hystrix.Do(name, runFunc, fallbackFunc)
	})
}
