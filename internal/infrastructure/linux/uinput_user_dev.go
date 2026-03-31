package linux

type UinputUserDev struct {
	Name       [80]byte
	ID         InputID
	EffectsMax uint32
	Absmax     [64]int32
	Absmin     [64]int32
	Absfuzz    [64]int32
	Absflat    [64]int32
}
