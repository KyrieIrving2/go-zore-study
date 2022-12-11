#！bin/sh
# target: dsn://law_case_platform:N8fELMmZ@tcp(9.134.39.99:18004)/law_case_platform?charset=utf8mb4&parseTime=True&loc=Local #请求服务地址

# 配置数据库
host=9.134.39.99
port=18004
user=law_case_platform
password=N8fELMmZ
database=law_case_platform

# 生成表
tables=lawyer,law_firm

# 生成脚本
echo "开始生成库：$database 的表：$tables"
goctl model mysql datasource -url="${user}:${password}@tcp(${host}:${port})/${database}" -table="${tables}"  -dir="./model"
