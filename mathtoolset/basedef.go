package mathtoolset

type SymbolType int

type InReelSymbolType int

const (
	IRSTypeAll               InReelSymbolType = 0
	IRSTypeSymbol            InReelSymbolType = 1
	IRSTypeWild              InReelSymbolType = 2
	IRSTypeSymbolAndWild     InReelSymbolType = 3
	IRSTypeNoSymbolAndNoWild InReelSymbolType = 4
	IRSTypeNoWild            InReelSymbolType = 5
	IRSTypeSymbol2           InReelSymbolType = 6
	IRSTypeSymbol2AndWild    InReelSymbolType = 7
)

func HasSymbol(symbols []SymbolType, symbol SymbolType) bool {
	for _, v := range symbols {
		if v == symbol {
			return true
		}
	}

	return false
}

func NewInReelSymbolTypeArr(num int) []InReelSymbolType {
	arr := make([]InReelSymbolType, num)

	return arr
}

func IsFirstWild(lst []InReelSymbolType, num int) bool {
	for i := 0; i < num; i++ {
		if lst[i] != IRSTypeWild {
			return false
		}
	}

	return true
}
