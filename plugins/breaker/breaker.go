package breaker

import (
	nethttp "net/http"

	"github.com/micro/go-micro/util/log"

	"github.com/afex/hystrix-go/hystrix"
)

//BreakerWrapper hystrix breaker
func BreakerWrapper(h nethttp.Handler) nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {

		name := r.Method + "-" + r.RequestURI
		log.Info("降级name ", name)
		log.Debug("name")
		runFunc := func() error {
			log.Info("开始降级 runFunc")
			//w.WriteHeader(nethttp.StatusOK)
			//w.Write([]byte("Ok"))
			h.ServeHTTP(w, r)
			return nil
			//return errors.New("hello")
		}
		//runFunc := func() error {
		//	sct := &http.StatusCodeTracker{ResponseWriter: w, Status: nethttp.StatusOK}
		//	h.ServeHTTP(sct.WrappedResponseWriter(), r)
		//
		//	if sct.Status >= nethttp.StatusInternalServerError {
		//		str := fmt.Sprintf("status code %d", sct.Status)
		//		return errors.New(str)
		//	}
		//	return nil
		//}

		fallbackFunc := func(e error) error {
			log.Info("开始降级 fallbackFunc")
			if e == hystrix.ErrCircuitOpen {
				//if e != nil {
				log.Infof("触发了降级：errr is %s\n", e.Error())
				w.WriteHeader(nethttp.StatusAccepted)
				w.Write([]byte("请稍后重试"))
				return nil
			} else {
				log.Infof("触发了降级，但不是ErrCircuitOpen，error is %s \n", e)
			}
			return e
		}

		hystrix.Do("GET-/user/login", runFunc, fallbackFunc)

	})
}
