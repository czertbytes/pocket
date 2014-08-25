package types

import "time"

type ClientToken struct {
	Value    string    `json:"value" datastore:"value"`
	ExpireAt time.Time `json:"expire_at" datastore:"expire_at"`
}
