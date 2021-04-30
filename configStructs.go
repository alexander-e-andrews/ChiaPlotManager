package main



//Explicit Declaration
//Decide what each and everyone will do, and not auto adjust to find the optimal configuration
// Yaml2Go
type ExplicitConfig struct {
	MaxRunning int     `yaml:"MaxRunning"`
	Tasks      []Tasks `yaml:"Tasks"`
	Loop       bool    `yaml:"Loop"`
}


// Tasks
type Tasks struct {
	PlotSize   int    `yaml:"PlotSize"`
	TempDir2   string `yaml:"TempDir2"`
	NumThreads int    `yaml:"NumThreads"`
	NumBuckets int    `yaml:"NumBuckets"`
	DisableBitfield   bool   `yaml:"DisableBitfield"`
	Ram        int    `yaml:"Ram"`
	TempDir1   string `yaml:"TempDir1"`
	FinalDir   string `yaml:"FinalDir"`
	StartAfter int    `yaml:"StartAfter"`
}












//General Declaration
//Set a range of parameters, and the process will try to discover which is best and go from there