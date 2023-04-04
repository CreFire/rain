package tools

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	// epoch 是开始时间戳，用于缩小时间戳范围，以便在后面生成ID时占用更少的位数。
	epoch int64 = 1616756528808 // 2021-03-26 08:22:08 GMT+8:00

	// nodeBits 表示节点ID的位数
	nodeBits uint8 = 10

	// sequenceBits 表示序列号的位数
	sequenceBits uint8 = 12

	// maxNode 表示节点ID的最大值
	maxNode uint16 = 1<<nodeBits - 1

	// maxSequence 表示序列号的最大值
	maxSequence uint16 = 1<<sequenceBits - 1

	// timeShift 表示时间戳向左的偏移量
	timeShift = nodeBits + sequenceBits

	// nodeShift 表示节点ID向左的偏移量
	nodeShift = sequenceBits
)

// Generator 是雪花算法的ID生成器
type Generator struct {
	mu        sync.Mutex
	node      uint16
	sequence  uint16
	lastStamp int64
}

// NewGenerator 创建一个新的雪花算法ID生成器
func NewGenerator(node uint16) (*Generator, error) {
	if node < 0 || node > maxNode {
		return nil, errors.New("invalid node id")
	}
	return &Generator{node: node}, nil
}

// Generate 生成一个新的唯一ID
func (g *Generator) Generate() uint64 {
	g.mu.Lock()
	defer g.mu.Unlock()

	// 获取当前时间戳
	now := time.Now().UnixNano()/1000000 - epoch
	if now < g.lastStamp {
		panic("clock moved backwards")
	}

	// 如果是同一毫秒内生成的ID，增加序列号
	if now == g.lastStamp {
		g.sequence++
		if g.sequence > maxSequence {
			time.Sleep(time.Millisecond)
			now = time.Now().UnixNano()/1000000 - epoch
			g.sequence = 0
		}
	} else {
		g.sequence = 0
	}

	// 更新最后时间戳
	g.lastStamp = now
	nowOne := uint64(time.Now().UnixNano())
	// 组装ID
	return (nowOne << timeShift) | (uint64(g.node) << nodeShift) | uint64(g.sequence)
}

// example 例子
func example() {
	// 创建一个新的ID生成器
	generator, err := NewGenerator(1)
	if err != nil {
		panic(err)
	}

	// 生成100个ID并输出
	for i := 0; i < 100; i++ {
		id := generator.Generate()
		fmt.Println(id)
	}
}
