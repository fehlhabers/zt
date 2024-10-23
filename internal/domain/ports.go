package domain

type Storer interface {
	GetZtream() (Ztream, error)
	StoreZtream(Ztream) error
}
