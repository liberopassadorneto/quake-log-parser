package parser

import (
	"github.com/liberopassadorneto/quake-log-parser/models"
	"testing"
)

func TestReadChunksFromLogFile(t *testing.T) {
	// Modifique com o caminho do arquivo de teste no seu sistema
	testFilePath := "../testdata/test.log"

	_, err := readChunksFromLogFile(testFilePath)
	if err != nil {
		t.Fatalf("Esperado nenhum erro, mas ocorreu: %v", err)
	}
}

func TestParseChunk(t *testing.T) {
	// Entrada de teste. Modifique conforme necessário.
	input := `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
	          InitGame: 
	          Kill: 1022 3 22: <world> killed Dono da Bola by MOD_TRIGGER_HURT`

	// O número do jogo para este teste
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
		t.Fatalf("Esperado nenhum erro, mas ocorreu: %v", err)
	}

	if game.GameNumber != expectedGame.GameNumber {
		t.Fatalf("Esperado número do jogo %v, mas obtido: %v", expectedGame.GameNumber, game.GameNumber)
	}

	// Verificações adicionais conforme necessário
	if len(game.Players) != len(expectedGame.Players) {
		t.Fatalf("Esperado número de jogadores %v, mas obtido: %v", len(expectedGame.Players), len(game.Players))
	}

	for name, expectedPlayer := range expectedGame.Players {
		player, exists := game.Players[name]
		if !exists {
			t.Fatalf("Jogador esperado %v não encontrado", name)
		}
		if player.Kills != expectedPlayer.Kills {
			t.Fatalf("Esperado %v kills para jogador %v, mas obtido: %v", expectedPlayer.Kills, name, player.Kills)
		}
	}

	if len(game.KillsByMeans) != len(expectedGame.KillsByMeans) {
		t.Fatalf("Esperado %v tipos de mortes, mas obtido: %v", len(expectedGame.KillsByMeans), len(game.KillsByMeans))
	}

	for means, count := range expectedGame.KillsByMeans {
		if game.KillsByMeans[means] != count {
			t.Fatalf("Esperado %v mortes por %v, mas obtido: %v", count, means, game.KillsByMeans[means])
		}
	}

	if game.TotalKills != expectedGame.TotalKills {
		t.Fatalf("Esperado total de %v mortes, mas obtido: %v", expectedGame.TotalKills, game.TotalKills)
	}
}

func TestParseLog(t *testing.T) {
	// Modifique com o caminho do arquivo de teste no seu sistema
	testFilePath := "../testdata/test.log"

	parser := NewParser()
	_, err := parser.ParseLog(testFilePath)
	if err != nil {
		t.Fatalf("Esperado nenhum erro, mas ocorreu: %v", err)
	}
}
