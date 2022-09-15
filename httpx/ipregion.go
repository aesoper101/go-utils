package httpx

import (
	"embed"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var (
	//go:embed region.xdb
	fs    embed.FS
	cBuff []byte
)

func init() {
	cBuff, _ = fs.ReadFile("region.xdb")
}

// GetRegionFromIP Get region from ip
func GetRegionFromIP(ip string) (string, error) {
	searcher, err := xdb.NewWithBuffer(cBuff)
	if err != nil {
		return "", err
	}
	defer searcher.Close()

	return searcher.SearchByStr(ip)
}

func GetRegionFromUint32IP(ip uint32) (string, error) {
	searcher, err := xdb.NewWithBuffer(cBuff)
	if err != nil {
		return "", err
	}
	defer searcher.Close()

	return searcher.Search(ip)
}
