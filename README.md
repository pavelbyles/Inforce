Just testing ...

1. Appengine
2. Testing (TDD)
3. Templates
UI
4. Bootstrap
5. Less

sudo npm install -g bower
sudo npm install -g vulcanize
bower init
bower install jquery
bower install font-awesome
bower install bootswatch
bower install less
bower install Happyjs
bower install polymer Polymer/polymer-elements Polymer/polymer-ui-elements
bower install Polymer/core-ajax

View / test api: http://localhost:8080/_ah/api/explorer

Generate client libs:
URL='http://localhost:8080/_ah/api/discovery/v1/apis/company/v1/rest'
curl -s $URL > inforce.rest.discovery
$GO_SDK/endpointscfg.py gen_client_lib java inforce.rest.discovery
unzip inforce.rest.zip

// npm update -g bower
