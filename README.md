# Close Milestone With Release

This action will close the milestone with the same name of the latest release title. For example, we have a milestone with name "1.5.0" and the latest release with the title "1.5.0", and when this action is triggered, the milestone will be closed.

It is recommended to trigger this action while publishing a release, as in:

```yaml
on: release
```

## Usage

```yaml
uses: adlerhsieh/prepare-release@v0.1.2
env: 
  GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
  IGNORE_MILESTONE_NOT_FOUND: false
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
