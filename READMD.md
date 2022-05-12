docker build -t docker_log -f ./Dockerfile .

docker run --name docker_log_ins -v /var/run/docker.sock:/var/run/docker.sock -d docker_log