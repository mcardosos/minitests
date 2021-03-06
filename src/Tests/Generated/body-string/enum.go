package stringgroup

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
)

// EnumClient is the test Infrastructure for AutoRest Swagger BAT
type EnumClient struct {
    ManagementClient
}
// NewEnumClient creates an instance of the EnumClient client.
func NewEnumClient() EnumClient {
  return EnumClient{New()}
}

// GetNotExpandable get enum value 'red color' from enumeration of 'red
// color', 'green-color', 'blue_color'.
func (client EnumClient) GetNotExpandable() (result String, err error) {
    req, err := client.GetNotExpandablePreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetNotExpandable", nil , "Failure preparing request")
    }

    resp, err := client.GetNotExpandableSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetNotExpandable", resp, "Failure sending request")
    }

    result, err = client.GetNotExpandableResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "GetNotExpandable", resp, "Failure responding to request")
    }

    return
}

// GetNotExpandablePreparer prepares the GetNotExpandable request.
func (client EnumClient) GetNotExpandablePreparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/enum/notExpandable"))
    return preparer.Prepare(&http.Request{})
}

// GetNotExpandableSender sends the GetNotExpandable request. The method will close the
// http.Response Body if it receives an error.
func (client EnumClient) GetNotExpandableSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetNotExpandableResponder handles the response to the GetNotExpandable request. The method always
// closes the http.Response Body.
func (client EnumClient) GetNotExpandableResponder(resp *http.Response) (result String, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// PutNotExpandable sends value 'red color' from enumeration of 'red color',
// 'green-color', 'blue_color'
//
// stringBody is possible values include: 'red color', 'green-color',
// 'blue_color'
func (client EnumClient) PutNotExpandable(stringBody Colors) (result autorest.Response, err error) {
    req, err := client.PutNotExpandablePreparer(stringBody)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutNotExpandable", nil , "Failure preparing request")
    }

    resp, err := client.PutNotExpandableSender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutNotExpandable", resp, "Failure sending request")
    }

    result, err = client.PutNotExpandableResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.EnumClient", "PutNotExpandable", resp, "Failure responding to request")
    }

    return
}

// PutNotExpandablePreparer prepares the PutNotExpandable request.
func (client EnumClient) PutNotExpandablePreparer(stringBody Colors) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/enum/notExpandable"),
                        autorest.WithJSON(stringBody))
    return preparer.Prepare(&http.Request{})
}

// PutNotExpandableSender sends the PutNotExpandable request. The method will close the
// http.Response Body if it receives an error.
func (client EnumClient) PutNotExpandableSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// PutNotExpandableResponder handles the response to the PutNotExpandable request. The method always
// closes the http.Response Body.
func (client EnumClient) PutNotExpandableResponder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

