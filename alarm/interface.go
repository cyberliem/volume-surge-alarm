package alarm

//Interface is the interface for firing an alarm.
type Interface interface {
	//Take a PriceList and fire alarm accordingly
	CheckAndFire(PriceList) error
}
