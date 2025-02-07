docker build . -t nowfeaturing:latest
docker stop nowfeaturing
docker rm nowfeaturing
docker run -d -p 8090:8090 --name nowfeaturing nowfeaturing:latest