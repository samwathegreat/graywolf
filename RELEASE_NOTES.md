# Graywolf 0.14.14-c.10 for Linux ARM64 and x86_64

This fork testing release retains all twelve selected changes from combined.9
and adds the separately maintained TX playback-completion fix. It is published
as Latest for custom-repository installers; hardware acceptance remains pending.
Both application and modem use the shorter `0.14.14-c.10` version and the same
source commit. This is not an official upstream Graywolf release.

New in this build:

- Wait for CPAL's predicted final-sample playback time before releasing PTT,
  rather than treating callback-consumed samples as physically played.
- Preserve the configured preamble and tail; no fixed extra tail is imposed.
  Low-latency devices wait for their own reported playback completion.
- Bound failed completion waits and discard timed-out output streams.
- Regression coverage includes large output latency, stereo sample positions,
  idle callbacks, repeated transmissions and PTT release on timeout.
- Anonymized old-binary controls: FX.25 validation recovered with longer tail;
  ordinary AX.25 went from 0/14 matched RF receptions at 100 ms tail to 10/10
  at 1000 ms. These establish the symptom/workaround, not acceptance of this fix.
  See docs/tx-playback-completion.md for methods, limitations and full evidence.
- Validate this build in both AX.25 and FX.25 at 100 ms tail before folding
  the fix into the standalone FX.25 PR. No automatic RF test is performed.

The `c.10` naming change sorts below `combined.9` in Debian version comparison.
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

The smaller FX.25/PTT build passed operator RF tests. Signal themes and
scrolling/resizing have operator acceptance, and upstream PRs #685/#686/#687
are open. The timestamp layout, GPS freshness styling and wrapping beacon-card layout
still need visual confirmation at 100%, narrow widths and other General
display-size settings. Native telemetry
remains a fork draft. Physical acceptance of the updated combination remains
pending on each architecture. No general RF performance advantage over Dire Wolf
is claimed.
