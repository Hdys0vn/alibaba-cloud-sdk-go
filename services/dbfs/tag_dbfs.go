package dbfs

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

// TagDbfs invokes the dbfs.TagDbfs API synchronously
func (client *Client) TagDbfs(request *TagDbfsRequest) (response *TagDbfsResponse, err error) {
	response = CreateTagDbfsResponse()
	err = client.DoAction(request, response)
	return
}

// TagDbfsWithChan invokes the dbfs.TagDbfs API asynchronously
func (client *Client) TagDbfsWithChan(request *TagDbfsRequest) (<-chan *TagDbfsResponse, <-chan error) {
	responseChan := make(chan *TagDbfsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TagDbfs(request)
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

// TagDbfsWithCallback invokes the dbfs.TagDbfs API asynchronously
func (client *Client) TagDbfsWithCallback(request *TagDbfsRequest, callback func(response *TagDbfsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TagDbfsResponse
		var err error
		defer close(result)
		response, err = client.TagDbfs(request)
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

// TagDbfsRequest is the request struct for api TagDbfs
type TagDbfsRequest struct {
	*requests.RpcRequest
	Tags   string `position:"Query" name:"Tags"`
	DbfsId string `position:"Query" name:"DbfsId"`
}

// TagDbfsResponse is the response struct for api TagDbfs
type TagDbfsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateTagDbfsRequest creates a request to invoke TagDbfs API
func CreateTagDbfsRequest() (request *TagDbfsRequest) {
	request = &TagDbfsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "TagDbfs", "dbfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTagDbfsResponse creates a response to parse from TagDbfs response
func CreateTagDbfsResponse() (response *TagDbfsResponse) {
	response = &TagDbfsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
