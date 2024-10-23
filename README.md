# zt
Quickly develop with multiple **Ztreamz** in your team!

A simple git-handover tool which focuses on **trunk based development**, where a team has one or multiple ensembles or
pairs in parallel.
A branch ends in a merge request, which is why the workflow of zt uses `main` or `master` as the origin and a new branch
is created as the start of a session.

### Installation
Simplest done with `go`

```bash
go install github.com/fehlhabers/zt@latest
```

### Workflow
`zt` is written using `cobra`, which allows for easy built help. The basic commands are:

- `create`
- `start`
- `next`
- `close`
