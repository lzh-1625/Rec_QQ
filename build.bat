@echo  off

chcp 65001
setlocal enabledelayedexpansion
set content=
for /f "tokens=*" %%c in (version.txt) do (set content=!content! %%c)

set operate=%1
set tagVersion=%2

if "%operate%" == "" (
    echo 构建代码
    set CGO_ENABLED=0
    set GOOS=linux
    set GOARCH=amd64
    make buildLinux
) else if "%operate%" == "-a" (
    if "%tagVersion%" == "" (
        echo 请输入版本号
    ) else (
        echo 打tag并推送到远程仓库
        git tag -a %tagVersion% -m "%content%"
        git push https://gitee.com/nnzn/simulation-server %tagVersion%
    )
) else if "%operate%" == "-d" (
    if "%tagVersion%" == "" (
        echo 请输入版本号
    ) else (
        echo 执行删除tag的操作并删除远程仓库的tag
        git tag -d %tagVersion%
        git push https://gitee.com/nnzn/simulation-server --delete %tagVersion%
    )
) else if "%operate%" == "-h" (
    echo -b 表示构建代码，如果带上版本号，可以新增tag版本
    echo -a 表示新增tag版本，要带上版本号，请在version.txt写入版本修改信息
    echo -d 表示删除tag版本，要带上版本号
) else (
    echo 请输入-h，查看命令帮助
)