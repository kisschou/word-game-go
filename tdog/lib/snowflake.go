package lib

import (
	"time"
)

// 雪花算法
// 0 : 固定位
// 1 - 41 : 时间戳
// 42 - 51 : 机器id
// 52 - 63 : 序列号

type (
	SnowFlake struct {
		MachineId     int64 // 机器 id 占10位, 十进制范围是 [ 0, 1023   ]
		SN            int64 // 序列号占 12 位,十进制范围是 [ 0, 4095   ]
		LastTimeStamp int64 // 上次的时间戳(毫秒级), 1秒=1000毫秒, 1毫秒=1000微秒,1微秒=1000纳秒
	}
)

func (sf *SnowFlake) New() int64 {
	curTimeStamp := time.Now().UnixNano() / 1000000

	// 把机器 id 左移 12 位,让出 12 位空间给序列号使用
	sf.MachineId = sf.MachineId << 12

	// 同一毫秒
	if curTimeStamp == sf.LastTimeStamp {
		sf.SN++
		// 序列号占 12 位,十进制范围是 [ 0, 4095   ]
		if sf.SN > 4095 {
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano() / 1000000
			sf.LastTimeStamp = curTimeStamp
			sf.SN = 0
		}

		// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1   )和时间戳进行并操作
		// 并结果( 右数   )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22
		id := rightBinValue | sf.MachineId | sf.SN
		return id
	}

	if curTimeStamp > sf.LastTimeStamp {
		sf.SN = 0
		sf.LastTimeStamp = curTimeStamp
		// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1   )和时间戳进行并操作
		// 并结果( 右数   )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22
		id := rightBinValue | sf.MachineId | sf.SN
		return id
	}

	if curTimeStamp < sf.LastTimeStamp {
		return 0
	}

	return 0
}
