package core

import (
	"github.com/xssnick/tonutils-go/ton"
	"testing"
)

func Test_isInShard(t *testing.T) {
	shard, _ := ParseShardID(int64(0x6080000000000000)) // 0x6080000000000000
	testCases := []struct {
		Shard  uint64
		Result bool
	}{
		{0x8000000000000000, true},  // 1000000000000000000000000000000000000000000000000000000000000000
		{0x6800000000000000, true},  // 0110100000000000000000000000000000000000000000000000000000000000
		{0x6080000000000000, true},  // 0110000010000000000000000000000000000000000000000000000000000000
		{0x6040000000000000, true},  // 0110000001000000000000000000000000000000000000000000000000000000 sub shard
		{0x60C0000000000000, true},  // 0110000011000000000000000000000000000000000000000000000000000000 sub shard
		{0x6020000000000000, true},  // 0110000000100000000000000000000000000000000000000000000000000000 sub shard
		{0x2800000000000000, false}, // 0010100000000000000000000000000000000000000000000000000000000000
		{0x6180000000000000, false}, // 0110000110000000000000000000000000000000000000000000000000000000
	}
	for _, c := range testCases {
		if !(shard.MatchBlockID(&ton.BlockIDExt{Shard: int64(c.Shard)}) == c.Result) {
			t.Fatalf("Invalid result for shard %x, must be %v", c.Shard, c.Result)
		}
	}
}