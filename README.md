# MokoGo

<div align="center">
  <img src="./static/image/logo.jpg" width="30%" alt="Logo Moko"><br>
  <a href="#"><img alt="MokoGo" src="https://img.shields.io/badge/MokoGo-blue?colorA=%23ff0000&colorB=%23017e40&style=for-the-badge"></a><br>
  <a href="https://github.com/Abuzzpoet/MokoGo/stargazers"><img alt="Stars" src="https://img.shields.io/github/stars/Abuzzpoet/MokoGo?style=flat-square"></a>
  <a href="https://github.com/Abuzzpoet/MokoGo/network/members"><img alt="Forks" src="https://img.shields.io/github/forks/Abuzzpoet/MokoGo?style=flat-square"></a>
  <a href="https://github.com/Abuzzpoet/MokoGo/watchers"><img alt="Watchers" src="https://img.shields.io/github/watchers/Abuzzpoet/MokoGo?style=flat-square"></a>
</div>

<br><br>
> <p>MokoGo adalah Bot WhatsApp yang dibuat menggunakan Golang dan package <a href="https://github.com/tulir/whatsmeow" target="_blank">whatsmeow</a>.</p>

> <p>Project Ini Support Linux/Windows</p>

___

## Cara Penggunaan Dan Penginstalan

1. **Langkah 1:** Unduh atau clone repositori ini.
3. **Langkah 2:** Instal Golang [disini](https://go.dev/doc/install)  >= 1.20
4. **Langkah 3:** Ubah file `.env` dengan informasi yang diperlukan (seperti jid owner bot dan nama bot).
5. **Langkah 4:** Jalankan bot dengan perintah:
```shell
# masuk ke direktori bot
cd MokoGo
# install dependencies
go get all 
# atau
go get ./src
# jalankan bot
go run src/moko.go
# atau
go build src/moko.go
./mao #for linux
mao.exe #for windows
```
7. **Langkah 5:** Buka WhatsApp dan scan QR yang muncul di log.

## Kontribusi

Jika Anda ingin berkontribusi pada pengembangan bot ini, berikut adalah langkah-langkah yang dapat Anda lakukan:
- Fork repositori ini.
- Buat branch baru: `git checkout -b fitur-baru`.
- Lakukan perubahan yang diperlukan.
- Commit perubahan Anda: `git commit -m 'Menambahkan fitur baru'`.
- Push ke branch yang Anda buat: `git push origin fitur-baru`.
- Buat pull request.

## Thanks To
- [Abuzzpoet](https://github.com/Abuzzpoet/MokoGo)
- [Fckvania](https://github.com/fckvania/MaoGo)

## Lisensi

[GPL-3.0 license](/LICENSE.txt)
