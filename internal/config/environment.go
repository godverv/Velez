// Code generated by RedSock CLI. DO NOT EDIT.

// Code generated by RedSock CLI. DO NOT EDIT.

package config

import (
	"time"
)

type EnvironmentConfig struct {
	AvailablePorts      []int
	CPUDefault          float64
	CustomPassToKey     string
	DisableAPISecurity  bool
	ExposeMatreshkaPort bool
	MakoshExposePort    bool
	MakoshImageName     string
	MakoshKey           string
	MakoshPort          int
	MakoshUrls          []string
	MakoshUrls          []string
	MatreshkaPort       int
	MatreshkaUrls       []string
	MemorySwapMb        int
	NodeMode            bool
	PortainerEnabled    bool
	RAMMbDefault        int
	ShutDownOnExit      bool
	WatchTowerEnabled   bool
	WatchTowerInterval  time.Duration
}
