graph TD
    subgraph "（一）嵌入生成模块 (Embed Generator)"
        A[1. 获取目标平台的独立 Python 发行版] --> B[2. 使用 pip 离线安装所需的 Python 第三方库];
        B --> C[3. 利用 Go 的 //go:embed 指令将 Python 文件打包嵌入至 Go 程序中];
    end

    subgraph "（二）运行解压模块 (Runtime Extractor)"
        D[4. 程序启动时调用解释器实例化函数] --> E[5. 解压嵌入的 Python 文件至临时目录];
        E --> F[6. 检查解压目录的完整性，避免重复释放];
    end

    subgraph "（三）命令构建模块 (Command Builder)"
        G[7. 构建指向临时目录下 Python 解释器的命令] --> H[8. 设置 PYTHONPATH 环境变量];
    end

    subgraph "（四）执行调度模块 (Executor)"
        I[9. 启动子进程执行 Python 脚本或模块] --> J[10. 捕获 stdout 和 stderr 输出，并向 Go 程序返回结果];
    end

    C --> D; // Embed Generator leads to Runtime Extractor
    F --> G; // Runtime Extractor leads to Command Builder
    H --> I; // Command Builder leads to Executor
