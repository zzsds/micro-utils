package nacos

import (
	"encoding/json"

	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/nacos_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/common/http_agent"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// Nacos ...
type Nacos struct {
	Endpoint    string
	NamespaceID string
	AccessKey   string
	SecretKey   string
}

var serverConfig = constant.ServerConfig{
	IpAddr:      "console.nacos.io",
	Port:        80,
	ContextPath: "/nacos",
}

// NewNacos ...
func NewNacos(endpoint, namespaceID, accessKey, secretKey string) *Nacos {
	return &Nacos{
		Endpoint:    endpoint,
		NamespaceID: namespaceID,
		AccessKey:   accessKey,
		SecretKey:   secretKey,
	}
}

// Client ...
func (n *Nacos) Client() (configClient config_client.ConfigClient, err error) {
	clientConfig := constant.ClientConfig{
		Endpoint:    n.Endpoint + ":8080",
		NamespaceId: n.NamespaceID,
		AccessKey:   n.AccessKey,
		SecretKey:   n.SecretKey,

		TimeoutMs:           5 * 1000,
		ListenInterval:      30 * 1000,
		BeatInterval:        5 * 1000,
		NotLoadCacheAtStart: true,
		// UpdateCacheWhenEmpty: true,
		// CacheDir:             "./data/nacos/cache",
		// LogDir:               "./data/nacos/log",
	}

	nc := nacos_client.NacosClient{}
	nc.SetServerConfig([]constant.ServerConfig{serverConfig})
	nc.SetClientConfig(clientConfig)
	nc.SetHttpAgent(&http_agent.HttpAgent{})
	return config_client.NewConfigClient(&nc)
}

// Get ...
func (n *Nacos) Get(dataID, group string, data interface{}) error {
	configClient, err := n.Client()
	if err != nil {
		return err
	}
	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group})
	return json.Unmarshal([]byte(content), &data)

}

// Listen ...
func (n *Nacos) Listen(dataID, group string, data interface{}) error {
	configClient, err := n.Client()
	if err != nil {
		return err
	}
	if err := n.Get(dataID, group, data); err != nil {
		return err
	}
	// 监听配置
	return configClient.ListenConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
		OnChange: func(namespace, group, dataId, content string) {
			// fmt.Println("ListenConfig namespace: " + namespace + " group:" + group + ", dataId:" + dataId + ", content:" + content)
			if err := json.Unmarshal([]byte(content), data); err != nil {
				panic(err)
			}
		},
	})
}
