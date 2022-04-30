#!/usr/bin/env bash
rm -rf .git && \
git clone https://github.com/PasteUs/CDN.git -b master pasteme_cdn && \
npm ci && \
npm run build --if-present && \
bash -x .github/shell/copy.sh && \
cd pasteme_cdn && \
git config user.name "Lucien Shui" && \
git config user.email "lucien@lucien.ink" && \
git add --all && \
set +x && \
git diff-index --quiet HEAD || (git commit -m "Pushed by github action $(TZ=UTC-8 date +'%Y-%m-%d %H:%M:%S')" && \
git push https://"${GH_TOKEN}"@github.com/PasteUs/CDN.git master) && \
set -x && \
cd ..
