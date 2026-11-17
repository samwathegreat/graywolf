# Combined testing 0.14.14-combined.4 for Linux ARM64 and x86_64

This is a testing release from samwathegreat/graywolf, not an official
upstream Graywolf release. It retains the twelve selected changes from combined.3
and updates the three recent feature PRs. Published as Latest for existing
custom-repository installers; operator acceptance is still pending.

New in this build:

- Packet-log dragging continues near the viewport edge and scrolls the page,
  so expansion is no longer limited by the bottom of the physical screen.
- Desktop logo/primary navigation stays visible while only Settings scrolls.
- Mobile enhanced packet accents override the legacy direction-only borders.
- The actual preference is now Enhanced packet view, with descriptive guidance;
  existing browser selections survive the rename. The feature remains off by default.
- Telemetry defaults to APRS-IS only, recommended for Internet-connected stations.
  Existing saved route choices are preserved. RF-only and IS-preferred/RF-fallback
  modes require an RF channel to save, including when telemetry is disabled.
- Operator handbook and PR documentation reflect these changes.

Telemetry sequencing is stored in Graywolf's main configuration database.
PacketRover's durable-database migration is a separate appliance update and is
not included in these packages. A host restoring older database snapshots can
still rewind sequences until its persistence update is installed.

Retained features:

- Native FX.25 transmit and receive at 1200 baud, with per-packet encoding,
  parity and repaired-symbol information. Auto is recommended when enabled;
  oversized frames use explicitly identified ordinary AX.25 fallback.
- Native station-wide T# activity telemetry, disabled by default, with ready-input
  ten-minute baseline, spaced definitions, delivery selection and recent status.
- Enhanced packet-flow colors/separators and browser-saved Dashboard log height,
  400 px default, 200 px minimum, content-aware maximum up to 1600 px.
- Independent RTS/DTR serial PTT polarity.
- Optional beacon display labels, preserving the APRS symbol.
- Optional suppression of tracker beacons when GPS updates become stale.
- Optional KISS interface display names.
- Exact-source APRS-IS-to-RF allow rules and clearer TX outcome counts.
- Preserve repeated digipeater markers when uploading RF packets to APRS-IS.
- Edit beacon comment commands in the Graywolf UI.

FX.25 transmit and GPS freshness suppression remain off by default. Review
channel audio/PTT, RF gating and beacon schedules before starting a station.
Comment commands run on the Graywolf host during beaconing; retain only your
intended commands. No mixer adjustment or automatic radio test is included.
The excluded digipeater-inherit-first-save change remains excluded.

Both architectures include matching Graywolf and native modem binaries.
Use linux_arm64/arm64 for a Pi 4 running 64-bit Raspberry Pi OS, or
linux_x86_64/amd64 for a 64-bit Intel/AMD Linux host. Both require glibc 2.36+
and ALSA (Debian Bookworm or later). There are no 32-bit or Android packages.

The source archive records every selected source head in
docs/combined-testing.md. BUILD-arm64.txt and BUILD-amd64.txt record the
actual built commit and tools; each binary archive contains BUILD.txt too.
Verify downloads using checksums.txt.

The smaller FX.25/PTT build passed operator RF tests. Physical acceptance of
this updated combination remains pending on each architecture. No general
RF performance advantage over Dire Wolf is claimed. Upstream FX.25 and
colorization/resizing/telemetry upstream submissions remain on hold.
