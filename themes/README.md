# Avro CLI Themes

Теми для Avro CLI з підтримкою кольорових схем у простому текстовому форматі.

## Формат файлу теми

```ini
# Theme Name
name = Tokyo Night
author = Your Name

[base]
normal = #a9b1d6
focus = #c0caf5
hot_normal = #bb9af7
hot_focus = #bb9af7
disabled = #565f89
background = #1a1b26

[menu]
normal = #c0caf5
focus = #1a1b26
hot_normal = #7dcfff
hot_focus = #7dcfff
disabled = #565f89
background = #24283b

[dialog]
normal = #c0caf5
focus = #1a1b26
hot_normal = #7dcfff
hot_focus = #7dcfff
disabled = #565f89
background = #24283b

[error]
normal = #ffffff
focus = #e0af68
hot_normal = #e0af68
hot_focus = #e0af68
disabled = #808080
background = #f7768e

[toplevel]
normal = #a9b1d6
focus = #c0caf5
hot_normal = #bb9af7
hot_focus = #bb9af7
disabled = #565f89
background = #1a1b26
```

## Секції

- `[base]` — основна колірна схема для вмісту вікон
- `[menu]` — кольори меню (верхня панель)
- `[dialog]` — кольори діалогових вікон
- `[error]` — кольори помилок і попереджень
- `[toplevel]` — кольори кореневого вікна

## Кольори

Кожна секція підтримує:
- `normal` — звичайний текст
- `focus` — текст у фокусі
- `hot_normal` — клавіші швидкого доступу (без фокусу)
- `hot_focus` — клавіші швидкого доступу (у фокусі)
- `disabled` — неактивний текст
- `background` — фон

Кольори вказуються у HEX форматі: `#RRGGBB`

## Наявні теми

- **default** — стандартна тема (вбудована)
- **tokyo-night** — Tokyo Night темна тема
- **dracula** — Dracula темна тема
- **nord** — Nord темна тема

## Використання

Теми завантажуються автоматично з папки `themes/`. Файли мають мати розширення `.theme`.

Перемикання теми (майбутня функціональність):
```
Tools → Preferences → Theme
```

## Створення власної теми

1. Скопіюйте існуючий `.theme` файл
2. Змініть `name` та `author`
3. Налаштуйте кольори у HEX форматі
4. Збережіть файл у папку `themes/`
5. Перезапустіть Avro CLI

## Примітка

Terminal.Gui v1 використовує 16-кольорову палітру. RGB кольори автоматично маплюються до найближчого кольору з палітри:
- Black, Red, Green, Blue
- Gray, BrightRed, BrightGreen, BrightBlue
- White, BrightYellow, BrightMagenta, BrightCyan
- DarkGray

Для підтримки True Color (24-bit RGB) потрібен Terminal.Gui v2.
