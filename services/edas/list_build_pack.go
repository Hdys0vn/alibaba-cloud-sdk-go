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

// ListBuildPack invokes the edas.ListBuildPack API synchronously
func (client *Client) ListBuildPack(request *ListBuildPackRequest) (response *ListBuildPackResponse, err error) {
	response = CreateListBuildPackResponse()
	err = client.DoAction(request, response)
	return
}

// ListBuildPackWithChan invokes the edas.ListBuildPack API asynchronously
func (client *Client) ListBuildPackWithChan(request *ListBuildPackRequest) (<-chan *ListBuildPackResponse, <-chan error) {
	responseChan := make(chan *ListBuildPackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBuildPack(request)
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

// ListBuildPackWithCallback invokes the edas.ListBuildPack API asynchronously
func (client *Client) ListBuildPackWithCallback(request *ListBuildPackRequest, callback func(response *ListBuildPackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBuildPackResponse
		var err error
		defer close(result)
		response, err = client.ListBuildPack(request)
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

// ListBuildPackRequest is the request struct for api ListBuildPack
type ListBuildPackRequest struct {
	*requests.RoaRequest
}

// ListBuildPackResponse is the response struct for api ListBuildPack
type ListBuildPackResponse struct {
	*responses.BaseResponse
	Code          int           `json:"Code" xml:"Code"`
	Message       string        `json:"Message" xml:"Message"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	BuildPackList BuildPackList `json:"BuildPackList" xml:"BuildPackList"`
}

// CreateListBuildPackRequest creates a request to invoke ListBuildPack API
func CreateListBuildPackRequest() (request *ListBuildPackRequest) {
	request = &ListBuildPackRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListBuildPack", "/pop/v5/app/build_pack_list", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListBuildPackResponse creates a response to parse from ListBuildPack response
func CreateListBuildPackResponse() (response *ListBuildPackResponse) {
	response = &ListBuildPackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
