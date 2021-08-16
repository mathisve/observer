![publish](https://github.com/mathisve/observer/actions/workflows/build-and-publish.yaml/badge.svg)
# observerBot

## Set-up
All variables are set using K8s secrets/env vars

- `docker build . -t observerbot`
- `docker tag observerbot mathisve/observerbot:latest`
- `docker push mathisve/observerbot:latest`

I am not liable or responsible for any damages incurred when using observerBot. Use on your own accord.
