#npm install @openapitools/openapi-generator-cli -g
cleanDirectory() {
  rm -rf $CLIENTPATH
  mkdir -p $CLIENTPATH
}

removeClientExtraFiles() {
  rm -rf ${CLIENTPATH}/docs
  rm -rf ${CLIENTPATH}/.openapi-generator
  find $CLIENTPATH -type f -not -name "model*" -exec rm -f {} \;
  rm -rf ${CLIENTPATH}/api
}

packageNames=( cowin )
for packageName in "${packageNames[@]}"
  do
    CLIENTPATH="model/generated/client/${packageName}"
    cleanDirectory
    openapi-generator-cli generate -g go-experimental -i resources/interface/ms/${packageName}/cowin.yml -o ${CLIENTPATH} --package-name ${packageName} -p enumClassPrefix=true
    removeClientExtraFiles
    go fmt "cowin-service/"$CLIENTPATH"/"
  done

