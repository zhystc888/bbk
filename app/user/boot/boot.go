package boot

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/config/nacos/v2"
	rnacos "github.com/gogf/gf/contrib/registry/nacos/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"os"
)

func init() {
	// 环境变量决定使用本地配置还是 Nacos
	env := os.Getenv("APP_ENV") // 例如 "local" 或 "dev"
	if env == "local" {
		// 本地环境：使用文件配置
		localAdapter, err := gcfg.NewAdapterFile("manifest/config/config.yaml")
		if err != nil {
			panic(err)
		}
		g.Cfg().SetAdapter(localAdapter)
	} else {
		ctx := gctx.GetInitCtx()
		baseCfgPath := "manifest/config/base.yaml"
		// 注册中心
		getNacosRegister(ctx, baseCfgPath)
		// 配置中心
		adapter, err := getNacosAdapter(ctx, baseCfgPath)
		if err != nil {
			panic(err)
		}
		g.Cfg().SetAdapter(adapter)
	}

	// 异步日志
	g.Log().SetAsync(true)
}

func getNacosClientConfig(ctx context.Context, baseCfgPath string) constant.ClientConfig {
	return constant.ClientConfig{
		CacheDir:    g.Cfg(baseCfgPath).MustGet(ctx, "nacos.cacheDir").String(),
		LogDir:      g.Cfg(baseCfgPath).MustGet(ctx, "nacos.logDir").String(),
		NamespaceId: g.Cfg(baseCfgPath).MustGet(ctx, "nacos.namespaceId").String(),
		Username:    g.Cfg(baseCfgPath).MustGet(ctx, "nacos.username").String(),
		Password:    g.Cfg(baseCfgPath).MustGet(ctx, "nacos.password").String(),
	}
}

func getNacosRegister(ctx context.Context, baseCfgPath string) {
	addr := g.Cfg(baseCfgPath).MustGet(ctx, "nacos.address").String()
	port := g.Cfg(baseCfgPath).MustGet(ctx, "nacos.port").Int()
	
	gsvc.SetRegistry(rnacos.New(fmt.Sprintf("%s:%d", addr, port), func(config *constant.ClientConfig) {
		*config = getNacosClientConfig(ctx, baseCfgPath) // 覆盖 config 指向的结构体内容
	}))
}

func getNacosAdapter(ctx context.Context, baseCfgPath string) (adapter gcfg.Adapter, err error) {
	var (
		serverConfig = constant.ServerConfig{
			IpAddr:      g.Cfg(baseCfgPath).MustGet(ctx, "nacos.address").String(),
			Port:        g.Cfg(baseCfgPath).MustGet(ctx, "nacos.port").Uint64(),
			GrpcPort:    g.Cfg(baseCfgPath).MustGet(ctx, "nacos.gprc.port").Uint64(),
			ContextPath: "/nacos",
		}
		clientConfig = getNacosClientConfig(ctx, baseCfgPath)
		configParam  = vo.ConfigParam{
			DataId: g.Cfg(baseCfgPath).MustGet(ctx, "nacos.dataID").String(),
			Group:  g.Cfg(baseCfgPath).MustGet(ctx, "nacos.group").String(),
		}
	)

	// 在 getNacosAdapter 中添加
	if err := os.MkdirAll(clientConfig.LogDir, 0755); err != nil {
		g.Log().Fatalf(ctx, "Failed to create nocos log dir %s: %v", clientConfig.LogDir, err)
	}
	if err := os.MkdirAll(clientConfig.CacheDir, 0755); err != nil {
		g.Log().Fatalf(ctx, "Failed to create nocos cache dir %s: %v", clientConfig.CacheDir, err)
	}

	// Create anacosClient that implements gcfg.Adapter.
	return nacos.New(ctx, nacos.Config{
		ServerConfigs: []constant.ServerConfig{serverConfig},
		ClientConfig:  clientConfig,
		ConfigParam:   configParam,
	})
}
