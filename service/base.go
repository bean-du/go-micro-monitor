package service

import (
	"github.com/micro/go-micro/v2/registry"
	"monitor/config"
	"sync"
)

var BaseSrv *baseSrv

func init()  {
	BaseSrv = &baseSrv{
		service: make(map[string][]*registry.Node),
	}
}

type baseSrv struct {
	mu      sync.Mutex
	service map[string][]*registry.Node
}


func (b *baseSrv) ServerList() []string {
	services, err := config.Micro.Options().Registry.ListServices()
	if err != nil {
		return nil
	}
	b.resetService()
	names := []string{}
	b.setServices(services)
	list := make(map[string]struct{})
	for _, s := range services {
		list[s.Name] = struct{}{}
	}
	for k := range list {
		names = append(names, k)
	}
	return names
}

func (b *baseSrv) NodeList(serverName string) []string {
	nodes := b.getServerNodes(serverName)
	nodeStr := []string{}
	if nodes != nil {
		for _, n := range nodes {
			nodeStr = append(nodeStr, n.Address)
		}
	}
	return nodeStr
}

func (b *baseSrv) setServices(service []*registry.Service) {
	b.mu.Lock()
	defer b.mu.Unlock()
	m := make(map[string][]*registry.Node)
	for _, s := range service{
		m[s.Name] = append(m[s.Name], s.Nodes...)
	}
	b.service = m
}

func (b *baseSrv) resetService() {
	b.mu.Lock()
	defer b.mu.Unlock()

	m := make(map[string][]*registry.Node)
	b.service = m
}

func (b *baseSrv) getServerNodes(name string) []*registry.Node {
	b.mu.Lock()
	defer b.mu.Unlock()

	if n, ok := b.service[name]; ok {
		return n
	}
	return nil
}


