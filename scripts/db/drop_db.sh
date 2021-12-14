USER=$(cat .env | grep POSTGRES_USER | awk -F"=" '{print $2}')
DB=$(cat .env | grep POSTGRES_DB | awk -F"=" '{print $2}')
CONTAINER_ID=$(docker ps | grep gochat_postgres | awk '{print $1}')

docker exec -it $CONTAINER_ID dropdb --username=$USER $DB
