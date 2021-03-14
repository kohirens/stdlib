#!/bin/sh

# You can get the latest commit ID by looking at the latest tagged commit here: https://github.com/repos/microsoft/vscode/releases/latest
commit_id="622cb03f7e070a9670c94bae1a45d78d7181fbd4"
hash="1612749334253"
archive="vscode-server-linux-x64.tar.gz"
# repo="microsoft/vscode"

# TODO: Find a way to get the latest commit ID via command line: https://api.github.com/repos/microsoft/vscode/releases/latest
# https://api.github.com/repos/microsoft/vscode/releases/latest
#
# get_latest_release() {
#   curl --silent "https://api.github.com/repos/$1/releases/latest" | # Get latest release from GitHub api
#     grep '"tag_name":' |                                            # Get tag line
#     sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
# }

# commit_tag=$(get_latest_release $repo)
# echo "commit id = ${commit_tag}"

# Download VS Code Server tarball to tmp directory.
curl -sSL "https://update.code.visualstudio.com/commit:${commit_id}/server-linux-x64/stable" -o /tmp/${archive}

# Make the parent directory where the server should live. NOTE: Ensure VS Code will have read/write access; namely the user running VScode or container user.
mkdir -p ~/.vscode-server/bin/${commit_id}_${hash}

# Extract the tarball to the right location.
tar --no-same-owner -xz --strip-components=1 -C ~/.vscode-server/bin/${commit_id}_${hash} -f /tmp/${archive}

mv -n ~/.vscode-server/bin/${commit_id}_${hash} /home/${USER_NAME}/.vscode-server/bin/${commit_id}
