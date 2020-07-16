package nacos

import (
	"os"
	"time"

	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/nacos_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/common/http_agent"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var (
	endpoint    = GetEnvDefault("CFG_ENDPOINT", "acm.aliyun.com:8080")
	namespaceID = GetEnvDefault("CFG_NAMESPACEID", "a0630038-0d1c-4002-8854-0c08c47fa3e3")
	access      = GetEnvDefault("CFG_ACCESSKEY", "")
	secret      = GetEnvDefault("CFG_SECRETKEY", "")
	data        = GetEnvDefault("CFG_DATA", "srv.user")
	group       = GetEnvDefault("CFG_GROUP", "dev")
)

// nacos ...
type nacos struct {
	client config_client.IConfigClient
	opts   source.Options
	err    error
}

func (n *nacos) Read() (*source.ChangeSet, error) {
	if n.err != nil {
		return nil, n.err
	}
	dataID, ok := n.opts.Context.Value(dataIDKey{}).(string)
	if !ok {
		dataID = data
	}
	groupID, ok := n.opts.Context.Value(groupKey{}).(string)
	if !ok {
		groupID = group
	}
	content, err := n.client.GetConfig(vo.ConfigParam{DataId: dataID, Group: groupID})
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

func (n *nacos) Write(cs *source.ChangeSet) error {
	if n.err != nil {
		return n.err
	}
	dataID, ok := n.opts.Context.Value(dataIDKey{}).(string)
	if !ok {
		dataID = data
	}
	groupID, ok := n.opts.Context.Value(groupKey{}).(string)
	if !ok {
		groupID = group
	}
	_, err := n.client.PublishConfig(vo.ConfigParam{
		DataId:  dataID,
		Group:   groupID,
		Content: string(cs.Data)})
	return err
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
	return newWatcher(n.client, cs, n.opts)
}

// NewAutoSource ...
func NewAutoSource(opts ...source.Option) source.Source {
	ss := file.NewSource(opts...)
	if _, err := ss.Read(); err != nil {
		ss = NewSource(opts...)
	}
	return ss
}

// NewSource ...
func NewSource(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)
	endpointKey, ok := options.Context.Value(endpointKey{}).(string)
	if !ok {
		endpointKey = endpoint
	}
	namespaceIDKey, ok := options.Context.Value(namespaceIDKey{}).(string)
	if !ok {
		namespaceIDKey = namespaceID
	}
	accessKey, ok := options.Context.Value(accessKey{}).(string)
	if !ok {
		accessKey = access
	}

	secretKey, ok := options.Context.Value(secretKey{}).(string)
	if !ok {
		secretKey = secret
	}

	clientConfig := constant.ClientConfig{
		Endpoint:    endpointKey,
		NamespaceId: namespaceIDKey,
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
		client: &c,
		opts:   options,
		err:    err,
	}
}

// GetEnvDefault ...
func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
