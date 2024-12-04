# Usage

## One ztream = one work package with one group
Ztreams is mainly used within one team. Should be possible to build out to multiple teams.

Whenever a part or the whole team starts with a ztream, the ztream is named and is kept throughout different
repositories.

```bash
# Configure static information
zt configure
# Create new zstream with metadata which will be appended to commits
# Creates a branch which is also pushed to remote
zt new feat/coolservice --meta 'Relates to issue: https://github.com/fehlhabers/zt/issues/1'

# Joins an existing ztream started by someone else
# Without the backend, it uses commit messages to get metadata. Branch is named after ztream
zt join feat/coolservice

# Joins the branch in new repository and starts timer
zt start

# Hands over the ztream to hext member. Commits and pushes WIP commit
zt next
```

