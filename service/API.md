####API列表

#####搜索接口
请求地址:`POST /search`<br>
参数详情:`Content-Type: application/json
Content-Length: 78
{
	"name":"golang",
	"language":"go",
	"description":"go语言",
	"stars":"1000"
}`<br>
返回结果:`{
        "Code": 200,
        "Data": {
          "SearchItems": [
            {
              "Name": "golang-china/gopl-zh",
              "Description": "📚 Go语言圣经中文版",
              "Starts": 3400,
              "Labels": null,
              "Language": "HTML"
            },
            {
              "Name": "studygolang/studygolang",
              "Description": "Go 语言中文网 | Golang中文社区 | Go语言学习园地 源码",
              "Starts": 1700,
              "Labels": [
                "studygolang",
                "go",
                "golang"
              ],
              "Language": "Go"
            },
            {
              "Name": "gopl-zh/gopl-zh.github.com",
              "Description": "Go语言圣经中文版(只接收PR, Issue请提交到golang-china/gopl-zh)",
              "Starts": 1900,
              "Labels": null,
              "Language": "Go"
            }
          ],
          "PageData": {
            "Total": 3,
            "TotalPage": 100,
            "CurrentPage": 0,
            "PageSize": 10
          }
        }
      }`
      
#####菜单接口
请求地址:`GET /menu `<br>
参数详情:`无`<br>
返回结果:`{
        "Code": 200,
        "Data": {
          "unitData": [
            {
              "unitEN": "name",
              "unitCN": "名称",
              "unitValue": [
                "golang",
                "php",
                "java"
              ]
            },
            {
              "unitEN": "readme",
              "unitCN": "说明",
              "unitValue": [
                "a"
              ]
            }
          ]
        }
      }`