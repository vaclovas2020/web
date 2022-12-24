build_cli:
	go build -o weblang/usr/local/bin/weblang main/main.go
	dpkg-deb --build weblang