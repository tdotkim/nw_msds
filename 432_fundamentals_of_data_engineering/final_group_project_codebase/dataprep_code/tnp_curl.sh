curl --header "X-App-Token: YOUR_APP_TOKEN" \
-X GET 'https://data.cityofchicago.org/resource/m6dm-c72p.json?$limit=50000000' \
-o 'tnp_full.json'