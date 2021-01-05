# observerBot

Custom discord bot that saves every message in a DynamoDB table.

## Set-up
- Paste discord token and AWS Secrets in `Dockerfile`.
- (OPTIONAL) Change the role and channel ID's in `consts.go`  
- Make a DynamoDB table with `authorId` (String) as a primary partition key and `messageId` (String) as the primary sort key.
- `docker build . -t observerbot`
- `docker run observerbot`

I am not liable or responsible for any damages incurred when using observerBot. Use on your own accord.