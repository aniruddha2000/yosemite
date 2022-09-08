FROM ubuntu

COPY ./bin/yosemite ./yosemite

ENTRYPOINT ["./yosemite"]
