[Ondrej Sika (sika.io)](https://sika.io) | [Sika Labs](https://sikalabs.com)

# supportmonkey

Simple service for sending random email for testing your support team

## Usage

### Docker

```
docker run ondrejsika/supportmonkey -h
```

### Kubernetes

Add `ondrejsika` charts:

```
helm repo add ondrejsika https://helm.oxs.cz
```

See values:

```
helm show values ondrejsika/supportmonkey
```

Deploy:

```
helm upgrade --install \
  supportmonkey \
  ondrejsika/supportmonkey \
  -f ./my-supportmonkey-values.yml
```
