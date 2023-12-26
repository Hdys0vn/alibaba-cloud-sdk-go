package cms

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

// CreateGroupMonitoringAgentProcess invokes the cms.CreateGroupMonitoringAgentProcess API synchronously
func (client *Client) CreateGroupMonitoringAgentProcess(request *CreateGroupMonitoringAgentProcessRequest) (response *CreateGroupMonitoringAgentProcessResponse, err error) {
	response = CreateCreateGroupMonitoringAgentProcessResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGroupMonitoringAgentProcessWithChan invokes the cms.CreateGroupMonitoringAgentProcess API asynchronously
func (client *Client) CreateGroupMonitoringAgentProcessWithChan(request *CreateGroupMonitoringAgentProcessRequest) (<-chan *CreateGroupMonitoringAgentProcessResponse, <-chan error) {
	responseChan := make(chan *CreateGroupMonitoringAgentProcessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGroupMonitoringAgentProcess(request)
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

// CreateGroupMonitoringAgentProcessWithCallback invokes the cms.CreateGroupMonitoringAgentProcess API asynchronously
func (client *Client) CreateGroupMonitoringAgentProcessWithCallback(request *CreateGroupMonitoringAgentProcessRequest, callback func(response *CreateGroupMonitoringAgentProcessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGroupMonitoringAgentProcessResponse
		var err error
		defer close(result)
		response, err = client.CreateGroupMonitoringAgentProcess(request)
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

// CreateGroupMonitoringAgentProcessRequest is the request struct for api CreateGroupMonitoringAgentProcess
type CreateGroupMonitoringAgentProcessRequest struct {
	*requests.RpcRequest
	AlertConfig                *[]CreateGroupMonitoringAgentProcessAlertConfig  `position:"Query" name:"AlertConfig"  type:"Repeated"`
	GroupId                    string                                           `position:"Query" name:"GroupId"`
	ProcessName                string                                           `position:"Query" name:"ProcessName"`
	MatchExpressFilterRelation string                                           `position:"Query" name:"MatchExpressFilterRelation"`
	MatchExpress               *[]CreateGroupMonitoringAgentProcessMatchExpress `position:"Query" name:"MatchExpress"  type:"Repeated"`
}

// CreateGroupMonitoringAgentProcessAlertConfig is a repeated param struct in CreateGroupMonitoringAgentProcessRequest
type CreateGroupMonitoringAgentProcessAlertConfig struct {
	Times               string                                                        `name:"Times"`
	Webhook             string                                                        `name:"Webhook"`
	NoEffectiveInterval string                                                        `name:"NoEffectiveInterval"`
	TargetList          *[]CreateGroupMonitoringAgentProcessAlertConfigTargetListItem `name:"TargetList" type:"Repeated"`
	SilenceTime         string                                                        `name:"SilenceTime"`
	Threshold           string                                                        `name:"Threshold"`
	ComparisonOperator  string                                                        `name:"ComparisonOperator"`
	EffectiveInterval   string                                                        `name:"EffectiveInterval"`
	EscalationsLevel    string                                                        `name:"EscalationsLevel"`
	Statistics          string                                                        `name:"Statistics"`
}

// CreateGroupMonitoringAgentProcessAlertConfigTargetListItem is a repeated param struct in CreateGroupMonitoringAgentProcessRequest
type CreateGroupMonitoringAgentProcessAlertConfigTargetListItem struct {
	Level      string `name:"Level"`
	Id         string `name:"Id"`
	Arn        string `name:"Arn"`
	JsonParams string `name:"JsonParams"`
}

// CreateGroupMonitoringAgentProcessMatchExpress is a repeated param struct in CreateGroupMonitoringAgentProcessRequest
type CreateGroupMonitoringAgentProcessMatchExpress struct {
	Function string `name:"Function"`
	Name     string `name:"Name"`
	Value    string `name:"Value"`
}

// CreateGroupMonitoringAgentProcessResponse is the response struct for api CreateGroupMonitoringAgentProcess
type CreateGroupMonitoringAgentProcessResponse struct {
	*responses.BaseResponse
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Success   bool     `json:"Success" xml:"Success"`
	Resource  Resource `json:"Resource" xml:"Resource"`
}

// CreateCreateGroupMonitoringAgentProcessRequest creates a request to invoke CreateGroupMonitoringAgentProcess API
func CreateCreateGroupMonitoringAgentProcessRequest() (request *CreateGroupMonitoringAgentProcessRequest) {
	request = &CreateGroupMonitoringAgentProcessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateGroupMonitoringAgentProcess", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateGroupMonitoringAgentProcessResponse creates a response to parse from CreateGroupMonitoringAgentProcess response
func CreateCreateGroupMonitoringAgentProcessResponse() (response *CreateGroupMonitoringAgentProcessResponse) {
	response = &CreateGroupMonitoringAgentProcessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
