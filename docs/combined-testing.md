# Combined testing build

This fork-only branch combines the selected changes for local station testing.
It is not an upstream Graywolf release and is not a permanent feature branch.
No upstream FX.25 PR is opened by this build.

Base: upstream/main `bde16a1752ef96e8907aea9dc18c698e3e96a9d7`.

| Change | Exact source head |
|---|---|
| #674 exact-source APRS-IS-to-RF allow rules | 3f64275a38ec5defccf308f453f90ae68853d260 |
| #675 beacon labels, including symbol retention | 04fae68220c2c655c754c828c03ab55cc794a9ff |
| #676 independent serial RTS/DTR PTT | d1e94fc447bce0890b7b0c80c36da385946b0efd |
| #680 source-attributed TX outcomes | d23d0af989656eb213a0c76d57728c219fa6488d |
| #682 KISS interface display names | 26f24b3fa7f837fe4d7115dfcc2f3a5de4f4c07f |
| #684 tracker GPS freshness suppression | a871b0e4cf239d49538589df773252f71a53dc7e |
| Native 1200-baud FX.25, awaiting upstream PR | 4d54af153e101b2dd99add49a82b32a99df22b68 |
| Fork #8 Signal Light/Dark themes, responsive Dashboard and packet presentation | 8b79e9db522afed15a37569afb9a766111088307 |
| Fork #9 edge-assisted packet-log resizing and pinned desktop navigation | 02550a3f7d3a94629bc01375cecab9086ecc2939 |
| Fork #7 native station-wide T# telemetry | 06d7691ba511b57e199c19a2eec53c9667254ffd |
| Upstream #657 preserve repeated RF path markers | 0883688836bf852ae166b16382ef239bbbe499c2 |
| Upstream #617 beacon comment-command UI | 90e7cda3ce987d7fd8ca1dee6c4a5fc7bd100160 |

This is twelve selected changes: six existing upstream PRs from this fork,
the unsubmitted FX.25 branch, fork drafts #7/#8/#9, and selected upstream
PRs #657 and #617. The
`ci/tracker-fresh-gps` test branch and `fix/digipeater-inherit-first-save`
are excluded. Digipeating retains upstream's explicit enablement behavior.

The updated feature heads include operator documentation, including the
1200-baud-only FX.25 limitation. #676 retains stable Linux
`/dev/serial/by-id/` choices and FX.25 retains unconditional badge wrapping.
Signal Light and Signal Dark are selected through Preferences > Theme.
They replace the separate Enhanced packet view selector. Graywolf remains the
default, and existing themes retain their original packet/monitor presentation.
The former browser toggle is ignored; it does not change the selected theme.
Signal bundles Inter and JetBrains Mono for offline mixed typography. Its exact
packet/interface palette is documented in [Signal themes](signal-themes.md).
Signal packet rows/cards use explicit flow labels, accents and separators;
badges grow across available screen width and wrap when needed. The Dashboard
and log pages expand on wide screens while settings forms remain bounded.
KISS-only channels show actual connection/activity state without fabricated
audio measurements. All enabled RF-capable beacon buttons use the existing
label helper and wrap into an action row when multiple beacons are configured.
APRS-IS-only buttons remain on the Beacons page. No simulated-TX indicator is
added. Payloads, exports, scheduling and RF routes retain their data/behavior.

The Dashboard log has a pointer/keyboard resize grip, browser-saved preferred
height, reset control, 200 px minimum and content-aware 1600 px maximum;
the default remains 400 px and the feed still fetches 20 packets. Holding
the pointer near the viewport edge continues resizing and scrolls the page.
Desktop primary navigation stays visible while Settings scrolls independently.
Native T# telemetry is disabled by default and has a dedicated Telemetry page.
APRS-IS-only is the default/recommended delivery for Internet-connected
stations. Existing delivery selections are preserved. RF-only and IS-preferred
with RF fallback require an RF channel even when saving a disabled reporter.
It reports station-wide RF Rx, Digi, RF Tx, RF-to-IS and IS-to-RF activity after
a full ten-minute ready-input baseline. Status/history and delivery settings
are documented in the telemetry handbook. Definition frames are spaced to
avoid a burst. This combined build includes forwarded Digi accounting; the
standalone telemetry branch retains its documented fallback for compatibility.

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
job merges any missing manifest heads with normal fast-forward-only branch
publication. Both matrix jobs check out that exact merged source SHA,
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
is published as a regular fork testing release and marked Latest so existing
custom-repository installers using `releases/latest` can discover it. This
publication flag does not establish operator acceptance. Its title, version
and notes explicitly
identify it as a combined testing build, not an upstream stable release.
This does not modify the PacketRover installer or its unfinished selector.
Do not assume compatibility with that selector until it is implemented/tested.

## Appliance database persistence

Graywolf saves telemetry's next sequence in its main configuration SQLite
file before submitting a report. Restoring an older database also restores
its older sequence. PacketRover's change to keep this database on persistent
storage is implemented separately, not packaged by this Graywolf build.
Until the appliance update is installed, a snapshot-restoring appliance can
still reset sequences after reboot. Ten-minute samples and recent submission
history intentionally remain process-local.

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
and Signal Light/Dark in both packet logs and the live monitor. Switch back to
existing themes and confirm their normal presentation. Check theme reload,
offline fonts, wide/narrow viewers, native/KISS-only channels, and long/multiple
beacon labels. Confirm grayscale/high-contrast and forced-color modes stay
readable. Check packet separation, resizing and saved-height reload/reset on
both desktop and mobile, including keyboard input and UI scaling. Test native
telemetry readiness, ten-minute baseline, route selection and recent status
with explicit operator-controlled sends. Each architecture needs its own
operator acceptance.
Use the sending station's actual position; initiate each RF packet manually.
No unattended RF test is included. Re-image the disposable station afterward
or reinstall the previous known-good paired package before restoring service.
