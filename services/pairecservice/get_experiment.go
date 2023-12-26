package pairecservice

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

// GetExperiment invokes the pairecservice.GetExperiment API synchronously
func (client *Client) GetExperiment(request *GetExperimentRequest) (response *GetExperimentResponse, err error) {
	response = CreateGetExperimentResponse()
	err = client.DoAction(request, response)
	return
}

// GetExperimentWithChan invokes the pairecservice.GetExperiment API asynchronously
func (client *Client) GetExperimentWithChan(request *GetExperimentRequest) (<-chan *GetExperimentResponse, <-chan error) {
	responseChan := make(chan *GetExperimentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetExperiment(request)
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

// GetExperimentWithCallback invokes the pairecservice.GetExperiment API asynchronously
func (client *Client) GetExperimentWithCallback(request *GetExperimentRequest, callback func(response *GetExperimentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetExperimentResponse
		var err error
		defer close(result)
		response, err = client.GetExperiment(request)
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

// GetExperimentRequest is the request struct for api GetExperiment
type GetExperimentRequest struct {
	*requests.RoaRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	ExperimentId string `position:"Path" name:"ExperimentId"`
}

// GetExperimentResponse is the response struct for api GetExperiment
type GetExperimentResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	AliasExperimentId string `json:"AliasExperimentId" xml:"AliasExperimentId"`
	ExperimentGroupId string `json:"ExperimentGroupId" xml:"ExperimentGroupId"`
	LayerId           string `json:"LayerId" xml:"LayerId"`
	LaboratoryId      string `json:"LaboratoryId" xml:"LaboratoryId"`
	SceneId           string `json:"SceneId" xml:"SceneId"`
	Name              string `json:"Name" xml:"Name"`
	Description       string `json:"Description" xml:"Description"`
	Type              string `json:"Type" xml:"Type"`
	FlowPercent       int    `json:"FlowPercent" xml:"FlowPercent"`
	Buckets           string `json:"Buckets" xml:"Buckets"`
	DebugUsers        string `json:"DebugUsers" xml:"DebugUsers"`
	DebugCrowdId      string `json:"DebugCrowdId" xml:"DebugCrowdId"`
	Config            string `json:"Config" xml:"Config"`
	Status            string `json:"Status" xml:"Status"`
	GmtCreateTime     string `json:"GmtCreateTime" xml:"GmtCreateTime"`
	GmtModifiedTime   string `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
}

// CreateGetExperimentRequest creates a request to invoke GetExperiment API
func CreateGetExperimentRequest() (request *GetExperimentRequest) {
	request = &GetExperimentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "GetExperiment", "/api/v1/experiments/[ExperimentId]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetExperimentResponse creates a response to parse from GetExperiment response
func CreateGetExperimentResponse() (response *GetExperimentResponse) {
	response = &GetExperimentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}