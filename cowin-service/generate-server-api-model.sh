SERVER_API_PATH="model/generated/server/api"

rm -rf $SERVER_API_PATH
swagger generate model -m=$SERVER_API_PATH --spec="swaggerui/swagger.json"
go fmt "cowin-service/${SERVER_API_PATH}"
