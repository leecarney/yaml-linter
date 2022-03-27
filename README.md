Yaml linter

Build image using:
docker build -t yaml-linter . --no-cache

To Use:
cd to your folder containing cicd-metadata.yaml
Run: docker run -v $(pwd):/tmp -it docker.io/library/yaml-linter# yaml-linter
