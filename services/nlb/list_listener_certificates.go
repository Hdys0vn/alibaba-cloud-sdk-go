package nlb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// ListListenerCertificates invokes the nlb.ListListenerCertificates API synchronously
func (client *Client) ListListenerCertificates(request *ListListenerCertificatesRequest) (response *ListListenerCertificatesResponse, err error) {
	response = CreateListListenerCertificatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListListenerCertificatesWithChan invokes the nlb.ListListenerCertificates API asynchronously
func (client *Client) ListListenerCertificatesWithChan(request *ListListenerCertificatesRequest) (<-chan *ListListenerCertificatesResponse, <-chan error) {
	responseChan := make(chan *ListListenerCertificatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListListenerCertificates(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListListenerCertificatesWithCallback invokes the nlb.ListListenerCertificates API asynchronously
func (client *Client) ListListenerCertificatesWithCallback(request *ListListenerCertificatesRequest, callback func(response *ListListenerCertificatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListListenerCertificatesResponse
		var err error
		defer close(result)
		response, err = client.ListListenerCertificates(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListListenerCertificatesRequest is the request struct for api ListListenerCertificates
type ListListenerCertificatesRequest struct {
	*requests.RpcRequest
	ListenerId string           `position:"Body" name:"ListenerId"`
	CertType   string           `position:"Body" name:"CertType"`
	NextToken  string           `position:"Body" name:"NextToken"`
	MaxResults requests.Integer `position:"Body" name:"MaxResults"`
}

// ListListenerCertificatesResponse is the response struct for api ListListenerCertificates
type ListListenerCertificatesResponse struct {
	*responses.BaseResponse
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	Success        bool          `json:"Success" xml:"Success"`
	Code           string        `json:"Code" xml:"Code"`
	Message        string        `json:"Message" xml:"Message"`
	HttpStatusCode int           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string        `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string        `json:"DynamicMessage" xml:"DynamicMessage"`
	TotalCount     int           `json:"TotalCount" xml:"TotalCount"`
	NextToken      string        `json:"NextToken" xml:"NextToken"`
	MaxResults     int           `json:"MaxResults" xml:"MaxResults"`
	CertificateIds []string      `json:"CertificateIds" xml:"CertificateIds"`
	Certificates   []Certificate `json:"Certificates" xml:"Certificates"`
}

// CreateListListenerCertificatesRequest creates a request to invoke ListListenerCertificates API
func CreateListListenerCertificatesRequest() (request *ListListenerCertificatesRequest) {
	request = &ListListenerCertificatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "ListListenerCertificates", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListListenerCertificatesResponse creates a response to parse from ListListenerCertificates response
func CreateListListenerCertificatesResponse() (response *ListListenerCertificatesResponse) {
	response = &ListListenerCertificatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
