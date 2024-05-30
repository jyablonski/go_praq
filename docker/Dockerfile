
FROM golang:1.22

# idk if this works yet chill dawg
RUN useradd --create-home appuser
WORKDIR /home/appuser
USER appuser

COPY ./src ./src

CMD ["go", "test", "./src", "-v"]