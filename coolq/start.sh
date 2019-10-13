docker stop cqhttp
docker rm cqhttp
docker run -d \
    --name cqhttp \
    -v $PWD/data:/home/user/coolq \
    -p 9123:9000 \
    -p 5767:5767 \
    -e VNC_PASSWD=12345678 \
    -e COOLQ_ACCOUNT=2951493615 \
    -e CQHTTP_SERVE_DATA_FILES=yes \
    richardchien/cqhttp