#!/bin/bash -e

OUTPUT=$(pwd)/prepared-release

cp version/number ${OUTPUT}/tag

echo "$(cat ${OUTPUT}/tag)-$(cat source/.git/ref)" | \
  tr -d '[:space:]' | \
  tee ${OUTPUT}/name

mkdir -p ${OUTPUT}/built

pushd built-plugins > /dev/null
  cat > ${OUTPUT}/body <<'EOF'
### Checksums
```
EOF
  for exe in download_droplet_plugin_*; do
    sha1sum $exe >> ${OUTPUT}/body
    cp $exe ${OUTPUT}/built
  done
popd > /dev/null

echo '```' >> ${OUTPUT}/body
