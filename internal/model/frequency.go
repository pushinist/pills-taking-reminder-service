package model

type Frequency string

const (
	OnceDaily   Frequency = "once_daily"
	TwiceDaily  Frequency = "twice_daily"
	ThriceDaily Frequency = "thrice_daily"
	FourthDaily Frequency = "fourth_daily"
	Hourly      Frequency = "hourly"
)
