package rank

import (
	"fmt"
	"sync"
)

type RankNode struct {
	Score     int64  `json:"score,omitempty"`
	Rank      int64  `json:"rank,omitempty"`
	MemberKey string `json:"member_key,omitempty"`
	Ctime     int64  `json:"ctime,omitempty"`
}

type RankListPo struct {
	RankKey  string `json:"rank_key,omitempty"`
	RankId   int64  `json:"rank_id,omitempty"`
	Count    int64  `json:"count,omitempty"`
	Capacity int64  `json:"capacity,omitempty"`
}
type RankList struct {
	RankListPo
	sync.Mutex
	RankNodes map[string]*RankNode //fixme 可以考虑采用三方库自带排序
}

func (rank *RankList) GetRankMajorKey() string {
	return GetRankMajorKey(rank.RankId, rank.RankKey)
}

func GetRankMajorKey(id int64, key string) string {
	return fmt.Sprintf("%d:%s", id, key)
}
