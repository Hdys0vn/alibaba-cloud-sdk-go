package dataworks_public

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

// DeleteMetaCollection invokes the dataworks_public.DeleteMetaCollection API synchronously
func (client *Client) DeleteMetaCollection(request *DeleteMetaCollectionRequest) (response *DeleteMetaCollectionResponse, err error) {
	response = CreateDeleteMetaCollectionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMetaCollectionWithChan invokes the dataworks_public.DeleteMetaCollection API asynchronously
func (client *Client) DeleteMetaCollectionWithChan(request *DeleteMetaCollectionRequest) (<-chan *DeleteMetaCollectionResponse, <-chan error) {
	responseChan := make(chan *DeleteMetaCollectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMetaCollection(request)
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

// DeleteMetaCollectionWithCallback invokes the dataworks_public.DeleteMetaCollection API asynchronously
func (client *Client) DeleteMetaCollectionWithCallback(request *DeleteMetaCollectionRequest, callback func(response *DeleteMetaCollectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMetaCollectionResponse
		var err error
		defer close(result)
		response, err = client.DeleteMetaCollection(request)
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

// DeleteMetaCollectionRequest is the request struct for api DeleteMetaCollection
type DeleteMetaCollectionRequest struct {
	*requests.RpcRequest
	QualifiedName string `position:"Query" name:"QualifiedName"`
}

// DeleteMetaCollectionResponse is the response struct for api DeleteMetaCollection
type DeleteMetaCollectionResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Status         bool   `json:"Status" xml:"Status"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDeleteMetaCollectionRequest creates a request to invoke DeleteMetaCollection API
func CreateDeleteMetaCollectionRequest() (request *DeleteMetaCollectionRequest) {
	request = &DeleteMetaCollectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeleteMetaCollection", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMetaCollectionResponse creates a response to parse from DeleteMetaCollection response
func CreateDeleteMetaCollectionResponse() (response *DeleteMetaCollectionResponse) {
	response = &DeleteMetaCollectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}