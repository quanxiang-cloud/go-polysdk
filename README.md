# polysdk-go

polysdk-go is a client for polyapi.

## Installation

To install go-polysdk:
```
go get -u github.com/quanxiang-cloud/go-polysdk
```


## Usage

1. create a poly client from the config file
```Go
	c, err := polysdk.NewPolyClient(&config.PolyConfig{
		RemoteURL:   "http://api.example.com/",
		Key:         config.PolyKeyConfig{AccessKeyID: "cxxx_xxxxxxxxxx", SecretKey: "xxxx-xxxxxxxxxc"},
		CreateAt:    "2020-01-01",
		Description: "test",
	})
	if err != nil {
		panic(err)
	}
	h := polysdk.Header{}
	h.Set(polysdk.HeaderContentType, "application/json")
	
	body := map[string]interface{}{
		"zone":       "pek3d",
	}
	polysdk.PrettyShow(body)

	uri := "/api/v1/polyapi/raw/request/system/app/jhdsk/customer/ns2/viewVM.r"
	r, err := c.DoRequestAPI(uri, polysdk.MethodPost, h, body)
	if err != nil {
		panic(err)
	}

	polysdk.PrettyShow(r)
```

~~ enjoy this sdk!