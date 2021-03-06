package formdatagroup

import (
	"io"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

// FormdataClient is the test Infrastructure for AutoRest Swagger BAT
type FormdataClient struct {
	ManagementClient
}

// NewFormdataClient creates an instance of the FormdataClient client.
func NewFormdataClient() FormdataClient {
	return FormdataClient{New()}
}

// UploadFile upload file
// fileContent is file to upload. fileContent will be closed upon successful
// return. Callers should ensure closure when receiving an error.fileName is
// file name to upload. Name has to be spelled exactly as written here.
func (client FormdataClient) UploadFile(fileContent io.ReadCloser, fileName string) (result ReadCloser, err error) {
	req, err := client.UploadFilePreparer(fileContent, fileName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFile", nil, "Failure preparing request")
	}

	resp, err := client.UploadFileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFile", resp, "Failure sending request")
	}

	result, err = client.UploadFileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFile", resp, "Failure responding to request")
	}

	return
}

// UploadFilePreparer prepares the UploadFile request.
func (client FormdataClient) UploadFilePreparer(fileContent io.ReadCloser, fileName string) (*http.Request, error) {
	formDataParameters := map[string]interface{}{
		"fileContent": fileContent,
		"fileName":    fileName,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/formdata/stream/uploadfile"),
		autorest.WithMultiPartFormData(formDataParameters))
	return preparer.Prepare(&http.Request{})
}

// UploadFileSender sends the UploadFile request. The method will close the
// http.Response Body if it receives an error.
func (client FormdataClient) UploadFileSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UploadFileResponder handles the response to the UploadFile request. The method always
// closes the http.Response Body.
func (client FormdataClient) UploadFileResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// UploadFileViaBody upload file
// fileContent is file to upload. fileContent will be closed upon successful
// return. Callers should ensure closure when receiving an error.
func (client FormdataClient) UploadFileViaBody(fileContent io.ReadCloser) (result ReadCloser, err error) {
	req, err := client.UploadFileViaBodyPreparer(fileContent)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFileViaBody", nil, "Failure preparing request")
	}

	resp, err := client.UploadFileViaBodySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFileViaBody", resp, "Failure sending request")
	}

	result, err = client.UploadFileViaBodyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formdatagroup.FormdataClient", "UploadFileViaBody", resp, "Failure responding to request")
	}

	return
}

// UploadFileViaBodyPreparer prepares the UploadFileViaBody request.
func (client FormdataClient) UploadFileViaBodyPreparer(fileContent io.ReadCloser) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/formdata/stream/uploadfile"),
		autorest.WithFile(fileContent))
	return preparer.Prepare(&http.Request{})
}

// UploadFileViaBodySender sends the UploadFileViaBody request. The method will close the
// http.Response Body if it receives an error.
func (client FormdataClient) UploadFileViaBodySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UploadFileViaBodyResponder handles the response to the UploadFileViaBody request. The method always
// closes the http.Response Body.
func (client FormdataClient) UploadFileViaBodyResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}
