@echo off
REM Download using a flat structure
go run main.go -urlfile ListOfAsciiSiteUrls.txt -dir downloads_flat -tree false -workers 20

REM Download using a tree structure
go run main.go -urlfile ListOfAsciiSiteUrls.txt -dir downloads_tree -tree true 