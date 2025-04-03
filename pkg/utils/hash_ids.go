package utils

import "github.com/speps/go-hashids/v2"

type Hasher struct {
	HashId *hashids.HashID
}

func NewHashIds(salt string, minLength int) *Hasher {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, _ := hashids.NewWithData(hd)
	return &Hasher{h}
}

func (h *Hasher) Encode(id int, db_type int) string {
	str, _ := h.HashId.Encode([]int{id, db_type})
	return str
}

func (h *Hasher) Decode(str string) (int, error) {
	ids, err := h.HashId.DecodeWithError(str)
	if err != nil {
		return 0, err
	}

	return ids[0], nil
}
