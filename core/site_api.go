package core

import (
	"errors"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/site"
)

var _ site.API = (*Site)(nil)

const (
	GridTariff    = "grid"
	FeedinTariff  = "feedin"
	PlannerTariff = "planner"
)

// GetPrioritySoc returns the PrioritySoc
func (site *Site) GetPrioritySoc() float64 {
	site.Lock()
	defer site.Unlock()
	return site.PrioritySoc
}

// SetPrioritySoc sets the PrioritySoc
func (site *Site) SetPrioritySoc(soc float64) error {
	site.Lock()
	defer site.Unlock()

	if len(site.batteryMeters) == 0 {
		return errors.New("battery not configured")
	}

	site.PrioritySoc = soc
	site.publish("prioritySoc", site.PrioritySoc)

	return nil
}

// GetBufferMin returns the BufferMin
func (site *Site) GetBufferMin() float64 {
	site.Lock()
	defer site.Unlock()
	return site.BufferMin
}

// SetBufferMin sets the BufferMin
func (site *Site) SetBufferMin(soc float64) error {
	site.Lock()
	defer site.Unlock()

	if len(site.batteryMeters) == 0 {
		return errors.New("battery not configured")
	}

	site.BufferMin = soc
	site.publish("bufferMin", site.BufferMin)

	return nil
}

// GetBufferMax returns the BufferMax
func (site *Site) GetBufferMax() float64 {
	site.Lock()
	defer site.Unlock()
	return site.BufferMax
}

// SetBufferMax sets the BufferMax
func (site *Site) SetBufferMax(soc float64) error {
	site.Lock()
	defer site.Unlock()

	if len(site.batteryMeters) == 0 {
		return errors.New("battery not configured")
	}

	site.BufferMax = soc
	site.publish("bufferMax", site.BufferMax)

	return nil
}

// GetResidualPower returns the ResidualPower
func (site *Site) GetResidualPower() float64 {
	site.Lock()
	defer site.Unlock()
	return site.ResidualPower
}

// SetResidualPower sets the ResidualPower
func (site *Site) SetResidualPower(power float64) error {
	site.Lock()
	defer site.Unlock()

	site.ResidualPower = power
	site.publish("residualPower", site.ResidualPower)

	return nil
}

// GetSmartCostLimit returns the SmartCostLimit
func (site *Site) GetSmartCostLimit() float64 {
	site.Lock()
	defer site.Unlock()
	return site.SmartCostLimit
}

// SetSmartCostLimit sets the SmartCostLimit
func (site *Site) SetSmartCostLimit(val float64) error {
	site.Lock()
	defer site.Unlock()

	site.SmartCostLimit = val
	site.publish("smartCostLimit", site.SmartCostLimit)

	return nil
}

// GetVehicles is the list of vehicles
func (site *Site) GetVehicles() []api.Vehicle {
	site.Lock()
	defer site.Unlock()
	return site.coordinator.GetVehicles()
}

// GetTariff returns the respective tariff if configured or nil
func (site *Site) GetTariff(tariff string) api.Tariff {
	site.Lock()
	defer site.Unlock()

	var t api.Tariff
	switch tariff {
	case GridTariff:
		t = site.tariffs.Grid
	case FeedinTariff:
		t = site.tariffs.FeedIn
	case PlannerTariff:
		if t = site.tariffs.Planner; t == nil {
			t = site.tariffs.Grid
		}
	}

	return t
}
