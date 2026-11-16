# Combined testing build

This fork-only branch combines the selected changes for local station testing.
It is not an upstream Graywolf release and is not a permanent feature branch.
No upstream FX.25 PR is opened by this build.

Base: upstream/main `bde16a1752ef96e8907aea9dc18c698e3e96a9d7`.

| Change | Exact source head |
|---|---|
| #674 exact-source APRS-IS-to-RF allow rules | 5b7e2d074da87cc90f06af8f8ab33fe20937d2f1 |
| #675 beacon labels, including symbol retention | 9e3a3c2feef2e48365089b89e69a614b5d86bbb9 |
| #676 independent serial RTS/DTR PTT | eee674e51747b4f00e04c42b481ec8b8de61f4f3 |
| #680 source-attributed TX outcomes | 6c5a3cb9cb1d31e28a4c4e88306bb0fe25cfd921 |
| #682 KISS interface display names | a4d8452ff0de2cab0342219d0a3ac4ef933e3cbe |
| #684 tracker GPS freshness suppression | a871b0e4cf239d49538589df773252f71a53dc7e |
| Native 1200-baud FX.25, awaiting upstream PR | 1f41127677210ccc817dc5342fee8f44880a0936 |
| Optional theme-aware packet colors, awaiting upstream PR | c3ed4198bfbb67fb5f948f27923c2672bfc85eb3 |
| Upstream #657 preserve repeated RF path markers | 0883688836bf852ae166b16382ef239bbbe499c2 |
| Upstream #617 beacon comment-command UI | 90e7cda3ce987d7fd8ca1dee6c4a5fc7bd100160 |

This is ten selected changes: six existing PRs from this fork, the unsubmitted
FX.25 and packet-color branches, and upstream PRs #657 and #617. The
`ci/tracker-fresh-gps` test branch and `fix/digipeater-inherit-first-save`
are excluded. Digipeating retains upstream's explicit enablement behavior.

The latest #676 head adds stable Linux `/dev/serial/by-id/` choices. The
latest FX.25 head wraps packet badges unconditionally. Packet colors are a
separate opt-in display feature: Preferences > General > Colorize packets,
immediately after Theme. It is off by default and remembered by this browser.
With colors off, existing packet colors and plain live-monitor text remain.
Graywolf/Night use the approved colors; grayscale/high-contrast themes use
matching monochrome treatments. All payloads and exports retain their data.

## Merge compatibility

Tracker freshness and beacon labels originally inserted fields at the same
locations in the beacon request and editor. The #684 source branch moves its
new fields alongside GPS settings; it also keeps default values independent
of #617's comment-command defaults. #680 reverts unrelated formatting in
SubmitSource that collided with FX.25's frame-ID addition. The affected pairs
merge cleanly in either order without making one feature depend on another.
Existing behavior and regression tests are retained. Original PR branches
are updated by ordinary commits, not rebases or force pushes.

The combined handbook OpenAPI copies are synchronized with the generated
schema. Protobuf regeneration must reproduce the merged IPC schema.

## Fork release

`INTEGRATION_VERSION` selects a distinct combined-test version. The upstream
VERSION and in-app release-note history are not bumped. The embedded note
parser accepts only stable x.y.z entries; combined-test notes instead ship in
RELEASE_NOTES.md and on the fork release. Both binaries are stamped with
the integration version and the same fixed eight-character source commit.
The fork-only `combined-release.yml` workflow runs on this branch. Its first
job merges the three pinned update heads above with normal fast-forward-only
branch publication. Both matrix jobs check out that exact merged source SHA,
then run web/API, Go race, Rust/PTT and pinned Dire Wolf interoperability checks.
They build natively in Debian Bookworm on separate ARM64 and amd64 runners.
Publication waits for both jobs. Supported targets are Pi 4 on 64-bit Raspberry
Pi OS and Linux x86_64, with glibc 2.36 or later and ALSA. These are not 32-bit,
Android, macOS or Windows packages.

Assets use the normal Graywolf names:

- `graywolf_<version>_linux_arm64.tar.gz`: paired app/modem binaries at archive root.
- `graywolf_<version>_linux_x86_64.tar.gz`: paired app/modem binaries for Linux x86_64.
- `graywolf_<version>_arm64.deb` and `graywolf_<version>_amd64.deb`:
  paired binaries, upstream service/udev files
  and upstream package hooks. Installation enables but does not start the
  service; it does not change ALSA levels.
- `graywolf_<version>_source.tar.gz`: exact source, including provenance and workflow.
- `BUILD-arm64.txt` and `BUILD-amd64.txt`: exact source SHA, tree and toolchains;
  each binary archive also contains its own `BUILD.txt`.
- `checksums.txt`: SHA256 for the release assets.

The workflow publishes only in `samwathegreat/graywolf`. The GitHub release
is published with its prerelease flag off so a custom-repository installer
using `releases/latest` can discover it. Its title, version and notes explicitly
identify it as a combined testing build, not an upstream stable release.
This does not modify the PacketRover installer or its unfinished selector.
Do not assume compatibility with that selector until it is implemented/tested.

## Installation and testing

Stage downloads under `/tmp`. Verify `checksums.txt`, the selected architecture,
and both component versions before replacing a test station's binaries.
Use a disposable installation for the initial combined-build acceptance.
An archive alone neither installs software nor starts radio transmission;
starting a previously configured service can activate its saved automation.
Keep scheduled beacons/gating off until reviewed. Do not run Dire Wolf and
native Graywolf against the same audio/PTT device simultaneously.

Existing FX.25-only physical results are in [FX.25 validation](fx25-validation.md).
They do not establish physical acceptance of this larger combination.
Recheck native RX/TX, PTT, duplicate handling, ordinary AX.25 and FX.25 fallback,
then exercise beacon labels/symbols, optional GPS freshness, KISS labels,
source allow rules, TX metrics, RF path preservation and comment commands.
Also check badge wrapping with several labels, stable serial by-id selection,
and the new color preference in both packet logs and the live monitor. Check
colors off, on, theme changes and reload; confirm grayscale/high-contrast stay
readable. Each architecture needs its own operator acceptance.
Use the sending station's actual position; initiate each RF packet manually.
No unattended RF test is included. Re-image the disposable station afterward
or reinstall the previous known-good paired package before restoring service.
