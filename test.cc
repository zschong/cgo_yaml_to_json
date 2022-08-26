#include <cstdio>
#include <string>
#include <cstring>
#include <stdlib.h>
#include "libyaml_to_json.h"

int main(int argc, const char **argv) {
	GoString path;

	if (argc > 1) {
		path.p = argv[1];
		path.n = strlen(path.p);
	} else {
		printf("usage: yaml_to_json.exe xxx.yaml\n");
	}
	char *p = yaml_to_json(path);
	if (p) {
		printf("%s\n", p);
		free(p);
	}

	return 0;
}
