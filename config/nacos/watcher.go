package nacos

import "github.com/micro/go-micro/v2/config/source"

type watcher struct {
}

func newWatcher() (source.Watcher, error) {
	return nil, nil
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	return nil, nil
}

func (w *watcher) Stop() error {
	return nil
}
