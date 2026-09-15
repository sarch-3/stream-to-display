package domain

type PlayerProperties struct {
	Volume int     `json:"volume"`
	Speed  float64 `json:"speed"`
}

type MPVRequest struct {
	Command []interface{} `json:"command"`
}

type MPVResponse struct {
	Data      interface{} `json:"data"`
	Error     string      `json:"error"`
	RequestID int         `json:"request_id"`
}

type AddVideoRequest struct {
	Source string `json:"source"`
}

// type SeekRequest struct {
// 	Seconds int `json:"seconds"`
// 	Delta   int `json:"delta"`
// }

// type VolumeRequest struct {
// 	Delta int `json:"delta"`
// }

// type SpeedRequest struct {
// 	Delta float64 `json:"delta"`
// }

type ActionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
