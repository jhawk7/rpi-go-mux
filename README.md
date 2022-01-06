# rpi-go-mux
Small golang application using gorilla-mux for ARM processors
UPDATE: Docker now supports multi-arch builds. This project can now be build using arm32v7 architecture and hosted on any machine.
* Run `docker buildx build --platform linux/arm/v7 -t <tag> .`