package alarm

import (
	"sort"
	"time"
)

//PriceList is the map of symbol-price return from fetcher service.
type PriceList map[string]float64

//PercentageList abstract the list of percentage change
type PercentageList []Percentage

//Len return the len of percentage list
func (pl PercentageList) Len() int {
	return len(pl)
}

//Swap swap two member of percentage list base on their indexes
func (pl PercentageList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

//Less determine if a percentage is less than the other percentage
func (pl PercentageList) Less(i, j int) bool {
	return pl[i].Percent < pl[j].Percent
}

//ChangePercentage is the map of symbol-change percentage from fetcher service
type ChangePercentage struct {
	Duration    time.Time
	Percentages PercentageList
}

//Percentage is the struct to hold percentage change of a pair.
type Percentage struct {
	Pair    string
	Percent float64
}

func (ch *ChangePercentage) sortByPercentages() {
	sort.Sort(ch.Percentages)
}
