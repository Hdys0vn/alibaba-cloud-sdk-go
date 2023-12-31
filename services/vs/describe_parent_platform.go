package vs

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

// DescribeParentPlatform invokes the vs.DescribeParentPlatform API synchronously
func (client *Client) DescribeParentPlatform(request *DescribeParentPlatformRequest) (response *DescribeParentPlatformResponse, err error) {
	response = CreateDescribeParentPlatformResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeParentPlatformWithChan invokes the vs.DescribeParentPlatform API asynchronously
func (client *Client) DescribeParentPlatformWithChan(request *DescribeParentPlatformRequest) (<-chan *DescribeParentPlatformResponse, <-chan error) {
	responseChan := make(chan *DescribeParentPlatformResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeParentPlatform(request)
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

// DescribeParentPlatformWithCallback invokes the vs.DescribeParentPlatform API asynchronously
func (client *Client) DescribeParentPlatformWithCallback(request *DescribeParentPlatformRequest, callback func(response *DescribeParentPlatformResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeParentPlatformResponse
		var err error
		defer close(result)
		response, err = client.DescribeParentPlatform(request)
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

// DescribeParentPlatformRequest is the request struct for api DescribeParentPlatform
type DescribeParentPlatformRequest struct {
	*requests.RpcRequest
	Id      string           `position:"Query" name:"Id"`
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeParentPlatformResponse is the response struct for api DescribeParentPlatform
type DescribeParentPlatformResponse struct {
	*responses.BaseResponse
	Status         string `json:"Status" xml:"Status"`
	ClientPort     int64  `json:"ClientPort" xml:"ClientPort"`
	ClientGbId     string `json:"ClientGbId" xml:"ClientGbId"`
	Protocol       string `json:"Protocol" xml:"Protocol"`
	Ip             string `json:"Ip" xml:"Ip"`
	Port           int64  `json:"Port" xml:"Port"`
	ClientPassword string `json:"ClientPassword" xml:"ClientPassword"`
	ClientUsername string `json:"ClientUsername" xml:"ClientUsername"`
	AutoStart      bool   `json:"AutoStart" xml:"AutoStart"`
	ClientAuth     bool   `json:"ClientAuth" xml:"ClientAuth"`
	GbId           string `json:"GbId" xml:"GbId"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Description    string `json:"Description" xml:"Description"`
	ClientIp       string `json:"ClientIp" xml:"ClientIp"`
	Name           string `json:"Name" xml:"Name"`
	CreatedTime    string `json:"CreatedTime" xml:"CreatedTime"`
	Id             string `json:"Id" xml:"Id"`
}

// CreateDescribeParentPlatformRequest creates a request to invoke DescribeParentPlatform API
func CreateDescribeParentPlatformRequest() (request *DescribeParentPlatformRequest) {
	request = &DescribeParentPlatformRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeParentPlatform", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeParentPlatformResponse creates a response to parse from DescribeParentPlatform response
func CreateDescribeParentPlatformResponse() (response *DescribeParentPlatformResponse) {
	response = &DescribeParentPlatformResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
