package common

import (
	"sort"
	"time"
)

//PriceList is the map of symbol-price return from fetcher service.
type PriceList map[string]float64

//Scores abstract the list of score change
type Scores []Score

//Len return the len of score list
func (pl Scores) Len() int {
	return len(pl)
}

//Swap swap two member of score list base on their indexes
func (pl Scores) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

//Less determine if a score is less than the other score
func (pl Scores) Less(i, j int) bool {
	return pl[i].Percent < pl[j].Percent
}

//ChangeCriteria is the map of symbol-change score from fetcher service
type ChangeCriteria struct {
	Duration time.Duration
	Scores   Scores
}

//Score is the struct to hold score change of a pair.
type Score struct {
	Pair    string
	Percent float64
}

//SortByscores sort changeCriteria by score
func (ch ChangeCriteria) SortByscores() {
	sort.Sort(ch.Scores)
}
