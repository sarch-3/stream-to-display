import signal
import subprocess

from src.core.config import settings


def main() -> None:
    command = [
        "mpv",
        settings.stream_url,
        "--vo=drm",
        "--drm-device=/dev/dri/card1",
        f"--drm-connector={settings.display_name}",
        "--loop=inf",
    ]

    process = subprocess.Popen(command)

    def stop_process(signum: int, _frame: object) -> None:
        process.terminate()

    signal.signal(signal.SIGTERM, stop_process)
    signal.signal(signal.SIGINT, stop_process)
    raise SystemExit(process.wait())


if __name__ == "__main__":
    main()
