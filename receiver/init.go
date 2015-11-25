package receiver

import (
	"github.com/eBayClassifiedsGroup/ammonitrix/backends/elastic"
	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

var (
	be_elastic *elastic.Elastic
)

func InitElastic(cfg *config.Config) {
	be_elastic = elastic.New(cfg.Elastic)
}
