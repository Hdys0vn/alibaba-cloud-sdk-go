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

// DeleteSimilarityLibrary invokes the green.DeleteSimilarityLibrary API synchronously
func (client *Client) DeleteSimilarityLibrary(request *DeleteSimilarityLibraryRequest) (response *DeleteSimilarityLibraryResponse, err error) {
	response = CreateDeleteSimilarityLibraryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSimilarityLibraryWithChan invokes the green.DeleteSimilarityLibrary API asynchronously
func (client *Client) DeleteSimilarityLibraryWithChan(request *DeleteSimilarityLibraryRequest) (<-chan *DeleteSimilarityLibraryResponse, <-chan error) {
	responseChan := make(chan *DeleteSimilarityLibraryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSimilarityLibrary(request)
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

// DeleteSimilarityLibraryWithCallback invokes the green.DeleteSimilarityLibrary API asynchronously
func (client *Client) DeleteSimilarityLibraryWithCallback(request *DeleteSimilarityLibraryRequest, callback func(response *DeleteSimilarityLibraryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSimilarityLibraryResponse
		var err error
		defer close(result)
		response, err = client.DeleteSimilarityLibrary(request)
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

// DeleteSimilarityLibraryRequest is the request struct for api DeleteSimilarityLibrary
type DeleteSimilarityLibraryRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// DeleteSimilarityLibraryResponse is the response struct for api DeleteSimilarityLibrary
type DeleteSimilarityLibraryResponse struct {
	*responses.BaseResponse
}

// CreateDeleteSimilarityLibraryRequest creates a request to invoke DeleteSimilarityLibrary API
func CreateDeleteSimilarityLibraryRequest() (request *DeleteSimilarityLibraryRequest) {
	request = &DeleteSimilarityLibraryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "DeleteSimilarityLibrary", "/green/similarity/library/delete", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteSimilarityLibraryResponse creates a response to parse from DeleteSimilarityLibrary response
func CreateDeleteSimilarityLibraryResponse() (response *DeleteSimilarityLibraryResponse) {
	response = &DeleteSimilarityLibraryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
