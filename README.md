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
		RemoteURL:   "https://apis.yunify.com/",
		Key:         config.PolyKeyConfig{AccessKeyID: "cGrm_Tu31AGh4HzIcCtyUg", SecretKey: "6hmAMl-YmVQATjCO5KDrj6FYebL3TIw2qJe0rSo0_Sc"},
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