from pydantic import Field
from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    """Settings populated from environment variables."""

    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        extra="ignore",
    )

    drm_device: str | None = Field(None, validation_alias="DRM_DEVICE")
    drm_connector: str | None = Field(None, validation_alias="DRM_CONNECTOR")
    audio_output: str | None = Field(None, validation_alias="AUDIO_OUTPUT")
    audio_device: str | None = Field(None, validation_alias="AUDIO_DEVICE")
    cache: bool = Field(True, validation_alias="CACHE")
    demuxer_max_bytes: str = Field("100M", validation_alias="DEMUXER_MAX_BYTES")
    demuxer_readahead_secs: str = Field("60", validation_alias="DEMUXER_READAHEAD_SECS")
    loop: str = Field("0", validation_alias="LOOP")  # int or 'inf'
    stream_resolution: str = Field("1080", validation_alias="STREAM_RESOLUTION")
    stream_url: str = Field(
        "https://www.w3schools.com/html/mov_bbb.mp4", validation_alias="STREAM_URL"
    )


settings = Settings()
