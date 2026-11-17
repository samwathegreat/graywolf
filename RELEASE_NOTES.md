# Combined testing 0.14.14-combined.3 for Linux ARM64 and x86_64

This is a testing release from samwathegreat/graywolf, not an official
upstream Graywolf release. It retains the selected changes from combined.2,
updates our feature heads, and adds native T# telemetry and log resizing.
This is a prerelease: installers using releases/latest stay on the previous
release. Select this version explicitly for acceptance testing.

New in this build:

- Native, optional station-wide T# activity telemetry with its own settings
  page, ten-minute baseline after RF input readiness, spaced definitions,
  delivery selection and recent collection/submission status. Disabled by default.
- Colored packets have clear gaps and subtle borders between complete records
  on desktop/mobile. Colorization remains optional and off by default.
- Drag the Dashboard packet log downward/upward, save the height in this browser,
  resize with the keyboard or reset to 400 px. Minimum 200 px; content-aware
  maximum up to 1600 px. The feed still fetches the latest 20 packets.
- Updated operator documentation across the selected feature branches. FX.25
  is explicitly limited to 1200-baud AFSK.
- Retain selected upstream #657 (RF path markers) and #617 (comment-command UI).

Retained features:

- Native FX.25 transmit and receive at 1200 baud, with per-packet encoding,
  parity and repaired-symbol information. Auto is recommended when enabled;
  oversized frames use explicitly identified ordinary AX.25 fallback.
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
