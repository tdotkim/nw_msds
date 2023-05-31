curl --header "X-App-Token: YOUR_APP_TOKEN" \
-X GET 'https://data.cityofchicago.org/resource/wrvz-psew.json?$limit=50000000' \
-o 'taxico_full.json'