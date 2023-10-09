## Auto Click插件

**简介**：自动化点击插件，选择页面 button，设定点击逻辑。

特征：

页面动态加载后根据加载的信息触发自动点击逻辑。

定时任务触发点击逻辑。



触发器：定时器，页面动态信息。

点击逻辑：顺序执行 button list





0.1.0 版本

添加按钮，通过类似 ADblock 选择屏蔽广告的形式选择按钮。

定时触发单个按钮。

官方文档。

https://developer.chrome.com/docs/extensions/mv3/

##### 扩展程序中文件的含义

manifest.json 必须且在根目录，

server worker 处理和监听浏览器事件，不直接与网页内容交互。

content scripts 在网页上下文中执行 Javascript。







