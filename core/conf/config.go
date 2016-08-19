package conf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const middle = "========="

type Config struct {
	Mymap   map[string]string
	Keylist []string
	strcet  string
}

func (c *Config) InitConfig(path string) {
	c.Keylist = make([]string, 0, 15)
	c.Mymap = make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//以＃开头的为注释，index 查找字符串，范围所在位置
		if strings.Index(s, "#") == 0 {
			continue
		}
		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		// 返回［］包裹的字符串
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			c.Keylist = append(c.Keylist, c.strcet)
			continue
		}
		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strcet + middle + frist
		//fmt.Println(key)
		c.Mymap[key] = strings.TrimSpace(second)
	}
}

func (c Config) Read(node, key string) string {
	key = node + middle + key
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}

func (c *Config) List() {
	for _, keys := range c.Keylist {
		fmt.Println(keys)
	}
	// 这里strcet 只会返回最后一个，需要改下
}
