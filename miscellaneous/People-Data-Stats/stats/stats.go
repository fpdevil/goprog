package stats

import (
	"sort"

	"github.com/fpdevil/goprog/miscellaneous/People-Data-Stats/types"

	log "github.com/sirupsen/logrus"
)

var (
	maleStats, femaleStats types.Stats
)

// Show fu	nction displays the statistics from the map
func Show(pmap types.PMap) {
	log.Printf("map: %v\n", len(pmap))
	if pmap == nil {
		log.Info("empty database, no statistics to compile...")
	}
	// log.Printf("map: %v\n", getStats(pmap['M']))
	// log.WithFields(log.Fields{"PMap": "values"}).Info(pmap)
	maleStats = getStats(pmap['M'])
	femaleStats = getStats(pmap['F'])

	maleStats.Display("Male")
	femaleStats.Display("Female")
}

func getStats(p types.Persons) (s types.Stats) {
	s.Total = len(p)
	s.MinSalary, s.MaxSalary, s.AvgSalary = getSalaries(p)
	return
}

func getSalaries(p types.Persons) (min, max, avg types.Currency) {

	sort.Sort(p)
	length := len(p)
	var salary types.Currency

	salaries := make([]types.Currency, 0, length)
	for _, v := range p {
		salary += v.Salary
		salaries = append(salaries, v.Salary)
	}
	min = salaries[0]
	max = salaries[length-1]
	avg = salary / types.Currency(length)

	return
}
