package loadlimits

type LoadLimits struct {
	AsymmetricLoadLimits []AsymmetricLoadLimit
	LinearLoadLimits     []LinearLoadLimit
	CumulativeLoadLimits []CumulativeLoadLimit
	CombinedLoadLimits   []CombinedLoadLimit
}
