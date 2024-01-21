# PTaaS FTP server

![](https://img.shields.io/badge/language-golang_v1.20-blue)
![GitHub release (with filter)](https://img.shields.io/github/v/release/ptaas-tool/ftp-server)

This is file manager component of our system. In this app we upload our attack scripts,
we execute attack scripts, and we download the results. This system provides an interface
in order to manage scripts and logs files from core app.

## Image

FTP server docker image address:

```shell
docker pull amirhossein21/ptaas-tool:ftp-v0.X.X
```

### environment variables

- ```HTTP_PORT``` which is the system port
- ```ACCESS_KEY``` which is used to block download access for files
- ```PRIVATE_KEY``` which is used to give execute script system call enable only for core
- ```MINIO_CLUSTER``` is the minio configs. ```access:secret@endpoint&bucket&true```

## Setup

Setup ftp server in docker container with following command:

```shell
docker run -d \
  -e HTTP_PORT=80 \
  -p 80:80 \
  amirhossein21/ptaas-tool:ftp-v0.X.X
```

## libatks

Put your new attacks in libatks directory. If you add ```draf``` at the end, it will be ignored.
Make sure to have only one ```main.go``` file for your attack with the following code in the beginning of
the ```main``` function.

```go
log.SetOutput(os.Stdout)

var (
  hostFlag      = flag.String("host", "localhost", "target host address")
  endpointsFlag = flag.String("endpoints", "/", "target specific endpoints")
  paramsFlag    = flag.String("params", "", "system parameters for testing")
)

flag.Parse()

endpoints := strings.Split(*endpointsFlag, ",")
paramSet := strings.Split(*paramsFlag, "&")

params := make(map[string]string)

for _, item := range paramSet {
  parts := strings.Split(item, "=")
  params[parts[0]] = parts[1]
}

log.Println(*hostFlag)
log.Println(endpoints)
log.Println(params)
```
