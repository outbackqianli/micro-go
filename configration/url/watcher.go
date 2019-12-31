package url

import (
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro/config/source"
)

type UrlWatcher struct {
	u    *UrlSource
	exit chan bool
}

func newWatcher(u *UrlSource) (*UrlWatcher, error) {
	return &UrlWatcher{
		u:    u,
		exit: make(chan bool),
	}, nil
}

func (u *UrlWatcher) Next() (*source.ChangeSet, error) {
	//select {
	//case <-u.exit:
	//	return nil, errors.New("url watcher stopped")
	//default:
	log.Info("进入next 方法 default")
	changeSet, err := u.u.Read()
	if err != nil {
		log.Info("UrlWatcher error ", err.Error())
		return nil, err
	}
	return changeSet, nil
}

//}

func (u *UrlWatcher) Stop() error {
	select {
	case <-u.exit:
	default:
	}
	return nil
}
