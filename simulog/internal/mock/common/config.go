package common

type Config struct {
	T          uint    `json:"t"`
	Normal     float32 `json:"normal"`
	Bruteforce float32 `json:"bruteforce"`
	DDoS       float32 `json:"ddos"`
}
