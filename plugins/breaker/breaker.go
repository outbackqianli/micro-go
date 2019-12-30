package breaker

import (
	"errors"
	"fmt"
	nethttp "net/http"
	"outback/micro-go/plugins/http"

	"github.com/afex/hystrix-go/hystrix"
)

//BreakerWrapper hystrix breaker
func BreakerWrapper(h nethttp.Handler) nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {

		name := r.Method + "-" + r.RequestURI

		//runFunc := func() error {
		//	fmt.Println("开始降级 runFunc")
		//	//w.WriteHeader(nethttp.StatusOK)
		//	//w.Write([]byte("Ok"))
		//	h.ServeHTTP(w, r)
		//	return nil
		//}
		runFunc := func() error {
			sct := &http.StatusCodeTracker{ResponseWriter: w, Status: nethttp.StatusOK}
			h.ServeHTTP(sct.WrappedResponseWriter(), r)

			if sct.Status >= nethttp.StatusInternalServerError {
				str := fmt.Sprintf("status code %d", sct.Status)
				return errors.New(str)
			}
			return nil
		}

		fallbackFunc := func(e error) error {
			fmt.Println("开始降级 fallbackFunc")
			if e == hystrix.ErrCircuitOpen {
				//if e != nil {
				fmt.Printf("触发了降级：errr is %s\n", e.Error())
				w.WriteHeader(nethttp.StatusAccepted)
				w.Write([]byte("请稍后重试"))
			}
			fmt.Println("触发了降级，但不是ErrCircuitOpen，error is ", e)
			return e
		}

		hystrix.Do(name, runFunc, fallbackFunc)

	})
}
