<a name="table-of-contents"></a>
## Table of contents

- [Get Started](#get-started)
- VAULT X API
	- [Register New User](#register-new-user)
	- [Pair Device](#pair-device)
	- [Setup PIN Code](#setup-pin-code)
	- [Repair Device](#repair-device)
	- [Issue User Login Request](#issue-user-login-request)
	- [Query User Status](#query-user-status)
	- [Backup User Key](#backup-user-key)
	- [Restore User Key](#restore-user-key)
	- [Sign Message](#sign-message)
	- [Sign Raw Transaction](#sign-raw-transaction)
	- [Decrypt Message](#decrypt-message)
	- [Query Callback Status](#query-callback-status)
- Testing
	- [Mock Server](#mock-server)
	- [CURL Testing Commands](#curl-testing-commands)
- Appendix
	- [Callback Definition](#callback-definition)
	- [API Callback List](#api-callback-list)
	- [Setup getnonce API](#setup-getnonce-api)

<a name="get-started"></a>
# Get Started

- Get API code and API secret on web admin console
- Setup callback URL and getnonce API URL
- Refer to [mock server](#mock-server) sample code 

# VAULTX API

<a name="register-new-user"></a>
## Register New User

Register a new user.

**`POST`** /v1/vaultx/users 

- [Sample curl command](#curl-register-new-user)

##### Request Format

An example of the request:

###### Post body

```json
{
  "name": "JOHN DOE",
  "email": "johndoe@example.com",
  "locale": "en"
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| name | string | User name |
| email | string | User email |
| locale | string | User preference locale (en, zh-TW, zh-CN, ja, ko) |

##### Response Format

An example of a successful response:


```json
{
  "email": "johndoe@example.com"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| email  | string | Registered user's email  |

#### Possible Errors

| HTTP Status Code | Error Code  | Error Message |
| :---             | :---        | :---          |
| 400  | 103 | Account already exists |

##### Refer to [Common Errors](#common-errors)

#####[Back to top](#table-of-contents)


<a name="pair-device"></a>
## Pair Device

Pair device with user. 
> This API has a callback.

**`POST`** /v1/vaultx/devices?email=`USER_EMAIL`&company_id=`COMPANY_ID`

- [Sample curl command](#curl-pair-device)

##### Request Format

An example of the reques:

###### API with query string

```
/v1/vaultx/devices?email=johndoe@example.com&company_id=1
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email      | string | Requester email      |
| company_id | number | Requester company ID |

##### Response Format

An example of a successful response:


```json
{
  "order_id": 40000000003,
  "url": "http://192.168.0.56:8080/v1/cybavob/otp/token?product=vaultx&token=oyZSzs6NX5Tx78_VLieH2rwCbCNIM4XsOJ-SvMqpCK4="
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | int64 | Callback ID. Use this ID to identify specific callback. |
| url   | string | URL to bind device with user. Convert to QR code then scan with CYBAVO Auth APP. |

#### Callback

An example of a callback:

```
{
  "behavior_result": 2,
  "behavior_type": 4,
  "company_id": 1,
  "order_id": 40000000003
 }
```
- [View callback definition](#callback-definition)

#### Possible Errors

| HTTP Status Code | Error Code  | Error Message |
| :---             | :---        | :---          |
| 400  | 177 | User already has paired device, please use re-pair api if you want to do re-pair action |

##### Refer to [Common Errors](#common-errors)

#####[Back to top](#table-of-contents)


<a name="setup-pin-code"></a>
## Setup PIN Code

Setup user PIN code.
> This API has a callback.
>
> The requester needs to set the PIN code on the CYBAVO Auth APP after issuing the API request.

**`POST`** /v1/vaultx/user/pin?email=`USER_EMAIL`&company_id=`COMPANY_ID`

- [Sample curl command](#curl-setup-pin-code)

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/users/pin?email=johndoe@example.com&company_id=1
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email | string | Requester email |
| company_id | number | Requester company ID |

##### Response Format

An example of a successful response:

```json
{
  "order_id": 50000000004
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | int64 | Callback ID. Use this ID to identify specific callback. |

#### Callback

An example of a callback:

```
{
  "behavior_result": 2,
  "behavior_type": 5,
  "company_id": 1,
  "order_id": 50000000004
}
```
- [View callback definition](#callback-definition)

#### Possible Errors

| HTTP Status Code | Error Code  | Error Message |
| :---             | :---        | :---          |
| 400  | 176 | User already has pin |

##### Refer to [Common Errors](#common-errors)

#####[Back to top](#table-of-contents)


<a name="repair-device"></a>
## Repair Device

Following steps below to complete device reparing process.
> This process has a callback.
 
- [Sample curl command](#curl-repair-device)

#### Step 1: Issue a device repairing request

**`POST`** /v1/vaultx/devices/repair?email=`USER_EMAIL`&company_id=`COMPANY_ID`

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/devices/repair?email=johndoe@example.com&company_id=1
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email      | string | Requester email |
| company_id | number | Requester company ID |

##### Response Format

An example of a successful response:

```json
{
  "token": "yMYoEfHenKLDAAMIdVLW5zGrn9E4ap0qE5VrGoXlQaE="
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| token | string | Submit this token in step 3 |

#### Step 2: Check requester's email to get 6-digit verification code

#### Step 3: Continue repairing process

**`POST`** /v1/vaultx/devices/repair?email=`USER_EMAIL`&company_id=`COMPANY_ID`

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/devices/repair?email=johndoe@example.com&company_id=1
```

###### Post body

```json
{
  "token": "yMYoEfHenKLDAAMIdVLW5zGrn9E4ap0qE5VrGoXlQaE=",
  "verify_num": 908520
}
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email      | string | Requester email |
| company_id | number | Requester company ID |

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| token | string | The token received in step 1 |
| verify_num | number | The 6-digit verification code received in requester's email |

##### Response Format

An example of a successful response:

```json
{
  "order_id": 40000000005,
  "url": "http://192.168.0.56:8080/v1/cybavob/otp/rebind?token=A3Rll4djJt5vphKN77wovhYs0XHSt1TKsV9996pwnlY="
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | int64 | Callback ID. Use this ID to identify specific callback. |
| url   | string | URL to bind device with user. Convert to QR code then scan with CYBAVO Auth APP. |

#### Callback

An example of a callback:

```
{
  "behavior_result": 2,
  "behavior_type": 4,
  "company_id": 1,
  "order_id": 40000000005
}
```
- [View callback definition](#callback-definition)

#### Possible Errors

| HTTP Status Code | Error Code  | Error Message |
| :---             | :---        | :---          |
| 400  | 173 | Pair device process not completed |

##### Refer to [Common Errors](#common-errors)

#####[Back to top](#table-of-contents)


<a name="issue-user-login-request"></a>
## Issue User Login Request

Issue a login request. Requester needs to confirm request on the CYBAVO Auth APP.
> This API has a callback.

**`POST`** /v1/vaultx/loginverify?email=`USER_EMAIL`&company_id=`COMPANY_ID`

- [Sample curl command](#curl-issue-user-login-request)

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/loginverify?email=johndoe@example.com&company_id=1
```

###### Post body

```json
{
  "ip": "192.168.0.101",
  "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36",
  "expires_at": 1572451200
}
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email      | string | Requester email |
| company_id | number | Requester company ID |

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| ip         | string | Requester IP |
| user_agent | string | User-Agent of the requester's browser |
| expires_at | number | Expiration time (in unix time, UTC) of this 2FA request |

##### Response Format

An example of a successful response:

```json
{
  "order_id": 10000000002,
  "expires_at": 1572451200
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | int64 | Callback ID. Use this ID to identify specific callback. |
| expires_at | int64 | The same value with request's `expires_at` field |

#### Possible Errors

##### Refer to [Common Errors](#common-errors)

#### Callback

An example of a callback:

```
{
  "behavior_type": 1,
  "behavior_result": 2,
  "company_id": 1,
  "order_id": 10000000002
}
```
- [View callback definition](#callback-definition)

#####[Back to top](#table-of-contents)


<a name="query-user-status"></a>
## Query User Status

Query specific user status.

**`GET`** /v1/vaultx/users/me?email=`USER_EMAIL`&company_id=`COMPANY_ID`

- [Sample curl command](#curl-query-user-status)

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/users/me?email=johndoe@example.com&company_id=1
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email      | string | Requester email |
| company_id | int64 | Requester company ID |

##### Response Format

An example of a successful response:

```json
{
  "user_email": "johndoe@example.com",
  "company_id": 1,
  "is_pair_device": true,
  "is_setup_pin": true,
  "is_do_backup": true,
  "wallets": [
    {
      "type": "quorum",
      "address": "0x69363AFef99DC6f52daA2C4D0731934d36f7844d",
      "public_key": "0x03055b24413ef30be430cf131114ab768f245803c98ef5bf3b3aaafde414f7670e"
    }
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| user_email | string | Requester's email |
| company_id | int64 | Requester's company ID |
| is\_pair_device | boolean | True if device paired, otherwise false |
| is\_setup_pin | boolean | True if PIN code setup, otherwise false |
| is\_do_backup | boolean | True if key has been backuped, otherwise false |
| wallets | array | Created wallets |

#### Possible Errors

##### Refer to [Common Errors](#common-errors)

#####[Back to top](#table-of-contents)


<a name="backup-user-key"></a>
## Backup User Key

Following steps below to complete key backup process.
> This process has a callback.

- [Sample curl command](#curl-backup-user-key)

#### Step 1: Issue a key backup request

**`POST`** /v1/vaultx/wallets/backup?email=`USER_EMAIL`&company_id=`COMPANY_ID`

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/wallets/backup?email=johndoe@example.com&company_id=1
```

###### Post body

```json
{
  "question": "YOUR QUESTION",
  "answer": "YOUR ANSWER"
}
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email | string | Requester email |
| company_id | int64 | Requester company ID |

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| question | string | The chanllenge question while restoring keys |
| answer | string | The answer of the chanllenge question |

##### Response Format

An example of a successful response:

```json
{
  "order_id": 60000000014,
  "token": "yMYoEfHenKLDAAMIdVLW5zGrn9E4ap0qE5VrGoXlQaE="
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | int64 | Callback ID. Use this ID to identify specific callback. |
| token | string | Submit this token in step 2 |

#### Step 2: Confirm request on CYBAVO Auth APP

Once request is replied the callback server will receive a corresponding callback. If request is accepted, continue to step 3 to download the key backup file.

> Specify the callback server on web admin console.

#### Callback

An example of a callback:

```
{
  "behavior_result": 2,
  "behavior_type": 6,
  "company_id": 1,
  "order_id": 60000000014
}
```

- [View callback definition](#callback-definition)

#### Step 3: Download key backup file

**`GET`** /v1/vaultx/wallets/backup/{TOKEN}?email=`USER_EMAIL`&company_id=`COMPANY_ID`

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/wallets/backup/yMYoEfHenKLDAAMIdVLW5zGrn9E4ap0qE5VrGoXlQaE=?email=johndoe@example.com&company_id=1
```

The request includes the following parameters:

###### API path param

| Field | Type  | Description |
| :---  | :---  | :---        |
| token | string | Token received in step 1 |

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email | string | Requester email |
| company_id | int64 | Requester company ID |

##### Response Format

Save response body as backup file.

#### Possible Errors

##### Refer to [Common Errors](#common-errors)

#####[Back to top](#table-of-contents)


<a name="restore-user-key"></a>
## Restore User Key

Following steps below to complete key restoring process.
> This process has a callback.

- [Sample curl command](#curl-restore-user-key)

#### Step 1: Upload key backup file to server

**`POST`** /v1/vaultx/wallets/upload?email=`USER_EMAIL`&company_id=`COMPANY_ID`

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/wallets/upload?email=johndoe@example.com&company_id=1
```

###### Post body

```
Content-Type:multipart/form-data; boundary=ZnGpDtePMx0KrHh_G0X99Yef9r8JZsRJSXC

--ZnGpDtePMx0KrHh_G0X99Yef9r8JZsRJSXC
Content-Disposition: form-data;name="backup"; filename="backup.dat"
Content-Type: application/octet-stream

.. binary data of the backup file ...

--ZnGpDtePMx0KrHh_G0X99Yef9r8JZsRJSXC
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email | string | Requester email |
| company_id | int64 | Requester company ID |

###### Post body

Upload backup file using **multipart/form-data** format with name **backup**.

##### Response Format

An example of a successful response:

```json
{
  "token": "yMYoEfHenKLDAAMIdVLW5zGrn9E4ap0qE5VrGoXlQaE=",
  "question": "YOUR QUESTION"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| token | string | Submit this token in step 2 |
| question | string | The chanllenge question submitted with `/wallets/backup` API |

#### Step 2: Submit the answer to restore keys

**`POST`** /v1/vaultx/wallets/restore?email=`USER_EMAIL`&company_id=`COMPANY_ID`

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/wallets/restore?email=johndoe@example.com&company_id=1
```

###### Post body

```json
{
  "token": "yMYoEfHenKLDAAMIdVLW5zGrn9E4ap0qE5VrGoXlQaE=",
  "question": "YOUR QUESTION",
  "answer": "YOUR ANSWER"
}
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email | string | Requester email |
| company_id | int64 | Requester company ID |

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| token | string | Token received in step 1 |
| question | string | The challenge question received in step 1 |
| answer | string | The answer of the chanllenge question |

##### Response Format

An example of a successful response:

```json
{
  "token": "Jd9_oDYP97U7o7DH-5lB8Vqbp4PkqxDCkGz8Df0r79I="
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| token | string |  |

#### Step 3: Confirm request on CYBAVO Auth APP

Accept restoring request to reset PIN code.

#### Possible Errors

##### Refer to [Common Errors](#common-errors)

#####[Back to top](#table-of-contents)


<a name="sign-message"></a>
## Sign Message

Sign message.
> This API has a callback.
> 
> After 2FA request confirmed on CYBAVO Auth APP, the corresponding callback will contain signed message.

**`POST`** /v1/vaultx/wallets/signature

- [Sample curl command](#curl-sign-message)

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/wallets/signature?email=johndoe@example.com&company_id=1
```

###### Post body

```json
{
  "message": "MESSAGE TO BE SIGNED"
}
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email | string | Requester email |
| company_id | int64 | Requester company ID |

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| message | string | Message to be signed |

##### Response Format

An example of a successful response:

```json
{
  "order_id": 30000000001
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | int64 | Callback ID. Use this ID to identify specific callback. |

#### Callback

An example of a callback:

```
{
  "behavior_result": 2,
  "behavior_type": 3,
  "company_id": 1,
  "input": "{\"message\":\"MESSAGE TO BE SIGNED\"}",
  "order_id": 30000000001,
  "output": "0xebd938d9c34531b2b847d70125cac4aed7a1177da0d0df0866c65bfc2157fa5b429bc59df60bb977fe451dcc460d6a9709901c25f724def7b491de95fb41837e01"
}
```
- [View callback definition](#callback-definition)

#### Possible Errors

##### Refer to [Common Errors](#common-errors)

#####[Back to top](#table-of-contents)


<a name="sign-raw-transaction"></a>
## Sign Raw Transaction

Sign raw transaction.

> This API has a callback.

> After 2FA request confirmed, server will call getnonce API to retrieve the nonce of the wallet address. Refer to [Setup getnonce API](#setup-getnonce-api) section.

**`POST`** /v1/vaultx/wallets/rawtx

- [Sample curl command](#curl-sign-raw-transaction)

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/wallets/rawtx?email=johndoe@example.com&company_id=1
```

###### Post body

```json
{
  "to": "0x9576e27257e0eceea565fce04ab1beedfc6f35e4",
  "gas_limit": 90000,
  "gas_price": 1000000000,
  "value": 1000000000000000000,
  "input": "0x5448495320495320412054455354494e4720535452494e47",
  "private": false
}
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email | string | Requester email |
| company_id | int64 | Requester company ID |

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| to | string | Receiver address. If <i>to</i> is empty, it is a contract deployment transaction. Otherwise, it is a normal transaction. |
| gas_limit | uint64 | Gas limitation |
| gas_price | int64 | Gas price in GWEI. 10<sup>9</sup> GWEI = 1 ETH |
| value | int64 | Amount to transfer in WEI. 10<sup>18</sup> WEI = 1 ETH |
| input | hex string | The input data of transaction |
| private | boolean | If private is true, use HomesteadSigner otherwise use NewEIP155Signer to sign transaction. |

##### Response Format

An example of a successful response:

```json
{
  "order_id": 20000000010
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | int64 | Callback ID. Use this ID to identify specific callback. |

#### Callback

An example of a callback:

```json
{
  "behavior_result": 2,
  "behavior_type": 2,
  "company_id": 1,
  "input": "{\"to\":\"0x9576e27257e0eceea565fce04ab1beedfc6f35e4\",\"gas_limit\":90000,\"gas_price\":1000000000,\"value\":1000000000000000000,\"input\":\"0x5448495320495320412054455354494e4720535452494e47\",\"private\":false}",
  "order_id": 20000000010,
  "output": "f88401843b9aca0083015f90949576e27257e0eceea565fce04ab1beedfc6f35e4880de0b6b3a7640000985448495320495320412054455354494e4720535452494e4729a0e7cb1b61aa52529adbdb0c45750f3ac8863882fd270fa387b15fc558a0e097cba0723d0f518fb60a345829c09479e37ab156d4a7c027276848a08ece5e75681f5c"
}
```

- [View callback definition](#callback-definition)

#### Possible Errors

##### Refer to [Common Errors](#common-errors)

#####[Back to top](#table-of-contents)


<a name="decrypt-message"></a>
## Decrypt Message

Decrypt the message which encrypted by requester's public key.

> This API has a callback.

> After 2FA request confirmed on CYBAVO Auth APP, the corresponding callback will contain decrypted message.

**`POST`** /v1/vaultx/wallets/decrypt

- [Sample curl command](#curl-decrypt-message)

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/wallets/decrypt?email=johndoe@example.com&company_id=1
```

###### Post body

```json
{
  "secret": "0x04c7a7db0953984c6b3127a65300a86732a1799f076137a2619d913b84bf3130ff5fb947ac5b01706d902cb4fee021148eb0299626fb7c69f5e763ba3657fcef51f2326178289ac4d63d424444abaac33ec3c66844e5cd19d6ce12dcb48abadb2cd0d6eba06d7a074ee4e8b2cdf6b2bea34a059c2fadedd23b6ba61e1a537e3fa47a0a7137fd8171b0ec894164bbb3418c53823475fcdca0fc41b2480cb28bb2a52bffaaddbd89143a5e5c63e46014e3a44f6c9176a6e1410aecd331ccaad062cb5e3cd9cf813e3a6fa1425b2d61fd904429c6a4b3bb5f5bca12364c46dc019f28766dc8058f116393cc57d556e80efa1bd784c05d8b6ddfe9a173d0161b5db7f50859cb908bca6e9abf05a8649dfd5997b70313f9fe48e882670ce4def568b8239910acf8ebae8089ee17d53c7222fa895d9ce98f7d43d4df84a3f4a3a41d9c55a85f395f8ee2bbf7c07b0aa5b664a38ef7c40a0ec5c3039833dafd3b75b3f83d6a58effc6b959a7e0af8a09a9b40d7ec54dbb2f86dfb2750a8df"
}
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email | string | Requester email |
| company_id | int64 | Requester company ID |

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| secret | hex string | The message encrypted by requester's public key |
> The public key could be retrieved by **/users/me** API

##### Response Format

An example of a successful response:

```json
{
  "order_id": 70000000001
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | int64 | Callback ID. Use this ID to identify specific callback. |

#### Callback

An example of a callback:

```json
{
  "behavior_result": 2,
  "behavior_type": 7,
  "company_id": 3,
  "input": "{\"secret\":\"0x04c7a7db0953984c6b3127a65300a86732a1799f076137a2619d913b84bf3130ff5fb947ac5b01706d902cb4fee021148eb0299626fb7c69f5e763ba3657fcef51f2326178289ac4d63d424444abaac33ec3c66844e5cd19d6ce12dcb48abadb2cd0d6eba06d7a074ee4e8b2cdf6b2bea34a059c2fadedd23b6ba61e1a537e3fa47a0a7137fd8171b0ec894164bbb3418c53823475fcdca0fc41b2480cb28bb2a52bffaaddbd89143a5e5c63e46014e3a44f6c9176a6e1410aecd331ccaad062cb5e3cd9cf813e3a6fa1425b2d61fd904429c6a4b3bb5f5bca12364c46dc019f28766dc8058f116393cc57d556e80efa1bd784c05d8b6ddfe9a173d0161b5db7f50859cb908bca6e9abf05a8649dfd5997b70313f9fe48e882670ce4def568b8239910acf8ebae8089ee17d53c7222fa895d9ce98f7d43d4df84a3f4a3a41d9c55a85f395f8ee2bbf7c07b0aa5b664a38ef7c40a0ec5c3039833dafd3b75b3f83d6a58effc6b959a7e0af8a09a9b40d7ec54dbb2f86dfb2750a8df\"}",
  "order_id": 70000000013,
  "output": "0x7b22636f6e74656e74223a2248656c6c6f20576f726c64222c2265787069726174696f6e223a22323031392d31312d30365431303a34313a30372e3133393837392b30383a3030222c2268617368223a22307862643662663637636638633130643833653330346339623833386263313064353565363930613439646332333730333063333862366539373231333237643833222c226d6573736167655f68617368223a22516d596674383674376934315663713379354476734455533333485a6a475a6d714b5a796577503145386732596a222c22737461747573223a302c2274696d65223a22323031392d31302d30375431303a34313a30372e3133393837392b30383a3030227d"
}
```

- [View callback definition](#callback-definition)

#### Possible Errors

##### Refer to [Common Errors](#common-errors)

#####[Back to top](#table-of-contents)


<a name="query-callback-status"></a>
## Query Callback Status

Query callbacks' status belong to requester's company.

**`POST`** /v1/vaultx/order/status

- [Sample curl command](#curl-query-callback-status)

##### Request Format

An example of the request:

###### API with query string

```
/v1/vaultx/order/status?email=johndoe@example.com&company_id=1
```

###### Post body

```json
{
  "order_ids": [
    10000000004,
    20000000004,
    30000000001
  ]
}
```

The request includes the following parameters:

###### Query string

| Field | Type  | Description |
| :---  | :---  | :---        |
| email | string | Requester email |
| company_id | int64 | Requester company ID |

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_ids | array | ID of callbacks |

##### Response Format

An example of a successful response:

```json
{
  "order_status": [
    {
      "is_exist": true,
      "order_id": 10000000004,
      "behavior_type": 1,
      "behavior_result": 2,
      "addon": {}
    },
    {
      "is_exist": true,
      "order_id": 20000000004,
      "behavior_type": 2,
      "behavior_result": 2,
      "addon": {
        "input": "{\"from\":\"0x81b7e08f65bdf5648606c89998a9cc8164397647\",\"to\":\"0x9576e27257e0eceea565fce04ab1beedfc6f35e4\",\"gas_limit\":90000,\"gas_price\":1000000000,\"value\":1000000000000000000,\"nonce\":0,\"data\":null,\"input\":\"0x414243\",\"private_from\":\"\",\"private_for\":null,\"private_tx_type\":\"\"}",
        "output": "f86582250f8083015f90949576e27257e0eceea565fce04ab1beedfc6f35e4808341424329a083c60d3f302c20390bb1fae0db919ef245542a903fcdeed95aecbcfc9ff844dea0613a4274ae4243ea2f2cc2dddcd64b8484b706f945c546b90eabbb0ac4cbe80f"
      }
    },
    {
      "is_exist": true,
      "order_id": 30000000001,
      "behavior_type": 3,
      "behavior_result": 2,
      "addon": {
        "input": "{\"message\":\"MESSAGE TO BE SIGNED\"}",
        "output": "0x06a2c198b7670bd47d1433bfe497ee0d5019fe2b78ba7b1fe6ae99c134deb5d23d045d5cf3bf314fa77cf8da060e141bf5e9d36c488000bab81e9f56c6a46f4f01"
      }
    }
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| is_exist | boolean | This order_id exists or not |
| order_id | int64 | ID of callback |
| behavior_type | int | [View callback definition](#callback-definition) |
| behavior_result | int | [View callback definition](#callback-definition) |
| addon | object | The input and output fields of callback |


<a name="mock-server"></a>
# Mock Server

## Setup Configuration
>	Set following configuration in mockserver.app.conf
>> Require API code and API secret on web admin console

```
api_server_url=""
api_code=""
api_secret=""
```

## Register mock server URL
>	Operate on web admin console

Callback URL

```
http://localhost:8890/v1/mock/callback
```

GetNonce URL

```
http://localhost:8890/v1/mock/getnonce
```

## How to compile
- Put sample code to {YOUR\_GO\_PATH}/github.com/cybavo/VAULTX\_MOCK\_SERVER
- Execute
	- glide install
	- go build ./mockserver.go
	- ./mockserver


<a name="curl-testing-commands"></a>
## CURL Testing Commands

<a name="curl-register-new-user"></a>
#### Register a new user
```
curl -X POST -d '{"name":"JOHN DOE","email":"johndoe@example.com","locale":"en"}' \
"http://localhost:8890/v1/mock/users"
```
- [API definition](#register-new-user)

<a name="curl-pair-device"></a>
#### Pair device
```
curl -X POST "http://localhost:8890/v1/mock/devices?email=johndoe@example.com&company_id=1"
```
- [API definition](#pair-device)

<a name="curl-setup-pin-code"></a>
#### Setup PIN code
```
curl -X POST "http://localhost:8890/v1/mock/users/pin?email=johndoe@example.com&company_id=1"
```
- [API definition](#setup-pin-code)

<a name="curl-repair-device"></a>
#### Repair device

###### Step 1: Issue a repair device request
```
curl -X POST "http://localhost:8890/v1/mock/devices/repair?email=johndoe@example.com&company_id=1"
```
###### step 2: Check email to get 6-digit verification code
###### Step 3: Continue repairing process
```
curl -X POST -d '{"token":"TOKEN_FROM_SETP1","verify_num":VERIFICATION_CODE_FROM_EMAIL}' \
"http://localhost:8890/v1/mock/devices/repair?email=johndoe@example.com&company_id=1"
```
- [API definition](#repair-device)

<a name="curl-issue-user-login-request"></a>
#### Issue a user login request
```
curl -X POST -d '{"ip":"192.168.0.1","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36","expires_at":1572451200}' \
"http://localhost:8890/v1/mock/loginverify?email=johndoe@example.com&company_id=1"
```
- [API definition](#issue-user-login-request)

<a name="curl-query-user-status"></a>
#### Query user status
```
curl -X GET "http://localhost:8890/v1/mock/users/me?email=johndoe@example.com&company_id=1"
```
- [API definition](#query-user-status)

<a name="curl-backup-user-key"></a>
#### Backup user key

###### Step 1: Issue a key backup request
```
curl -X POST -d '{"question":"YOUR QUESTION","answer":"YOUR ANSWER"}' \
"http://localhost:8890/v1/mock/wallets/backup?email=johndoe@example.com&company_id=1"
```
###### Step 2: Confirm request on CYBAVO Auth APP
###### Step 3: Download key backup file
```
curl -X GET "http://localhost:8890/v1/mock/wallets/backup/{TOKEN}?email=johndoe@example.com&company_id=1"
```
- [API definition](#backup-user-key)

<a name="curl-restore-user-key"></a>
#### Restore user key

###### Step 1: Upload key backup file to server
```
curl -X POST "http://localhost:8890/v1/mock/wallets/upload?email=johndoe@example.com&company_id=1"
```
###### Step 2: Submit the answer to restore keys
```
curl -X POST -d '{"token":"TOKEN FROM SERVER","question":"QUESTION FROM SERVER","answer":"YOUR ANSWER"}' \
"http://localhost:8890/v1/mock/wallets/restore?email=johndoe@example.com&company_id=1"
```
###### Step 3: Confirm request on CYBAVO Auth APP
- [API definition](#restore-user-key)

<a name="curl-sign-message"></a>
#### Sign Message
```
curl -X POST -d '{"message":"MESSAGE TO BE SIGNED"}' \
"http://localhost:8890/v1/mock/wallets/signature?email=johndoe@example.com&company_id=1"
```
- [API definition](#sign-message)

<a name="curl-sign-raw-transaction"></a>
#### Sign Transaction
```
curl -X POST -d '{"to":"0x9576e27257e0eceea565fce04ab1beedfc6f35e4","gas_limit":90000,"gas_price":1000000000,"value":1000000000000000000,"input":"0x5448495320495320412054455354494e4720535452494e47","private":false}' \
"http://localhost:8890/v1/mock/wallets/rawtx?email=johndoe@example.com&company_id=1"
```
- [API definition](#sign-raw-transaction)

<a name="curl-decrypt-message"></a>
#### Decrypt Message
```
curl -X POST -d '{"secret":"0x04c7a7db0953984c6b3127a65300a86732a1799f076137a2619d913b84bf3130ff5fb947ac5b01706d902cb4fee021148eb0299626fb7c69f5e763ba3657fcef51f2326178289ac4d63d424444abaac33ec3c66844e5cd19d6ce12dcb48abadb2cd0d6eba06d7a074ee4e8b2cdf6b2bea34a059c2fadedd23b6ba61e1a537e3fa47a0a7137fd8171b0ec894164bbb3418c53823475fcdca0fc41b2480cb28bb2a52bffaaddbd89143a5e5c63e46014e3a44f6c9176a6e1410aecd331ccaad062cb5e3cd9cf813e3a6fa1425b2d61fd904429c6a4b3bb5f5bca12364c46dc019f28766dc8058f116393cc57d556e80efa1bd784c05d8b6ddfe9a173d0161b5db7f50859cb908bca6e9abf05a8649dfd5997b70313f9fe48e882670ce4def568b8239910acf8ebae8089ee17d53c7222fa895d9ce98f7d43d4df84a3f4a3a41d9c55a85f395f8ee2bbf7c07b0aa5b664a38ef7c40a0ec5c3039833dafd3b75b3f83d6a58effc6b959a7e0af8a09a9b40d7ec54dbb2f86dfb2750a8df"}' \
"http://localhost:8890/v1/mock/wallets/decrypt?email=johndoe@example.com&company_id=1"
```
- [API definition](#decrypt-message)


<a name="curl-query-callback-status"></a>
#### Query Callback Status
```
curl -X POST -d '{"order_ids":[10000000002,10000000003]}' \
"http://localhost:8890/v1/mock/order/status?email=johndoe@example.com&company_id=1"
```
- [API definition](#query-callback-status)

#####[Back to top](#table-of-contents)


# Appendix

<a name="callback-definition"></a>
## Callback Definition

<table>
  <tr>
    <td>Field</td>
    <td>Type</td>
    <td>Description</td>
  </tr>
  <tr>
    <td>order_id</td>
    <td> int64 </td>
    <td>ID of callback</td>
  </tr>
  <tr>
    <td>company_id</td>
    <td>int64</td>
    <td>ID of requester's company</td>
  </tr>
  <tr>
    <td>behavior_type</td>
    <td>number</td>
    <td rowspan="7">
      <b>1</b> - Login<br>
      <b>2</b> - Sign raw tx<br>
      <b>3</b> - Sign signature<br>
      <b>4</b> - Pair device<br>
      <b>5</b> - Setup PIN code<br>
      <b>6</b> - Backup user key<br>
      <b>7</b> - Decrypt message<br>
    </td>
  </tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr>
    <td>behavior_result</td>
    <td>number</td>
    <td rowspan="4">
      <b>0</b> - Pending<br>
      <b>1</b> - Rejected<br>
      <b>2</b> - Accepted<br>
      <b>3</b> - Expired<br>
      <b>4</b> - Failed
    </td>
  </tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr>
    <td>input</td>
    <td>json string</td>
    <td rowspan=3>
      behavior_type<br>
      <b>2</b> - the raw transaction to sign<br>
      <b>3</b> - the message to sign<br>
      <b>7</b> - the message to decrypt (hex string)<br>
    </td>
  </tr>
  <tr></tr>
  <tr></tr>
  <tr>
    <td>output</td>
    <td>hex string</td>
    <td rowspan=4>
      behavior_type<br>
      <b>2</b> - the signed transaction<br>
      <b>3</b> - the signed message<br>
      <b>7</b> - the decrypted message (hex string)<br>
    </td>
  </tr>
</table>

#####[Back to top](#table-of-contents)


<a name="api-callback-list"></a>
## API Callback List

| API  | Callback Type | Input/Output |
| :--- | :---          | :---         |
| /loginverify | 1 | - |
| /wallets/rawtx | 2 | yes |
| /wallets/signature | 3 | yes |
| /devices | 4 | - |
| /devices/repair | 4 | - |
| /users/pin | 5 | - |
| /wallets/backup | 6 | - |
| /wallets/decrypt | 7 | yes |

#####[Back to top](#table-of-contents)


<a name="common-errors"></a>
## Common Errors

| HTTP Status Code | Error Code  | Error Message |
| :---             | :---        | :---          |
| 400  | 112 | Invalid parameter |
| 400  | - | JSON unmarshal failure reason |
| 403  | - | Forbidden |
| 403  | 703 | Operation failed |
| 503  | 903 | KMS out of serivce. Please try again later. |

#####[Back to top](#table-of-contents)


<a name="setup-getnonce-api"></a>
## Setup getnonce API

During the raw transation signing process, the CYBAVO server will invoke getnonce API to retrieve the nonce of the wallet address.
> Set the getnonce API URL on web admin console.

##### Request Format

An example of the request:

###### Post body

```json
{
	"address": "0x94b3d188620bf633f6bff5f8c718b4bfbaca0c55",
	"company_id": 1
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| address | string | The address of the requester's wallet |
| company_id | int64 | Requester company ID |

##### Response Format

An example of a successful response:

```json
{
  "nonce": 1
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| nonce | int64 | The nonce of given address |

> Refer to [CallbackController.go](controllers/CallbackController.go)
