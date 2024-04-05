package config

type SearchSettings struct {
	ScoreThreshold float32 `yaml:"scorethreshold"`
	TopKResults    int     `yaml:"topkresults"`
}
