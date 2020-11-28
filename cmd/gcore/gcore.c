#define _GNU_SOURCE
#include <errno.h>
#include <fcntl.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int pid;

static void
usage() {
	fprintf(stderr, "usage: %s pid | gzip >core.gz\n", program_invocation_short_name);
}

static int
readlink_ns(int pid, const char *ns, char *link, size_t sz) {
	char path[32];
	int i;
	ssize_t n;

	if (pid) {
		i = snprintf(path, sizeof(path), "/proc/%d/ns/%s", pid, ns);
	} else {
		i = snprintf(path, sizeof(path), "/proc/self/ns/%s", ns);
	}
	if (i >= sizeof(path)) {
		fprintf(stderr, "snprintf: overflow\n");
		return -1;
	}

	n = readlink(path, link, sz - 1);
	if (n < 0) {
		perror("readlink");
		return -1;
	}

	link[n] = '\0';
	return 0;
}

static int
set_ns(int pid, const char *ns) {
	char path[32];
	int fd;

	if (snprintf(path, sizeof(path), "/proc/%d/ns/%s", pid, ns) >= sizeof(path)) {
		fprintf(stderr, "snprintf: overflow\n");
		return -1;
	}

	fd = open(path, O_RDONLY);
	if (fd < 0) {
		perror("open");
		return -1;
	}

	if (setns(fd, 0) < 0) {
		close(fd);
		perror("setns");
		return -1;
	}

	close(fd);
	return 0;
}

static int
nscmp(int pid, const char *ns) {
	char ourns[32];
	char theirns[32];

	if (readlink_ns(0 /* self */, ns, ourns, sizeof(ourns)) < 0) {
		return -1;
	}

	if (readlink_ns(pid, ns, theirns, sizeof(theirns)) < 0) {
		return -1;
	}

	return !!strcmp(ourns, theirns);
}

static int
nspid(int pid) {
	char path[32];
	FILE *f;
	char *line = NULL;
	size_t n = 0;
	ssize_t ct;
	int rv = -1;

	if (snprintf(path, sizeof(path), "/proc/%d/status", pid) >= sizeof(path)) {
		fprintf(stderr, "snprintf: overflow\n");
		return -1;
	}

	f = fopen(path, "r");
	if (!f) {
		perror("fopen");
		return -1;
	}

	while (1) {
		ct = getline(&line, &n, f);
		if (ct == -1) {
			break;
		}

		if (!strncmp(line, "NSpid:\t", sizeof("NSpid:\t") - 1)) {
			rv = atoi(strrchr(line, '\t') + 1);
			break;
		}
	}

	if (line) {
		free(line);
	}

	fclose(f);

	return rv;
}

__attribute__((constructor)) void
init(int argc, const char **argv) {
	char *endptr;
	int _pid;
	int need_setns_pid;
	int need_setns_mnt;

	if (argc != 2 || !*argv[1]) {
		usage();
		exit(1);
	}

	_pid = strtol(argv[1], &endptr, 10);
	if (*endptr || _pid < 1) {
		usage();
		exit(1);
	}

	need_setns_pid = nscmp(_pid, "pid");
	if (need_setns_pid == -1) {
		exit(1);
	}

	need_setns_mnt = nscmp(_pid, "mnt");
	if (need_setns_mnt == -1) {
		exit(1);
	}

	if (need_setns_pid) {
		pid = nspid(_pid);
		if (pid < 0) {
			exit(1);
		}

		if (set_ns(_pid, "pid") < 0) {
			exit(1);
		}

		if (fork()) {
			exit(0);
		}

	} else {
		pid = _pid;
	}

	if (need_setns_mnt) {
		if (set_ns(_pid, "mnt") < 0) {
			exit(1);
		}
	}

	return;
}
