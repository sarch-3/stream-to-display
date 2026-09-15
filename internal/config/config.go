package config

import (
	"os"
	"strconv"
)

type Config struct {
	Host       string // Cant be changed at runtime
	Port       int    // Cant be changed at runtime
	SocketPath string // Cant be changed at runtime
	// SeekIntervalSeconds  int    // Cant be changed at runtime
	VideoOutput          string // Cant be changed at runtime
	DRMDevice            string // Cant be changed at runtime
	DRMConnector         string // Cant be changed at runtime
	AudioOutput          string // Cant be changed at runtime
	AudioDevice          string
	HWDEC                string
	Cache                string
	DemuxerMaxBytes      int
	DemuxerReadaheadSecs int
}

func Load() Config {
	return Config{
		Host:                 getEnv("APP_HOST", "0.0.0.0"),
		Port:                 getEnvInt("APP_PORT", 8080),
		SocketPath:           getEnv("SOCKET_PATH", "/tmp/mpvsocket"),
		VideoOutput:          getEnv("VIDEO_OUTPUT", "drm"),
		DRMDevice:            getEnv("DRM_DEVICE", "/dev/dri/card0"),
		DRMConnector:         getEnv("DRM_CONNECTOR", "HDMI-A-1"),
		AudioOutput:          getEnv("AUDIO_OUTPUT", "alsa"),
		AudioDevice:          getEnv("AUDIO_DEVICE", "default"),
		HWDEC:                getEnv("HWDEC", "vaapi"),
		Cache:                getEnv("CACHE", "yes"),
		DemuxerMaxBytes:      getEnvInt("DEMUXER_MAX_BYTES", 100*1024),
		DemuxerReadaheadSecs: getEnvInt("DEMUXER_READAHEAD_SECS", 60),
		// SeekIntervalSeconds:  getEnvInt("SEEK_INTERVAL_SECONDS", 10),
	}
}

func (c Config) Addr() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}
