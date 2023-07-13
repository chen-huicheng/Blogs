### 原理

基于 **PicGo** 实现 **typora** 插入图片自动上传 **github**

1.   typora 设置支持插入图片时**上传图片**
2.   基于 PicGo + GitHub 搭建一个图床 

### 实现

1.   新建一个GitHub 厂库  https://github.com/new

2.   创建 GitHub 鉴权 token  

     1.   https://github.com/settings/tokens  or 头像 -> setting -> Developer Settings -> Personal access tokens -> tokens 

          ![image-20230710114122138](https://raw.githubusercontent.com/chen-huicheng/ImageHub/main/typora_img/202307101248978.png)

     >   GitHub token 鉴权用于给第三方授予厂库权限。

3.   PicGo 安装配置

     1.   应用简述 https://picgo.github.io/PicGo-Doc/zh/guide/#picgo-is-here

     2.   下载地址 https://github.com/Molunerfinn/PicGo/releases 

     3.   安装成功后配置 GitHub 图床

          ![image-20230710115346305](https://raw.githubusercontent.com/chen-huicheng/ImageHub/main/typora_img/202307101248134.png)

4.   Typora 配置插入图片时上传图片，并设置上传服务为 PicGo.app

     ![image-20230710112238748](https://raw.githubusercontent.com/chen-huicheng/ImageHub/main/typora_img/202307101249592.png)

