![QR Icon](https://cdn.iconscout.com/icon/free/png-256/qr-code-1851030-1569017.png)
# FILE QRS

[![Docker Pulls](https://img.shields.io/docker/pulls/maguilag/file-qrs)](https://hub.docker.com/r/maguilag/file-qrs)

A simple golang script to publish files of a local folder via http and generate QR codes for published urls.

![Index Sample](https://github.com/rsierra/file-qrs/blob/master/index.png)

## ⭐ Features

* Simple html service.
* Generate QR codes for file urls.
* Allow subfolder navigation.
* Control access by .htpasswd file.

## 📜 How it works

Command:

```bash
HTPASSWD_FILE=path-to-htpasswd file-qrs -d path-to-files -p port
```

Options:

* 📁 **-d** => local path to the directory to publish. Is the current directory by default.
* 🔌 **-p** => port for local server. 8100 by default.
* 🔑 **HTPASSWD_FILE** => optinal environment variable for htpasswd file if you need basic http auth to control access.

NOTE: if you doesn't have a `htpasswd` file in your server, you can create one with `htpasswd` command from `apache2-utils` package or you can add users to a file with an online generator like [this](https://hostingcanada.org/htpasswd-generator/).

## 🐳 Docker

Run with docker:

```bash
docker run -d \
  --name file-qrs \
  -v /local-path:/files \
  -v /local-path-to-htpasswd:/.htpasswd
  maguilag/file-qrs
```

Run with docker-compose:

```yml
file-qrs:
  image: maguilag/file-qrs
  container_name: file-qrs
  environment:
    - HTPASSWD_FILE=.htpasswd
  volumes:
    - <path to data>:/files
    - <path to htpasswd>/.htpasswd
  ports:
    - 8100:8100
  restart: unless-stopped
```

## Compiling SCSS

1. Install globally `nodemon` & `node-sass`
```bash
npm install -g nodemon node-sass
```

2. Start watching src/css folder and output compiled css to app.css
```bash
nodemon --watch statics/css/src -e css,scss -x "node-sass --output-style compressed statics/css/src/app.scss statics/css/app.css"
```

## 🔨 Build binary

Install golang, download code and build with:

```bash
go build -o file-qrs
```

## 🚧 Working on

* Automatic binaries release.
* Add html styles.
* Tests.
