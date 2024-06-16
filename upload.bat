git add .
git commit -m "Ultimo Commit"
git push 
set GOARCH=amd64 
set GOOS=linux go build -o bootstrap main.go
zip deployment.zip bootstrap
