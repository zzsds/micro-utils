package nacos

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/micro/go-micro/v2/config/source"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type watcher struct {
	opts source.Options
	name string

	sync.RWMutex
	cs *source.ChangeSet

	ch   chan *source.ChangeSet
	exit chan bool
}

func newWatcher(client config_client.IConfigClient, cs *source.ChangeSet, opts source.Options) (source.Watcher, error) {
	dataID, ok := opts.Context.Value(dataIDKey{}).(string)
	if !ok {
		return nil, errors.New("accessKey not null")
	}
	group, ok := opts.Context.Value(groupKey{}).(string)
	if !ok {
		return nil, errors.New("group not null")
	}
	w := &watcher{
		opts: opts,
		name: "nacos",
		cs:   cs,
		ch:   make(chan *source.ChangeSet),
		exit: make(chan bool),
	}

	err := client.ListenConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
			cs := &source.ChangeSet{
				Timestamp: time.Now(),
				Source:    w.name,
				Data:      []byte(data),
				Format:    opts.Encoder.String(),
			}
			cs.Checksum = cs.Sum()
			w.ch <- cs
		},
	})
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	select {
	case cs := <-w.ch:
		return cs, nil
	case <-w.exit:
		return nil, errors.New("watcher stopped")
	}
}

func (w *watcher) Stop() error {
	select {
	case <-w.exit:
		return nil
	default:
		close(w.exit)
	}
	return nil
}
