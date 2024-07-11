# Tech Challenge 1 - Mercado Livre Webhook payment

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)


## Description

In the Fast Food application, the customer can pay the order in two way: **QR Code** or **Credit**.
If the customer choose *QR Code* payment type, the entire payment flow will be:

 - The customer needs to generate the QR Code via `POST /api/qrcode/generate`. This endpoint will generate an **Order** in the database.
 - With this QR Code, the customer needs to open `Mercado Pago` app and scan the QR Code to make the payment.
 - When the payment completes, the Webhook `POST /api/webhook/ml/payment` will receive the Mercado Livre data.
 - After the Webhook completes, the *Order* will be finihsed.

The `POST /api/qrcode/generate` will return the data like this:

```
"data":"00020101021243650016COM.MERCADOLIBRE020130636d1377de2-422e-4068-9317-6fdd9f626b295204000053039865802BR5909Test Test6009SAO PAULO62070503***630416A3"
```

> [!NOTE]  
> After `POST /api/qrcode/generate` retrieve the **QR Code data**, to test the order payment, the developer needs to transform the *QR Code data* in a QR Code image. This site can be used to generate is [Generate QR Code image](https://br.qr-code-generator.com/)
