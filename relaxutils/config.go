package relaxutils

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type FloatList struct {
	Vals []float32 `xml:"item"`
}

type StringList struct {
	Vals []string `xml:"item"`
}

type Symbol struct {
	XMLName xml.Name `xml:"symbol"`
	Name    string   `xml:"name,attr"`
	Val     int      `xml:"value,attr"`
}

type Symbols struct {
	XMLName xml.Name  `xml:"symbols"`
	Symbols []*Symbol `xml:"symbol"`
}

type PayoutWin struct {
	XMLName xml.Name `xml:"win"`
	Count   int      `xml:"count,attr"`
	Payout  int      `xml:"payout,attr"`
}

type Payout struct {
	XMLName xml.Name `xml:"symbol"`
	Name    string   `xml:"name,attr"`
	Wins    []*PayoutWin
}

type Payouts struct {
	XMLName xml.Name `xml:"payouts"`
	Payouts []*Payout
}

type Table struct {
	XMLName      xml.Name      `xml:"table"`
	TableComment string        `xml:",comment"`
	Reel         []*StringList `xml:"reel"`
}

type Reels struct {
	Tables []*Table
}

type Config struct {
	XMLName            xml.Name    `xml:"config"`
	GeneralComment     string      `xml:",comment"`
	SD                 float32     `xml:"sd"`
	RTP                float32     `xml:"rtp"`
	ID                 int         `xml:"id"`
	Name               string      `xml:"name"`
	ConfigVersion      string      `xml:"configVersion"`
	RTPsComment        string      `xml:",comment"`
	RTPs               *FloatList  `xml:"rtps"`
	ConfigComment      string      `xml:",comment"`
	Symbols            *Symbols    `xml:"symbols"`
	Wilds              *StringList `xml:"wilds"`
	CreditsPerBet      int         `xml:"CREDITS_PER_BET"`
	FSAwarded          int         `xml:"FS_AWARDED"`
	FSAwardedRetrigger int         `xml:"FS_AWARDED_RETRIGGER"`
	Payouts            *Payouts    `xml:"payouts"`
}

func SaveConfig(fn string, cfg interface{}) error {
	output, err := xml.MarshalIndent(cfg, "  ", "    ")
	if err != nil {
		goutils.Error("SaveConfig:MarshalIndent",
			zap.Error(err))

		return err
	}

	xmlhead := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	buf := []byte(xmlhead)
	buf = append(buf, output...)

	err = ioutil.WriteFile(fn, buf, 0644)
	if err != nil {
		goutils.Error("SaveConfig:WriteFile",
			zap.Error(err))

		return err
	}

	return nil
}
