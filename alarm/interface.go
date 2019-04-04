package alarm

import (
	"github.com/cyberliem/volume-surge-alarm/common"
)

//Interface is the interface for firing an alarm.
type Interface interface {
	//Take a PriceList and fire alarm accordingly
	CheckAndFire(common.PriceList) error
}

//Firer define the interface for fire method, that is. to send a change Score to
//specific end
type Firer interface {
	Fire(cp common.ChangeCriteria) error
}

//Checker define interafce for check method
type Checker interface {
	Check(common.PriceList) (common.ChangeCriteria, error)
}
