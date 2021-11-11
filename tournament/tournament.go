package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Rank struct {
	TeamName string
	Matches  int
	Points   int
	Wins     int
	Losses   int
	Draws    int
}

const (
	winPoints  int = 3
	drawPoints int = 1
)

func Tally(r io.Reader, b io.Writer) error {
	matchOutcomes := map[string]Rank{}

	lines := bufio.NewScanner(r)
	lines.Split(bufio.ScanLines)

	for lines.Scan() {
		// test for blank lines and comment lines, ignore them
		line := lines.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		match := strings.Split(line, ";")

		if len(match) != 3 {
			return fmt.Errorf("invalid match: %s", line)
		}

		homeTeam, awayTeam := matchOutcomes[match[0]], matchOutcomes[match[1]]

		homeTeam.TeamName, awayTeam.TeamName = match[0], match[1]
		homeTeam.Matches++
		awayTeam.Matches++

		switch match[2] {
		case "draw":
			homeTeam.Draws++
			homeTeam.Points += drawPoints
			awayTeam.Draws++
			awayTeam.Points += drawPoints
		case "win":
			homeTeam.Wins += 1
			homeTeam.Points += winPoints
			awayTeam.Losses += 1
		case "loss":
			awayTeam.Wins++
			awayTeam.Points += winPoints
			homeTeam.Losses++
		default:
			return fmt.Errorf("needs to be win/loss/draw: %s", match[2])
		}
		matchOutcomes[match[0]], matchOutcomes[match[1]] = homeTeam, awayTeam
	}

	// Output - write to io.Writer
	fmt.Fprint(b, "Team                           | MP |  W |  D |  L |  P\n")

	teamBufferLen := 30 // length of Team field in output
	teamsInOrder := orderTeams(matchOutcomes)
	for _, v := range teamsInOrder {
		fmt.Fprint(b, fmt.Sprintf("%-*s |  %d |  %d |  %d |  %d |  %d\n", teamBufferLen, v.TeamName, v.Matches, v.Wins, v.Draws, v.Losses, v.Points))
	}

	return nil
}

// orderTeams orders the teams by point and then alphabetical order
func orderTeams(teams map[string]Rank) []Rank {
	var teamSlice []Rank
	for _, team := range teams {
		teamSlice = append(teamSlice, team)
	}

	sort.Slice(teamSlice, func(i,j int) bool {
		if teamSlice[i].Points != teamSlice[j].Points {
			return teamSlice[i].Points > teamSlice[j].Points
		}
		return teamSlice[i].TeamName < teamSlice[j].TeamName
	})

	return teamSlice
}
