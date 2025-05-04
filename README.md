# golang-csharp-bridge

## 项目简介

通过CGO桥接golang和C#实现热更新，服务端使用客户端的热更DLL的动态加载和卸载实现热更新功能。

## 安装部署

1. 编译csharp项目需要安装 .Net 9 环境，通过Native AOT 编译出CalculatorV1.dll和CalculatorV2.dll，分别代表热更前的代码和热更后的代码，编译命令如何：
    
    ```bash
    dotnet publish -c Release -r win-x64 --self-contained
    ```
    
2. 验证编译出的CalculatorV1.dll和CalculatorV2.dll的有效性，通过命令查看DLL符号：
    
    ```bash
    dumpbin /EXPORTS CalculatorV1.dll
    ```
    
3. 将编译出的CalculatorV1.dll和CalculatorV2.dll拷贝到golang项目的根目录，通过如何命令编译运行，分别输出热更前和热更后的计算结果：
    
    ```bash
    go build main.go -o main.exe
    .\main.exe
    ```