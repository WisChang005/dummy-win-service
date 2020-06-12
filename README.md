# Fake Windows Service

## Usage

* Run Service

```batch
dummy_service.exe
```

* Install

```batch
set SERVICE_NAME="dummyclient"
set SERVICE_DISPLAY_NAME="My Dummy Service"

dummy_service.exe /install %SERVICE_NAME% %SERVICE_DISPLAY_NAME%
```

* Uninstall

```batch
set SERVICE_NAME="dummyclient"
set SERVICE_DISPLAY_NAME="My Dummy Service"

dummy_service.exe /uninstall %SERVICE_NAME% %SERVICE_DISPLAY_NAME%
```
