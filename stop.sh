docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
rm fabcar.tar.gz