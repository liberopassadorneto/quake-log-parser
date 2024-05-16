package parser

import (
	"github.com/liberopassadorneto/quake-log-parser/models"
	"testing"
)

func TestReadChunksFromLogFile(t *testing.T) {
	// Modify with the path to the test file on your system
	testFilePath := "../test/test.log"

	_, err := readChunksFromLogFile(testFilePath)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
}

func TestParseChunk(t *testing.T) {
	// Test input. Modify as needed.
	input := `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
	          InitGame: 
	          Kill: 1022 3 22: <world> killed Dono da Bola by MOD_TRIGGER_HURT`

	// The game number for this test
	gameNumber := 1

	expectedGame := &models.Game{
		GameNumber: gameNumber,
		Players: map[string]*models.Player{
			"Isgalamido":   {Name: "Isgalamido", Kills: -1},
			"Dono da Bola": {Name: "Dono da Bola", Kills: -1},
		},
		KillsByMeans: map[string]int{"MOD_TRIGGER_HURT": 2},
		TotalKills:   2,
	}

	game, err := parseChunk(input, gameNumber)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if game.GameNumber != expectedGame.GameNumber {
		t.Fatalf("Expected game number %v, but got: %v", expectedGame.GameNumber, game.GameNumber)
	}

	// Additional checks as needed
	if len(game.Players) != len(expectedGame.Players) {
		t.Fatalf("Expected number of players %v, but got: %v", len(expectedGame.Players), len(game.Players))
	}

	for name, expectedPlayer := range expectedGame.Players {
		player, exists := game.Players[name]
		if !exists {
			t.Fatalf("Expected player %v not found", name)
		}
		if player.Kills != expectedPlayer.Kills {
			t.Fatalf("Expected %v kills for player %v, but got: %v", expectedPlayer.Kills, name, player.Kills)
		}
	}

	if len(game.KillsByMeans) != len(expectedGame.KillsByMeans) {
		t.Fatalf("Expected %v kill types, but got: %v", len(expectedGame.KillsByMeans), len(game.KillsByMeans))
	}

	for means, count := range expectedGame.KillsByMeans {
		if game.KillsByMeans[means] != count {
			t.Fatalf("Expected %v kills by %v, but got: %v", count, means, game.KillsByMeans[means])
		}
	}

	if game.TotalKills != expectedGame.TotalKills {
		t.Fatalf("Expected total kills %v, but got: %v", expectedGame.TotalKills, game.TotalKills)
	}
}

func TestParseLog(t *testing.T) {
	// Modify with the path to the test file on your system
	testFilePath := "../test/test.log"

	parser := NewParser()
	_, err := parser.ParseLog(testFilePath)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
}
