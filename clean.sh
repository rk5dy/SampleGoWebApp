docker container stop $(docker container ls -q)
docker container prune -f
docker rmi $(docker images | grep \<none\> | awk '{ print $3 }')
docker volume rm $(docker volume ls -q --dangling=true)
