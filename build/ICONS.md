# Hermes icons: what's here and how to use them (Wails v3.1.0)

One design, every platform. The single source of truth is `appicon.png` (1024x1024).
Editable vector sources are in `icon-src/` (`hermes-icon.svg`, `hermes-glyph.svg`).

| Platform | Files | Notes |
|---|---|---|
| All (source) | `appicon.png` | Rounded-square artwork with margin |
| Windows | `windows/icon.ico` | 16, 32, 48, 64, 128, 256 px; tighter margin so it reads at 16-48 px |
| macOS | `darwin/icons.icns` | 16 to 1024 px (incl. @2x) |
| Linux | `appicon.png`, `linux/hermes-{32..512}.png` | PNG sizes for `.desktop` / hicolor theme |
| Android | `android/app/src/main/res/mipmap-*/` | Legacy, round, adaptive foreground + background; monochrome layer for themed icons |
| Android store | `android/play-store-icon.png` | 512 px Play Console listing icon |
| iPhone + iPad | `ios/Assets.xcassets/AppIcon.appiconset/` | All required sizes + 1024 App Store icon + `Contents.json`. Opaque and square; iOS rounds the corners itself |
| Tray | `tray/tray.png`, `tray/tray-template.png` | Colour (Windows/Linux) and black template image (macOS) |

## 0. Install the files

Copy this whole `build/` folder into your Wails project's `build/` (merge, overwrite when asked).
Regenerate desktop formats any time you change `appicon.png`:

```bash
wails3 generate icons -input build/appicon.png
```

The template's desktop build runs this step too, so `appicon.png` wins over the pre-generated `.ico`/`.icns`.
(The `.ico` I ship has a slightly tighter margin than what the tool will regenerate; both are fine.)

## 1. Window / dock icon (all desktop OSes), in `main.go`

```go
//go:embed build/appicon.png
var appIcon []byte

app := application.New(application.Options{
    Name: "Hermes",
    Icon: appIcon,
    // ...
})
```

## 2. Windows

`build/windows/icon.ico` is embedded into the `.exe` and used by the installer by the template's Windows tasks
(search `build/windows/Taskfile.yml` for `icon.ico` to see where). Build with `wails3 build` / `wails3 package`.
If Explorer shows an old icon, rebuild and clear the icon cache (or rename the exe).

## 3. macOS

`build/darwin/icons.icns` is copied into the `.app` bundle by the darwin package task and referenced from
`build/darwin/Info.plist`. Build with `wails3 package`. Icons in the Dock during `wails3 dev` come from the
runtime icon in step 1 rather than the bundle.

## 4. Linux

`build/appicon.png` is what the AppImage and deb/rpm packaging tasks use. If your `.desktop` file says `Icon=hermes`,
install the PNGs from `build/linux/` as `hermes.png` in the matching `hicolor/<size>x<size>/apps/` folders
(nfpm config: `build/linux/nfpm/`). Package with `wails3 package`.

## 5. Android

Copy `android/app/src/main/res/` into your project's `build/android/app/src/main/res/` (merge).
Check `build/android/app/src/main/AndroidManifest.xml` contains:

```xml
android:icon="@mipmap/ic_launcher"
android:roundIcon="@mipmap/ic_launcher_round"
```

and delete any template launcher icons that clash (for example leftover `drawable` icons with the same name).
Build with `wails3 task android:package`. For the Play Store listing upload `android/play-store-icon.png`.

## 6. iPhone / iPad

iOS support in Wails v3 is experimental, and it needs macOS with full Xcode.

1. Generate the iOS project once: `wails3 task ios:run`. This creates `build/ios/`.
2. Put the icons in the project's app-icon asset catalog. Look in `build/ios/` for `Assets.xcassets/AppIcon.appiconset`
   and replace its contents with the files from `ios/Assets.xcassets/AppIcon.appiconset/` (keep the set name `AppIcon`).
3. If your generated project has no asset catalog yet, run `wails3 task ios:xcode`, open `Assets.xcassets > AppIcon`
   in Xcode and drag in `icon-1024.png` (Xcode 14+ fills every size from one 1024 px image), or drag in the whole
   folder contents. Make sure the target's App Icon setting is `AppIcon`.
4. Rebuild (`wails3 task ios:run` for the simulator, `ios:package IOS_PLATFORM=device ...` for a device).
   Delete the old app from the simulator/device first, because iOS caches icons.
5. Bundle id and display name live in `build/config.yml` under `ios:`.

The 1024 px App Store icon is `icon-1024.png` in that folder (also `ios/appicon-ios-1024.png`). Apple requires it to
be opaque with no rounded corners, which these are.

## 7. Tray (optional)

Embed `tray/tray.png` and `tray/tray-template.png` and pass them to the system tray's icon and template-icon setters
(macOS uses the template image so it adapts to light/dark menu bars). Mobile has no tray.

## What I could and couldn't verify

- Verified: file formats and sizes (ICO and ICNS contents, iOS images opaque RGB, Android densities), and how the
  artwork looks at each size.
- Not verified: building in Xcode, Android Studio or Wails itself (not available here). Exact template file paths
  (Taskfiles, `build/ios/` layout) can differ between Wails releases; the Wails docs don't spell out where the iOS icon
  catalog lives, which is why step 6 says to look inside `build/ios/` first.
- Not included: iOS 18 dark and tinted icon variants and macOS 26 Icon Composer (`.icon`) files.
  `wails3 generate icons` has `-iconcomposerinput` for the latter if you want to add it later.
