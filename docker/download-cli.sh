#! /bin/sh

URL=$(curl -s https://api.github.com/repos/Jokers13/tmc/releases/latest | jq -r '.assets | .[] | select(.name == "tmc-linux-amd64") | .browser_download_url')

curl -OL $URL
mv ./tmc-linux-amd64 ./tmc
