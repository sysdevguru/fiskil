package dcl

import (
	"reflect"
	"testing"

	"github.com/sysdevguru/fiskil/models"
)

func TestGetSeverities(t *testing.T) {
	logsData := []struct {
		name   string
		logs   []models.Log
		wanted map[string]*models.Severity
	}{
		{
			name:   "empty logs",
			logs:   []models.Log{},
			wanted: map[string]*models.Severity{},
		},
		{
			name: "one log",
			logs: []models.Log{
				{
					ServiceName: "service_1",
					Severity:    "debug",
					Payload:     "log content",
				},
			},
			wanted: map[string]*models.Severity{
				"service_1-debug": {
					ServiceName: "service_1",
					Severity:    "debug",
					Count:       1,
				},
			},
		},
		{
			name: "multiple logs",
			logs: []models.Log{
				{
					ServiceName: "service_1",
					Severity:    "debug",
					Payload:     "log content",
				},
				{
					ServiceName: "service_2",
					Severity:    "debug",
					Payload:     "log content",
				},
				{
					ServiceName: "service_1",
					Severity:    "info",
					Payload:     "log content",
				},
				{
					ServiceName: "service_1",
					Severity:    "warn",
					Payload:     "log content",
				},
				{
					ServiceName: "service_2",
					Severity:    "warn",
					Payload:     "log content",
				},
				{
					ServiceName: "service_3",
					Severity:    "error",
					Payload:     "log content",
				},
				{
					ServiceName: "service_4",
					Severity:    "info",
					Payload:     "log content",
				},
				{
					ServiceName: "service_3",
					Severity:    "debug",
					Payload:     "log content",
				},
				{
					ServiceName: "service_2",
					Severity:    "fatal",
					Payload:     "log content",
				},
				{
					ServiceName: "service_1",
					Severity:    "info",
					Payload:     "log content",
				},
				{
					ServiceName: "service_3",
					Severity:    "error",
					Payload:     "log content",
				},
			},
			wanted: map[string]*models.Severity{
				"service_1-debug": {
					ServiceName: "service_1",
					Severity:    "debug",
					Count:       1,
				},
				"service_1-info": {
					ServiceName: "service_1",
					Severity:    "info",
					Count:       2,
				},
				"service_1-warn": {
					ServiceName: "service_1",
					Severity:    "warn",
					Count:       1,
				},
				"service_2-debug": {
					ServiceName: "service_2",
					Severity:    "debug",
					Count:       1,
				},
				"service_3-error": {
					ServiceName: "service_3",
					Severity:    "error",
					Count:       2,
				},
				"service_2-warn": {
					ServiceName: "service_2",
					Severity:    "warn",
					Count:       1,
				},
				"service_2-fatal": {
					ServiceName: "service_2",
					Severity:    "fatal",
					Count:       1,
				},
				"service_3-debug": {
					ServiceName: "service_3",
					Severity:    "debug",
					Count:       1,
				},
				"service_4-info": {
					ServiceName: "service_4",
					Severity:    "info",
					Count:       1,
				},
			},
		},
	}

	for _, logData := range logsData {
		t.Run(logData.name, func(t *testing.T) {
			result, err := GetSeverities(&logData.logs)
			if err != nil {
				t.Errorf("got %v, want nil", err)
			}

			if !reflect.DeepEqual(*result, logData.wanted) {
				t.Errorf("got %v, want %v", *result, logData.wanted)
			}
		})
	}
}
