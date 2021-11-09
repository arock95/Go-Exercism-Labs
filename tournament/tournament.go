package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Rank struct {
	name string
	mp   int
	p    int
	w    int
	l    int
	d    int
}

const (
	win  int = 3
	draw int = 1
)

func Tally(r io.Reader, b io.Writer) error {
	matchOutcomes := map[string]*Rank{}

	lines := bufio.NewScanner(r)
	lines.Split(bufio.ScanLines)

	for lines.Scan() {
		// test for blank lines and comment lines, ignore them
		line := lines.Text()
		if len(line) == 0 || line[0] == '#'{
			continue
		}

		match := strings.Split(line, ";")

		if len(match) != 3 {
			return fmt.Errorf("invalid number of parameters - should include home team;away team;outcome")
		}

		homeTeam := match[0]
		awayTeam := match[1]

		// if there's no entry in the dict for a particular team, add it
		// to initialize the Rank struct and add the team's name
		_, ok := matchOutcomes[homeTeam]
		if !ok {
			matchOutcomes[homeTeam] = &Rank{name: homeTeam}
		}
		_, ok = matchOutcomes[awayTeam]
		if !ok {
			matchOutcomes[awayTeam] = &Rank{name: awayTeam}
		}

		matchOutcomes[homeTeam].mp += 1
		matchOutcomes[awayTeam].mp += 1

		switch match[2] {
		case "draw":
			matchOutcomes[homeTeam].d += 1
			matchOutcomes[homeTeam].p += draw
			matchOutcomes[awayTeam].d += 1
			matchOutcomes[awayTeam].p += draw
		case "win":
			matchOutcomes[homeTeam].w += 1
			matchOutcomes[homeTeam].p += win
			matchOutcomes[awayTeam].l += 1
		case "loss":
			matchOutcomes[awayTeam].w += 1
			matchOutcomes[awayTeam].p += win
			matchOutcomes[homeTeam].l += 1
		default:
			return fmt.Errorf("needs to be win/loss/draw")
		}
	}

	// Output - write to io.Writer
	fmt.Fprint(b, "Team                           | MP |  W |  D |  L |  P\n")

	teamBufferLen := 30 // length of Team field in output
	teamsInOrder := orderTeams(matchOutcomes)
	for _, v := range teamsInOrder {
		fmt.Fprint(b, fmt.Sprintf("%-*s |  %d |  %d |  %d |  %d |  %d\n", teamBufferLen, v.name, v.mp, v.w, v.d, v.l, v.p))
	}

	return nil
}

// orderTeams orders the teams by point and then alphabetical order
func orderTeams (teams map[string] *Rank) []Rank {
	var teamSlice []Rank
	for _, team := range teams {
		teamSlice = append(teamSlice, *team)
	}

	sort.SliceStable(teamSlice, func(i,j int) bool {
		return teamSlice[i].name < teamSlice[j].name
	})
	sort.SliceStable(teamSlice, func(i,j int) bool {
		return teamSlice[i].p > teamSlice[j].p
	})
	
	return teamSlice
}
