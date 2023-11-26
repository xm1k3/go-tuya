# go-tuya
Golang Tuya integration

## Installation

To install your tool, you can use the go get command:

```
go install -v github.com/xm1k3/go-tuya@latest
```

Make sure you have Go installed on your machine.

## Usage

After installation, you can use your tool from the command line. Here are some example commands:

## Options

Your tool supports the following options:

- **-c, --clientid**: Client ID parameter (required).
- **-s, --secret**: Secret parameter (required).
- **-d, --deviceid**: Device ID parameter (required).

### Get a Token

```
go-tuya token -c {CLIENT-ID} -s {SECRET} 
```

This command will return a token after authenticating the client with the specified parameters.

Example output: 

```json
{
  "result": {
    "access_token": "96268*****************6275e",
    "expire_time": 5626,
    "refresh_token": "35fe6**********************79bbe",
    "uid": "bay169***********04VX4T"
  },
  "success": true,
  "t": 1700995257134,
  "tid": "..."
}
```

### Get a Device

```
go-tuya device -c {CLIENT-ID} -s {SECRET} -d {DEVICE-ID} 
```

This command will return device information after authenticating the client with the specified parameters.

Example output: 

```json
{
  "result": {
    "active_time": 1699111647,
    "biz_type": 0,
    "category": "ms",
    "create_time": 1699108124,
    "icon": "...",
    "id": "{ID}",
    "ip": "",
    "lat": "...",
    "local_key": "...",
    "lon": "...",
    "model": "...",
    "name": "{NAME}",
    "node_id": "{NODE-ID}",
    "online": false,
    "owner_id": "...",
    "product_id": "...",
    "product_name": "...",
    "status": [
      {
        "code": "unlock_method_create",
        "value": "..."
      },
      {
        "code": "unlock_offline_pd",
        "value": ""
      }
      ...
    ],
    "sub": true,
    "time_zone": "+02:00",
    "uid": "...",
    "update_time": 1700985606,
    "uuid": "..."
  },
  "success": true,
  "t": 1700995312944,
  "tid": "..."
}
```

# License

go-tuya is distributed under Apache-2.0 License