#! /bin/bash
curl -s https://api.github.com/users/aelkarou | jq ' .id'