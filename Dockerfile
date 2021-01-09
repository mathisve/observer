FROM golang:1.14

WORKDIR /app
COPY . .

ENV TOKEN=DISCORD_TOKEN
ENV AWS_ACCESS_KEY_ID=AKID
ENV AWS_SECRET_ACCESS_KEY=SAK
ENV AWS_DEFAULT_REGION=REGION

# DynamoDB table where messages will be stored (gets created automatically)
ENV TABLE_NAME=TABLE_NAME

# Lamba function name of contentLambda
ENV FUNCTION=LAMBDAFUNCTION

# Channel where messages will be deleted within 5 minutes
ENV GARBAGE=CHANNELID

# Role where `--` deletion immunity prefix is allowed
ENV IMMUNE=ROLEID

RUN go build -o main .
CMD ["/app/main"]