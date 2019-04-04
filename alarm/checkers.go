package alarm

import (
	"math"
	"strings"

	"github.com/cyberliem/volume-surge-alarm/common"
)

//PriceStepChecker is a simple checker which check if a pair has change more than
// %threshold percent since its last price
type PriceStepChecker struct {
	lastPrice common.PriceList
	threshold float64
	bases     []string
}

func calChangePercentage(old, new float64) float64 {
	if old == 0.0 {
		return 0
	}
	return 100.0 * (new - old) / old
}

func (psc *PriceStepChecker) hasBase(pair string) bool {
	for _, base := range psc.bases {
		if strings.HasSuffix(pair, base) {
			return true
		}
	}
	return false
}

//Check implement checker.Check
func (psc *PriceStepChecker) Check(pl common.PriceList) (common.ChangeCriteria, error) {
	var (
		scores common.Scores
		cc     common.ChangeCriteria
	)

	if len(psc.lastPrice) == 0 {
		psc.lastPrice = pl
		return cc, nil
	}
	for pair, newPrice := range pl {
		if !psc.hasBase(pair) {
			continue
		}
		oldPrice, avail := psc.lastPrice[pair]
		if avail {
			change := calChangePercentage(oldPrice, newPrice)
			if math.Abs(change) >= psc.threshold {
				scores = append(scores, common.Score{
					Pair:    pair,
					Percent: change,
				})
			}
		}
	}
	psc.lastPrice = pl
	cc.Scores = scores
	return cc, nil
}
