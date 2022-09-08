FROM ubuntu

COPY ./bin/do_client_go ./do_client_go

ENTRYPOINT ["./do_client_go"]
