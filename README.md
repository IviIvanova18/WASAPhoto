# WASAPhoto
## Introduction 
Stay connected with your friends and share those unforgettable moments with the power of WASAPhoto! With WASAPhoto, you can effortlessly upload your cherished photos from your PC and instantly make them accessible to all your followers. This project is your gateway to a seamless photo-sharing experience, ensuring that your special memories are easily shared and treasured.

WASAPhoto contains the source code for [Web and Software Application](http://gamificationlab.uniroma1.it/en/wasa/) course @ university La Sapienza.

## Design Specifications
1. API Design Using OpenAPI Standard:
2. Backend Development in Go, using database SQLite
3. Frontend Development in JavaScript, using Vue.js framework
4. Docker Container for Deployment
## How to build

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

## How to build container images
### Backend
```
$ docker build -t wasa-photos-backend:latest -f Dockerfile.backend .
```
### Frontend
```
$ docker build -t wasa-photos-frontend:latest -f Dockerfile.frontend .
```
## How to run container images
### Backend
```
$ docker run -it --rm -p 3000:3000 wasa-photos-backend:latest
```
### Frontend
```
$ docker run -it --rm -p 8081:80 wasa-photos-frontend:latest
```
## License
Based on [Fantastic Coffee Decaffeinated](https://github.com/sapienzaapps/fantastic-coffee-decaffeinated/) project
