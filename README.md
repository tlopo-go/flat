# Flat

Flat is a command line tool to flatten json and yaml structures

## Why? 

I wrote this tool a while ago when dealing with a great deal of databags for Chef which are basically json files,  I wanted look inside those files and grep for a specify path, I could not do it, I would need to flatten the structure before I could grep. 

It also helps me  a lot when writting kubernetes, helm charts and cloudformation files. 

## How to install

1. Go the [Releases Page](https://github.com/tlopo-go/flat/releases)
2. Download the desired version
3. Extract
4. Save the flat binary to a directory in your path

With one-liner: 
```
curl 'https://github.com/tlopo-go/flat/releases/download/v0.1.2/flat_0.1.2_darwin_amd64.tar.gz' -s -L | tar -C /usr/local/bin  -x flat
```

## Usage 

Flattening using the default separator ` | `: 
```
$ curl https://postman-echo.com/get  -s | flat
args = {}
headers | x-forwarded-proto = https
headers | host = postman-echo.com
headers | accept = */*
headers | user-agent = curl/7.54.0
headers | x-forwarded-port = 443
url = https://postman-echo.com/get
```

Flattening specifying `.` as separator: 

```
$ curl https://postman-echo.com/get  -s | flat -s .
args = {}
headers.x-forwarded-proto = https
headers.host = postman-echo.com
headers.accept = */*
headers.user-agent = curl/7.54.0
headers.x-forwarded-port = 443
url = https://postman-echo.com/get
```
