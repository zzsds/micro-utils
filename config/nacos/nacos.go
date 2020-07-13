package nacos

import (
	"errors"
	"time"

	"github.com/micro/go-micro/v2/config/source"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/nacos_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/common/http_agent"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// nacos ...
type nacos struct {
	client config_client.ConfigClient
	opts   source.Options
	err    error
}

func (n *nacos) Read() (*source.ChangeSet, error) {
	if n.err != nil {
		return nil, n.err
	}
	dataID, ok := n.opts.Context.Value(dataIDKey{}).(string)
	if !ok {
		return nil, errors.New("accessKey not null")
	}
	group, ok := n.opts.Context.Value(groupKey{}).(string)
	if !ok {
		return nil, errors.New("group not null")
	}
	content, err := n.client.GetConfig(vo.ConfigParam{DataId: dataID, Group: group})
	if err != nil {
		return nil, err
	}
	cs := &source.ChangeSet{
		Timestamp: time.Now(),
		Source:    n.String(),
		Data:      []byte(content),
		Format:    n.opts.Encoder.String(),
	}
	cs.Checksum = cs.Sum()
	return cs, nil
}

func (n *nacos) Write(*source.ChangeSet) error {
	return nil
}

func (n *nacos) String() string {
	return "nacos"
}

func (n *nacos) Watch() (source.Watcher, error) {
	if n.err != nil {
		return nil, n.err
	}
	cs, err := n.Read()
	if err != nil {
		return nil, err
	}
	return newWatcher(n.client.Watcher, cs, n.opts)
}

// NewSource ...
func NewSource(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)
	endpointKey, ok := options.Context.Value(endpointKey{}).(string)
	if !ok {
		endpointKey = "acm.aliyun.com:8080"
	}
	namespaceID, ok := options.Context.Value(namespaceIDKey{}).(string)
	if !ok {
		panic("namespaceID not null")
	}
	accessKey, ok := options.Context.Value(accessKey{}).(string)
	if !ok {
		panic("accessKey not null")
	}

	secretKey, ok := options.Context.Value(secretKey{}).(string)
	if !ok {
		panic("secretKey not null")
	}

	clientConfig := constant.ClientConfig{
		Endpoint:    endpointKey,
		NamespaceId: namespaceID,
		AccessKey:   accessKey,
		SecretKey:   secretKey,

		TimeoutMs:           5 * 1000,
		ListenInterval:      30 * 1000,
		BeatInterval:        5 * 1000,
		NotLoadCacheAtStart: true,
		// UpdateCacheWhenEmpty: true,
		// CacheDir:             "./data/nacos/cache",
		// LogDir:               "./data/nacos/log",
	}

	serverConfig := constant.ServerConfig{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}

	nc := nacos_client.NacosClient{}
	nc.SetServerConfig([]constant.ServerConfig{serverConfig})
	nc.SetClientConfig(clientConfig)
	nc.SetHttpAgent(&http_agent.HttpAgent{})
	c, err := config_client.NewConfigClient(&nc)
	return &nacos{
		client: c,
		opts:   options,
		err:    err,
	}
}
