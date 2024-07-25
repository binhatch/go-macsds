# \WiringDiagramsImageAPIAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Image2WiringDiagrams**](WiringDiagramsImageAPIAPI.md#Image2WiringDiagrams) | **Get** /api/v1/image/wiring-diagrams | Get image wiring-diagrams



## Image2WiringDiagrams

> *os.File Image2WiringDiagrams(ctx).DiagramId(diagramId).Execute()

Get image wiring-diagrams



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	diagramId := "diagramId_example" // string | Diagram Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiringDiagramsImageAPIAPI.Image2WiringDiagrams(context.Background()).DiagramId(diagramId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiringDiagramsImageAPIAPI.Image2WiringDiagrams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Image2WiringDiagrams`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `WiringDiagramsImageAPIAPI.Image2WiringDiagrams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImage2WiringDiagramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diagramId** | **string** | Diagram Id | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
