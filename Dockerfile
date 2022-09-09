FROM ubuntu

COPY ./bin/yosemite ./yosemite
COPY ./scripts/entrypoint.sh ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
