# bron

<img src="http://beforeitsnews.com/mediadrop/uploads/2013/38/1589ea296eb82ead17d6b307601025236ec0b738.jpg">

## Dependencies

```
git
```

If using the `-viz` option, you'll need to have [blessed-contrib](https://github.com/yaronn/blessed-contrib) installed.  `bron` expects the `blessed-contrib` src directory to be a cousin to the `bron` src directory.

For example if your `bron` src directory is located at:

```
/go/src/github.com/gophergala/bron
```

Then your `blessed-contrib` src directory should be located somewhere like this:

```
/go/src/github.com/yaronn/blessed-contrib
```

After installing `blessed-contrib` be sure and copy over the dashboards directory from `bron` like so:

```
cp -R dashboards ../../yaronn/blessed-contrib/
```

## Installation

#### Run on Nitrous:

[![Hack gophergala/bron on Nitrous](https://d3o0mnbgv6k92a.cloudfront.net/assets/hack-l-v1-d464cf470a5da050619f6f247a1017ec.png)](https://www.nitrous.io/hack_button?source=embed&runtime=go&repo=gophergala%2Fbron)

#### Run on Docker:

Note, the Docker image (from the Hub or from source) will already include all of the above mentioned dependencies, so you can skip those steps.

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

#### Install using Go:

```
go get github.com/gophergala/bron
bron -h
```

#### Using Docker Volumes to Persist Results

```
docker run -it -v /home/go/src/github/gophergala/bron/db:/go/src/github.com/gophergala/bron/db bron -viz
```
