package service

type Base struct {
	aUid *uint
}

func (b *Base) GetAUid() *uint {
	return b.aUid
}
