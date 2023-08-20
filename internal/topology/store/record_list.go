package store

import (
	"ClusterWatcher/internal/pkg/log"
	"ClusterWatcher/internal/pkg/model"
	"sync"
)

type RecordList struct {
	Records []model.RecordM `json:"records"`
}

func (rl *RecordList) render() model.NodeSummaries {
	return rl.mergeRecordList()
}

func (rl *RecordList) mergeRecordList() model.NodeSummaries {
	var wg sync.WaitGroup

	var returns []model.RecordM
	// 创建通道来接收每个协程的结果

	totalRecords := len(rl.Records)

	if totalRecords == 0 {
		return model.NodeSummaries{}
	}

	if totalRecords <= 100 {
		log.Debugf("[store] merge records: %d", totalRecords)
		result := rl.Records[0]
		for i := 1; i < totalRecords; i++ {
			recordMerge(result.Nodes, rl.Records[i].Nodes)
		}
		return result.Nodes
	}
	log.Debugf("[store] merge records: %d", totalRecords)

	results := make(chan model.RecordM)

	// 定义每个协程的任务
	task := func(works []model.RecordM) {
		defer wg.Done()
		length := len(works)
		for j := length - 2; j >= 0; j-- {
			recordMerge(works[length-1].Nodes, works[j].Nodes)
		}
		// 将结果发送到通道
		results <- works[length-1]
	}

	chunkSize := 10
	numWorker := totalRecords / chunkSize

	// 启动协程执行任务
	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > totalRecords {
			end = totalRecords
			if end-start <= 1 {
				returns = append(returns, rl.Records[end-1])
			}
		}
		go task(rl.Records[start:end])
	}

	// 等待所有协程完成任务
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集每个协程的结果
	for partialResult := range results {
		returns = append(returns, partialResult)
	}

	for i := 1; i < len(returns); i++ {
		if returns[0].Timestamp.After(returns[i].Timestamp) {
			recordMerge(returns[0].Nodes, returns[i].Nodes)
		} else {
			recordMerge(returns[i].Nodes, returns[0].Nodes)
		}
	}
	return returns[0].Nodes
}

func recordMerge(dest model.NodeSummaries, other model.NodeSummaries) {
	for key, Node := range other {
		if value, ok := dest[key]; ok {
			nodeMerge(&value, &Node)
			// 一定要用键值进行更新，不然对象不会被更新
			dest[key] = value
		} else {
			dest[key] = Node
		}
	}
}

func nodeMerge(origin *model.NodeSummary, other *model.NodeSummary) {
	if other.Label != "" {
		origin.Label = other.Label
	}
	if other.LabelMinor != "" {
		origin.LabelMinor = other.LabelMinor
	}
	if other.Rank != "" {
		origin.Rank = other.Rank
	}
	if other.Tag != "" {
		origin.Tag = other.Tag
	}
	origin.Stack = other.Stack
	copy(origin.Metadata, other.Metadata)
	copy(origin.Parents, other.Parents)
	copy(origin.Tables, other.Tables)
	copy(origin.Adjacency, origin.Adjacency.Merge(other.Adjacency))
}
