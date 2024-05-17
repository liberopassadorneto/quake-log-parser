package parser

import (
	"bufio"
	"fmt"
	"github.com/liberopassadorneto/quake/logger"
	"github.com/liberopassadorneto/quake/models"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseLog(filePath string) ([]*models.Game, error) {
	chunks, err := readChunksFromLogFile(filePath)
	if err != nil {
		logger.Log.Fatalf("Error reading log file: %v", err)
		return nil, err
	}

	// Channel for collectGames
	gamesChannel := make(chan *models.Game, len(chunks))

	// Channel to send chunks to workers
	taskChannel := make(chan chunkTask, len(chunks))

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(taskChannel, gamesChannel, &wg)
	}

	for i, chunk := range chunks {
		taskChannel <- chunkTask{chunk: chunk, gameNumber: i + 1}
	}
	close(taskChannel)

	go func() {
		wg.Wait()
		close(gamesChannel)
	}()

	return collectGames(gamesChannel), nil
}

type chunkTask struct {
	chunk      string
	gameNumber int
}

func worker(tasks <-chan chunkTask, games chan<- *models.Game, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		game, err := parseChunk(task.chunk, task.gameNumber)
		if err != nil {
			logger.Log.Printf("Error parsing chunk: %v", err)
			continue
		}
		games <- game
	}
}

func readChunksFromLogFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var chunks []string
	var currentChunk strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "InitGame:") && currentChunk.Len() > 0 {
			chunks = append(chunks, currentChunk.String())
			currentChunk.Reset()
		}
		currentChunk.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	if currentChunk.Len() > 0 {
		chunks = append(chunks, currentChunk.String())
	}

	return chunks, nil
}

func parseChunk(chunk string, gameNumber int) (*models.Game, error) {
	lines := strings.Split(chunk, "\n")
	game := &models.Game{
		GameNumber:   gameNumber,
		Players:      make(map[string]*models.Player),
		KillsByMeans: make(map[string]int),
	}

	killRegexp := regexp.MustCompile(`Kill: \d+ \d+ \d+: (.+) killed (.+) by (.+)`)

	for _, line := range lines {
		if matches := killRegexp.FindStringSubmatch(line); len(matches) == 4 {
			killerName, victimName, meansOfDeath := matches[1], matches[2], matches[3]
			updateGameStats(game, killerName, victimName, meansOfDeath)
		}
	}

	return game, nil
}

func updateGameStats(game *models.Game, killerName, victimName, meansOfDeath string) {
	if _, exists := game.Players[victimName]; !exists {
		game.Players[victimName] = &models.Player{Name: victimName}
	}

	if killerName == "<world>" {
		game.Players[victimName].Kills--
	} else if killerName != victimName {
		if _, exists := game.Players[killerName]; !exists {
			game.Players[killerName] = &models.Player{Name: killerName}
		}
		game.Players[killerName].Kills++
	}
	game.KillsByMeans[meansOfDeath]++
	game.TotalKills++
}

func collectGames(gamesChannel <-chan *models.Game) []*models.Game {
	var games []*models.Game
	for game := range gamesChannel {
		games = append(games, game)
	}
	sort.Slice(games, func(i, j int) bool {
		return games[i].GameNumber < games[j].GameNumber
	})
	return games
}
