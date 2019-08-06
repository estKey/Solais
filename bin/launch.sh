docker build --rm -f "../Dockerfile" -t solais:latest ../
docker run --rm -it -p 8080:8080/tcp solais:latest
