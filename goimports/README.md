## goimports 工具
需求：项目的的导入包只需要区分官方和非官方的包；  
官方出版的 [goimports](https://github.com/golang/tools/tree/master/cmd/goimports) 对于后面导入的包存在多个空行；  
所以此项目是对官方的 goimports 进行修改，对于导入的包只分开 官方和非官方 的包。  

地址: [goimports](https://github.com/workwb/tools)

使用方式不会发生改变，通过该项目进行编译后，覆盖原来的 goimports 即可！
