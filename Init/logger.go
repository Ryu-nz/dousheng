package Init

import (
	"dousheng/global"
	"dousheng/utils"
	"fmt"

	"go.uber.org/zap"
)

// InitLogger 初始化Logger
func Logger() {
	// 实例化zap 配置
	zapConfig := zap.NewDevelopmentConfig()
	// 注意global.Settings.LogsAddress是在settings-dev.yaml配置过的
	// 配置日志的输出地址
	zapConfig.OutputPaths = []string{
		fmt.Sprintf("%slog-%s.log", global.Settings.LogsAddress, utils.GetNowFormatTodayTime()),
		"stdout",
	}
	// 创建logger实例
	logger, _ := zapConfig.Build()
	zap.ReplaceGlobals(logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	global.Lg = logger         // 注册到全局变量中
}
