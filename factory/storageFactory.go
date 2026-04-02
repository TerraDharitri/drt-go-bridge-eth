package factory

import (
	"path"

	"github.com/TerraDharitri/drt-go-bridge-eth/core"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/storage/factory"
	"github.com/TerraDharitri/drt-go-chain/storage/storageunit"
)

// CreateUnitStorer based on the config and the working directory
func CreateUnitStorer(config config.StorageConfig, workingDir string) (core.Storer, error) {
	persisterCreator, err := factory.NewPersisterFactory(config.DB)
	if err != nil {
		return nil, err
	}

	statusMetricsDbConfig := factory.GetDBFromConfig(config.DB)
	dbPath := path.Join(workingDir, config.DB.FilePath)
	statusMetricsDbConfig.FilePath = dbPath

	return storageunit.NewStorageUnitFromConf(
		factory.GetCacherFromConfig(config.Cache),
		statusMetricsDbConfig,
		persisterCreator)
}
