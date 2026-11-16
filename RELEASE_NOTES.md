# Combined testing 0.14.14-combined.2 for Linux ARM64 and x86_64

This is a testing release from samwathegreat/graywolf, not an official
upstream Graywolf release. It retains the previous nine selected changes and
adds optional packet colorization, with updated serial PTT and badge wrapping.

New in this build:

- Packet badges wrap so labels remain visible when several apply.
- Linux serial PTT can use stable /dev/serial/by-id/ device names.
- Optional theme-aware packet highlights distinguish reception, transmission,
  gating and digipeating in packet logs and the live monitor. Colorize packets
  is after Theme in General preferences and is OFF by default. With it off,
  the existing colors and plain monitor text remain. Badge wrapping is always on.
- Matching app/modem packages are available for both Linux ARM64 and x86_64.

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
colorization PRs remain on hold.
