package domain

type Branch struct {
	Ref string `json:"ref"`
	SHA string `json:"sha"`
}