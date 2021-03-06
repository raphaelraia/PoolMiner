package config

import (
	"sync"
)

type Task struct {
	Header     string
	Nonce      string
	Solution   string
	Difficulty string
}
type ReqObj struct {
        Id      int      `json:"id"`
        Jsonrpc string   `json:"jsonrpc"`
        Method  string   `json:"method"`
        Params  []string `json:"params"`
}


type TaskWrapper struct {
	Lock  sync.Mutex
	TaskQ Task
}

type StreamData struct{
	Nedges uint32
	ThreadId uint32
	Difficulty string
	Nonce uint64
	Header []byte
}


type Param struct {
	Server [3]string
	Account string
	VerboseLevel    uint
	Algorithm int
	Threads int
	Cpu bool
	Cuda bool
	Opencl bool
	Worker_name string
	ServerIndex int
}

type DeviceInfo struct {
	Lock           sync.Mutex
	DeviceId       uint32
	Start_time	int64
	Use_time       int64
	Solution_count int64
	Hash_rate      float32
	Gps            int64
}

func (param Param) New(server string, server1 string, server2 string, account string, worker_name string, verboseLevel uint, algorithm int, threads int, cpu bool, cuda bool, opencl bool) Param{
	param.Server[0] = server
	param.Server[1] = server1
	param.Server[2] = server2
	param.Account = account
	param.VerboseLevel = verboseLevel
	param.Algorithm = algorithm
	param.Threads = threads
	param.Cpu = cpu
	param.Cuda = cuda
	param.Opencl = opencl
	param.Worker_name = worker_name
	param.ServerIndex = 0
	return param
}
func (deviceInfo DeviceInfo) New(_lock sync.Mutex, _deviceId uint32, _start_time int64, _use_time int64, _solution_count int64, _hash_rate float32, _gps int64) DeviceInfo{
	deviceInfo.Lock = _lock
	deviceInfo.DeviceId = _deviceId
	deviceInfo.Start_time = _start_time
	deviceInfo.Use_time = _use_time
	deviceInfo.Solution_count = _solution_count
	deviceInfo.Hash_rate = _hash_rate
	deviceInfo.Gps = _gps
	return deviceInfo
}

//global task variable, need sync accessed
var CurrentTask TaskWrapper

func (streamData StreamData) New(nedges uint32, threadId uint32, Difficulty string, nonce uint64, header []byte) StreamData {
	streamData.Nedges = nedges
	streamData.ThreadId = threadId
	streamData.Difficulty = Difficulty
	streamData.Nonce = nonce
	streamData.Header = header
	return streamData
}
