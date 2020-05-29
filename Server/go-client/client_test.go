package client

import "testing"

func TestDO(t *testing.T) {
	str, err := DO(
		"./config.toml",
		"CSV2JSON",
		Args{
			File:      "../../_data/data.csv",
			WholeDump: true,
			ToNATS:    false,
		})
	fPln(str)
	fPln(err)
}
