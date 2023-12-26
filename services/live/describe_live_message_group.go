package live

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

// DescribeLiveMessageGroup invokes the live.DescribeLiveMessageGroup API synchronously
func (client *Client) DescribeLiveMessageGroup(request *DescribeLiveMessageGroupRequest) (response *DescribeLiveMessageGroupResponse, err error) {
	response = CreateDescribeLiveMessageGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveMessageGroupWithChan invokes the live.DescribeLiveMessageGroup API asynchronously
func (client *Client) DescribeLiveMessageGroupWithChan(request *DescribeLiveMessageGroupRequest) (<-chan *DescribeLiveMessageGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveMessageGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveMessageGroup(request)
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

// DescribeLiveMessageGroupWithCallback invokes the live.DescribeLiveMessageGroup API asynchronously
func (client *Client) DescribeLiveMessageGroupWithCallback(request *DescribeLiveMessageGroupRequest, callback func(response *DescribeLiveMessageGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveMessageGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveMessageGroup(request)
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

// DescribeLiveMessageGroupRequest is the request struct for api DescribeLiveMessageGroup
type DescribeLiveMessageGroupRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	DataCenter string `position:"Query" name:"DataCenter"`
	AppId      string `position:"Query" name:"AppId"`
}

// DescribeLiveMessageGroupResponse is the response struct for api DescribeLiveMessageGroup
type DescribeLiveMessageGroupResponse struct {
	*responses.BaseResponse
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	GroupId          string                 `json:"GroupId" xml:"GroupId"`
	CreatorId        string                 `json:"CreatorId" xml:"CreatorId"`
	Createtime       int64                  `json:"Createtime" xml:"Createtime"`
	GroupName        string                 `json:"GroupName" xml:"GroupName"`
	GroupInfo        string                 `json:"GroupInfo" xml:"GroupInfo"`
	Delete           bool                   `json:"Delete" xml:"Delete"`
	TotalTimes       int64                  `json:"TotalTimes" xml:"TotalTimes"`
	OnlineUserCounts int64                  `json:"OnlineUserCounts" xml:"OnlineUserCounts"`
	MsgAmount        map[string]interface{} `json:"MsgAmount" xml:"MsgAmount"`
	Deletatime       int64                  `json:"Deletatime" xml:"Deletatime"`
	Deletor          string                 `json:"Deletor" xml:"Deletor"`
	AdminList        []string               `json:"AdminList" xml:"AdminList"`
}

// CreateDescribeLiveMessageGroupRequest creates a request to invoke DescribeLiveMessageGroup API
func CreateDescribeLiveMessageGroupRequest() (request *DescribeLiveMessageGroupRequest) {
	request = &DescribeLiveMessageGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveMessageGroup", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveMessageGroupResponse creates a response to parse from DescribeLiveMessageGroup response
func CreateDescribeLiveMessageGroupResponse() (response *DescribeLiveMessageGroupResponse) {
	response = &DescribeLiveMessageGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
