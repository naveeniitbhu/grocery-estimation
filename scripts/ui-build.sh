#!/bin/bash

cd ../web
# npm install
npm run build
cd ..
rm -rf static
mkdir static
cp web/build/* static -r

