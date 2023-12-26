package ivision

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

// DeleteFaceGroup invokes the ivision.DeleteFaceGroup API synchronously
func (client *Client) DeleteFaceGroup(request *DeleteFaceGroupRequest) (response *DeleteFaceGroupResponse, err error) {
	response = CreateDeleteFaceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFaceGroupWithChan invokes the ivision.DeleteFaceGroup API asynchronously
func (client *Client) DeleteFaceGroupWithChan(request *DeleteFaceGroupRequest) (<-chan *DeleteFaceGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteFaceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaceGroup(request)
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

// DeleteFaceGroupWithCallback invokes the ivision.DeleteFaceGroup API asynchronously
func (client *Client) DeleteFaceGroupWithCallback(request *DeleteFaceGroupRequest, callback func(response *DeleteFaceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFaceGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaceGroup(request)
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

// DeleteFaceGroupRequest is the request struct for api DeleteFaceGroup
type DeleteFaceGroupRequest struct {
	*requests.RpcRequest
	ShowLog string           `position:"Query" name:"ShowLog"`
	GroupId string           `position:"Query" name:"GroupId"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteFaceGroupResponse is the response struct for api DeleteFaceGroup
type DeleteFaceGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	GroupId   string `json:"GroupId" xml:"GroupId"`
}

// CreateDeleteFaceGroupRequest creates a request to invoke DeleteFaceGroup API
func CreateDeleteFaceGroupRequest() (request *DeleteFaceGroupRequest) {
	request = &DeleteFaceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "DeleteFaceGroup", "ivision", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDeleteFaceGroupResponse creates a response to parse from DeleteFaceGroup response
func CreateDeleteFaceGroupResponse() (response *DeleteFaceGroupResponse) {
	response = &DeleteFaceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
