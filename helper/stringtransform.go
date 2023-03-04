package helper

import (
	"strings"
)

func DashString(name string) string {
	TobeLowered := strings.ToLower(name)
	TobeDash := strings.Replace(TobeLowered, " ", "-", -1)
	return TobeDash

}

func LowerAndDash(name string) []string {
	TobeLowered := strings.ToLower(name)
	slugToLower := strings.ToLower(TobeLowered)
	dash := strings.Replace(slugToLower, " ", "-", -1)
	TobeSplitted := strings.Split(dash, ",")
	return TobeSplitted
}
