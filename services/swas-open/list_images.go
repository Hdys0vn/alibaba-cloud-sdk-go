package swas_open

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

// ListImages invokes the swas_open.ListImages API synchronously
func (client *Client) ListImages(request *ListImagesRequest) (response *ListImagesResponse, err error) {
	response = CreateListImagesResponse()
	err = client.DoAction(request, response)
	return
}

// ListImagesWithChan invokes the swas_open.ListImages API asynchronously
func (client *Client) ListImagesWithChan(request *ListImagesRequest) (<-chan *ListImagesResponse, <-chan error) {
	responseChan := make(chan *ListImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListImages(request)
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

// ListImagesWithCallback invokes the swas_open.ListImages API asynchronously
func (client *Client) ListImagesWithCallback(request *ListImagesRequest, callback func(response *ListImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListImagesResponse
		var err error
		defer close(result)
		response, err = client.ListImages(request)
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

// ListImagesRequest is the request struct for api ListImages
type ListImagesRequest struct {
	*requests.RpcRequest
	ImageType string `position:"Query" name:"ImageType"`
	ImageIds  string `position:"Query" name:"ImageIds"`
}

// ListImagesResponse is the response struct for api ListImages
type ListImagesResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Images    []Image `json:"Images" xml:"Images"`
}

// CreateListImagesRequest creates a request to invoke ListImages API
func CreateListImagesRequest() (request *ListImagesRequest) {
	request = &ListImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "ListImages", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListImagesResponse creates a response to parse from ListImages response
func CreateListImagesResponse() (response *ListImagesResponse) {
	response = &ListImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}