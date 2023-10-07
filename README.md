# PTaaS FTP server

![](https://img.shields.io/badge/language-golang_v1.20-blue)
![](https://img.shields.io/badge/version-v0.2.1-green)

This is file transfer part of our system. In this app we upload our attack scripts,
we execute attack scripts, and we download the results. This system provides an interface
in order to manage scripts and logs files from core app.

## Image

FTP server docker image address:

```shell
docker pull amirhossein21/ptaas-tool:ftp-v0.2.1
```

## Environment Variables

- ```HTTP_PORT``` which is the system port
- ```ACCESS_KEY``` which is used to block download access for files
- ```PRIVATE_KEY``` which is used to give execute script system call enable only for core
- ```MINIO_CLUSTER``` is the minio configs. ```access:secret@endpoint&bucket&true```

## Setup

Setup ftp server in docker container with following command:

```shell
docker run -d \
  -p 80:80 \
  amirhossein21/ptaas-tool:ftp-v0.2.1
```
