package hbase

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

// EnableHBaseueModule invokes the hbase.EnableHBaseueModule API synchronously
func (client *Client) EnableHBaseueModule(request *EnableHBaseueModuleRequest) (response *EnableHBaseueModuleResponse, err error) {
	response = CreateEnableHBaseueModuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableHBaseueModuleWithChan invokes the hbase.EnableHBaseueModule API asynchronously
func (client *Client) EnableHBaseueModuleWithChan(request *EnableHBaseueModuleRequest) (<-chan *EnableHBaseueModuleResponse, <-chan error) {
	responseChan := make(chan *EnableHBaseueModuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableHBaseueModule(request)
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

// EnableHBaseueModuleWithCallback invokes the hbase.EnableHBaseueModule API asynchronously
func (client *Client) EnableHBaseueModuleWithCallback(request *EnableHBaseueModuleRequest, callback func(response *EnableHBaseueModuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableHBaseueModuleResponse
		var err error
		defer close(result)
		response, err = client.EnableHBaseueModule(request)
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

// EnableHBaseueModuleRequest is the request struct for api EnableHBaseueModule
type EnableHBaseueModuleRequest struct {
	*requests.RpcRequest
	ClientToken        string           `position:"Query" name:"ClientToken"`
	ModuleTypeName     string           `position:"Query" name:"ModuleTypeName"`
	HbaseueClusterId   string           `position:"Query" name:"HbaseueClusterId"`
	BdsId              string           `position:"Query" name:"BdsId"`
	ModuleClusterName  string           `position:"Query" name:"ModuleClusterName"`
	AutoRenewPeriod    requests.Integer `position:"Query" name:"AutoRenewPeriod"`
	Period             requests.Integer `position:"Query" name:"Period"`
	DiskSize           requests.Integer `position:"Query" name:"DiskSize"`
	MasterInstanceType string           `position:"Query" name:"MasterInstanceType"`
	DiskType           string           `position:"Query" name:"DiskType"`
	VswitchId          string           `position:"Query" name:"VswitchId"`
	PeriodUnit         string           `position:"Query" name:"PeriodUnit"`
	CoreInstanceType   string           `position:"Query" name:"CoreInstanceType"`
	VpcId              string           `position:"Query" name:"VpcId"`
	NodeCount          requests.Integer `position:"Query" name:"NodeCount"`
	ZoneId             string           `position:"Query" name:"ZoneId"`
	PayType            string           `position:"Query" name:"PayType"`
}

// EnableHBaseueModuleResponse is the response struct for api EnableHBaseueModule
type EnableHBaseueModuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ClusterId string `json:"ClusterId" xml:"ClusterId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateEnableHBaseueModuleRequest creates a request to invoke EnableHBaseueModule API
func CreateEnableHBaseueModuleRequest() (request *EnableHBaseueModuleRequest) {
	request = &EnableHBaseueModuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "EnableHBaseueModule", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableHBaseueModuleResponse creates a response to parse from EnableHBaseueModule response
func CreateEnableHBaseueModuleResponse() (response *EnableHBaseueModuleResponse) {
	response = &EnableHBaseueModuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
