name: CI
on:
  push:
    branches:
      - master

jobs:
  deploy:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v1
      - name: Get Node.js
        uses: actions/setup-node@v1
        with:
          node-version: 12
      - name: Install dependencies
        run: npm install
      - name: Testing
        run: CI=true npm test
      - name: Build
        run: npm run build
