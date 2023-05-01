go test ./... -cover -v -coverprofile cover.out
go tool cover -html cover.out -o cover.html
rm cover.out
open cover.html
rm -i cover.html