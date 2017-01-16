package util

type RegisterRQ struct {
	UserID        string   `json:"user_id"`
	Role          int      `json:"role"`
	Affiliation   string   `json:"affiliation"`
	Authorization []string `json:"authorization"`
	RegistarId    string   `json:"registar_id"`
	RegistarToken string   `json:"registar_token"`
}

type RegisterRS struct {
	Token     string `json:"token"`
	ErrorCode int    `json:"err_code"`
	ErrorMsg  string `json:"err_msg"`
}

type DeployRQ struct {
	UserID    string   `json:"user_id"`
	UserToken string   `json:"user_token"`
	Path      string   `json:"path"`
	Args      []string `json:"args"`
	Meta      string   `json:"meta"`
}

type DeployRS struct {
	ChaincodeID string `json:"chain_id"`
	ErrorCode   int    `json:"err_code"`
	ErrorMsg    string `json:"err_msg"`
}

type InvokeRQ struct {
	UserID    string   `json:"user_id"`
	UserToken string   `json:"user_token"`
	ChainID   string   `json:"chain_id"`
	Args      []string `json:"args"`
	TxID      string   `json:"txid"`
	Meta      string   `json:"meta"`
}

type InvokeRS struct {
	Msg       string `json:"msg"`
	ErrorCode int    `json:"err_code"`
	ErrorMsg  string `json:"err_msg"`
}

type QueryRQ struct {
	UserID    string   `json:"user_id"`
	UserToken string   `json:"user_token"`
	ChainID   string   `json:"chain_id"`
	Args      []string `json:"args"`
	TxID      string   `json:"txid"`
	Meta      string   `json:"meta"`
}

type QueryRS struct {
	Msg       string `json:"msg"`
	ErrorCode int    `json:"err_code"`
	ErrorMsg  string `json:"err_msg"`
}
