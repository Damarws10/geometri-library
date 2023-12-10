package shape

type BangunDatar struct{
	width float32
	height float32
}

func (Bdatar *BangunDatar) Luas() float32 {
	return (Bdatar.width * Bdatar.height) + 2
}

func (Bdatar *BangunDatar) Keliling() float32 {
	return (Bdatar.width + Bdatar.height) / 2
}