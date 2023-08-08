package rank

import (
	"encoding/json"
	"fmt"
	redis2 "github.com/go-redis/redis"
	"go-demo/src/logic/database"
	"sync"
)

// 基于redis 的排行榜系统
func init() {
	// 1：需要加载db
	//2:更新/删除操作，异步落地
	load()
	syncData()
}

var RanKeyListKey = "s:rank:major:key:list"
var RanInfoListKey = "h:rank:major:info:list"

var RakMemberKey = "z:rank:member:list:%s"

var cmd chan *Cmd
var close chan int
var wg = sync.WaitGroup{}
var supervisor *Supervisor

type Supervisor struct {
	sync.Mutex
	rankPool map[string]*RankList
}

type Cmd struct {
	Op       string
	RankInfo *RankList
}

func getMemberKey(id int64, key string) string {
	mk := GetRankMajorKey(id, key)
	return fmt.Sprintf(RakMemberKey, mk)
}
func getMemberKey2(mk string) string {
	return fmt.Sprintf(RakMemberKey, mk)
}

func getRankKeyListKey() string {
	return RanKeyListKey
}

func load() {
	supervisor = &Supervisor{
		Mutex:    sync.Mutex{},
		rankPool: make(map[string]*RankList),
	}
	redis := database.GetDefaultClient()
	mk := redis.SMembers(getRankKeyListKey())
	ks, err := mk.Result()
	if err != nil {
		fmt.Println("load rank info not found")
		panic(err)
	}
	for _, key := range ks {
		getRankInfo(key)
	}
}

func getRankInfo(majorKey string) *RankList {
	rankInfo, ok := supervisor.rankPool[majorKey]
	if !ok {
		redis := database.GetDefaultClient()
		hget := redis.HGet(RanInfoListKey, majorKey)

		data, err := hget.Result()
		if err != nil {
			panic(fmt.Errorf("rank info not found,key=%s", majorKey))
		}
		if data == "" {
			return nil
		}
		rankInfo = &RankList{}
		err = json.Unmarshal([]byte(data), rankInfo)
		if err != nil {
			panic(fmt.Errorf("rank info unmarshal error,key=%s", majorKey))
		}
		loadMember(rankInfo)
		supervisor.Lock()
		defer supervisor.Unlock()
		supervisor.rankPool[majorKey] = rankInfo
	}
	return rankInfo
}

func loadMember(list *RankList) {
	redis := database.GetDefaultClient()
	ret := redis.ZRangeByScoreWithScores(getMemberKey(list.RankId, list.RankKey), redis2.ZRangeBy{
		Min:    "0",
		Max:    "-1",
		Offset: 0,
		Count:  list.Capacity,
	})
	z, err := ret.Result()
	if err != nil {
		panic("load member error")
	}
	rank := int64(0)
	for _, mz := range z {
		m := mz.Member
		ctime := int64(mz.Score) << 32
		s := int64(mz.Score) >> 32
		node := &RankNode{
			Score:     s,
			Rank:      rank,
			MemberKey: m.(string),
			Ctime:     ctime,
		}
		list.RankNodes[node.MemberKey] = node
		list.Count++
	}

}
func updateNode(list *RankList) {
	list.Lock()
	defer list.Unlock()
	redis := database.GetDefaultClient()
	redis.HSet(RanInfoListKey, list.GetRankMajorKey(), list)
	nodes := list.RankNodes
	for mk, node := range nodes {
		ret := node.Ctime
		ret <<= 32
		ret = ret | node.Score
		redis.ZAdd(getMemberKey(list.RankId, list.RankKey), redis2.Z{
			Score:  float64(ret),
			Member: mk,
		})
	}
}
func deleteRank(rank *RankList) {
	supervisor.Lock()
	defer supervisor.Unlock()
	majorKey := rank.GetRankMajorKey()
	delete(supervisor.rankPool, majorKey)
	redis := database.GetDefaultClient()
	redis.HDel(RanInfoListKey, majorKey)
	redis.SRem(getRankKeyListKey(), majorKey)
	redis.Del(getMemberKey2(majorKey))
}

func syncData() {
	go func() {
		for {
			select {
			case cmd := <-cmd:
				switch cmd.Op {
				case "update":
					up := func() {
						wg.Add(1)
						defer wg.Done()
						updateNode(cmd.RankInfo)
					}
					up()
					break
				case "delete":
					del := func() {
						wg.Add(1)
						defer wg.Done()
						deleteRank(cmd.RankInfo)
					}
					del()
				}
				break

			case <-close:
				wg.Wait()
			}
		}
	}()

}
