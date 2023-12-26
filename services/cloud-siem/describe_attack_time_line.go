package cloud_siem

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

// DescribeAttackTimeLine invokes the cloud_siem.DescribeAttackTimeLine API synchronously
func (client *Client) DescribeAttackTimeLine(request *DescribeAttackTimeLineRequest) (response *DescribeAttackTimeLineResponse, err error) {
	response = CreateDescribeAttackTimeLineResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAttackTimeLineWithChan invokes the cloud_siem.DescribeAttackTimeLine API asynchronously
func (client *Client) DescribeAttackTimeLineWithChan(request *DescribeAttackTimeLineRequest) (<-chan *DescribeAttackTimeLineResponse, <-chan error) {
	responseChan := make(chan *DescribeAttackTimeLineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAttackTimeLine(request)
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

// DescribeAttackTimeLineWithCallback invokes the cloud_siem.DescribeAttackTimeLine API asynchronously
func (client *Client) DescribeAttackTimeLineWithCallback(request *DescribeAttackTimeLineRequest, callback func(response *DescribeAttackTimeLineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAttackTimeLineResponse
		var err error
		defer close(result)
		response, err = client.DescribeAttackTimeLine(request)
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

// DescribeAttackTimeLineRequest is the request struct for api DescribeAttackTimeLine
type DescribeAttackTimeLineRequest struct {
	*requests.RpcRequest
	AssetName    string           `position:"Body" name:"AssetName"`
	StartTime    requests.Integer `position:"Body" name:"StartTime"`
	EndTime      requests.Integer `position:"Body" name:"EndTime"`
	IncidentUuid string           `position:"Body" name:"IncidentUuid"`
}

// DescribeAttackTimeLineResponse is the response struct for api DescribeAttackTimeLine
type DescribeAttackTimeLineResponse struct {
	*responses.BaseResponse
	Success   bool       `json:"Success" xml:"Success"`
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeAttackTimeLineRequest creates a request to invoke DescribeAttackTimeLine API
func CreateDescribeAttackTimeLineRequest() (request *DescribeAttackTimeLineRequest) {
	request = &DescribeAttackTimeLineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeAttackTimeLine", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAttackTimeLineResponse creates a response to parse from DescribeAttackTimeLine response
func CreateDescribeAttackTimeLineResponse() (response *DescribeAttackTimeLineResponse) {
	response = &DescribeAttackTimeLineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
