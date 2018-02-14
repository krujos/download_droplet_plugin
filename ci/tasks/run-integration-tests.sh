#!/bin/bash -e

for env_var in $(env | grep '^CF_' | cut -d '=' -f1); do
  if [ -z $(eval echo "\$${env_var}") ]; then
    echo "${env_var} is empty, skipping integration tests"
    exit 0
  fi
done

apt-get -y update
apt-get -y install wget

wget 'https://cli.run.pivotal.io/stable?release=linux64-binary&source=github' -O cli.tgz
tar -xvf cli.tgz
chmod +x ./cf

./cf api ${CF_API}
./cf auth ${CF_USER} "${CF_PASS}"
./cf target -s "${CF_SPACE}" -o "${CF_ORG}"
./cf install-plugin -f built-plugins/download_droplet_plugin_linux

installed_version=$(./cf plugins | grep '^download-droplet' | awk '{print $2}')
if [[ "${installed_version}" != "$(cat version/number)" ]]; then
  echo "Installed version '${installed_version}' does not match expected version '$(cat version/number)'"
  exit 1
fi

./cf download-droplet ${CF_APP} ./droplet.tar
tar -xvmf droplet.tar
