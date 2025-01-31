package lowcode

import (
	"os"

	"github.com/zhs007/goutils"
	sgc7game "github.com/zhs007/slotsgamecore7/game"
	"github.com/zhs007/slotsgamecore7/mathtoolset"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type StatsConfig struct {
	Name      string         `yaml:"name"`
	Component string         `yaml:"component"`
	Children  []*StatsConfig `yaml:"children"`
}

type ComponentConfig struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	Config string `yaml:"config"`
}

type GameModConfig struct {
	Type       string             `yaml:"type"`
	Components []*ComponentConfig `yaml:"components"`
}

type Config struct {
	Width            int                            `yaml:"width"`
	Height           int                            `yaml:"height"`
	Linedata         map[string]string              `yaml:"linedata"`
	MapLinedate      map[string]*sgc7game.LineData  `yaml:"-"`
	Paytables        map[string]string              `yaml:"paytables"`
	MapPaytables     map[string]*sgc7game.PayTables `yaml:"-"`
	Reels            map[string]string              `yaml:"reels"`
	MapReels         map[string]*sgc7game.ReelsData `yaml:"-"`
	SymbolsViewer    string                         `yaml:"symbolsViewer"`
	DefaultScene     string                         `yaml:"defaultScene"`
	DefaultPaytables string                         `yaml:"defaultPaytables"`
	DefaultLinedata  string                         `yaml:"defaultLinedata"`
	Bets             []int                          `yaml:"bets"`
	GameMods         []*GameModConfig               `yaml:"gamemods"`
	StatsSymbols     []string                       `yaml:"statsSymbols"`
	StatsSymbolCodes []mathtoolset.SymbolType       `yaml:"-"`
	Stats            *StatsConfig                   `yaml:"stats"`
}

func (cfg *Config) BuildStatsSymbolCodes(paytables *sgc7game.PayTables) error {
	cfg.StatsSymbolCodes = nil
	for _, v := range cfg.StatsSymbols {
		symbolCode, isok := paytables.MapSymbols[v]
		if !isok {
			goutils.Error("Config.BuildStatsSymbolCodes",
				zap.String("symbol", v),
				zap.Error(ErrIvalidStatsSymbolsInConfig))

			return ErrIvalidStatsSymbolsInConfig
		}

		cfg.StatsSymbolCodes = append(cfg.StatsSymbolCodes, mathtoolset.SymbolType(symbolCode))
	}

	return nil
}

func LoadConfig(fn string) (*Config, error) {
	data, err := os.ReadFile(fn)
	if err != nil {
		goutils.Error("LoadConfig:ReadFile",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		goutils.Error("LoadConfig:Unmarshal",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	if len(cfg.Linedata) > 0 {
		cfg.MapLinedate = make(map[string]*sgc7game.LineData)

		for k, v := range cfg.Linedata {
			ld, err := sgc7game.LoadLineDataFromExcel(v)
			if err != nil {
				goutils.Error("LoadConfig:LoadLineDataFromExcel",
					zap.String("key", k),
					zap.String("linedatafn", v),
					zap.String("fn", fn),
					zap.Error(err))

				return nil, err
			}

			cfg.MapLinedate[k] = ld
		}
	}

	cfg.MapPaytables = make(map[string]*sgc7game.PayTables)

	for k, v := range cfg.Paytables {
		pt, err := sgc7game.LoadPaytablesFromExcel(v)
		if err != nil {
			goutils.Error("LoadConfig:LoadPaytablesFromExcel",
				zap.String("key", k),
				zap.String("paytablesfn", v),
				zap.String("fn", fn),
				zap.Error(err))

			return nil, err
		}

		cfg.MapPaytables[k] = pt
	}

	pt, isok := cfg.MapPaytables["main"]
	if !isok {
		if err != nil {
			goutils.Error("LoadConfig",
				zap.String("fn", fn),
				zap.Error(ErrMustHaveMainPaytables))

			return nil, ErrMustHaveMainPaytables
		}
	}

	if len(cfg.Reels) > 0 {
		cfg.MapReels = make(map[string]*sgc7game.ReelsData)

		for k, v := range cfg.Reels {
			rd, err := sgc7game.LoadReelsFromExcel2(v, pt)
			if err != nil {
				goutils.Error("LoadConfig:LoadPaytablesFromExcel",
					zap.String("key", k),
					zap.String("paytablesfn", v),
					zap.String("fn", fn),
					zap.Error(err))

				return nil, err
			}

			cfg.MapReels[k] = rd
		}
	}

	return cfg, nil
}
