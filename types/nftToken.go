package types

type SmartContract struct {
	ContractAddress string `json:"contract_address"`
	ContractName    string `json:"contract_name"`
}

type NftToken struct {
	TokenId       string        `json:"token_id"`
	CurrentHolder string        `json:"current_holder"`
	SmartContract SmartContract `json:"smart_contract"`
	Balance       float64       `json:"balance"`
	Metadata      any           `json:"metadata"`
}
