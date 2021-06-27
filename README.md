# observerBot

Custom discord bot that saves every message in a DynamoDB table.

## Set-up
- Paste discord token and AWS Secrets in `Dockerfile`.
- (OPTIONAL) Change the role and channel ID's in `consts.go`
- `docker build . -t observerbot`
- `docker tag observerbot mathisco/observerbot:latest`
- `docker push mathisco/observerbot:latest`

I am not liable or responsible for any damages incurred when using observerBot. Use on your own accord.