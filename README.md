# Close Milestone With Tagging

## Usage

```yaml
uses: adlerhsieh/prepare-release@v0.1.0
env: 
  GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
with:
  repo_owner: ''
  repo_name: ''
```

## Development

Run locally:

```shell
GITHUB_TOKEN=mytoken REPO_OWNER=owner REPO_NAME=name go run main.go
```

Build:

```shell
docker build .
```
