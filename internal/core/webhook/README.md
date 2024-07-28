# FastFood API - Mercado Livre Webhook payment

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)


## Description

In the Fast Food application, the customer can pay the order in two way: **QR Code** or **Credit**.
If the customer choose *QR Code* payment type, the entire payment flow will be:

 - The customer needs to generate the QR Code via `POST /api/qrcode/generate`. This endpoint will generate an **Order** in the database.
 - With this QR Code, the customer needs to open `Mercado Pago` app and scan the QR Code to make the payment.
 - When the payment completes, the Webhook `POST /api/webhook/ml/payment` will receive the Mercado Livre data.
 - After the Webhook completes, the *Order* will finish the payment and will change the status to `Pago`.

The `POST /api/qrcode/generate` will return the data like this:

```
"data":"00020101021243650016COM.MERCADOLIBRE020130636d1377de2-422e-4068-9317-6fdd9f626b295204000053039865802BR5909Test Test6009SAO PAULO62070503***630416A3"
```

> [!NOTE]  
> After `POST /api/qrcode/generate` retrieves the **QR Code data**, to test the order payment, the developer needs to transform the *QR Code data* in a QR Code image. This site can be used to generate is [Generate QR Code image](https://br.qr-code-generator.com/)

> [!WARNING]
> Sometimes the **Mercado Livre** server returns `500 Internal Server Error` for unknown reason. The error returned by the server is: `{"error":"alias_obtainment_error","message":"Get aliases for user failed","status":500,"causes":[]}`. When this occurs, **IS NOT possible to proceed with QR Code Payment**. The main reason for this is on `Weekend the Mercado Livre development environment does not work`

After the payment has completed, the Webhook Endpoint (`POST /api/webhook/ml/payment`) will be called and it will receive the following data:

```
{
  "resource": "https://api.mercadolibre.com/merchant_orders/20203112410",
  "topic": "merchant_order"
}
```

With the `resource` node, we can call it with `GET` **method** and receives the following response:

```
{"id":19961356837,"status":"closed","external_reference":"123|1245","preference_id":"1865158750-e089c0ab-be88-4591-9d4f-43fe938a76c7","payments":[{"id":81030262220,"transaction_amount":150,"total_paid_amount":150,"shipping_cost":0,"currency_id":"BRL","status":"rejected","status_detail":"cc_rejected_other_reason","operation_type":"regular_payment","date_approved":"0001-01-01T00:00:00.000+00:00","date_created":"2024-06-20T20:07:56.000-04:00","last_modified":"2024-06-20T20:08:00.000-04:00","amount_refunded":0},{"id":81030312426,"transaction_amount":150,"total_paid_amount":150,"shipping_cost":0,"currency_id":"BRL","status":"approved","status_detail":"accredited","operation_type":"regular_payment","date_approved":"2024-06-20T20:09:10.000-04:00","date_created":"2024-06-20T20:09:10.000-04:00","last_modified":"2024-06-20T20:09:10.000-04:00","amount_refunded":0}],"shipments":[],"payouts":[],"collector":{"id":1865158750,"email":"","nickname":"TESTUSER97284132"},"marketplace":"NONE","notification_url":"https://webhook-test.com/983f20261b344f0aec95305f78e57bb8","date_created":"2024-06-20T20:07:06.421-04:00","last_updated":"2024-06-20T20:09:10.472-04:00","sponsor_id":null,"shipping_cost":0,"total_amount":150,"site_id":"MLB","paid_amount":150,"refunded_amount":0,"payer":{"id":1862647967,"email":""},"items":[{"id":"","category_id":"marketplace","currency_id":"BRL","description":"This is the Point Mini","picture_url":null,"title":"Point Mini","quantity":1,"unit_price":150}],"cancelled":false,"additional_info":"","application_id":null,"is_test":true,"order_status":"paid","client_id":"4523867654733557"}
```

By getting the `external_reference` we can **split** by slash (/) and with these 2 IDs we can reference the internal Order and Payment ID.
After setting the correct status for both we **finish** the entire QR Code payment process with `Mercado Livre`
