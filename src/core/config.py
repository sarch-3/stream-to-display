"""Application configuration loaded from environment variables."""

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
    display_name: str = Field("HDMI-A-1", validation_alias="DISPLAY_NAME")
    log_level: str = Field("INFO", validation_alias="LOG_LEVEL")


settings = Settings()
