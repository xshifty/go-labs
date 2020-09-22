package graph

type Color uint32

func (c Color) Red() uint8 {
	return uint8(c>>24) & 0xff
}

func (c Color) Green() uint8 {
	return uint8(c>>16) & 0xff
}

func (c Color) Blue() uint8 {
	return uint8(c>>8) & 0xff
}

func (c Color) Alpha() uint8 {
	return uint8(c & 0xff)
}

func (c Color) RedF() float32 {
	return float32(c.Red() / 255.0)
}

func (c Color) GreenF() float32 {
	return float32(c.Green() / 255.0)
}

func (c Color) BlueF() float32 {
	return float32(c.Blue() / 255.0)
}

func (c Color) AlphaF() float32 {
	return float32(c.Alpha() / 255.0)
}
