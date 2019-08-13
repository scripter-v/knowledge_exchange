package main

import (
	"encoding/json"
	"fmt"
)

type device struct {
	name   string
	role   string
	tenant string
}

type Results struct {
	Devices []*struct {
		Name       string
		DeviceRole *struct {
			Slug string
		}
		Tenant *struct {
			Slug string
		}
	}
}

func (r *Results) getDevices() []device {
	devices := make([]device, 0, len(r.Devices))
	for _, d := range r.Devices {
		devices = append(devices, device{
			name:   d.Name,
			role:   d.DeviceRole.Slug,
			tenant: d.Tenant.Slug,
		})
	}

	return devices
}

func main() {
	var results Results
	if err := json.Unmarshal([]byte(``), &results); err != nil {
		panic(err)
	}

	devices := results.getDevices()
	fmt.Printf("%+v", devices)
}
