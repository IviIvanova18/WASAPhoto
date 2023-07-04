# WASAPhoto
##How to build

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:
```
go build ./cmd/webapi/
```
If you're using the WebUI and you want to embed it into the final executable:
```
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
```

##How to build container images
###Backend
```
$ docker build -t wasa-photos-backend:latest -f Dockerfile.backend .
```
###Frontend
```
$ docker build -t wasa-photos-frontend:latest -f Dockerfile.frontend .
```
##How to run container images
###Backend
```
$ docker run -it --rm -p 3000:3000 wasa-photos-backend:latest
```
###Frontend
```
$ docker run -it --rm -p 8081:80 wasa-photos-frontend:latest
```
##License

See LICENSE.
