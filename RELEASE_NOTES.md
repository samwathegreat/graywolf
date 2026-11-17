# Combined testing 0.14.14-combined.6 for Linux ARM64 and x86_64

This is a testing release from samwathegreat/graywolf, not an official
upstream Graywolf release. It retains all twelve selected changes from combined.5
and refines the Signal Dashboard audio meter and status layout. Published as Latest for existing
custom-repository installers; operator acceptance is still pending.

New in this build:

- Native DCD/RX/TX indicators sit beside the channel name. The numeric dBFS
  reading sits immediately before the RX Audio meter.
- A visible empty track and -60/-40/-20/0 dBFS scale make the level readable.
  Meter width responds to available space, capped at 480 px, and moves to its
  own row on narrow screens. KISS-only channels retain their TNC status.
- PR #9 scrolling/resizing has operator acceptance; its source is unchanged.

Retained Signal theme updates:

- Signal Light and Signal Dark in the existing Theme selector, replacing the
  separate Enhanced packet view switch. The existing default/themes remain
  available; the old packet-toggle value is ignored.
- Bundled Inter interface text and JetBrains Mono packet/terminal data, with
  offline font files and licenses. Exact palettes are in docs/signal-themes.md.
- Responsive wide Dashboard/log layouts, flexible badge/address columns and
  wrapping mobile packet cards with flow colors and clear packet separators.
- Theme-matched native DCD/RX/TX and audio presentation. KISS-only channels show
  actual TNC state/activity, without a fabricated audio meter or DCD reading.
- Multiple enabled RF-capable Beacon Now buttons use custom beacon labels and
  wrap below the channel status. APRS-IS-only beacons remain on the Beacons page.

Included scrolling and telemetry updates:

- PR #9: drag the log continuously near the viewport edge; the page scrolls
  automatically. The fixed upper desktop navigation keeps common links visible
  while Settings scrolls independently. Browser-saved height and reset remain.
- Telemetry defaults to APRS-IS only, recommended for Internet-connected stations.
  Existing saved delivery choices are preserved. RF-only or IS-preferred with
  RF fallback requires an RF channel to save, including while disabled.
- All other selected feature and upstream PRs remain included.

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
- Theme-driven enhanced packet-flow colors/separators and browser-saved log height,
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
RF performance advantage over Dire Wolf is claimed. Theme and telemetry
upstream submissions remain on hold. PR #9 and FX.25 are
operator-accepted and prepared for separate upstream submissions.
