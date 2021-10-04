#!/bin/bash
set -e

echo "# ================"
echo "# =  POLL APP    ="
echo "# ================"
printf "\n"

echo "> Ping Go Service.."

ENDPOINT_URI="http://google.com"
while true; do
  printf "\n > "
  curl -XGET ${ENDPOINT_URI}
  sleep 3s
done
