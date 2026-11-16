# Combined testing build for 64-bit Raspberry Pi

This is a testing release from samwathegreat/graywolf, not an official
upstream Graywolf release. It combines our six existing upstream PRs and
native 1200-baud FX.25 with upstream PRs #657 and #617.

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

The ARM64 archive and Debian package contain both Graywolf and its matching
native modem. The source archive records every included PR head in
`docs/combined-testing.md`; BUILD.txt records the exact built commit and tools.
Use checksums.txt to verify downloads. This package targets 64-bit Raspberry
Pi OS Bookworm or later (glibc 2.36+ and ALSA), including Pi 4.

The smaller FX.25/PTT build passed operator RF tests. Physical acceptance of
this combined release remains pending. No general RF performance advantage
over Dire Wolf is claimed. The upstream FX.25 PR remains on hold.
