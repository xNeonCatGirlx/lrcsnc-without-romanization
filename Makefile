VERSION ?= ${shell git tag --sort=-version:refname | head -n 1}
COMMIT ?= ${shell git log -n 1 | head -n 1 | cut -d' ' -f2 | cut -b 1-16}

GO ?= "go"
BIN ?= "lrcsnc"
DESTDIR ?= 
PREFIX ?= "/usr/local"
PACKAGEDIR ?= "${BIN}_${VERSION}"
PACKAGENAME := "${PACKAGEDIR}.tar.gz"

LDFLAGS_VERSION ?= -X lrcsnc/internal/setup.version=${VERSION}-${COMMIT}
LDFLAGS ?= \
	${LDFLAGS_VERSION}

default: build
all: build install clean

build:
	${GO} build -ldflags="${LDFLAGS}" -o ${BIN}
install: build
	install -Dm755 ${BIN} ${DESTDIR}${PREFIX}/bin/${BIN}
package: build
	strip ${BIN}
	cp -t ${PACKAGEDIR} ${BIN}
	tar -czvf ${PACKAGE} ${PACKAGEDIR}
clean:
	rm -f lrcsnc