build_cli:
	go build -o weblang-deb/usr/local/bin/weblang weblang/weblang.go
	dpkg-deb --build weblang-deb