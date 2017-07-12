#!/bin/bash

CLIENT_ID="uzR6MXvX8qczmMSq0s1LoA"
CLIENT_SECRET=""

curl -s -H "Content-Type: application/x-www-form-urlencoded" \
        -XPOST "https://api.yelp.com/oauth2/token" \
        -d "client_id=$CLIENT_ID&client_secret=$CLIENT_SECRET"
