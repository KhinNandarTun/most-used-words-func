FROM golang:1.17 

ENV APP_USER="knt" \
    APP_WORKDIR="/app"

WORKDIR ${APP_WORKDIR}

COPY go.mod /${APP_WORKDIR}
#create app user
RUN useradd --create-home ${APP_USER}

#go mod download
RUN go mod download

# Copy go files
COPY *.go /${APP_WORKDIR}