package trueskill

// Team is a composition of players that play together. The skill of a team
// (µ and σ) is determined by the skills of the players that form the team.
type Team struct {
	players []Player
}

// NewEmptyTeam creates a team without any players. To add players use
// AddPlayers.
func NewEmptyTeam() Team {
	return Team{
		players: make([]Player, 0),
	}
}

// NewTeam creates a team from a slice of players.
func NewTeam(p []Player) (t Team) {
	t.players = p
	return
}

// AddPlayers adds players to the team.
func (t *Team) AddPlayers(p []Player) {
	t.players = append(t.players, p...)
	return
}

// GetPlayers returns the players the team is composed of.
func (t *Team) GetPlayers() (p []Player) {
	return t.players
}
