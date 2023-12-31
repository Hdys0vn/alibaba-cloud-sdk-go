package green

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

// AddFaces invokes the green.AddFaces API synchronously
func (client *Client) AddFaces(request *AddFacesRequest) (response *AddFacesResponse, err error) {
	response = CreateAddFacesResponse()
	err = client.DoAction(request, response)
	return
}

// AddFacesWithChan invokes the green.AddFaces API asynchronously
func (client *Client) AddFacesWithChan(request *AddFacesRequest) (<-chan *AddFacesResponse, <-chan error) {
	responseChan := make(chan *AddFacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddFaces(request)
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

// AddFacesWithCallback invokes the green.AddFaces API asynchronously
func (client *Client) AddFacesWithCallback(request *AddFacesRequest, callback func(response *AddFacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddFacesResponse
		var err error
		defer close(result)
		response, err = client.AddFaces(request)
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

// AddFacesRequest is the request struct for api AddFaces
type AddFacesRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// AddFacesResponse is the response struct for api AddFaces
type AddFacesResponse struct {
	*responses.BaseResponse
}

// CreateAddFacesRequest creates a request to invoke AddFaces API
func CreateAddFacesRequest() (request *AddFacesRequest) {
	request = &AddFacesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "AddFaces", "/green/sface/face/add", "", "")
	request.Method = requests.POST
	return
}

// CreateAddFacesResponse creates a response to parse from AddFaces response
func CreateAddFacesResponse() (response *AddFacesResponse) {
	response = &AddFacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
