package boot

import (
	"github.com/gogf/gf/contrib/config/nacos/v2"
	"github.com/gogf/gf/v2/frame/g"
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
		adapter, err := getNacosAdapter()
		if err != nil {
			panic(err)
		}
		g.Cfg().SetAdapter(adapter)
	}

	// 异步日志
	g.Log().SetAsync(true)
}

func getNacosAdapter() (adapter gcfg.Adapter, err error) {
	baseCfgPath := "manifest/config/base.yaml"
	var (
		ctx          = gctx.GetInitCtx()
		serverConfig = constant.ServerConfig{
			IpAddr:      g.Cfg(baseCfgPath).MustGet(ctx, "nacos.address").String(),
			Port:        g.Cfg(baseCfgPath).MustGet(ctx, "nacos.port").Uint64(),
			GrpcPort:    g.Cfg(baseCfgPath).MustGet(ctx, "nacos.gprc.port").Uint64(),
			ContextPath: "/nacos",
		}
		clientConfig = constant.ClientConfig{
			CacheDir:    g.Cfg(baseCfgPath).MustGet(ctx, "nacos.cacheDir").String(),
			LogDir:      g.Cfg(baseCfgPath).MustGet(ctx, "nacos.logDir").String(),
			NamespaceId: g.Cfg(baseCfgPath).MustGet(ctx, "nacos.namespaceId").String(),
			Username:    g.Cfg(baseCfgPath).MustGet(ctx, "nacos.username").String(),
			Password:    g.Cfg(baseCfgPath).MustGet(ctx, "nacos.password").String(),
		}

		configParam = vo.ConfigParam{
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
