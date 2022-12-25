build_cli:
	go build -o weblang/usr/local/bin/weblang main/weblang.go
	dpkg-deb --build weblang