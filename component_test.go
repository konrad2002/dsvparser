package main

import (
	"bytes"
	"fmt"
	"github.com/konrad2002/dsvparser/model"
	"github.com/konrad2002/dsvparser/model/elements"
	"github.com/konrad2002/dsvparser/model/types"
	"github.com/konrad2002/dsvparser/parser"
	"github.com/stretchr/testify/assert"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func Test_StandardExample_Definitionsliste(t *testing.T) {
	dat, err := os.ReadFile("assets/definition.dsv7")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	res, _ := r.Read()
	def := res.(*model.Wettkampfdefinitionsliste)
	assert.Equal(t, 7, def.Format.Version)
	assert.Equal(t, "JAHRGANG 1990", def.Wertungen[2].Wertungsname)
	assert.Equal(t, strings.ToLower(types.EINZEL), strings.ToLower(def.Meldegelder[1].MeldegeldTyp))
}

func Test_StandardExample_Ergebnisliste(t *testing.T) {
	dat, err := os.ReadFile("assets/ergebnis.dsv7")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	res, err := r.Read()
	if err != nil {
		fmt.Printf(err.Error())
	}
	erg := res.(*model.Wettkampfergebnisliste)
	assert.Equal(t, 7, erg.Format.Version)
	assert.Equal(t, "Duisburg", erg.Ausrichter.Ort)
	assert.Equal(t, "123440", erg.PNErgebnisse[4].DsvId)
}

func Test_BeispielMarienberg_Ergebnisliste(t *testing.T) {
	dat, err := os.ReadFile("assets/2023-05-14-Marienbe-Pr.DSV6")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	res, err := r.Read()
	if err != nil {
		fmt.Printf(err.Error())
	}
	erg := res.(*model.Wettkampfergebnisliste)
	assert.Equal(t, 6, erg.Format.Version)
	assert.Equal(t, "Olbernhau", erg.Ausrichter.Ort)
	assert.Equal(t, "429663", erg.PNErgebnisse[4].DsvId)
}

func Test_BeispielIESC_Ergebnisliste(t *testing.T) {
	dat, err := os.ReadFile("assets/Ergebnisdatei.dsv6")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	res, err := r.Read()
	if err != nil {
		fmt.Printf(err.Error())
	}
	erg := res.(*model.Wettkampfergebnisliste)
	assert.Equal(t, 6, erg.Format.Version)
	assert.Equal(t, "", erg.Ausrichter.Ort)
}

func getPointsForRank(r int) int {
	switch r {
	case 1:
		return 9
	case 2:
		return 7
	case 3:
		return 6
	case 4:
		return 5
	case 5:
		return 4
	case 6:
		return 3
	case 7:
		return 2
	case 8:
		return 1
	default:
		return 0
	}
}

func getGroupTeamByTeamName(n string) string {
	re := regexp.MustCompile(`\((.*?)\)`)

	matches := re.FindStringSubmatch(n)

	if len(matches) > 1 {
		return matches[1]
	}
	return n
}

func getGroupTeamByRelayName(n string) string {
	if strings.Contains(n, "Westsachsen") {
		return "TW"
	}
	if strings.Contains(n, "Elbe-Elster") {
		return "TEE"
	}

	return n
}

func Test_PointsEsbjerg(t *testing.T) {
	println("----====[ Esbjerg 2024 ]====----")
	calcEsbjergPoints("assets/2024-05-12-Esbjerg-Pk.DSV7", [2]int{0, 2011}, [2]int{0, 2010})
	println()
	println("----====[ Esbjerg 2025 ]====----")
	calcEsbjergPoints("assets/2025-06-01-Esbjerg-Pr.DSV7", [2]int{0, 2012}, [2]int{0, 2011})
}

func calcEsbjergPoints(file string, preAgesGirls [2]int, preAgesBoys [2]int) {
	dat, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	res, err := r.Read()
	if err != nil {
		fmt.Printf("failed to read: %s\n", err.Error())
	}

	pointsPerTeam := map[string]int{}
	teams := map[string]elements.Verein{}

	erg := res.(*model.Wettkampfergebnisliste)

	for _, verein := range erg.Vereine {
		teams[verein.Vereinsbezeichnung] = verein
		pointsPerTeam[verein.Vereinsbezeichnung] = 0
	}

	for _, wettkampf := range erg.Wettkaempfe {
		//fmt.Printf("%d: %d %c\n", i, wettkampf.Wettkampfnummer, wettkampf.Technik)

		// check if final

		//if !(wettkampf.Wettkampfart == 'E' || wettkampf.Wettkampfart == 'F') && wettkampf.Einzelstrecke != 50 {
		//fmt.Printf("skip because not final\n")
		//	continue
		//}

		switch wettkampf.Wettkampfnummer {
		case 91, 92, 93, 94, 291, 292, 293, 294:
			continue
		}

		// get age groups

		// get places for athlete
		for _, ergebnis := range erg.PNErgebnisse {
			if ergebnis.Wettkampfnummer != wettkampf.Wettkampfnummer {
				continue
			}

			if wettkampf.Wettkampfart == 'V' {
				if wettkampf.Einzelstrecke == 50 {
					if wettkampf.Geschlecht == 'W' {
						if !(ergebnis.Jahrgang > preAgesGirls[0] && ergebnis.Jahrgang <= preAgesGirls[1]) {
							continue
						}
					} else if wettkampf.Geschlecht == 'M' {
						if !(ergebnis.Jahrgang > preAgesBoys[0] && ergebnis.Jahrgang <= preAgesBoys[1]) {
							//fmt.Printf("skip in %d: %s -> %d\n", ergebnis.Wettkampfnummer, ergebnis.Name, ergebnis.Jahrgang)
							continue
						}
					}
				} else if wettkampf.Einzelstrecke == 100 {
					if wettkampf.Geschlecht == 'W' {
						if ergebnis.Jahrgang > preAgesGirls[0] && ergebnis.Jahrgang <= preAgesGirls[1] {
							continue
						}
					} else if wettkampf.Geschlecht == 'M' {
						if ergebnis.Jahrgang > preAgesBoys[0] && ergebnis.Jahrgang <= preAgesBoys[1] {
							//fmt.Printf("skip in %d: %s -> %d\n", ergebnis.Wettkampfnummer, ergebnis.Name, ergebnis.Jahrgang)
							continue
						}
					}
				} else {
					continue
				}
			}

			p := getPointsForRank(ergebnis.Platz)
			pointsPerTeam[ergebnis.Verein] += p
		}

		for _, ergebnis := range erg.STErgebnisse {
			if ergebnis.Wettkampfnummer.Int() != wettkampf.Wettkampfnummer {
				continue
			}
			p := getPointsForRank(ergebnis.Platz.Int()) * 2
			pointsPerTeam[getGroupTeamByRelayName(ergebnis.Verein)] += p
		}

		// calc points
	}

	pointsByGroupedTeams := map[string]int{}

	for team, _ := range pointsPerTeam {
		groupTeam := getGroupTeamByTeamName(team)
		pointsByGroupedTeams[groupTeam] = 0
	}

	for team, points := range pointsPerTeam {
		groupTeam := getGroupTeamByTeamName(team)
		pointsByGroupedTeams[groupTeam] += points
	}

	PrintSortedMap(pointsByGroupedTeams)

	//for v, p := range pointsByGroupedTeams {
	//	fmt.Printf("%s\t%d\n", v, p)
	//}

}

// PrintSortedMap prints a map[string]int sorted by value descending,
// with fixed-width columns and total sum at the bottom.
func PrintSortedMap(m map[string]int) {
	const colWidth = 25

	// Convert map to slice of key-value pairs
	type kv struct {
		Key   string
		Value int
	}

	var sorted []kv
	total := 0
	for k, v := range m {
		sorted = append(sorted, kv{k, v})
		total += v
	}

	// Sort by value descending, then key alphabetically
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value == sorted[j].Value {
			return sorted[i].Key < sorted[j].Key
		}
		return sorted[i].Value > sorted[j].Value
	})

	// Header
	fmt.Printf("%-*s | %s\n", colWidth, "Team", "Points")
	fmt.Printf("%s-|-%s\n", repeat("-", colWidth), "------")

	// Rows
	for _, kv := range sorted {
		fmt.Printf("%-*s | %d\n", colWidth, kv.Key, kv.Value)
	}

	// Total
	fmt.Printf("%s-|-%s\n", repeat("-", colWidth), "------")
	fmt.Printf("%-*s | %d\n", colWidth, "Total", total)
}

// Helper to repeat characters
func repeat(s string, count int) string {
	out := ""
	for i := 0; i < count; i++ {
		out += s
	}
	return out
}
