import signal
import subprocess

from src.core.config import settings


def build_command() -> list[str]:
    """Build the command to run mpv with the specified settings."""

    command = ["mpv"]

    command.append("--vo=drm")

    if settings.drm_device:
        command.append(f"--drm-device={settings.drm_device}")

    if settings.drm_connector:
        command.append(f"--drm-connector={settings.drm_connector}")

    if settings.audio_output:
        command.append(f"--ao={settings.audio_output}")

    if settings.audio_device:
        command.append(f"--audio-device={settings.audio_device}")

    if settings.cache:
        command.append("--cache=yes")
    else:
        command.append("--cache=no")

    command.append(f"--loop={settings.loop}")

    command.append(
        f"--ytdl-format=bestvideo[height<=?{settings.stream_resolution}]+bestaudio/best"
    )

    command.append("--demuxer-max-bytes=100M")

    command.append("--demuxer-readahead-secs=60")

    command.append(settings.stream_url)

    return command


def main() -> None:
    command = build_command()

    process = subprocess.Popen(command)

    def stop_process(signum: int, _frame: object) -> None:
        process.terminate()

    signal.signal(signal.SIGTERM, stop_process)
    signal.signal(signal.SIGINT, stop_process)
    raise SystemExit(process.wait())


if __name__ == "__main__":
    main()
