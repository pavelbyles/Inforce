URL='http://127.0.0.1:8080/_ah/api/discovery/v1/apis/inforce/v1/rest'
curl -s $URL > inforce.rest.discovery
$GO_SDK/endpointscfg.py gen_client_lib java inforce.rest.discovery
unzip inforce.rest.zip
rm inforce.rest.zip
