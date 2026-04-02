package factory

import (
	"io"

	"github.com/TerraDharitri/drt-go-bridge-eth/api/gin"
	"github.com/TerraDharitri/drt-go-bridge-eth/config"
	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-bridge-eth/facade"
)

// StartWebServer creates and starts a web server able to respond with the metrics holder information
func StartWebServer(configs config.Configs, metricsHolder core.MetricsHolder) (io.Closer, error) {
	argsFacade := facade.ArgsRelayerFacade{
		MetricsHolder: metricsHolder,
		ApiInterface:  configs.FlagsConfig.RestApiInterface,
		PprofEnabled:  configs.FlagsConfig.EnablePprof,
	}

	relayerFacade, err := facade.NewRelayerFacade(argsFacade)
	if err != nil {
		return nil, err
	}

	httpServerArgs := gin.ArgsNewWebServer{
		Facade:          relayerFacade,
		ApiConfig:       configs.ApiRoutesConfig,
		AntiFloodConfig: configs.GeneralConfig.WebAntiflood,
	}

	httpServerWrapper, err := gin.NewWebServerHandler(httpServerArgs)
	if err != nil {
		return nil, err
	}

	err = httpServerWrapper.StartHttpServer()
	if err != nil {
		return nil, err
	}

	return httpServerWrapper, nil
}
