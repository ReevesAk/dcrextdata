package pow

type PowDataSource struct {
	Source string
}

type PowData struct {
	Time              int64
	NetworkHashrate   float64
	PoolHashrate      float64
	Workers           int64
	CoinPrice         float64
	BtcPrice          float64
	Source            string
}

type PowDataDto struct {
	Time              string  `json:"time"`
	NetworkHashrateTh float64 `json:"network_hashrate_th"`
	PoolHashrateTh    float64 `json:"pool_hashrate_th"`
	Workers           int64   `json:"workers"`
	CoinPrice         float64 `json:"coin_price"`
	BtcPrice          float64 `json:"btc_price"`
	Source            string  `json:"source"`
}

type luxorPowData struct {
	Time              string  `json:"time"`
	NetworkHashrate   float64 `json:"network_hashrate"`
	PoolHashrate      float64 `json:"pool_hashrate"`
	Workers           int64   `json:"workers"`
	CoinPrice         string  `json:"coin_price"`
	BtcPrice          string  `json:"btc_price"`
}

type luxorAPIResponse struct {
	GlobalStats []luxorPowData `json:"globalStats"`
}

type f2poolPowData map[string]float64

type f2poolAPIResponse struct {
	Hashrate f2poolPowData `json:"hashrate_history"`
}

type coinmineAPIResponse struct {
	NetworkHashrate float64 `json:"network_hashrate"`
	PoolHashrate    float64 `json:"hashrate"`
	Workers         int64   `json:"workers"`
}

type btcData struct {
	NetworkHashrate     string              `json:"network_hashrate"`
	NetworkHashrateUnit string              `json:"network_hashrate_unit"`
	PoolHashrate        string              `json:"pool_hashrate"`
	PoolHashrateUnit    string              `json:"pool_hashrate_unit"`
	Rates               btcExchangeRateData `json:"exchange_rate"`
}

type btcExchangeRateData struct {
	CoinPrice float64 `json:"DCR2USD"`
}

type btcAPIResponse struct {
	BtcData btcData `json:"data"`
}

type uupoolData struct {
	PoolHashrate  float64 `json:"hr1"`
	OnlineWorkers int64   `json:"onlineWorkers"`
}

type uunetworkData struct {
	NetworkHashrate   float64 `json:"networkhashps"`
}

type uupoolAPIResponse struct {
	Pool    uupoolData    `json: "pool"`
	Network uunetworkData `json: "network"`
}
