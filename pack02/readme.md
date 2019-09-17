### 定义go程序package

属于同一个package的go代码程序们中声明的对外内容，都能被package调用  

同级文件夹下的go代码文件为同一个package  

同级文件夹下的go代码文件必须声明同一个package

同级文件夹下的go代码文件因为机制的原因，在同一个packege，他们的对外内容能互相引用而不需要import

即这样定义了能被其他程序import的packege
