package postgresql

type PostgresqlRequest struct {
	Avatar  string `json:"avatar"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
