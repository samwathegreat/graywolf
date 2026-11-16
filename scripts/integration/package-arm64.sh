#!/usr/bin/env bash
# Package already-built paired ARM64 binaries. Does not install or start them.
set -euo pipefail
version=$(cat INTEGRATION_VERSION)
[[ "$version" =~ ^[0-9]+\.[0-9]+\.[0-9]+-combined\.[0-9]+$ ]]
bin_dir=${1:?usage: package-arm64.sh BINARY_DIRECTORY OUTPUT_DIRECTORY}
out_dir=${2:?usage: package-arm64.sh BINARY_DIRECTORY OUTPUT_DIRECTORY}
bin_dir=$(realpath "$bin_dir")
mkdir -p "$out_dir"
out_dir=$(realpath "$out_dir")
for binary in graywolf graywolf-modem; do
  test -x "$bin_dir/$binary"
  readelf -h "$bin_dir/$binary" | grep -q 'Machine:.*AArch64'
done
stage=$(mktemp -d)
trap 'rm -rf "$stage"' EXIT
commit=$(git rev-parse HEAD)
epoch=$(git show -s --format=%ct HEAD)
archive="graywolf_${version}_linux_arm64.tar.gz"
source="graywolf_${version}_source.tar.gz"
package="graywolf_${version}_arm64.deb"
mkdir -p "$stage/archive" "$stage/deb/DEBIAN" "$stage/deb/usr/bin" \
  "$stage/deb/usr/lib/systemd/system" "$stage/deb/etc/udev/rules.d" \
  "$stage/deb/usr/share/doc/graywolf"
install -m 0755 "$bin_dir/graywolf" "$bin_dir/graywolf-modem" "$stage/archive/"
cp LICENSE README.md RELEASE_NOTES.md docs/combined-testing.md "$stage/archive/"
{
  echo 'Graywolf combined testing build; not an official upstream release'
  echo "Version: $version"
  echo "Commit: $commit"
  echo "Tree: $(git rev-parse HEAD^{tree})"
  echo 'Target: Linux ARM64, glibc >= 2.36, ALSA'
  echo "Go: $(go version)"
  echo "Rust: $(rustc --version)"
  echo 'Exact selected PR heads: combined-testing.md'
} > "$stage/archive/BUILD.txt"
cp "$stage/archive/BUILD.txt" "$out_dir/BUILD.txt"
cp RELEASE_NOTES.md "$out_dir/RELEASE_NOTES.md"
tar --sort=name --mtime="@$epoch" --owner=0 --group=0 --numeric-owner \
  -C "$stage/archive" -cf - . | gzip -n > "$out_dir/$archive"
git archive --format=tar --prefix="graywolf-${version}/" HEAD | gzip -n > "$out_dir/$source"
install -m 0755 "$bin_dir/graywolf" "$bin_dir/graywolf-modem" "$stage/deb/usr/bin/"
install -m 0644 packaging/systemd/graywolf.service "$stage/deb/usr/lib/systemd/system/"
install -m 0644 packaging/udev/99-graywolf-cm108.rules "$stage/deb/etc/udev/rules.d/"
install -m 0755 packaging/scripts/postinstall.sh "$stage/deb/DEBIAN/postinst"
install -m 0755 packaging/scripts/preremove.sh "$stage/deb/DEBIAN/prerm"
printf '/etc/udev/rules.d/99-graywolf-cm108.rules\n' > "$stage/deb/DEBIAN/conffiles"
cp LICENSE RELEASE_NOTES.md docs/combined-testing.md "$stage/archive/BUILD.txt" "$stage/deb/usr/share/doc/graywolf/"
cat > "$stage/deb/DEBIAN/control" <<CONTROL
Package: graywolf
Version: $version
Section: hamradio
Priority: optional
Architecture: arm64
Maintainer: samwathegreat <samwathegreat@users.noreply.github.com>
Depends: libc6 (>= 2.36), libasound2 | libasound2t64
Homepage: https://github.com/samwathegreat/graywolf
Description: Graywolf combined testing build with native 1200-baud FX.25
 Fork integration build for supervised testing on 64-bit Raspberry Pi OS.
 Includes matching Graywolf application and native modem binaries.
Installed-Size: $(du -sk "$stage/deb/usr" | cut -f1)
CONTROL
SOURCE_DATE_EPOCH="$epoch" dpkg-deb --root-owner-group --build "$stage/deb" "$out_dir/$package"
(
  cd "$out_dir"
  sha256sum "$archive" "$package" "$source" BUILD.txt RELEASE_NOTES.md > checksums.txt
  sha256sum -c checksums.txt
)
