HERO_ID="1"
curl -s https://zone01normandie.org/assets/superhero/all.json | jq ' .[] | select( .id == $HERO_ID )' | jq ' .connections.relatives' | sed 's/"//g'