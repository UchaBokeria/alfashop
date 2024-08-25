package newvideos

type NewVideoDto struct {
	Video   []byte `form:"file"`
}
