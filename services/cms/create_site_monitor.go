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

// CreateSiteMonitor invokes the cms.CreateSiteMonitor API synchronously
func (client *Client) CreateSiteMonitor(request *CreateSiteMonitorRequest) (response *CreateSiteMonitorResponse, err error) {
	response = CreateCreateSiteMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSiteMonitorWithChan invokes the cms.CreateSiteMonitor API asynchronously
func (client *Client) CreateSiteMonitorWithChan(request *CreateSiteMonitorRequest) (<-chan *CreateSiteMonitorResponse, <-chan error) {
	responseChan := make(chan *CreateSiteMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSiteMonitor(request)
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

// CreateSiteMonitorWithCallback invokes the cms.CreateSiteMonitor API asynchronously
func (client *Client) CreateSiteMonitorWithCallback(request *CreateSiteMonitorRequest, callback func(response *CreateSiteMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSiteMonitorResponse
		var err error
		defer close(result)
		response, err = client.CreateSiteMonitor(request)
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

// CreateSiteMonitorRequest is the request struct for api CreateSiteMonitor
type CreateSiteMonitorRequest struct {
	*requests.RpcRequest
	ReportProject   string           `position:"Query" name:"ReportProject"`
	TaskName        string           `position:"Query" name:"TaskName"`
	AlertIds        string           `position:"Query" name:"AlertIds"`
	Address         string           `position:"Query" name:"Address"`
	TaskType        string           `position:"Query" name:"TaskType"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	AgentGroup      string           `position:"Query" name:"AgentGroup"`
	IspCities       string           `position:"Query" name:"IspCities"`
	OptionsJson     string           `position:"Query" name:"OptionsJson"`
	IspCityFailRate requests.Float   `position:"Query" name:"IspCityFailRate"`
	IntervalUnit    string           `position:"Query" name:"IntervalUnit"`
	Interval        string           `position:"Query" name:"Interval"`
}

// CreateSiteMonitorResponse is the response struct for api CreateSiteMonitor
type CreateSiteMonitorResponse struct {
	*responses.BaseResponse
	Code             string                              `json:"Code" xml:"Code"`
	Message          string                              `json:"Message" xml:"Message"`
	RequestId        string                              `json:"RequestId" xml:"RequestId"`
	Success          string                              `json:"Success" xml:"Success"`
	Data             Data                                `json:"Data" xml:"Data"`
	CreateResultList CreateResultListInCreateSiteMonitor `json:"CreateResultList" xml:"CreateResultList"`
}

// CreateCreateSiteMonitorRequest creates a request to invoke CreateSiteMonitor API
func CreateCreateSiteMonitorRequest() (request *CreateSiteMonitorRequest) {
	request = &CreateSiteMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateSiteMonitor", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSiteMonitorResponse creates a response to parse from CreateSiteMonitor response
func CreateCreateSiteMonitorResponse() (response *CreateSiteMonitorResponse) {
	response = &CreateSiteMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
