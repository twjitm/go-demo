package rank

import (
	"fmt"
	"go-demo/src/logic/database"
	"sync"
	"time"
)

func InitRank(key string, id, capacity int64) *RankList {
	rank := &RankList{
		RankListPo: RankListPo{RankId: id, Capacity: capacity, RankKey: key},
		Mutex:      sync.Mutex{},
		RankNodes:  map[string]*RankNode{},
	}
	redis := database.GetDefaultClient()
	redis.SAdd(getRankKeyListKey(), rank.GetRankMajorKey())
	redis.HSet(RanInfoListKey, rank.GetRankMajorKey(), rank.RankListPo)
	return rank
}

func GetRankInfo(id int64, key string) *RankList {
	return getRankInfo(GetRankMajorKey(id, key))
}

func GetMemberData(id int64, key string, memberKey string) *RankNode {
	rank := GetRankInfo(id, key)
	if rank == nil {
		return nil
	}
	member, ok := rank.RankNodes[memberKey]
	if !ok {
		return nil
	}
	return member
}

func GetRankList(id int64, key string, start int64, end int64) []*RankNode {
	rank := GetRankInfo(id, key)
	if rank == nil {
		return nil
	}

	nodes := rank.RankNodes
	var ret []*RankNode
	for _, node := range nodes {
		if node.Rank >= start && node.Rank < end {
			ret = append(ret, node)
		}
	}
	return ret
}

func DeleteRank(id int64, key string) error {
	rank := GetRankInfo(id, key)
	if rank == nil {
		return nil
	}
	cmd <- &Cmd{
		Op:       "delete",
		RankInfo: rank,
	}
	return nil
}

func AppendNode(id int64, key string, memberKey string, score int64) error {
	rank := GetRankInfo(id, key)
	if rank == nil {
		return fmt.Errorf("rank info nod found")
	}
	rank.Lock()
	defer rank.Unlock()
	count := rank.Count
	rank.RankNodes[memberKey] = &RankNode{
		Score:     score,
		Rank:      count,
		MemberKey: memberKey,
		Ctime:     time.Now().Unix(),
	}
	rank.Count++
	cmd <- &Cmd{
		Op:       "update",
		RankInfo: rank,
	}
	return nil
}

func DeleteNode(id int64, key string, memberKey string) error {
	rank := GetRankInfo(id, key)
	if rank == nil {
		return fmt.Errorf("rank info nod found")
	}
	rank.Lock()
	defer rank.Unlock()
	_, ok := rank.RankNodes[memberKey]
	if !ok {
		return fmt.Errorf("member not found")
	}
	delete(rank.RankNodes, memberKey)
	cmd <- &Cmd{
		Op:       "update",
		RankInfo: rank,
	}
	return nil
}
