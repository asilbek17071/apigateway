package models

type PySys struct {
	Name string `json:"name"`
}

type PySysResp struct {
	PysysId   string `json:"pysys_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PySysUpdate struct {
	Name string `json:"name"`
}

type PySyssList struct {
	PySyss []PySysResp `json:"payment_systems"`
	Count  int64       `json:"count"`
}
