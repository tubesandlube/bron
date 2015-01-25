# bron

<p align="center">
<img src="http://beforeitsnews.com/mediadrop/uploads/2013/38/1589ea296eb82ead17d6b307601025236ec0b738.jpg">
</p>

This hack was written by [@defermat](https://github.com/defermat) and [@schvin](https://github.com/schvin) as a submission in the first [Gopher Gala](http://gophergala.com/) over the weekend of 24-25 January, 2015. Please give it a whirl and let us know what you think.

Analysis of large source code repositories is always interesting, especially over long periods of time. Fun and useful to see various characteristics by language, who was actively contributing in a project, or when activity peaked. After coming across [blessed](https://github.com/yaronn/blessed-contrib), we knew we had a quick way visualize the analysis with no fuss or complications.

Target audience is other developers or consumers of development projects. Easy to glean more information about a project by looking at the source from a high-level viewpoint.

## Dependencies for manual installation

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

## Automated installation methods

#### Run on Nitrous:

[![Hack gophergala/bron on Nitrous](https://d3o0mnbgv6k92a.cloudfront.net/assets/hack-l-v1-d464cf470a5da050619f6f247a1017ec.png)](https://www.nitrous.io/hack_button?source=embed&runtime=go&repo=gophergala%2Fbron)

#### Run on Docker:

Note, the Docker image (from the Hub or from source) will already include all of the above mentioned dependencies, so you can skip those steps.

Get it from the [Docker Hub](https://registry.hub.docker.com/u/tubesandlube/bron/):

```
docker pull tubesandlube/bron
docker run -it bron -h
```

Or, build it yourself:

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
