package httpInfrastructuregroup

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
)

// HTTPFailureClient is the test Infrastructure for AutoRest
type HTTPFailureClient struct {
    ManagementClient
}
// NewHTTPFailureClient creates an instance of the HTTPFailureClient client.
func NewHTTPFailureClient() HTTPFailureClient {
  return HTTPFailureClient{New()}
}

// GetEmptyError get empty error form server
func (client HTTPFailureClient) GetEmptyError() (result Bool, err error) {
    req, err := client.GetEmptyErrorPreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpInfrastructuregroup.HTTPFailureClient", "GetEmptyError", nil , "Failure preparing request")
    }

    resp, err := client.GetEmptyErrorSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "httpInfrastructuregroup.HTTPFailureClient", "GetEmptyError", resp, "Failure sending request")
    }

    result, err = client.GetEmptyErrorResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpInfrastructuregroup.HTTPFailureClient", "GetEmptyError", resp, "Failure responding to request")
    }

    return
}

// GetEmptyErrorPreparer prepares the GetEmptyError request.
func (client HTTPFailureClient) GetEmptyErrorPreparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/failure/emptybody/error"))
    return preparer.Prepare(&http.Request{})
}

// GetEmptyErrorSender sends the GetEmptyError request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPFailureClient) GetEmptyErrorSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetEmptyErrorResponder handles the response to the GetEmptyError request. The method always
// closes the http.Response Body.
func (client HTTPFailureClient) GetEmptyErrorResponder(resp *http.Response) (result Bool, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// GetNoModelError get empty error form server
func (client HTTPFailureClient) GetNoModelError() (result Bool, err error) {
    req, err := client.GetNoModelErrorPreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "httpInfrastructuregroup.HTTPFailureClient", "GetNoModelError", nil , "Failure preparing request")
    }

    resp, err := client.GetNoModelErrorSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "httpInfrastructuregroup.HTTPFailureClient", "GetNoModelError", resp, "Failure sending request")
    }

    result, err = client.GetNoModelErrorResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "httpInfrastructuregroup.HTTPFailureClient", "GetNoModelError", resp, "Failure responding to request")
    }

    return
}

// GetNoModelErrorPreparer prepares the GetNoModelError request.
func (client HTTPFailureClient) GetNoModelErrorPreparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/http/failure/nomodel/error"))
    return preparer.Prepare(&http.Request{})
}

// GetNoModelErrorSender sends the GetNoModelError request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPFailureClient) GetNoModelErrorSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetNoModelErrorResponder handles the response to the GetNoModelError request. The method always
// closes the http.Response Body.
func (client HTTPFailureClient) GetNoModelErrorResponder(resp *http.Response) (result Bool, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}
