# alexadoom


Alexa skill to have Alexa whisper the 'I Dogma' intro to the Doom 2016 game.  
https://www.youtube.com/watch?v=nMd--Wyfw0g


## build and deploy

build the go file into an executable.  
`GOOS=linux GOARCH=amd64 go build -o alexa alexa.go`

zip up the go executable.  
`zip alexa.zip alexa`

update the lambda zip in S3.  