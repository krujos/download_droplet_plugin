#!/bin/bash -e

OUTPUT=$(pwd)/prepared-release

cp source/.git/ref ${OUTPUT}/commit
cp version/number ${OUTPUT}/tag

echo "$(cat ${OUTPUT}/tag)-$(cat ${OUTPUT}/commit)" | \
  tr -d '[:space:]' | \
  tee ${OUTPUT}/name

mkdir -p ${OUTPUT}/built

pushd built-plugins > /dev/null
  for exe in download_droplet_plugin_*; do
    sha1sum $exe >> ${OUTPUT}/body
    cp $exe ${OUTPUT}/built
  done
popd > /dev/null
