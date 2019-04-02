package alarm

import (
	"go.uber.org/zap"
)

// Alarm is the implementation of alarm's interface
type Alarm struct {
	sugar     *zap.SugaredLogger
	lastPrice PriceList
}

//check do all the check necesary to determine if the price meet the firing condition
func (al *Alarm) check() bool {
	//TODO: implement this
	return true
}

//fire send notifications and return error if crahed into err.
func (al *Alarm) fire() error {
	//TODO: implement this
	return nil
}

//CheckAndFire looks through a price list to
func (al *Alarm) CheckAndFire(PriceList) error {
	return nil
}
