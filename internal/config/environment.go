// Code generated by RedSock CLI. DO NOT EDIT.

package config

import (
	"time"
)

type EnvironmentConfig struct {
	AvailablePorts     []int
	CPUDefault         float64
	CustomPassToKey    string
	DisableAPISecurity bool
	MakoshExposePort   bool
	MakoshImage        string
	MakoshKey          string
	MakoshPort         int
	MakoshURL          string
	MatreshkaImage     string
	MatreshkaPort      int
	MemorySwapMb       int
	NodeMode           bool
	PortainerEnabled   bool
	RAMMbDefault       int
	ShutDownOnExit     bool
	WatchTowerEnabled  bool
	WatchTowerInterval time.Duration
}
