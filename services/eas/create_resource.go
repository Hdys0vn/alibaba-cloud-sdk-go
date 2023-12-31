package eas

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

// CreateResource invokes the eas.CreateResource API synchronously
func (client *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
	response = CreateCreateResourceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateResourceWithChan invokes the eas.CreateResource API asynchronously
func (client *Client) CreateResourceWithChan(request *CreateResourceRequest) (<-chan *CreateResourceResponse, <-chan error) {
	responseChan := make(chan *CreateResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateResource(request)
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

// CreateResourceWithCallback invokes the eas.CreateResource API asynchronously
func (client *Client) CreateResourceWithCallback(request *CreateResourceRequest, callback func(response *CreateResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateResourceResponse
		var err error
		defer close(result)
		response, err = client.CreateResource(request)
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

// CreateResourceRequest is the request struct for api CreateResource
type CreateResourceRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateResourceResponse is the response struct for api CreateResource
type CreateResourceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ResourceID   string `json:"ResourceID" xml:"ResourceID"`
	ResourceName string `json:"ResourceName" xml:"ResourceName"`
	ClusterId    string `json:"ClusterId" xml:"ClusterId"`
	OwnerUid     string `json:"OwnerUid" xml:"OwnerUid"`
}

// CreateCreateResourceRequest creates a request to invoke CreateResource API
func CreateCreateResourceRequest() (request *CreateResourceRequest) {
	request = &CreateResourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2021-07-01", "CreateResource", "/api/v2/resources", "eas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateResourceResponse creates a response to parse from CreateResource response
func CreateCreateResourceResponse() (response *CreateResourceResponse) {
	response = &CreateResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
