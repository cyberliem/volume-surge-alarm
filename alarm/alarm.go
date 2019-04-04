package alarm

import (
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/cyberliem/volume-surge-alarm/common"
)

// Alarm is the implementation of alarm's interface
type Alarm struct {
	sugar    *zap.SugaredLogger
	firers   []Firer
	checkers []Checker
	duration time.Duration
}

//check do all the check necesary to determine if the price meet the firing condition
func (al *Alarm) check(pl common.PriceList) (common.ChangeCriteria, error) {
	var result = common.ChangeCriteria{
		Duration: al.duration,
		Scores:   common.Scores{},
	}

	for _, check := range al.checkers {
		cp, err := check.Check(pl)
		if err != nil {
			return cp, err
		}
		al.sugar.Debugw("got a list of change", "len", len(cp.Scores))
		result.Scores = append(result.Scores, cp.Scores...)
	}
	result.SortByscores()
	return result, nil
}

//fire send notifications and return error if crahed into err.
func (al *Alarm) fire(cc common.ChangeCriteria) error {
	var g errgroup.Group

	for _, fire := range al.firers {
		g.Go(func() error {
			return fire.Fire(cc)
		})
	}
	return g.Wait()
}

//CheckAndFire looks through a price list to
func (al *Alarm) CheckAndFire(pl common.PriceList) error {
	cc, err := al.check(pl)
	if err != nil {
		return err
	}
	return al.fire(cc)
}

//Option defines initial behaviour for alarm.
type Option func(*Alarm) error

//NewAlarm return an Alarm instance and error if occur
func NewAlarm(sugar *zap.SugaredLogger, duration time.Duration, options ...Option) (*Alarm, error) {
	alarm := Alarm{
		sugar:    sugar,
		duration: duration,
	}
	for _, opt := range options {
		if err := opt(&alarm); err != nil {
			return nil, err
		}
	}
	return &alarm, nil
}

//WithChecker return an option to include checker in
func WithChecker(ck Checker) Option {
	return func(al *Alarm) error {
		al.checkers = append(al.checkers, ck)
		return nil
	}
}

//WithFirer return an option to include firer in
func WithFirer(firer Firer) Option {
	return func(al *Alarm) error {
		al.firers = append(al.firers, firer)
		return nil
	}
}
