# Fake Windows Service

## Usage

* Insall

```batch
set SERVICE_NAME="csagent"
set SERVICE_DISPLAY_NAME="CrowdStrike Falcon"

dummy_service.exe /install %SERVICE_NAME% %SERVICE_DISPLAY_NAME%
```

* Uninstall

```batch
set SERVICE_NAME="csagent"
set SERVICE_DISPLAY_NAME="CrowdStrike Falcon"

dummy_service.exe /uninstall %SERVICE_NAME% %SERVICE_DISPLAY_NAME%
```
