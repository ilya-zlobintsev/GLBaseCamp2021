It's possible to build the app with an arbitrary version used both by the image tag and from within the app:

`VERSION=0.2 docker build -t app:$VERSION --build-arg VERSION=$VERSION .`

Image on a registry: https://ghcr.io/ilyazzz/app