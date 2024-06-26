package polysdk

import (
	"github.com/quanxiang-cloud/go-polysdk/internal/apipath"
)

// PolyAPIRequest request poly api from polyapi
// fullNamespace is the full namespace of poly api. eg: /system/poly/sample_poly_api
func (c *PolyClient) PolyAPIRequest(fullNamespace string, method string, header Header, data interface{}) (*HTTPResponse, error) {
	uri := apipath.Join(apipath.APIPolyRequest, fullNamespace)
	return c.DoRequestAPI(uri, method, header, data)
}

// PolyAPIDoc request poly api from polyapi
// fullNamespace is the full namespace of poly api. eg: /system/poly/sample_poly_api
func (c *PolyClient) PolyAPIDoc(fullNamespace string, docType string, titleFirst bool) (*HTTPResponse, error) {
	d := apiDocReq{
		BodyBase:   BodyBase{},
		DocType:    docType,
		TitleFirst: titleFirst,
	}
	header := Header{
		HeaderContentType: []string{MIMEJSON},
	}
	uri := apipath.Join(apipath.APIPolyDoc, fullNamespace)
	return c.DoRequestAPI(uri, MethodPost, header, d)
}
