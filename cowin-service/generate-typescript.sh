npm install @openapitools/openapi-generator-cli -g
TYPESCRIPT_PATH="resources/generated/client/typescript/cowin-service"

cleanDirectory () {
  rm -rf $TYPESCRIPT_PATH
  mkdir -p $TYPESCRIPT_PATH
}
removeTypescriptExtraFiles () {
  mv ${TYPESCRIPT_PATH}/models /tmp
  rm -r ${TYPESCRIPT_PATH}/*
  rm -rf ${TYPESCRIPT_PATH}/apis
  rm -rf ${TYPESCRIPT_PATH}/.openapi-generator
  rm -r ${TYPESCRIPT_PATH}/.openapi-generator-ignore
  rm -r ${TYPESCRIPT_PATH}/.gitignore
  mv /tmp/models ${TYPESCRIPT_PATH}
  chmod +x ${TYPESCRIPT_PATH}
}
cleanDirectory
openapi-generator-cli generate -g typescript-rxjs -i swaggerui/swagger.json -o $TYPESCRIPT_PATH -p enumPropertyNaming=UPPERCASE
removeTypescriptExtraFiles
