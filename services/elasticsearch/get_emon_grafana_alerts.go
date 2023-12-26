package elasticsearch

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

// GetEmonGrafanaAlerts invokes the elasticsearch.GetEmonGrafanaAlerts API synchronously
func (client *Client) GetEmonGrafanaAlerts(request *GetEmonGrafanaAlertsRequest) (response *GetEmonGrafanaAlertsResponse, err error) {
	response = CreateGetEmonGrafanaAlertsResponse()
	err = client.DoAction(request, response)
	return
}

// GetEmonGrafanaAlertsWithChan invokes the elasticsearch.GetEmonGrafanaAlerts API asynchronously
func (client *Client) GetEmonGrafanaAlertsWithChan(request *GetEmonGrafanaAlertsRequest) (<-chan *GetEmonGrafanaAlertsResponse, <-chan error) {
	responseChan := make(chan *GetEmonGrafanaAlertsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEmonGrafanaAlerts(request)
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

// GetEmonGrafanaAlertsWithCallback invokes the elasticsearch.GetEmonGrafanaAlerts API asynchronously
func (client *Client) GetEmonGrafanaAlertsWithCallback(request *GetEmonGrafanaAlertsRequest, callback func(response *GetEmonGrafanaAlertsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEmonGrafanaAlertsResponse
		var err error
		defer close(result)
		response, err = client.GetEmonGrafanaAlerts(request)
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

// GetEmonGrafanaAlertsRequest is the request struct for api GetEmonGrafanaAlerts
type GetEmonGrafanaAlertsRequest struct {
	*requests.RoaRequest
	ProjectId string `position:"Path" name:"ProjectId"`
}

// GetEmonGrafanaAlertsResponse is the response struct for api GetEmonGrafanaAlerts
type GetEmonGrafanaAlertsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateGetEmonGrafanaAlertsRequest creates a request to invoke GetEmonGrafanaAlerts API
func CreateGetEmonGrafanaAlertsRequest() (request *GetEmonGrafanaAlertsRequest) {
	request = &GetEmonGrafanaAlertsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "GetEmonGrafanaAlerts", "/openapi/emon/projects/[ProjectId]/grafana/proxy/api/alerts", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetEmonGrafanaAlertsResponse creates a response to parse from GetEmonGrafanaAlerts response
func CreateGetEmonGrafanaAlertsResponse() (response *GetEmonGrafanaAlertsResponse) {
	response = &GetEmonGrafanaAlertsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
