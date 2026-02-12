package class

type Unity int

const (
	UNITY_ECLAIREUR = iota
	UNITY_ECLAIREUSE
	UNITY_PION
	UNITY_PIONNE
	UNITY_ANIM
)

type Order int

const (
	ORDER_TEMPLAR = iota
	ORDER_TEUTONIC
	ORDER_HOSPITAL
	ORDER_LAZARITE
	ORDER_SECULAR
)

type CacheID string

type User struct {
	ID          string    `json:"id"`
	Player_name string    `json:"player_name"`
	Unity       Unity     `json:"unity"`
	Order       Order     `json:"order"`
	Caches      []CacheID `json:"caches"`
}

func NewCache(id string) CacheID {
	return CacheID(id)
}

func NewUser(id string, order Order) User {
	return User{
		ID:     id,
		Order:  order,
		Caches: nil,
	}
}
