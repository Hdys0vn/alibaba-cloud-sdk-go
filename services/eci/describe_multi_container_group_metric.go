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

package eci

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeMultiContainerGroupMetric invokes the eci.DescribeMultiContainerGroupMetric API synchronously
// api document: https://help.aliyun.com/api/eci/describemulticontainergroupmetric.html
func (client *Client) DescribeMultiContainerGroupMetric(request *DescribeMultiContainerGroupMetricRequest) (response *DescribeMultiContainerGroupMetricResponse, err error) {
	response = CreateDescribeMultiContainerGroupMetricResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMultiContainerGroupMetricWithChan invokes the eci.DescribeMultiContainerGroupMetric API asynchronously
// api document: https://help.aliyun.com/api/eci/describemulticontainergroupmetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMultiContainerGroupMetricWithChan(request *DescribeMultiContainerGroupMetricRequest) (<-chan *DescribeMultiContainerGroupMetricResponse, <-chan error) {
	responseChan := make(chan *DescribeMultiContainerGroupMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMultiContainerGroupMetric(request)
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

// DescribeMultiContainerGroupMetricWithCallback invokes the eci.DescribeMultiContainerGroupMetric API asynchronously
// api document: https://help.aliyun.com/api/eci/describemulticontainergroupmetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMultiContainerGroupMetricWithCallback(request *DescribeMultiContainerGroupMetricRequest, callback func(response *DescribeMultiContainerGroupMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMultiContainerGroupMetricResponse
		var err error
		defer close(result)
		response, err = client.DescribeMultiContainerGroupMetric(request)
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

// DescribeMultiContainerGroupMetricRequest is the request struct for api DescribeMultiContainerGroupMetric
type DescribeMultiContainerGroupMetricRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	RegionId             string           `position:"Query" name:"RegionId"`
	ContainerGroupIds    string           `position:"Query" name:"ContainerGroupIds"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	MetricType           string           `position:"Query" name:"MetricType"`
}

// DescribeMultiContainerGroupMetricResponse is the response struct for api DescribeMultiContainerGroupMetric
type DescribeMultiContainerGroupMetricResponse struct {
	*responses.BaseResponse
	RequestId    string                                     `json:"RequestId" xml:"RequestId"`
	MonitorDatas []DescribeMultiContainerGroupMetricRecord0 `json:"MonitorDatas" xml:"MonitorDatas"`
}

type DescribeMultiContainerGroupMetricRecord0 struct {
	ContainerGroupId string                                            `json:"ContainerGroupId" xml:"ContainerGroupId"`
	Records          []DescribeMultiContainerGroupMetricPodStat1       `json:"Records" xml:"Records"`
	ContainerInfos   []DescribeMultiContainerGroupMetricContainerInfo1 `json:"ContainerInfos" xml:"ContainerInfos"`
}

type DescribeMultiContainerGroupMetricPodStat1 struct {
	Timestamp  string                                        `json:"Timestamp" xml:"Timestamp"`
	Containers []DescribeMultiContainerGroupMetricContainer2 `json:"Containers" xml:"Containers"`
	CPU        DescribeMultiContainerGroupMetricCPU2         `json:"CPU" xml:"CPU"`
	Memory     DescribeMultiContainerGroupMetricMemory2      `json:"Memory" xml:"Memory"`
	Network    DescribeMultiContainerGroupMetricNetwork2     `json:"Network" xml:"Network"`
}

type DescribeMultiContainerGroupMetricContainer2 struct {
	Name   string                                   `json:"Name" xml:"Name"`
	CPU    DescribeMultiContainerGroupMetricCPU3    `json:"CPU" xml:"CPU"`
	Memory DescribeMultiContainerGroupMetricMemory3 `json:"Memory" xml:"Memory"`
}

type DescribeMultiContainerGroupMetricCPU3 struct {
	Limit                int64 `json:"Limit" xml:"Limit"`
	Load                 int64 `json:"Load" xml:"Load"`
	UsageCoreNanoSeconds int64 `json:"UsageCoreNanoSeconds" xml:"UsageCoreNanoSeconds"`
	UsageNanoCores       int64 `json:"UsageNanoCores" xml:"UsageNanoCores"`
}

type DescribeMultiContainerGroupMetricMemory3 struct {
	AvailableBytes int64 `json:"AvailableBytes" xml:"AvailableBytes"`
	UsageBytes     int64 `json:"UsageBytes" xml:"UsageBytes"`
	Cache          int64 `json:"Cache" xml:"Cache"`
	WorkingSet     int64 `json:"WorkingSet" xml:"WorkingSet"`
	Rss            int64 `json:"Rss" xml:"Rss"`
}

type DescribeMultiContainerGroupMetricCPU2 struct {
	Limit                int64 `json:"Limit" xml:"Limit"`
	Load                 int64 `json:"Load" xml:"Load"`
	UsageCoreNanoSeconds int64 `json:"UsageCoreNanoSeconds" xml:"UsageCoreNanoSeconds"`
	UsageNanoCores       int64 `json:"UsageNanoCores" xml:"UsageNanoCores"`
}

type DescribeMultiContainerGroupMetricMemory2 struct {
	AvailableBytes int64 `json:"AvailableBytes" xml:"AvailableBytes"`
	UsageBytes     int64 `json:"UsageBytes" xml:"UsageBytes"`
	Cache          int64 `json:"Cache" xml:"Cache"`
	WorkingSet     int64 `json:"WorkingSet" xml:"WorkingSet"`
	Rss            int64 `json:"Rss" xml:"Rss"`
}

type DescribeMultiContainerGroupMetricNetwork2 struct {
	Interfaces []DescribeMultiContainerGroupMetricInterface3 `json:"Interfaces" xml:"Interfaces"`
}

type DescribeMultiContainerGroupMetricInterface3 struct {
	TxBytes  int64  `json:"TxBytes" xml:"TxBytes"`
	RxBytes  int64  `json:"RxBytes" xml:"RxBytes"`
	TxErrors int64  `json:"TxErrors" xml:"TxErrors"`
	RxErrors int64  `json:"RxErrors" xml:"RxErrors"`
	Name     string `json:"Name" xml:"Name"`
}

type DescribeMultiContainerGroupMetricContainerInfo1 struct {
	Id             string                                             `json:"Id" xml:"Id"`
	Name           string                                             `json:"Name" xml:"Name"`
	Namespace      string                                             `json:"Namespace" xml:"Namespace"`
	Labels         string                                             `json:"Labels" xml:"Labels"`
	ContainerStats []DescribeMultiContainerGroupMetricContainerStats2 `json:"ContainerStats" xml:"ContainerStats"`
	Aliases        []string                                           `json:"Aliases" xml:"Aliases"`
	ContainerSpec  DescribeMultiContainerGroupMetricContainerSpec2    `json:"ContainerSpec" xml:"ContainerSpec"`
}

type DescribeMultiContainerGroupMetricContainerStats2 struct {
	Timestamp        string                                               `json:"Timestamp" xml:"Timestamp"`
	FsStats          []DescribeMultiContainerGroupMetricFsStats3          `json:"FsStats" xml:"FsStats"`
	AcceleratorStats []DescribeMultiContainerGroupMetricAcceleratorStats3 `json:"AcceleratorStats" xml:"AcceleratorStats"`
	CpuStats         DescribeMultiContainerGroupMetricCpuStats3           `json:"CpuStats" xml:"CpuStats"`
	DiskIoStats      DescribeMultiContainerGroupMetricDiskIoStats3        `json:"DiskIoStats" xml:"DiskIoStats"`
	MemoryStats      DescribeMultiContainerGroupMetricMemoryStats3        `json:"MemoryStats" xml:"MemoryStats"`
	NetworkStats     DescribeMultiContainerGroupMetricNetworkStats3       `json:"NetworkStats" xml:"NetworkStats"`
	TaskStats        DescribeMultiContainerGroupMetricTaskStats3          `json:"TaskStats" xml:"TaskStats"`
}

type DescribeMultiContainerGroupMetricFsStats3 struct {
	Device          string `json:"Device" xml:"Device"`
	Type            string `json:"Type" xml:"Type"`
	Limit           int64  `json:"Limit" xml:"Limit"`
	Usage           int64  `json:"Usage" xml:"Usage"`
	BaseUsage       int64  `json:"BaseUsage" xml:"BaseUsage"`
	Available       int64  `json:"Available" xml:"Available"`
	HasInodes       bool   `json:"HasInodes" xml:"HasInodes"`
	Inodes          int64  `json:"Inodes" xml:"Inodes"`
	InodesFree      int64  `json:"InodesFree" xml:"InodesFree"`
	ReadsCompleted  int64  `json:"ReadsCompleted" xml:"ReadsCompleted"`
	ReadsMerged     int64  `json:"ReadsMerged" xml:"ReadsMerged"`
	SectorsRead     int64  `json:"SectorsRead" xml:"SectorsRead"`
	ReadTime        int64  `json:"ReadTime" xml:"ReadTime"`
	WritesCompleted int64  `json:"WritesCompleted" xml:"WritesCompleted"`
	WritesMerged    int64  `json:"WritesMerged" xml:"WritesMerged"`
	SectorsWritten  int64  `json:"SectorsWritten" xml:"SectorsWritten"`
	WriteTime       int64  `json:"WriteTime" xml:"WriteTime"`
	IoInProgress    int64  `json:"IoInProgress" xml:"IoInProgress"`
	IoTime          int64  `json:"IoTime" xml:"IoTime"`
	WeightedIoTime  int64  `json:"WeightedIoTime" xml:"WeightedIoTime"`
}

type DescribeMultiContainerGroupMetricAcceleratorStats3 struct {
	Id          string `json:"Id" xml:"Id"`
	Make        string `json:"Make" xml:"Make"`
	Model       string `json:"Model" xml:"Model"`
	MemoryTotal int64  `json:"MemoryTotal" xml:"MemoryTotal"`
	MemoryUsed  int64  `json:"MemoryUsed" xml:"MemoryUsed"`
	DutyCycle   int64  `json:"DutyCycle" xml:"DutyCycle"`
}

type DescribeMultiContainerGroupMetricCpuStats3 struct {
	LoadAverage int64                                      `json:"LoadAverage" xml:"LoadAverage"`
	CpuUsage    DescribeMultiContainerGroupMetricCpuUsage4 `json:"CpuUsage" xml:"CpuUsage"`
	CpuCFS      DescribeMultiContainerGroupMetricCpuCFS4   `json:"CpuCFS" xml:"CpuCFS"`
}

type DescribeMultiContainerGroupMetricCpuUsage4 struct {
	Total        int64    `json:"Total" xml:"Total"`
	User         int64    `json:"User" xml:"User"`
	System       int64    `json:"System" xml:"System"`
	PerCpuUsages []string `json:"PerCpuUsages" xml:"PerCpuUsages"`
}

type DescribeMultiContainerGroupMetricCpuCFS4 struct {
	Periods          int64 `json:"Periods" xml:"Periods"`
	ThrottledPeriods int64 `json:"ThrottledPeriods" xml:"ThrottledPeriods"`
	ThrottledTime    int64 `json:"ThrottledTime" xml:"ThrottledTime"`
}

type DescribeMultiContainerGroupMetricDiskIoStats3 struct {
	IoServiceBytes []DescribeMultiContainerGroupMetricIoServiceByte4 `json:"IoServiceBytes" xml:"IoServiceBytes"`
	IoServiced     []DescribeMultiContainerGroupMetricIoServiced4    `json:"IoServiced" xml:"IoServiced"`
	IoQueued       []DescribeMultiContainerGroupMetricIoQueued4      `json:"IoQueued" xml:"IoQueued"`
	Sectors        []DescribeMultiContainerGroupMetricSector4        `json:"Sectors" xml:"Sectors"`
	IoServiceTime  []DescribeMultiContainerGroupMetricIoServiceTime4 `json:"IoServiceTime" xml:"IoServiceTime"`
	IoWaitTime     []DescribeMultiContainerGroupMetricIoWaitTime4    `json:"IoWaitTime" xml:"IoWaitTime"`
	IoMerged       []DescribeMultiContainerGroupMetricIoMerged4      `json:"IoMerged" xml:"IoMerged"`
	IoTime         []DescribeMultiContainerGroupMetricIoTime4        `json:"IoTime" xml:"IoTime"`
}

type DescribeMultiContainerGroupMetricIoServiceByte4 struct {
	Device string `json:"Device" xml:"Device"`
	Major  int64  `json:"Major" xml:"Major"`
	Minor  int64  `json:"Minor" xml:"Minor"`
	Stats  string `json:"Stats" xml:"Stats"`
}

type DescribeMultiContainerGroupMetricIoServiced4 struct {
	Device string `json:"Device" xml:"Device"`
	Major  int64  `json:"Major" xml:"Major"`
	Minor  int64  `json:"Minor" xml:"Minor"`
	Stats  string `json:"Stats" xml:"Stats"`
}

type DescribeMultiContainerGroupMetricIoQueued4 struct {
	Device string `json:"Device" xml:"Device"`
	Major  int64  `json:"Major" xml:"Major"`
	Minor  int64  `json:"Minor" xml:"Minor"`
	Stats  string `json:"Stats" xml:"Stats"`
}

type DescribeMultiContainerGroupMetricSector4 struct {
	Device string `json:"Device" xml:"Device"`
	Major  int64  `json:"Major" xml:"Major"`
	Minor  int64  `json:"Minor" xml:"Minor"`
	Stats  string `json:"Stats" xml:"Stats"`
}

type DescribeMultiContainerGroupMetricIoServiceTime4 struct {
	Device string `json:"Device" xml:"Device"`
	Major  int64  `json:"Major" xml:"Major"`
	Minor  int64  `json:"Minor" xml:"Minor"`
	Stats  string `json:"Stats" xml:"Stats"`
}

type DescribeMultiContainerGroupMetricIoWaitTime4 struct {
	Device string `json:"Device" xml:"Device"`
	Major  int64  `json:"Major" xml:"Major"`
	Minor  int64  `json:"Minor" xml:"Minor"`
	Stats  string `json:"Stats" xml:"Stats"`
}

type DescribeMultiContainerGroupMetricIoMerged4 struct {
	Device string `json:"Device" xml:"Device"`
	Major  int64  `json:"Major" xml:"Major"`
	Minor  int64  `json:"Minor" xml:"Minor"`
	Stats  string `json:"Stats" xml:"Stats"`
}

type DescribeMultiContainerGroupMetricIoTime4 struct {
	Device string `json:"Device" xml:"Device"`
	Major  int64  `json:"Major" xml:"Major"`
	Minor  int64  `json:"Minor" xml:"Minor"`
	Stats  string `json:"Stats" xml:"Stats"`
}

type DescribeMultiContainerGroupMetricMemoryStats3 struct {
	Usage            int64                                              `json:"Usage" xml:"Usage"`
	MaxUsage         int64                                              `json:"MaxUsage" xml:"MaxUsage"`
	Cache            int64                                              `json:"Cache" xml:"Cache"`
	Rss              int64                                              `json:"Rss" xml:"Rss"`
	Swap             int64                                              `json:"Swap" xml:"Swap"`
	WorkingSet       int64                                              `json:"WorkingSet" xml:"WorkingSet"`
	FailCnt          int64                                              `json:"FailCnt" xml:"FailCnt"`
	ContainerData    DescribeMultiContainerGroupMetricContainerData4    `json:"ContainerData" xml:"ContainerData"`
	HierarchicalData DescribeMultiContainerGroupMetricHierarchicalData4 `json:"HierarchicalData" xml:"HierarchicalData"`
}

type DescribeMultiContainerGroupMetricContainerData4 struct {
	PgFault    int64 `json:"PgFault" xml:"PgFault"`
	PgmajFault int64 `json:"PgmajFault" xml:"PgmajFault"`
}

type DescribeMultiContainerGroupMetricHierarchicalData4 struct {
	PgFault    int64 `json:"PgFault" xml:"PgFault"`
	PgmajFault int64 `json:"PgmajFault" xml:"PgmajFault"`
}

type DescribeMultiContainerGroupMetricNetworkStats3 struct {
	Name           string                                             `json:"Name" xml:"Name"`
	RxBytes        int64                                              `json:"RxBytes" xml:"RxBytes"`
	RxPackets      int64                                              `json:"RxPackets" xml:"RxPackets"`
	RxErrors       int64                                              `json:"RxErrors" xml:"RxErrors"`
	RxDropped      int64                                              `json:"RxDropped" xml:"RxDropped"`
	TxBytes        int64                                              `json:"TxBytes" xml:"TxBytes"`
	TxPackets      int64                                              `json:"TxPackets" xml:"TxPackets"`
	TxDropped      int64                                              `json:"TxDropped" xml:"TxDropped"`
	TxErrors       int64                                              `json:"TxErrors" xml:"TxErrors"`
	InterfaceStats []DescribeMultiContainerGroupMetricInterfaceStats4 `json:"InterfaceStats" xml:"InterfaceStats"`
	Tcp            DescribeMultiContainerGroupMetricTcp4              `json:"Tcp" xml:"Tcp"`
	Tcp6           DescribeMultiContainerGroupMetricTcp64             `json:"Tcp6" xml:"Tcp6"`
	Udp            DescribeMultiContainerGroupMetricUdp4              `json:"Udp" xml:"Udp"`
	Udp6           DescribeMultiContainerGroupMetricUdp64             `json:"Udp6" xml:"Udp6"`
}

type DescribeMultiContainerGroupMetricInterfaceStats4 struct {
	Name      string `json:"Name" xml:"Name"`
	RxBytes   int64  `json:"RxBytes" xml:"RxBytes"`
	RxPackets int64  `json:"RxPackets" xml:"RxPackets"`
	RxErrors  int64  `json:"RxErrors" xml:"RxErrors"`
	RxDropped int64  `json:"RxDropped" xml:"RxDropped"`
	TxBytes   int64  `json:"TxBytes" xml:"TxBytes"`
	TxPackets int64  `json:"TxPackets" xml:"TxPackets"`
	TxDropped int64  `json:"TxDropped" xml:"TxDropped"`
	TxErrors  int64  `json:"TxErrors" xml:"TxErrors"`
}

type DescribeMultiContainerGroupMetricTcp4 struct {
	Established int64 `json:"Established" xml:"Established"`
	SynSent     int64 `json:"SynSent" xml:"SynSent"`
	SynRecv     int64 `json:"SynRecv" xml:"SynRecv"`
	FinWait1    int64 `json:"FinWait1" xml:"FinWait1"`
	FinWait2    int64 `json:"FinWait2" xml:"FinWait2"`
	TimeWait    int64 `json:"TimeWait" xml:"TimeWait"`
	Close       int64 `json:"Close" xml:"Close"`
	CloseWait   int64 `json:"CloseWait" xml:"CloseWait"`
	LastAck     int64 `json:"LastAck" xml:"LastAck"`
	Listen      int64 `json:"Listen" xml:"Listen"`
	Closing     int64 `json:"Closing" xml:"Closing"`
}

type DescribeMultiContainerGroupMetricTcp64 struct {
	Established int64 `json:"Established" xml:"Established"`
	SynSent     int64 `json:"SynSent" xml:"SynSent"`
	SynRecv     int64 `json:"SynRecv" xml:"SynRecv"`
	FinWait1    int64 `json:"FinWait1" xml:"FinWait1"`
	FinWait2    int64 `json:"FinWait2" xml:"FinWait2"`
	TimeWait    int64 `json:"TimeWait" xml:"TimeWait"`
	Close       int64 `json:"Close" xml:"Close"`
	CloseWait   int64 `json:"CloseWait" xml:"CloseWait"`
	LastAck     int64 `json:"LastAck" xml:"LastAck"`
	Listen      int64 `json:"Listen" xml:"Listen"`
	Closing     int64 `json:"Closing" xml:"Closing"`
}

type DescribeMultiContainerGroupMetricUdp4 struct {
	Listen   int64 `json:"Listen" xml:"Listen"`
	Dropped  int64 `json:"Dropped" xml:"Dropped"`
	RxQueued int64 `json:"RxQueued" xml:"RxQueued"`
	TxQueued int64 `json:"TxQueued" xml:"TxQueued"`
}

type DescribeMultiContainerGroupMetricUdp64 struct {
	Listen   int64 `json:"Listen" xml:"Listen"`
	Dropped  int64 `json:"Dropped" xml:"Dropped"`
	RxQueued int64 `json:"RxQueued" xml:"RxQueued"`
	TxQueued int64 `json:"TxQueued" xml:"TxQueued"`
}

type DescribeMultiContainerGroupMetricTaskStats3 struct {
	NrSleeping        int64 `json:"NrSleeping" xml:"NrSleeping"`
	NrRunning         int64 `json:"NrRunning" xml:"NrRunning"`
	NrStopped         int64 `json:"NrStopped" xml:"NrStopped"`
	NrUninterruptible int64 `json:"NrUninterruptible" xml:"NrUninterruptible"`
	NrIoWait          int64 `json:"NrIoWait" xml:"NrIoWait"`
}

type DescribeMultiContainerGroupMetricContainerSpec2 struct {
	CreationTime     string                                            `json:"CreationTime" xml:"CreationTime"`
	HasCpu           bool                                              `json:"HasCpu" xml:"HasCpu"`
	HasMemory        bool                                              `json:"HasMemory" xml:"HasMemory"`
	HasNetwork       bool                                              `json:"HasNetwork" xml:"HasNetwork"`
	HasFilesystem    bool                                              `json:"HasFilesystem" xml:"HasFilesystem"`
	HasDiskIo        bool                                              `json:"HasDiskIo" xml:"HasDiskIo"`
	HasCustomMetrics bool                                              `json:"HasCustomMetrics" xml:"HasCustomMetrics"`
	Image            string                                            `json:"Image" xml:"Image"`
	Labels           string                                            `json:"Labels" xml:"Labels"`
	Envs             string                                            `json:"Envs" xml:"Envs"`
	ContainerCpu     DescribeMultiContainerGroupMetricContainerCpu3    `json:"ContainerCpu" xml:"ContainerCpu"`
	ContainerMemory  DescribeMultiContainerGroupMetricContainerMemory3 `json:"ContainerMemory" xml:"ContainerMemory"`
}

type DescribeMultiContainerGroupMetricContainerCpu3 struct {
	Limit    int64  `json:"Limit" xml:"Limit"`
	MaxLimit int64  `json:"MaxLimit" xml:"MaxLimit"`
	Mask     string `json:"Mask" xml:"Mask"`
	Quota    int64  `json:"Quota" xml:"Quota"`
	Period   int64  `json:"Period" xml:"Period"`
}

type DescribeMultiContainerGroupMetricContainerMemory3 struct {
	Limit       int64 `json:"Limit" xml:"Limit"`
	Reservation int64 `json:"Reservation" xml:"Reservation"`
	SwapLimit   int64 `json:"SwapLimit" xml:"SwapLimit"`
}

// CreateDescribeMultiContainerGroupMetricRequest creates a request to invoke DescribeMultiContainerGroupMetric API
func CreateDescribeMultiContainerGroupMetricRequest() (request *DescribeMultiContainerGroupMetricRequest) {
	request = &DescribeMultiContainerGroupMetricRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eci", "2018-08-08", "DescribeMultiContainerGroupMetric", "eci", "openAPI")
	return
}

// CreateDescribeMultiContainerGroupMetricResponse creates a response to parse from DescribeMultiContainerGroupMetric response
func CreateDescribeMultiContainerGroupMetricResponse() (response *DescribeMultiContainerGroupMetricResponse) {
	response = &DescribeMultiContainerGroupMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
