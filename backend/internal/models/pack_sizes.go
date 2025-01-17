package models

type PackSize struct {
	Size int
}

func DefaultPackSizes() []PackSize {
	return []PackSize{
		{Size: 250},
		{Size: 500},
		{Size: 1000},
		{Size: 2000},
		{Size: 5000},
	}
}
