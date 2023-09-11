package test_data

var TestPositiveData = []byte(`{
"order_uid": "1",
"track_number": "WBILMTESTTRACK",
"entry": "WBILl",
"delivery": {
"name": "Test Testov",
"phone": "+79211552101",
"zip": "2639809",
"city": "Kiryat Mozkin",
"address": "Ploshad Mira 15",
"region": "Kraiot",
"email": "test@gmail.com"
},
"payment": {
"transaction": "b563feb7b2b84b6test",
"request_id": "12345",
"currency": "USD",
"provider": "wbpay",
"amount": 1817,
"payment_dt": 1637907627,
"bank": "alpha",
"delivery_cost": 1500,
"goods_total": 317,
"custom_fee": 0
},
"items": [
{
"chrt_id": 1,
"track_number": "WBILMTESTTRACK",
"price": 453,
"rid": "ab4319087a764ae0btest",
"name": "Mascaras",
"sale": 30,
"size": "0",
"total_price": 317,
"nm_id": 2389232,
"brand": "Vivienne Sabo",
"status": 202
}
],
"locale": "en",
"internal_signature": "sad",
"customer_id": "test",
"delivery_service": "meest",
"shardkey": "9",
"sm_id": 99,
"date_created": "2021-11-26T06:22:19Z",
"oof_shard": "1"
}`)

// empty field track_number
var TestNegativeData1 = []byte(`{
"order_uid": "1",
"track_number": "",
"entry": "WBILl",
"delivery": {
"name": "Test Testov",
"phone": "+79211552101",
"zip": "2639809",
"city": "Kiryat Mozkin",
"address": "Ploshad Mira 15",
"region": "Kraiot",
"email": "test@gmail.com"
},
"payment": {
"transaction": "b563feb7b2b84b6test",
"request_id": "12345",
"currency": "USD",
"provider": "wbpay",
"amount": 1817,
"payment_dt": 1637907627,
"bank": "alpha",
"delivery_cost": 1500,
"goods_total": 317,
"custom_fee": 0
},
"items": [
{
"chrt_id": 1,
"track_number": "WBILMTESTTRACK",
"price": 453,
"rid": "ab4319087a764ae0btest",
"name": "Mascaras",
"sale": 30,
"size": "0",
"total_price": 317,
"nm_id": 2389232,
"brand": "Vivienne Sabo",
"status": 202
}
],
"locale": "en",
"internal_signature": "sad",
"customer_id": "test",
"delivery_service": "meest",
"shardkey": "9",
"sm_id": 99,
"date_created": "2021-11-26T06:22:19Z",
"oof_shard": "1"
}`)

// invalid field email
var TestNegativeData2 = []byte(`{
"order_uid": "1",
"track_number": "WBILMTESTTRACK",
"entry": "WBILl",
"delivery": {
"name": "Test Testov",
"phone": "+79211552101",
"zip": "2639809",
"city": "Kiryat Mozkin",
"address": "Ploshad Mira 15",
"region": "Kraiot",
"email": "test@gmail"
},
"payment": {
"transaction": "b563feb7b2b84b6test",
"request_id": "12345",
"currency": "USD",
"provider": "wbpay",
"amount": 1817,
"payment_dt": 1637907627,
"bank": "alpha",
"delivery_cost": 1500,
"goods_total": 317,
"custom_fee": 0
},
"items": [
{
"chrt_id": 1,
"track_number": "WBILMTESTTRACK",
"price": 453,
"rid": "ab4319087a764ae0btest",
"name": "Mascaras",
"sale": 30,
"size": "0",
"total_price": 317,
"nm_id": 2389232,
"brand": "Vivienne Sabo",
"status": 202
}
],
"locale": "en",
"internal_signature": "sad",
"customer_id": "test",
"delivery_service": "meest",
"shardkey": "9",
"sm_id": 99,
"date_created": "2021-11-26T06:22:19Z",
"oof_shard": "1"
}`)

// empty field items.price
var TestNegativeData3 = []byte(`{
"order_uid": "1",
"track_number": "WBILMTESTTRACK",
"entry": "WBILl",
"delivery": {
"name": "Test Testov",
"phone": "+79211552101",
"zip": "2639809",
"city": "Kiryat Mozkin",
"address": "Ploshad Mira 15",
"region": "Kraiot",
"email": "test@gmail.com"
},
"payment": {
"transaction": "b563feb7b2b84b6test",
"request_id": "12345",
"currency": "USD",
"provider": "wbpay",
"amount": 1817,
"payment_dt": 1637907627,
"bank": "alpha",
"delivery_cost": 1500,
"goods_total": 317,
"custom_fee": 0
},
"items": [
{
"chrt_id": 1,
"track_number": "WBILMTESTTRACK",
"price": "",
"rid": "ab4319087a764ae0btest",
"name": "Mascaras",
"sale": 30,
"size": "0",
"total_price": 317,
"nm_id": 2389232,
"brand": "Vivienne Sabo",
"status": 202
}
],
"locale": "en",
"internal_signature": "sad",
"customer_id": "test",
"delivery_service": "meest",
"shardkey": "9",
"sm_id": 99,
"date_created": "2021-11-26T06:22:19Z",
"oof_shard": "1"
}`)

// empty field payment.currency field
var TestNegativeData4 = []byte(`{
"order_uid": "1",
"track_number": "WBILMTESTTRACK",
"entry": "WBILl",
"delivery": {
"name": "Test Testov",
"phone": "+79211552101",
"zip": "2639809",
"city": "Kiryat Mozkin",
"address": "Ploshad Mira 15",
"region": "Kraiot",
"email": "test@gmail.com"
},
"payment": {
"transaction": "b563feb7b2b84b6test",
"request_id": "12345",
"currency": "USDS",
"provider": "wbpay",
"amount": 1817,
"payment_dt": 1637907627,
"bank": "alpha",
"delivery_cost": 1500,
"goods_total": 317,
"custom_fee": 0
},
"items": [
{
"chrt_id": 1,
"track_number": "WBILMTESTTRACK",
"price": 453,
"rid": "ab4319087a764ae0btest",
"name": "Mascaras",
"sale": 30,
"size": "0",
"total_price": 317,
"nm_id": 2389232,
"brand": "Vivienne Sabo",
"status": 202
}
],
"locale": "en",
"internal_signature": "sad",
"customer_id": "test",
"delivery_service": "meest",
"shardkey": "9",
"sm_id": 99,
"date_created": "2021-11-26T06:22:19Z",
"oof_shard": "1"
}`)
