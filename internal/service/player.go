package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/sarch-3/stream-to-display/internal/domain"
)

type PlayerService struct {
	state      domain.PlayerProperties
	socketPath string
}

func NewPlayerService(socketPath string) *PlayerService {
	return &PlayerService{
		state: domain.PlayerProperties{
			Volume: 100,
			Speed:  1,
		},
		socketPath: socketPath,
	}
}

func (s *PlayerService) AddVideo(source string) domain.ActionResponse {
	if source == "" {
		return domain.ActionResponse{
			Success: false,
			Message: "source is required",
		}
	}

	command := []interface{}{"loadfile", source, "append-play"}

	resp := sendCommand(command, s.socketPath)

	if resp.Error == "success" {
		return domain.ActionResponse{
			Success: true,
			Message: "video added",
		}
	} else {
		return domain.ActionResponse{
			Success: true,
			Message: fmt.Sprintf("mpv returned error: %s", resp.Error),
		}
	}
}

// func (s *PlayerService) Playback() domain.ActionResponse {
// 	s.state.IsPlaying = !s.state.IsPlaying
// 	return domain.ActionResponse{
// 		Success: true,
// 		Message: "video playback toggled",
// 	}
// }

// func (s *PlayerService) Seek(deltaSeconds int) domain.ActionResponse {
// 	s.state.PositionSeconds += deltaSeconds
// 	if s.state.PositionSeconds < 0 {
// 		s.state.PositionSeconds = 0
// 	}

// 	return domain.ActionResponse{
// 		Success: true,
// 		Message: fmt.Sprintf("position updated by %d seconds", deltaSeconds),
// 	}
// }

// func (s *PlayerService) SkipVideo() domain.ActionResponse {
// 	if len(s.state.Queue) == 0 {
// 		return domain.ActionResponse{
// 			Success: false,
// 			Message: "queue is empty",
// 		}
// 	}

// 	if len(s.state.Queue) > 1 {
// 		s.state.Queue = append([]string{}, s.state.Queue[1:]...)
// 	}

// 	if len(s.state.Queue) > 0 {
// 		s.state.CurrentSource = s.state.Queue[0]
// 		s.state.PositionSeconds = 0
// 		s.state.IsPlaying = true
// 		return domain.ActionResponse{
// 			Success: true,
// 			Message: "video skipped",
// 		}
// 	}

// 	s.state.CurrentSource = ""
// 	s.state.Queue = nil
// 	s.state.PositionSeconds = 0
// 	s.state.IsPlaying = false
// 	return domain.ActionResponse{
// 		Success: true,
// 		Message: "queue cleared after skip",
// 	}
// }

// func (s *PlayerService) AdjustVolume(delta int) domain.ActionResponse {
// 	s.state.Volume += delta
// 	if s.state.Volume < 0 {
// 		s.state.Volume = 0
// 	}
// 	if s.state.Volume > 200 {
// 		s.state.Volume = 200
// 	}

// 	return domain.ActionResponse{
// 		Success: true,
// 		Message: fmt.Sprintf("volume changed by %d", delta),
// 	}
// }

// func (s *PlayerService) AdjustSpeed(delta float64) domain.ActionResponse {
// 	s.state.Speed += delta
// 	if s.state.Speed < 0.25 {
// 		s.state.Speed = 0.25
// 	}
// 	if s.state.Speed > 3.0 {
// 		s.state.Speed = 3.0
// 	}

// 	return domain.ActionResponse{
// 		Success: true,
// 		Message: fmt.Sprintf("speed changed by %.2f", delta),
// 	}
// }

// func (s *PlayerService) ClearQueue() domain.ActionResponse {
// 	s.state.Queue = nil
// 	s.state.CurrentSource = ""
// 	s.state.PositionSeconds = 0
// 	s.state.IsPlaying = false

// 	return domain.ActionResponse{
// 		Success: true,
// 		Message: "queue cleared",
// 	}
// }

// func contains(items []string, target string) bool {
// 	for _, item := range items {
// 		if item == target {
// 			return true
// 		}
// 	}
// 	return false
// }

func sendCommand(command []interface{}, socket string) domain.MPVResponse {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		log.Fatalf("error at connecting to mpv socket: %v", err)
	}
	defer conn.Close()

	req := domain.MPVRequest{
		Command: command,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("marshaling error: %v", err)
	}

	jsonData = append(jsonData, '\n')

	_, err = conn.Write(jsonData)
	if err != nil {
		log.Fatalf("writing error: %v", err)
	}

	reader := bufio.NewReader(conn)
	rawResponse, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("reading error: %v", err)
	}

	var resp domain.MPVResponse
	if err := json.Unmarshal([]byte(rawResponse), &resp); err != nil {
		log.Fatalf("Ошибка парсинга ответа: %v", err)
	}

	return resp
}
