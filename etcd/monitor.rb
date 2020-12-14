require 'etcdv3'

def init_monitor(cli)
  val = <<flag
    {
        "service_name": "monitor",
        "version": "v1.0.0",
        "env": "dev",

        "port": 8001,

        "redis": {
            "addr": "127.0.0.1:6379",
            "pwd": ""
        }
    }
flag
  cli.put('/config/monitor', val)
end

#启动
if __FILE__ == $0
  client = Etcdv3.new(endpoints: 'http://127.0.0.1:2379')

  init_monitor client
end
