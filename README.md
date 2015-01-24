# bron

## Installation

Run on nitrous:

[![Hack gophergala/bron on Nitrous](https://d3o0mnbgv6k92a.cloudfront.net/assets/hack-l-v1-d464cf470a5da050619f6f247a1017ec.png)](https://www.nitrous.io/hack_button?source=embed&runtime=go&repo=gophergala%2Fbron)

Run on docker:

Get it from the Docker Hub:

```
docker pull tubesandlube/bron
docker run -it bron -h
```

or build yourself:
```
git clone https://github.com/gophergala/bron.git
cd bron
docker build -t bron .
docker run -it bron -h
```

Install using Go:

```
go get github.com/gophergala/bron
```

