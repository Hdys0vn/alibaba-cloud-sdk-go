package edas

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

// ListApplication invokes the edas.ListApplication API synchronously
func (client *Client) ListApplication(request *ListApplicationRequest) (response *ListApplicationResponse, err error) {
	response = CreateListApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// ListApplicationWithChan invokes the edas.ListApplication API asynchronously
func (client *Client) ListApplicationWithChan(request *ListApplicationRequest) (<-chan *ListApplicationResponse, <-chan error) {
	responseChan := make(chan *ListApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListApplication(request)
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

// ListApplicationWithCallback invokes the edas.ListApplication API asynchronously
func (client *Client) ListApplicationWithCallback(request *ListApplicationRequest, callback func(response *ListApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListApplicationResponse
		var err error
		defer close(result)
		response, err = client.ListApplication(request)
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

// ListApplicationRequest is the request struct for api ListApplication
type ListApplicationRequest struct {
	*requests.RoaRequest
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	AppIds                string           `position:"Query" name:"AppIds"`
	AppName               string           `position:"Query" name:"AppName"`
	LogicalRegionId       string           `position:"Query" name:"LogicalRegionId"`
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage           requests.Integer `position:"Query" name:"CurrentPage"`
	ClusterId             string           `position:"Query" name:"ClusterId"`
	LogicalRegionIdFilter string           `position:"Query" name:"LogicalRegionIdFilter"`
}

// ListApplicationResponse is the response struct for api ListApplication
type ListApplicationResponse struct {
	*responses.BaseResponse
	Code            int                              `json:"Code" xml:"Code"`
	Message         string                           `json:"Message" xml:"Message"`
	RequestId       string                           `json:"RequestId" xml:"RequestId"`
	ApplicationList ApplicationListInListApplication `json:"ApplicationList" xml:"ApplicationList"`
}

// CreateListApplicationRequest creates a request to invoke ListApplication API
func CreateListApplicationRequest() (request *ListApplicationRequest) {
	request = &ListApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListApplication", "/pop/v5/app/app_list", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListApplicationResponse creates a response to parse from ListApplication response
func CreateListApplicationResponse() (response *ListApplicationResponse) {
	response = &ListApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
