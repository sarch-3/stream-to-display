from pydantic import Field
from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    """Settings populated from environment variables."""

    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        extra="ignore",
    )

    stream_url: str = Field(
        "https://www.w3schools.com/html/mov_bbb.mp4", validation_alias="STREAM_URL"
    )
    drm_device: str | None = Field(None, validation_alias="DRM_DEVICE")
    drm_connector: str | None = Field(None, validation_alias="DRM_CONNECTOR")
    audio_output: str | None = Field(None, validation_alias="AUDIO_OUTPUT")
    audio_device: str | None = Field(None, validation_alias="AUDIO_DEVICE")
    loop: str = Field("0", validation_alias="LOOP")  # int or 'inf'
    log_level: str = Field("INFO", validation_alias="LOG_LEVEL")


settings = Settings()