func (b *baseSrv) Index() string {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <!-- import CSS -->
    <link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css">
    <script src="https://cdn.bootcdn.net/ajax/libs/axios/0.21.0/axios.min.js"></script>
</head>
<body>
<div id="app">
    <el-container>
        <el-header style="background-color: #2b2f3a; margin-bottom: 20px;">
            <div class="my_menu">
                <el-row style="height: 60px">
                    <el-col :span="4" style="height: 60px; font-size:40px; ">
                        <span style="line-height: 60px;">Monitor</span>
                    </el-col>
                    <el-col :span="16" style="height: 60px">
                        <span></span>
                    </el-col>
                    <el-col :span="4" style="height: 60px">
                        <el-dropdown>
                            <el-button type="primary" style="margin-top: 11px;">
                                Providers<i class="el-icon-arrow-down el-icon--right"></i>
                            </el-button>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item>GOPS</el-dropdown-item>
                                <el-dropdown-item>PPROF</el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </el-col>
                </el-row>
            </div>
        </el-header>
        <el-container style="width: 1024px;height: 100%; margin:0 auto">
            <el-aside width="250px">
                <div style="margin-left: 20px;">
                    <span style=" line-height: 30px; font-size: 25px;">Service</span>
                </div>
                <el-menu v-for="(item,index) in serviceList">
                    <el-submenu :index="leftIndex.toString()" :collapse="isCollapse">
                        <template slot="title">
                            <div @click="getNodes(item.name, index)">
                                <span style="font-size: 14px;" slot="title">{{ item.name }}</span>
                            </div>
                        </template>
                        <el-menu-item-group v-for="node in nodes" v-show="index == leftIndex">
                            <el-menu-item index="1-1" @click="onClickNode(node.addr)">{{ node.addr }}</el-menu-item>
                        </el-menu-item-group>
                    </el-submenu>
                </el-menu>
            </el-aside>
            <el-main>
                <div class="current" style="margin-left: 10px;">
                    <span>当前服务名称：<el-tag v-if="currentService != ''" type="success"
                                         style="font-size: 16px;">{{ currentService }}</el-tag>   当前节点: <el-tag
                                v-if="currentNode != ''" type="success">{{ currentNode }}</el-tag> </span>
                </div>
                <div class="function_button">
                    <el-row>
                        <el-tooltip class="item" effect="dark" content="基本状态" placement="top">
                            <el-button @click="handleFunction('stats')" round>Stats</el-button>
                        </el-tooltip>
                        <el-tooltip class="item" effect="dark" content="Stack 状态" placement="top">
                            <el-button @click="handleFunction('stack')" round>Stack</el-button>
                        </el-tooltip>
                        <el-tooltip class="item" effect="dark" content="内存状态" placement="top">
                            <el-button @click="handleFunction('memStats')" round>MemStats</el-button>
                        </el-tooltip>
                        <el-tooltip class="item" effect="dark" content="CpuProfile 文件下载" placement="top">
                            <el-button @click="handleFunction('cpuProfiles', true)" round>CpuProfiles</el-button>
                        </el-tooltip>
                        <el-tooltip class="item" effect="dark" content="heapProfiles 文件下载" placement="top">
                            <el-button @click="handleFunction('heapProfiles', true)" round>HeapProfiles</el-button>
                        </el-tooltip>
                        <el-tooltip class="item" effect="dark" content="binaryDump 文件下载" placement="top">
                            <el-button @click="handleFunction('binaryDump', true)" round>BinaryDump</el-button>
                        </el-tooltip>
                    </el-row>
                </div>
                <div class="content" style="margin-top: 20px;">
                    <div v-if="info != ''" style="margin-left: 10px;"><pre>{{ info }}</pre></div>
                </div>
            </el-main>
        </el-container>
    </el-container>
</div>
</body>
<!-- import Vue before Element -->
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<!-- import JavaScript -->
<script src="https://unpkg.com/element-ui/lib/index.js"></script>
<script>
    new Vue({
        el: '#app',
        data: function () {
            return {
                leftIndex: 0,
                isCollapse: true,
                currentService: "",
                currentNode: "",
                visible: false,
                serviceList: [],
                nodes: [],
                info: ''
            }
        },
        methods: {
            leftTab(index) {
                this.leftIndex = index
                console.log(index)
            },
            getNodes(name, index) {
                let self = this
                self.leftIndex = index
                console.log(self.leftIndex)
                self.nodes = []
                self.currentService = name
                axios.get('/service/nodes?name=' + name).then(function (res) {
                    if (res.status == 200) {
                        let data = res.data.data
                        for (let i = 0; i < data.length; i++) {
                            let node = {addr: data[i]}
                            self.nodes.push(node)
                        }
                    }
                }).catch(function (res) {
                    console.log(res)
                })
            },
            onClickNode(addr) {
                this.currentNode = addr
            },
            handleFunction(name, isDownload) {
                if (this.currentNode == '') {
                    this.$message({
                        showClose: true,
                        message: '请先选择需要查询的节点',
                        type: 'error'
                    });
                    return
                }
                let url = '/'
                url = url + name + "?addr=" + this.currentNode + "&service=" + this.currentService
                let self = this
                if (isDownload) {
                    axios({
                        method: 'post',
                        url: url,
                        responseType: 'arraybuffer'
                    }).then(function (res) {
                        if (res.status == 200) {
                            if (res.data.code != undefined && res.data.code != 0) {
                                self.$message({
                                    showClose: true,
                                    message: '下载文件失败：' + res.data.msg,
                                    type: 'info'
                                });
                            }
                            console.log(res.data)
                            const content = res.data
                            const blob = new Blob([content], {type: 'application/octet-stream'})
                            const fileName = self.currentService + '_' + name
                            if ('download' in document.createElement('a')) { // 非IE下载
                                const elink = document.createElement('a')
                                elink.download = fileName
                                elink.style.display = 'none'
                                elink.href = URL.createObjectURL(blob)
                                document.body.appendChild(elink)
                                elink.click()
                                URL.revokeObjectURL(elink.href) // 释放URL 对象
                                document.body.removeChild(elink)
                            } else { // IE10+下载
                                navigator.msSaveBlob(blob, fileName)
                            }
                        }
                    }).catch(function (err) {
                        console.log(err)
                    })
                } else {
                    axios.get(url).then(function (res) {
                        if (res.status == 200) {
                            self.info = res.data.data
                            console.log(self.info)
                        }
                    }).catch(function (err) {
                        console.log(err)
                    })
                }
            },
            getServerList() {
                let self = this
                axios.get("/service").then(function (res) {
                    if (res.status == 200) {
                        let data = res.data.data
                        for (let i = 0; i < data.length; i++) {
                            let server = {name: data[i]}
                            self.serviceList.push(server)
                        }
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
        },
        mounted() {
            this.getServerList()
        },
    })
</script>

<style>
    body {
        background-color: white;
    }

    .my_menu {
        width: 1024px;
        height: 60px;
        margin: 0 auto;
    }

    .my_menu span {
        color: #dfe4ed;
    }

    .function_button {
        margin-top: 20px;
    }

    #app {
        margin: -10px;
    }
</style>
</html>
`
	return tmpl
}
