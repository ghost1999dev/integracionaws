git add .
git commit -m "Ultimo Commit"
git push 
set GOARCH=amd64 
set GOOS=linux 
go build main.go
del deployment.zip bootstrap
tar.exe -a -cf deployment.zip bootstrap
