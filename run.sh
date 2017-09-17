#!/bin/bash
go build
kill -9 $(lsof -ti tcp:8000)
GO_ENV=production PORT=8000 ./mitty-server >> release.log &
