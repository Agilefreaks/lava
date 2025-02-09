package badgegenerator

import "github.com/lavanet/lava/x/pairing/types"

type ProjectConfiguration struct {
	ProjectPublicKey  string                         `json:"project_public_key"`
	ProjectPrivateKey string                         `json:"private_key"`
	EpochsMaxCu       int64                          `json:"epochs_max_cu"`
	UpdatedEpoch      uint64                         `json:"update_epoch,omitempty"`
	PairingList       *types.QueryGetPairingResponse `json:"pairing_list,omitempty"`
}

type UserBadgeItem struct {
	AllowedCu int64
	Epoch     uint64
	Signature string
	PublicKey string
	UserId    string
}

type IpData struct {
	FromIp      int64
	ToIP        int64
	CountryCode string
	Geolocation int
}
