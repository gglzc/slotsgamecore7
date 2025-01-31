package lowcode

import "errors"

var (
	// ErrUnkonow - unknow error
	ErrUnkonow = errors.New("unknow error")

	// ErrMustHaveMainPaytables - must have main paytables
	ErrMustHaveMainPaytables = errors.New("must have main paytables")

	// ErrInvalidGameMod - invalid gamemod
	ErrInvalidGameMod = errors.New("invalid gamemod")

	// ErrInvalidComponent - invalid component
	ErrInvalidComponent = errors.New("invalid component")

	// ErrInvalidReels - invalid reels
	ErrInvalidReels = errors.New("invalid reels")

	// ErrInvalidSymbol - invalid symbol
	ErrInvalidSymbol = errors.New("invalid symbol")

	// ErrInvalidPaytables - invalid paytables
	ErrInvalidPaytables = errors.New("invalid paytables")

	// ErrInvalidGamePropertyString - invalid gameProperty string
	ErrInvalidGamePropertyString = errors.New("invalid gameProperty string")

	// ErrParseScript - parse script error
	ErrParseScript = errors.New("parse script error")
	// ErrNoFunctionInScript - no function in script
	ErrNoFunctionInScript = errors.New("no function in script")
	// ErrWrongFunctionInScript - wrong function in script
	ErrWrongFunctionInScript = errors.New("wrong function in script")

	// ErrIvalidComponentName - invalid component name
	ErrIvalidComponentName = errors.New("invalid component name")

	// ErrIvalidCurGameModParams - invalid CurGameModParams
	ErrIvalidCurGameModParams = errors.New("invalid CurGameModParams")

	// ErrIvalidPlayResultLength - invalid PlayResult Length
	ErrIvalidPlayResultLength = errors.New("invalid PlayResult Length")

	// ErrIvalidMultiLevelReelsConfig - invalid MultiLevelReels config
	ErrIvalidMultiLevelReelsConfig = errors.New("invalid MultiLevelReels config")

	// ErrIvalidStatsSymbolsInConfig - invalid StatsSymbols in config
	ErrIvalidStatsSymbolsInConfig = errors.New("invalid StatsSymbols in config")
	// ErrIvalidStatsComponentInConfig - invalid Stats's component in config
	ErrIvalidStatsComponentInConfig = errors.New("invalid Stats's component in config")
)
