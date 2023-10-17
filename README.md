# Easy QR-Code API
Implement QR Codes into your app without any effort!

---

Use this endpoint to generate qr-code images. The data is provided by the ``data`` query parameter.

```https://qr.niclas.lol/api/generate?data=https%3A%2F%2Fniclas.lol```

To change the size of the QR Code you can use the size parameter. The default size is 256.

```https://qr.niclas.lol/api/generate?data=https%3A%2F%2Fniclas.lol?size=512```

You can also directly embed the QR code into HTML using the ``img`` tag.

```<img src="https://qr.niclas.lol/api/generate?data=https%3A%2F%2Fniclas.lol>```
