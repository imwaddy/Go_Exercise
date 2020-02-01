// wrong command

go build -ldflags="-X 'main.mainpackagevariable=main variable' -X 'cronjobs.otherpackagevariable=Other package variable'" main.go

// right command

go build -ldflags="-X 'main.mainpackagevariable=main variable' -X '<WORKSPACE_DIR>/reportCron/cronjobs.otherpackagevariable=Other package variable' " main.go


In my case I have created directory /src/Go_Exercise. If you have other named directory then you have to change import of a file.

eg. If your workspce is "RobertsProject" under /src then import in main like this "RobertsProject/reportCron/cronjobs"

See my blog https://medium.com/@mayurwadekar2/golangs-go-build-command-f471a5e8535d
