yaml_to_json.exe: test.cc 
	g++ $+ -o $@ -L./ -lyaml_to_json -lpthread
libyaml_to_json.so: yaml_to_json.go
	go build -o $@ -buildmode=c-shared $+
strip:
	strip *.so *.exe
clean:
	rm -rf *.so *.exe
