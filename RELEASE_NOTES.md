# Graywolf 0.14.14-c.13 for Linux ARM64 and x86_64

Retains the complete c.12 combined feature set and corrects the channel column
for locally generated APRS-IS telemetry. IS reports and definitions display
**APRS-IS**, rather than the configured RF delivery channel. RF telemetry retains
its actual RF channel name. Historical IS telemetry entries also display correctly.
New IS entries have no RF channel association; measurement scope remains the
entire station. No modem processing, sample rates, audio levels, PTT timing or
station settings are changed by this update.

Both application and modem use version `0.14.14-c.13` and the same source commit.
The existing workflow builds ARM64 and AMD64 DEB/tarball artifacts and publishes
the fork release as Latest only after both architectures pass. This is not an
official upstream Graywolf release. Installed channel-column acceptance is pending.

# Graywolf 0.14.14-c.12 for Linux ARM64 and x86_64

This fork testing release retains every selected c.11 feature, including the TX
playback-completion fix and FX.25 channel configuration badge, while updating
telemetry traffic presentation and FX.25 Auto explanatory wording.
Both application and modem use version `0.14.14-c.12` and the same source commit.
It is published as Latest for custom-repository installers after both Linux
architectures pass the existing release checks. This is not an official
upstream Graywolf release.

New in c.12:

- Identify locally generated telemetry data and definitions as TX IS or TX RF
  from their event context. APRS-IS submissions use the standard route annotation;
  received telemetry is still labeled as reception. Failed uploads do not create
  successful-send entries. Valid uploads previously could appear Unspecified in
  the enhanced packet-flow view.
- Synchronize generated API schema, handbook copies and TypeScript client with
  the telemetry source annotation.
- Clarify FX.25 Auto wording: parity selection is based on protected frame length,
  not measured link quality. No encoding behavior changes.
- Preserve the c.11 blue configured FX.25 badge and wrapping channel headers,
  with the same shared badge size/spacing and disabled-channel semantics.

Local combined-source preparation passed 585 web tests, API/schema drift checks,
API ID lint, TypeScript checks, the production web build and focused Go race
regressions. All thirteen manifest heads are present.

The updated standalone telemetry branch passed hosted CI, Android and armhf-t64
validation at c393fdf. The release workflow independently checks the combined
source on both Linux architectures before publication. Installed-build telemetry
label acceptance remains pending; no automatic test RF is generated.

Retained from c.10:

- Wait for CPAL's predicted final-sample playback time before releasing PTT.
  Preserve the configured preamble and tail without fixed extra padding.
- Bound failed completion waits and discard timed-out output streams.
- Preserve the same original twelve feature selections plus the independently
  submitted playback-completion fix (upstream PR #689).
- Operator AX.25 and FX.25 field acceptance at 100 ms tail is documented with
  generic station identifiers in docs/tx-playback-completion.md on the fix PR.
  This badge update does not change that timing implementation.

The `c.N` naming convention sorts below `combined.N` in Debian version comparison.
Conventional Debian package tools may require an explicit downgrade allowance
when moving from the old naming convention; verify the intended release and
source commit. Existing reviewed paired-binary deployment remains available.

Retained from combined.9:

- Beacon cards show the symbol and full label on their own top row, wrapping
  long labels instead of truncating them beside status badges.
- Callsign/object identity and wrapping status badges appear beneath the label.
- The beacon grid fits narrow screens. All themes use the improved layout.
- The beacon handbook documents the new presentation.

Retained from combined.8:

- Consistent tracker GPS freshness checkbox, label and hint styling, including
  narrow-screen wrapping, with unchanged behavior and defaults.
- 1200-baud FX.25 airtime documentation with timing assumptions and limitations.
- The experimental 300-baud FX.25 branch and draft PR remain excluded.
  Native FX.25 in this build remains limited to 1200 baud.

Retained from combined.7:

- Complete Signal packet timestamps at 100% and other General display sizes,
  with wrapping between the date and clock on narrow cards.
- Correct APRS-IS server-filter guidance and generic documentation examples.

Retained audio-meter layout:

- Native DCD/RX/TX sit beside the channel name. Numeric dBFS precedes an outlined
  RX Audio meter with -60/-40/-20/0 scale marks and a responsive 480 px width cap.
- KISS-only channels retain actual TNC state/activity without fabricated audio.

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

- PR #686: drag the log continuously near the viewport edge; the page scrolls
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
  parity and repaired-symbol information. Auto selects parity from frame length and remains recommended when enabled;
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

The smaller FX.25/PTT build passed operator RF tests. Signal themes and
scrolling/resizing have operator acceptance, and upstream PRs #685/#686/#687
are open. The timestamp layout, GPS freshness styling and wrapping beacon-card layout
still need visual confirmation at 100%, narrow widths and other General
display-size settings. Native telemetry
remains a fork draft. Physical acceptance of the updated combination remains
pending on each architecture. No general RF performance advantage over Dire Wolf
is claimed.
