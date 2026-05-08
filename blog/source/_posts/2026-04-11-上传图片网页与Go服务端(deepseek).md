---
title: 上传图片网页与Go服务端(deepseek)
date: 2026-04-11 10:00:00
tags:
- DeepSeek
- web
- Gin
- Go
categories:
- 学习【资料整理记录】

---

## 部署最终版本项目地址：[点击跳转](http://154.58.233.77:8000/ "上传任意文件到服务器")

## DeepSeek 原对话内容
```
帮我生成一个可以上传图片的网页界面和对应的服务端程序，服务端语言使用GO+gin框架实现
以下是您所需的上传图片网页界面及Go+Gin服务端程序的完整实现。

html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=yes">
    <title>图片上传工具 - Go Gin 服务</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: 'Segoe UI', 'Roboto', 'Helvetica Neue', sans-serif;
            background: linear-gradient(135deg, #f5f7fa 0%, #e9edf2 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 24px;
        }

        .upload-container {
            max-width: 680px;
            width: 100%;
            background: rgba(255, 255, 255, 0.95);
            border-radius: 32px;
            box-shadow: 0 20px 35px -12px rgba(0, 0, 0, 0.2), 0 1px 3px rgba(0, 0, 0, 0.05);
            backdrop-filter: blur(2px);
            overflow: hidden;
            transition: all 0.2s ease;
        }

        .header {
            background: #2c3e66;
            padding: 24px 28px;
            color: white;
        }

        .header h1 {
            font-size: 1.9rem;
            font-weight: 600;
            letter-spacing: -0.3px;
            margin-bottom: 8px;
        }

        .header p {
            opacity: 0.85;
            font-size: 0.95rem;
        }

        .card-body {
            padding: 32px 28px;
        }

        .upload-area {
            border: 2px dashed #b9c4d1;
            border-radius: 28px;
            background: #fefefe;
            padding: 40px 20px;
            text-align: center;
            cursor: pointer;
            transition: all 0.25s;
            margin-bottom: 28px;
        }

        .upload-area:hover {
            border-color: #3b6e9e;
            background: #f4f9ff;
            transform: scale(0.99);
        }

        .upload-icon {
            font-size: 48px;
            margin-bottom: 16px;
            color: #5a6e8a;
        }

        .upload-area p {
            color: #2c3e50;
            font-size: 1.1rem;
            margin: 8px 0;
        }

        .upload-area .small-hint {
            font-size: 0.8rem;
            color: #6c7a91;
            margin-top: 12px;
        }

        #fileInput {
            display: none;
        }

        .preview-section {
            background: #f8fafc;
            border-radius: 24px;
            padding: 16px;
            margin: 20px 0 20px;
            display: none;
            border: 1px solid #e2e8f0;
        }

        .preview-title {
            font-weight: 500;
            margin-bottom: 12px;
            color: #1e2a3e;
            display: flex;
            align-items: center;
            gap: 8px;
        }

        .preview-img {
            max-width: 100%;
            max-height: 240px;
            border-radius: 16px;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
            object-fit: contain;
            background: #fff;
            display: block;
            margin: 0 auto;
        }

        .file-info {
            font-size: 0.85rem;
            color: #3b4b62;
            margin-top: 12px;
            word-break: break-all;
            text-align: center;
        }

        .action-buttons {
            display: flex;
            gap: 16px;
            justify-content: space-between;
            margin-top: 8px;
            flex-wrap: wrap;
        }

        .btn {
            border: none;
            padding: 12px 24px;
            border-radius: 48px;
            font-weight: 600;
            font-size: 0.95rem;
            cursor: pointer;
            transition: all 0.2s ease;
            background: #edf2f7;
            color: #1e2a3e;
            flex: 1;
            text-align: center;
        }

        .btn-primary {
            background: #2c5f8a;
            color: white;
            box-shadow: 0 2px 5px rgba(0,0,0,0.1);
        }

        .btn-primary:hover {
            background: #1f496b;
            transform: translateY(-1px);
        }

        .btn-primary:active {
            transform: translateY(1px);
        }

        .btn-outline {
            background: transparent;
            border: 1px solid #cbd5e1;
        }

        .btn-outline:hover {
            background: #f1f5f9;
            border-color: #94a3b8;
        }

        .status-message {
            margin-top: 24px;
            padding: 14px 18px;
            border-radius: 20px;
            background: #f1f5f9;
            font-size: 0.9rem;
            border-left: 5px solid #3b6e9e;
            transition: 0.2s;
        }

        .status-success {
            background: #e0f2e9;
            border-left-color: #2c7a4d;
            color: #1e4620;
        }

        .status-error {
            background: #fee9e6;
            border-left-color: #c2412c;
            color: #9b2c1d;
        }

        .status-loading {
            background: #eef2ff;
            border-left-color: #3b6e9e;
        }

        .footer {
            background: #f1f5f9;
            padding: 16px 28px;
            font-size: 0.75rem;
            color: #5c6e82;
            text-align: center;
            border-top: 1px solid #e2edf7;
        }

        @media (max-width: 540px) {
            .card-body {
                padding: 24px 18px;
            }
            .btn {
                padding: 10px 16px;
            }
            .header h1 {
                font-size: 1.5rem;
            }
        }
    </style>
</head>
<body>
<div class="upload-container">
    <div class="header">
        <h1>📸 图片上传服务</h1>
        <p>基于 Go + Gin 的高效上传 | 支持 JPG / PNG / GIF / WEBP</p>
    </div>
    <div class="card-body">
        <!-- 拖拽上传区域 -->
        <div class="upload-area" id="dropArea">
            <div class="upload-icon">🖼️⬆️</div>
            <p>点击或拖拽图片至此区域</p>
            <p class="small-hint">支持 .jpg, .png, .gif, .webp 格式，最大 10MB</p>
            <input type="file" id="fileInput" accept="image/jpeg,image/png,image/gif,image/webp">
        </div>

        <!-- 预览区 -->
        <div class="preview-section" id="previewSection">
            <div class="preview-title">
                <span>✨ 图片预览</span>
            </div>
            <img id="previewImage" class="preview-img" alt="预览图">
            <div class="file-info" id="fileInfo"></div>
        </div>

        <div class="action-buttons">
            <button class="btn btn-outline" id="clearBtn">清空</button>
            <button class="btn btn-primary" id="uploadBtn">📤 上传到服务器</button>
        </div>

        <div id="statusMessage" class="status-message">
            💡 等待操作，选择图片后即可上传
        </div>
    </div>
    <div class="footer">
        ⚡ Gin 后端 · 图片存储于 ./uploads 目录
    </div>
</div>

<script>
    (function() {
        // DOM 元素
        const dropArea = document.getElementById('dropArea');
        const fileInput = document.getElementById('fileInput');
        const previewSection = document.getElementById('previewSection');
        const previewImage = document.getElementById('previewImage');
        const fileInfo = document.getElementById('fileInfo');
        const clearBtn = document.getElementById('clearBtn');
        const uploadBtn = document.getElementById('uploadBtn');
        const statusDiv = document.getElementById('statusMessage');

        let currentFile = null;   // 存储当前选中的文件对象

        // 显示状态消息
        function setStatus(type, text) {
            statusDiv.innerHTML = text;
            statusDiv.className = 'status-message';
            if (type === 'success') {
                statusDiv.classList.add('status-success');
            } else if (type === 'error') {
                statusDiv.classList.add('status-error');
            } else if (type === 'loading') {
                statusDiv.classList.add('status-loading');
            } else {
                statusDiv.classList.remove('status-success', 'status-error', 'status-loading');
            }
        }

        // 清空所有选择和预览
        function clearAll() {
            currentFile = null;
            fileInput.value = '';
            previewSection.style.display = 'none';
            previewImage.src = '';
            fileInfo.innerText = '';
            setStatus('info', '🗑️ 已清空，可以重新选择图片');
        }

        // 渲染预览区
        function previewFile(file) {
            if (!file) {
                previewSection.style.display = 'none';
                return;
            }

            // 检查类型
            if (!file.type.startsWith('image/')) {
                setStatus('error', '❌ 无效文件类型，请选择图片文件 (jpg, png, gif, webp)');
                clearAll();
                return false;
            }

            // 大小限制 10MB
            const maxSize = 10 * 1024 * 1024;
            if (file.size > maxSize) {
                setStatus('error', '⚠️ 图片超过 10MB，请压缩后再上传');
                clearAll();
                return false;
            }

            const reader = new FileReader();
            reader.onload = function(e) {
                previewImage.src = e.target.result;
                const sizeKB = (file.size / 1024).toFixed(1);
                const sizeMB = (file.size / (1024 * 1024)).toFixed(2);
                const displaySize = file.size > 1024 * 1024 ? `${sizeMB} MB` : `${sizeKB} KB`;
                fileInfo.innerText = `文件名: ${file.name}  |  大小: ${displaySize}  |  类型: ${file.type}`;
                previewSection.style.display = 'block';
                setStatus('info', '✅ 图片已加载，可以上传');
            };
            reader.onerror = () => {
                setStatus('error', '读取图片失败，请重试');
                clearAll();
            };
            reader.readAsDataURL(file);
            return true;
        }

        // 处理选择文件
        function handleFiles(files) {
            if (!files || files.length === 0) return;
            const file = files[0];
            if (!file.type.startsWith('image/')) {
                setStatus('error', '❌ 请选择有效的图片文件 (jpg/png/gif/webp)');
                return;
            }
            currentFile = file;
            previewFile(currentFile);
        }

        // 触发文件选择
        dropArea.addEventListener('click', () => {
            fileInput.click();
        });

        fileInput.addEventListener('change', (e) => {
            if (e.target.files && e.target.files.length > 0) {
                handleFiles(e.target.files);
            }
        });

        // 拖拽事件
        dropArea.addEventListener('dragover', (e) => {
            e.preventDefault();
            dropArea.style.borderColor = '#2c5f8a';
            dropArea.style.backgroundColor = '#f0f7fe';
        });

        dropArea.addEventListener('dragleave', (e) => {
            e.preventDefault();
            dropArea.style.borderColor = '#b9c4d1';
            dropArea.style.backgroundColor = '#fefefe';
        });

        dropArea.addEventListener('drop', (e) => {
            e.preventDefault();
            dropArea.style.borderColor = '#b9c4d1';
            dropArea.style.backgroundColor = '#fefefe';
            const dt = e.dataTransfer;
            if (dt.files && dt.files.length) {
                handleFiles(dt.files);
            } else {
                setStatus('error', '拖拽无效，请直接选择图片');
            }
        });

        // 清空按钮
        clearBtn.addEventListener('click', () => {
            clearAll();
        });

        // 上传按钮: 发送至 Gin 后端
        uploadBtn.addEventListener('click', async () => {
            if (!currentFile) {
                setStatus('error', '⚠️ 没有选择图片，请先选择或拖拽图片');
                return;
            }

            // 再次检查大小和类型
            if (!currentFile.type.startsWith('image/')) {
                setStatus('error', '无效图片类型');
                return;
            }
            if (currentFile.size > 10 * 1024 * 1024) {
                setStatus('error', '图片超过10MB限制');
                return;
            }

            const formData = new FormData();
            formData.append('image', currentFile);

            setStatus('loading', '⏳ 上传中，请稍候...');
            uploadBtn.disabled = true;
            uploadBtn.textContent = '上传中...';

            try {
                // 注意: 后端API路径为 /api/upload，默认使用同源地址。若前后端不同端口请修改。
                // 此处假定前端与后端运行在同一地址（如后端也是:8080 提供静态和API）
                const response = await fetch('/api/upload', {
                    method: 'POST',
                    body: formData
                });

                const result = await response.json();

                if (response.ok && result.code === 200) {
                    let imageUrl = result.data.url;
                    if (!imageUrl.startsWith('http')) {
                        // 相对路径转绝对
                        imageUrl = imageUrl;
                    }
                    setStatus('success', `🎉 上传成功！<br> 文件: ${result.data.filename}<br> 访问地址: <a href="${imageUrl}" target="_blank" style="color:#2c7a4d;">${imageUrl}</a>`);
                    // 可选：上传成功后清空预览？保留但不自动清，让用户决定
                    // 但为避免重复上传同一个文件，保留展示，用户可手动清空
                } else {
                    setStatus('error', `❌ 上传失败: ${result.message || '未知错误'}`);
                }
            } catch (err) {
                console.error('上传异常:', err);
                setStatus('error', `❌ 网络或服务器错误: ${err.message}. 请确保后端服务运行在相同地址下。`);
            } finally {
                uploadBtn.disabled = false;
                uploadBtn.textContent = '📤 上传到服务器';
            }
        });

        // 初始化提示
        setStatus('info', '✨ 就绪，点击上方区域选择图片或拖拽图片至此');
    })();
</script>
</body>
</html>
go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
)

// 上传配置常量
const (
    maxUploadSize   = 10 << 20                 // 10MB
    uploadDir       = "./uploads"              // 存储目录
    allowedExts     = ".jpg.jpeg.png.gif.webp" // 允许扩展名
    staticRoutePath = "/uploads"               // 静态访问路由
)

// UploadResponse 定义上传返回结构
type UploadResponse struct {
    Code     int    `json:"code"`
    Message  string `json:"message"`
    Data     *FileInfo `json:"data,omitempty"`
}

type FileInfo struct {
    Filename string `json:"filename"`
    URL      string `json:"url"`
    Size     int64  `json:"size"`
}

func main() {
    // 设置 Gin 模式 (调试模式下输出详细日志)
    gin.SetMode(gin.ReleaseMode) // 生产可改为 ReleaseMode，调试用 DebugMode
    r := gin.Default()

    // 确保上传目录存在
    if err := ensureUploadDir(uploadDir); err != nil {
        log.Fatalf("创建上传目录失败: %v", err)
    }

    // 提供静态文件服务: 使上传的图片可以通过 /uploads/xxx 访问
    r.Static(staticRoutePath, uploadDir)

    // 处理前端页面 (简单返回嵌入的HTML，实际开发也可以单独文件，这里为了演示直接使用字符串或读取文件)
    // 为了代码整洁，定义一个函数返回HTML内容 (上面提供的UI界面)
    r.GET("/", serveUploadPage)
    // 上传API
    r.POST("/api/upload", handleImageUpload)

    // 可选：健康检查
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().String()})
    })

    // 启动服务
    port := ":8080"
    log.Printf("🚀 Gin 图片上传服务启动在 http://localhost%s", port)
    log.Printf("📁 上传目录: %s (绝对路径: %s)", uploadDir, getAbsPath(uploadDir))
    if err := r.Run(port); err != nil {
        log.Fatal("服务启动失败:", err)
    }
}

// serveUploadPage 返回图片上传界面HTML
func serveUploadPage(c *gin.Context) {
    // 注意: 在实际项目中可以将html单独存为文件然后使用 c.File()
    // 此处为了演示直接写HTML字符串（从上面前端代码嵌入）
    // 由于上面HTML代码较长，采用多行字符串返回
    htmlContent := getUploadHTML()
    c.Header("Content-Type", "text/html; charset=utf-8")
    c.String(http.StatusOK, htmlContent)
}

// handleImageUpload 处理图片上传
func handleImageUpload(c *gin.Context) {
    // 限制请求体大小
    c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)
    // 从表单获取文件，字段名为 'image' (和前端FormData保持一致)
    file, header, err := c.Request.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: "获取上传文件失败: " + err.Error(),
        })
        return
    }
    defer file.Close()

    // 校验文件大小 (gin默认也有，但再确认)
    if header.Size > maxUploadSize {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: fmt.Sprintf("文件超过最大限制 %dMB", maxUploadSize>>20),
        })
        return
    }

    // 获取文件名并校验扩展名
    filename := header.Filename
    ext := strings.ToLower(filepath.Ext(filename))
    if !isAllowedExtension(ext) {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: "不支持的文件类型，请上传 JPG/PNG/GIF/WEBP 格式图片",
        })
        return
    }

    // 生成安全文件名: 时间戳+随机数+原扩展名，避免重名和路径攻击
    safeFilename := generateSafeFilename(ext)
    savePath := filepath.Join(uploadDir, safeFilename)

    // 保存文件
    if err := c.SaveUploadedFile(header, savePath); err != nil {
        log.Printf("保存文件失败: %v", err)
        c.JSON(http.StatusInternalServerError, UploadResponse{
            Code:    500,
            Message: "保存图片失败，请稍后重试",
        })
        return
    }

    // 构建可访问的URL
    // 例如: /uploads/20241215_163022_abcd1234.png
    fileURL := fmt.Sprintf("%s/%s", staticRoutePath, safeFilename)

    // 返回成功信息
    c.JSON(http.StatusOK, UploadResponse{
        Code:    200,
        Message: "上传成功",
        Data: &FileInfo{
            Filename: safeFilename,
            URL:      fileURL,
            Size:     header.Size,
        },
    })
    log.Printf("上传成功: %s -> %s (大小: %d bytes)", filename, savePath, header.Size)
}

// 检查扩展名是否在白名单内
func isAllowedExtension(ext string) bool {
    ext = strings.ToLower(ext)
    allowedList := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
    for _, allowed := range allowedList {
        if ext == allowed {
            return true
        }
    }
    return false
}

// 生成唯一文件名: 时间戳_纳秒+随机后缀 (简单避免冲突)
func generateSafeFilename(ext string) string {
    // 使用纳秒时间戳 + 随机字符串，保证唯一性且安全
    timestamp := time.Now().UnixNano() // 纳秒级别
    // 附加随机字符串提高安全性 (也可以用uuid，简单方式)
    randomPart := fmt.Sprintf("%d", time.Now().UnixNano()%1000000)
    // 文件名类似: 1702645123456789_123456.png
    return fmt.Sprintf("%d_%s%s", timestamp, randomPart, ext)
}

// 确保上传目录存在
func ensureUploadDir(dir string) error {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        return os.MkdirAll(dir, 0755)
    }
    return nil
}

// 获取绝对路径，辅助日志
func getAbsPath(dir string) string {
    abs, err := filepath.Abs(dir)
    if err != nil {
        return dir
    }
    return abs
}

// getUploadHTML 返回前端html界面（上面提供的完整UI代码）
// 这里以多行字符串嵌入，保证单个文件即可运行
func getUploadHTML() string {
    return `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=yes">
    <title>图片上传工具 - Go Gin 服务</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: 'Segoe UI', 'Roboto', 'Helvetica Neue', sans-serif;
            background: linear-gradient(135deg, #f5f7fa 0%, #e9edf2 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 24px;
        }

        .upload-container {
            max-width: 680px;
            width: 100%;
            background: rgba(255, 255, 255, 0.95);
            border-radius: 32px;
            box-shadow: 0 20px 35px -12px rgba(0, 0, 0, 0.2), 0 1px 3px rgba(0, 0, 0, 0.05);
            backdrop-filter: blur(2px);
            overflow: hidden;
            transition: all 0.2s ease;
        }

        .header {
            background: #2c3e66;
            padding: 24px 28px;
            color: white;
        }

        .header h1 {
            font-size: 1.9rem;
            font-weight: 600;
            letter-spacing: -0.3px;
            margin-bottom: 8px;
        }

        .header p {
            opacity: 0.85;
            font-size: 0.95rem;
        }

        .card-body {
            padding: 32px 28px;
        }

        .upload-area {
            border: 2px dashed #b9c4d1;
            border-radius: 28px;
            background: #fefefe;
            padding: 40px 20px;
            text-align: center;
            cursor: pointer;
            transition: all 0.25s;
            margin-bottom: 28px;
        }

        .upload-area:hover {
            border-color: #3b6e9e;
            background: #f4f9ff;
            transform: scale(0.99);
        }

        .upload-icon {
            font-size: 48px;
            margin-bottom: 16px;
            color: #5a6e8a;
        }

        .upload-area p {
            color: #2c3e50;
            font-size: 1.1rem;
            margin: 8px 0;
        }

        .upload-area .small-hint {
            font-size: 0.8rem;
            color: #6c7a91;
            margin-top: 12px;
        }

        #fileInput {
            display: none;
        }

        .preview-section {
            background: #f8fafc;
            border-radius: 24px;
            padding: 16px;
            margin: 20px 0 20px;
            display: none;
            border: 1px solid #e2e8f0;
        }

        .preview-title {
            font-weight: 500;
            margin-bottom: 12px;
            color: #1e2a3e;
            display: flex;
            align-items: center;
            gap: 8px;
        }

        .preview-img {
            max-width: 100%;
            max-height: 240px;
            border-radius: 16px;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
            object-fit: contain;
            background: #fff;
            display: block;
            margin: 0 auto;
        }

        .file-info {
            font-size: 0.85rem;
            color: #3b4b62;
            margin-top: 12px;
            word-break: break-all;
            text-align: center;
        }

        .action-buttons {
            display: flex;
            gap: 16px;
            justify-content: space-between;
            margin-top: 8px;
            flex-wrap: wrap;
        }

        .btn {
            border: none;
            padding: 12px 24px;
            border-radius: 48px;
            font-weight: 600;
            font-size: 0.95rem;
            cursor: pointer;
            transition: all 0.2s ease;
            background: #edf2f7;
            color: #1e2a3e;
            flex: 1;
            text-align: center;
        }

        .btn-primary {
            background: #2c5f8a;
            color: white;
            box-shadow: 0 2px 5px rgba(0,0,0,0.1);
        }

        .btn-primary:hover {
            background: #1f496b;
            transform: translateY(-1px);
        }

        .btn-primary:active {
            transform: translateY(1px);
        }

        .btn-outline {
            background: transparent;
            border: 1px solid #cbd5e1;
        }

        .btn-outline:hover {
            background: #f1f5f9;
            border-color: #94a3b8;
        }

        .status-message {
            margin-top: 24px;
            padding: 14px 18px;
            border-radius: 20px;
            background: #f1f5f9;
            font-size: 0.9rem;
            border-left: 5px solid #3b6e9e;
            transition: 0.2s;
        }

        .status-success {
            background: #e0f2e9;
            border-left-color: #2c7a4d;
            color: #1e4620;
        }

        .status-error {
            background: #fee9e6;
            border-left-color: #c2412c;
            color: #9b2c1d;
        }

        .status-loading {
            background: #eef2ff;
            border-left-color: #3b6e9e;
        }

        .footer {
            background: #f1f5f9;
            padding: 16px 28px;
            font-size: 0.75rem;
            color: #5c6e82;
            text-align: center;
            border-top: 1px solid #e2edf7;
        }

        @media (max-width: 540px) {
            .card-body {
                padding: 24px 18px;
            }
            .btn {
                padding: 10px 16px;
            }
            .header h1 {
                font-size: 1.5rem;
            }
        }
    </style>
</head>
<body>
<div class="upload-container">
    <div class="header">
        <h1>📸 图片上传服务</h1>
        <p>基于 Go + Gin 的高效上传 | 支持 JPG / PNG / GIF / WEBP</p>
    </div>
    <div class="card-body">
        <div class="upload-area" id="dropArea">
            <div class="upload-icon">🖼️⬆️</div>
            <p>点击或拖拽图片至此区域</p>
            <p class="small-hint">支持 .jpg, .png, .gif, .webp 格式，最大 10MB</p>
            <input type="file" id="fileInput" accept="image/jpeg,image/png,image/gif,image/webp">
        </div>

        <div class="preview-section" id="previewSection">
            <div class="preview-title">
                <span>✨ 图片预览</span>
            </div>
            <img id="previewImage" class="preview-img" alt="预览图">
            <div class="file-info" id="fileInfo"></div>
        </div>

        <div class="action-buttons">
            <button class="btn btn-outline" id="clearBtn">清空</button>
            <button class="btn btn-primary" id="uploadBtn">📤 上传到服务器</button>
        </div>

        <div id="statusMessage" class="status-message">
            💡 等待操作，选择图片后即可上传
        </div>
    </div>
    <div class="footer">
        ⚡ Gin 后端 · 图片存储于 ./uploads 目录
    </div>
</div>

<script>
    (function() {
        const dropArea = document.getElementById('dropArea');
        const fileInput = document.getElementById('fileInput');
        const previewSection = document.getElementById('previewSection');
        const previewImage = document.getElementById('previewImage');
        const fileInfo = document.getElementById('fileInfo');
        const clearBtn = document.getElementById('clearBtn');
        const uploadBtn = document.getElementById('uploadBtn');
        const statusDiv = document.getElementById('statusMessage');

        let currentFile = null;

        function setStatus(type, text) {
            statusDiv.innerHTML = text;
            statusDiv.className = 'status-message';
            if (type === 'success') {
                statusDiv.classList.add('status-success');
            } else if (type === 'error') {
                statusDiv.classList.add('status-error');
            } else if (type === 'loading') {
                statusDiv.classList.add('status-loading');
            } else {
                statusDiv.classList.remove('status-success', 'status-error', 'status-loading');
            }
        }

        function clearAll() {
            currentFile = null;
            fileInput.value = '';
            previewSection.style.display = 'none';
            previewImage.src = '';
            fileInfo.innerText = '';
            setStatus('info', '🗑️ 已清空，可以重新选择图片');
        }

        function previewFile(file) {
            if (!file) {
                previewSection.style.display = 'none';
                return;
            }
            if (!file.type.startsWith('image/')) {
                setStatus('error', '❌ 无效文件类型，请选择图片文件');
                clearAll();
                return false;
            }
            const maxSize = 10 * 1024 * 1024;
            if (file.size > maxSize) {
                setStatus('error', '⚠️ 图片超过 10MB，请压缩后再上传');
                clearAll();
                return false;
            }
            const reader = new FileReader();
            reader.onload = function(e) {
                previewImage.src = e.target.result;
                const sizeKB = (file.size / 1024).toFixed(1);
                const sizeMB = (file.size / (1024 * 1024)).toFixed(2);
                const displaySize = file.size > 1024 * 1024 ? \`\${sizeMB} MB\` : \`\${sizeKB} KB\`;
                fileInfo.innerText = \`文件名: \${file.name}  |  大小: \${displaySize}  |  类型: \${file.type}\`;
                previewSection.style.display = 'block';
                setStatus('info', '✅ 图片已加载，可以上传');
            };
            reader.onerror = () => {
                setStatus('error', '读取图片失败，请重试');
                clearAll();
            };
            reader.readAsDataURL(file);
            return true;
        }

        function handleFiles(files) {
            if (!files || files.length === 0) return;
            const file = files[0];
            if (!file.type.startsWith('image/')) {
                setStatus('error', '❌ 请选择有效的图片文件');
                return;
            }
            currentFile = file;
            previewFile(currentFile);
        }

        dropArea.addEventListener('click', () => {
            fileInput.click();
        });
        fileInput.addEventListener('change', (e) => {
            if (e.target.files && e.target.files.length > 0) {
                handleFiles(e.target.files);
            }
        });
        dropArea.addEventListener('dragover', (e) => {
            e.preventDefault();
            dropArea.style.borderColor = '#2c5f8a';
            dropArea.style.backgroundColor = '#f0f7fe';
        });
        dropArea.addEventListener('dragleave', (e) => {
            e.preventDefault();
            dropArea.style.borderColor = '#b9c4d1';
            dropArea.style.backgroundColor = '#fefefe';
        });
        dropArea.addEventListener('drop', (e) => {
            e.preventDefault();
            dropArea.style.borderColor = '#b9c4d1';
            dropArea.style.backgroundColor = '#fefefe';
            if (e.dataTransfer.files && e.dataTransfer.files.length) {
                handleFiles(e.dataTransfer.files);
            } else {
                setStatus('error', '拖拽无效，请直接选择图片');
            }
        });
        clearBtn.addEventListener('click', () => {
            clearAll();
        });
        uploadBtn.addEventListener('click', async () => {
            if (!currentFile) {
                setStatus('error', '⚠️ 没有选择图片，请先选择或拖拽图片');
                return;
            }
            if (!currentFile.type.startsWith('image/')) {
                setStatus('error', '无效图片类型');
                return;
            }
            if (currentFile.size > 10 * 1024 * 1024) {
                setStatus('error', '图片超过10MB限制');
                return;
            }
            const formData = new FormData();
            formData.append('image', currentFile);
            setStatus('loading', '⏳ 上传中，请稍候...');
            uploadBtn.disabled = true;
            uploadBtn.textContent = '上传中...';
            try {
                const response = await fetch('/api/upload', {
                    method: 'POST',
                    body: formData
                });
                const result = await response.json();
                if (response.ok && result.code === 200) {
                    let imageUrl = result.data.url;
                    setStatus('success', \`🎉 上传成功！<br> 文件: \${result.data.filename}<br> 访问地址: <a href="\${imageUrl}" target="_blank" style="color:#2c7a4d;">\${imageUrl}</a>\`);
                } else {
                    setStatus('error', \`❌ 上传失败: \${result.message || '未知错误'}\`);
                }
            } catch (err) {
                console.error('上传异常:', err);
                setStatus('error', \`❌ 网络或服务器错误: \${err.message}. 请确保后端服务运行在相同地址下。\`);
            } finally {
                uploadBtn.disabled = false;
                uploadBtn.textContent = '📤 上传到服务器';
            }
        });
        setStatus('info', '✨ 就绪，点击上方区域选择图片或拖拽图片至此');
    })();
</script>
</body>
</html>`
}
图片上传服务部署与使用
您可以通过以下步骤快速启动并体验这个图片上传工具。

服务端运行：将提供的Go代码保存为 main.go，在终端执行 go run main.go。服务默认会在 http://localhost:8080 启动，并自动创建 ./uploads 文件夹用于存储图片。

前端交互：启动服务后，用浏览器访问 http://localhost:8080 即可看到上传界面。您可以通过点击或拖拽的方式选择图片（支持JPG、PNG、GIF、WEBP格式，最大10MB），图片会实时预览。

上传流程：选择图片后点击“上传到服务器”，前端会将图片通过 FormData 发送到后端的 /api/upload 接口。后端会校验文件类型和大小，并为文件生成唯一名称保存到 uploads 目录。

结果反馈：上传成功后，页面会显示服务器返回的图片访问URL，您可以直接点击查看。所有上传的图片都会通过 /uploads/ 路由对外提供静态访问。

优化建议： 您可以根据需要调整 maxUploadSize（最大文件大小）和 allowedExts（允许的文件类型）常量。生产环境下，建议将前端HTML代码分离为独立文件，并通过 c.File() 提供，而不是嵌入在Go代码中。

对应实现一个查看预览已上传图片列表的接口
以下是增强后的服务端程序，新增了查看已上传图片列表的接口，并对前端界面进行了扩展以展示图片列表。

go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "sort"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
)

// 上传配置常量
const (
    maxUploadSize   = 10 << 20                 // 10MB
    uploadDir       = "./uploads"              // 存储目录
    allowedExts     = ".jpg.jpeg.png.gif.webp" // 允许扩展名
    staticRoutePath = "/uploads"               // 静态访问路由
)

// UploadResponse 定义上传返回结构
type UploadResponse struct {
    Code    int       `json:"code"`
    Message string    `json:"message"`
    Data    *FileInfo `json:"data,omitempty"`
}

// FileInfo 文件信息结构
type FileInfo struct {
    Filename  string    `json:"filename"`
    URL       string    `json:"url"`
    Size      int64     `json:"size"`
    SizeKB    float64   `json:"size_kb"`
    SizeMB    float64   `json:"size_mb"`
    UploadTime time.Time `json:"upload_time"`
    Ext       string    `json:"ext"`
}

// ImageListResponse 图片列表返回结构
type ImageListResponse struct {
    Code    int        `json:"code"`
    Message string     `json:"message"`
    Data    *ImageListData `json:"data,omitempty"`
}

// ImageListData 图片列表数据
type ImageListData struct {
    Total   int         `json:"total"`
    Images  []*FileInfo `json:"images"`
    BaseURL string      `json:"base_url"`
}

func main() {
    // 设置 Gin 模式
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()

    // 添加CORS中间件（可选，方便跨域调试）
    r.Use(corsMiddleware())

    // 确保上传目录存在
    if err := ensureUploadDir(uploadDir); err != nil {
        log.Fatalf("创建上传目录失败: %v", err)
    }

    // 提供静态文件服务: 使上传的图片可以通过 /uploads/xxx 访问
    r.Static(staticRoutePath, uploadDir)

    // 处理前端页面
    r.GET("/", serveUploadPage)

    // 上传图片接口
    r.POST("/api/upload", handleImageUpload)

    // 获取图片列表接口
    r.GET("/api/images", handleGetImageList)

    // 删除图片接口（可选）
    r.DELETE("/api/images/:filename", handleDeleteImage)

    // 健康检查
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().String()})
    })

    // 启动服务
    port := ":8080"
    log.Printf("🚀 Gin 图片上传服务启动在 http://localhost%s", port)
    log.Printf("📁 上传目录: %s (绝对路径: %s)", uploadDir, getAbsPath(uploadDir))
    log.Printf("📋 图片列表接口: http://localhost%s/api/images", port)
    if err := r.Run(port); err != nil {
        log.Fatal("服务启动失败:", err)
    }
}

// corsMiddleware CORS中间件，允许前端跨域请求
func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

// serveUploadPage 返回图片上传界面HTML
func serveUploadPage(c *gin.Context) {
    htmlContent := getUploadHTML()
    c.Header("Content-Type", "text/html; charset=utf-8")
    c.String(http.StatusOK, htmlContent)
}

// handleImageUpload 处理图片上传
func handleImageUpload(c *gin.Context) {
    // 限制请求体大小
    c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)
    // 从表单获取文件，字段名为 'image'
    file, header, err := c.Request.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: "获取上传文件失败: " + err.Error(),
        })
        return
    }
    defer file.Close()

    // 校验文件大小
    if header.Size > maxUploadSize {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: fmt.Sprintf("文件超过最大限制 %dMB", maxUploadSize>>20),
        })
        return
    }

    // 获取文件名并校验扩展名
    filename := header.Filename
    ext := strings.ToLower(filepath.Ext(filename))
    if !isAllowedExtension(ext) {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: "不支持的文件类型，请上传 JPG/PNG/GIF/WEBP 格式图片",
        })
        return
    }

    // 生成安全文件名
    safeFilename := generateSafeFilename(ext)
    savePath := filepath.Join(uploadDir, safeFilename)

    // 保存文件
    if err := c.SaveUploadedFile(header, savePath); err != nil {
        log.Printf("保存文件失败: %v", err)
        c.JSON(http.StatusInternalServerError, UploadResponse{
            Code:    500,
            Message: "保存图片失败，请稍后重试",
        })
        return
    }

    // 构建可访问的URL
    fileURL := fmt.Sprintf("%s/%s", staticRoutePath, safeFilename)

    // 返回成功信息
    c.JSON(http.StatusOK, UploadResponse{
        Code:    200,
        Message: "上传成功",
        Data: &FileInfo{
            Filename:   safeFilename,
            URL:        fileURL,
            Size:       header.Size,
            SizeKB:     float64(header.Size) / 1024,
            SizeMB:     float64(header.Size) / (1024 * 1024),
            UploadTime: time.Now(),
            Ext:        ext,
        },
    })
    log.Printf("上传成功: %s -> %s (大小: %d bytes)", filename, savePath, header.Size)
}

// handleGetImageList 获取已上传图片列表
func handleGetImageList(c *gin.Context) {
    // 读取上传目录
    files, err := os.ReadDir(uploadDir)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ImageListResponse{
            Code:    500,
            Message: "读取图片列表失败: " + err.Error(),
        })
        return
    }

    var imageInfos []*FileInfo
    for _, file := range files {
        if file.IsDir() {
            continue
        }
        
        filename := file.Name()
        ext := strings.ToLower(filepath.Ext(filename))
        
        // 只显示图片文件
        if !isAllowedExtension(ext) {
            continue
        }
        
        // 获取文件详细信息
        filePath := filepath.Join(uploadDir, filename)
        fileInfo, err := os.Stat(filePath)
        if err != nil {
            log.Printf("获取文件信息失败 %s: %v", filename, err)
            continue
        }
        
        // 获取文件修改时间作为上传时间
        uploadTime := fileInfo.ModTime()
        
        // 构建文件信息
        imageInfo := &FileInfo{
            Filename:   filename,
            URL:        fmt.Sprintf("%s/%s", staticRoutePath, filename),
            Size:       fileInfo.Size(),
            SizeKB:     float64(fileInfo.Size()) / 1024,
            SizeMB:     float64(fileInfo.Size()) / (1024 * 1024),
            UploadTime: uploadTime,
            Ext:        ext,
        }
        imageInfos = append(imageInfos, imageInfo)
    }
    
    // 按上传时间倒序排序（最新的在前面）
    sort.Slice(imageInfos, func(i, j int) bool {
        return imageInfos[i].UploadTime.After(imageInfos[j].UploadTime)
    })
    
    c.JSON(http.StatusOK, ImageListResponse{
        Code:    200,
        Message: "获取成功",
        Data: &ImageListData{
            Total:   len(imageInfos),
            Images:  imageInfos,
            BaseURL: staticRoutePath,
        },
    })
}

// handleDeleteImage 删除图片
func handleDeleteImage(c *gin.Context) {
    filename := c.Param("filename")
    
    // 安全检查，防止路径遍历攻击
    if strings.Contains(filename, "..") || strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    400,
            "message": "非法文件名",
        })
        return
    }
    
    // 检查文件扩展名
    ext := strings.ToLower(filepath.Ext(filename))
    if !isAllowedExtension(ext) {
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    400,
            "message": "只能删除图片文件",
        })
        return
    }
    
    filePath := filepath.Join(uploadDir, filename)
    
    // 检查文件是否存在
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{
            "code":    404,
            "message": "文件不存在",
        })
        return
    }
    
    // 删除文件
    if err := os.Remove(filePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "code":    500,
            "message": "删除文件失败: " + err.Error(),
        })
        return
    }
    
    log.Printf("删除成功: %s", filename)
    c.JSON(http.StatusOK, gin.H{
        "code":    200,
        "message": "删除成功",
    })
}

// 检查扩展名是否在白名单内
func isAllowedExtension(ext string) bool {
    ext = strings.ToLower(ext)
    allowedList := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
    for _, allowed := range allowedList {
        if ext == allowed {
            return true
        }
    }
    return false
}

// 生成唯一文件名
func generateSafeFilename(ext string) string {
    timestamp := time.Now().UnixNano()
    randomPart := fmt.Sprintf("%d", time.Now().UnixNano()%1000000)
    return fmt.Sprintf("%d_%s%s", timestamp, randomPart, ext)
}

// 确保上传目录存在
func ensureUploadDir(dir string) error {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        return os.MkdirAll(dir, 0755)
    }
    return nil
}

// 获取绝对路径
func getAbsPath(dir string) string {
    abs, err := filepath.Abs(dir)
    if err != nil {
        return dir
    }
    return abs
}

// getUploadHTML 返回前端html界面（包含图片列表功能）
func getUploadHTML() string {
    return `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=yes">
    <title>图片管理 - Go Gin 服务</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: 'Segoe UI', 'Roboto', 'Helvetica Neue', sans-serif;
            background: linear-gradient(135deg, #f5f7fa 0%, #e9edf2 100%);
            min-height: 100vh;
            padding: 24px;
        }

        .app-container {
            max-width: 1400px;
            margin: 0 auto;
        }

        .upload-container {
            background: rgba(255, 255, 255, 0.95);
            border-radius: 32px;
            box-shadow: 0 20px 35px -12px rgba(0, 0, 0, 0.2);
            overflow: hidden;
            margin-bottom: 32px;
        }

        .header {
            background: #2c3e66;
            padding: 24px 28px;
            color: white;
        }

        .header h1 {
            font-size: 1.9rem;
            font-weight: 600;
            letter-spacing: -0.3px;
            margin-bottom: 8px;
        }

        .header p {
            opacity: 0.85;
            font-size: 0.95rem;
        }

        .card-body {
            padding: 32px 28px;
        }

        .upload-area {
            border: 2px dashed #b9c4d1;
            border-radius: 28px;
            background: #fefefe;
            padding: 40px 20px;
            text-align: center;
            cursor: pointer;
            transition: all 0.25s;
            margin-bottom: 28px;
        }

        .upload-area:hover {
            border-color: #3b6e9e;
            background: #f4f9ff;
            transform: scale(0.99);
        }

        .upload-icon {
            font-size: 48px;
            margin-bottom: 16px;
            color: #5a6e8a;
        }

        .upload-area p {
            color: #2c3e50;
            font-size: 1.1rem;
            margin: 8px 0;
        }

        .upload-area .small-hint {
            font-size: 0.8rem;
            color: #6c7a91;
            margin-top: 12px;
        }

        #fileInput {
            display: none;
        }

        .preview-section {
            background: #f8fafc;
            border-radius: 24px;
            padding: 16px;
            margin: 20px 0 20px;
            display: none;
            border: 1px solid #e2e8f0;
        }

        .preview-title {
            font-weight: 500;
            margin-bottom: 12px;
            color: #1e2a3e;
            display: flex;
            align-items: center;
            gap: 8px;
        }

        .preview-img {
            max-width: 100%;
            max-height: 240px;
            border-radius: 16px;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
            object-fit: contain;
            background: #fff;
            display: block;
            margin: 0 auto;
        }

        .file-info {
            font-size: 0.85rem;
            color: #3b4b62;
            margin-top: 12px;
            word-break: break-all;
            text-align: center;
        }

        .action-buttons {
            display: flex;
            gap: 16px;
            justify-content: space-between;
            margin-top: 8px;
            flex-wrap: wrap;
        }

        .btn {
            border: none;
            padding: 12px 24px;
            border-radius: 48px;
            font-weight: 600;
            font-size: 0.95rem;
            cursor: pointer;
            transition: all 0.2s ease;
            background: #edf2f7;
            color: #1e2a3e;
            flex: 1;
            text-align: center;
        }

        .btn-primary {
            background: #2c5f8a;
            color: white;
            box-shadow: 0 2px 5px rgba(0,0,0,0.1);
        }

        .btn-primary:hover {
            background: #1f496b;
            transform: translateY(-1px);
        }

        .btn-primary:active {
            transform: translateY(1px);
        }

        .btn-outline {
            background: transparent;
            border: 1px solid #cbd5e1;
        }

        .btn-outline:hover {
            background: #f1f5f9;
            border-color: #94a3b8;
        }

        .btn-danger {
            background: #dc2626;
            color: white;
        }

        .btn-danger:hover {
            background: #b91c1c;
        }

        .status-message {
            margin-top: 24px;
            padding: 14px 18px;
            border-radius: 20px;
            background: #f1f5f9;
            font-size: 0.9rem;
            border-left: 5px solid #3b6e9e;
            transition: 0.2s;
        }

        .status-success {
            background: #e0f2e9;
            border-left-color: #2c7a4d;
            color: #1e4620;
        }

        .status-error {
            background: #fee9e6;
            border-left-color: #c2412c;
            color: #9b2c1d;
        }

        .status-loading {
            background: #eef2ff;
            border-left-color: #3b6e9e;
        }

        /* 图片列表样式 */
        .gallery-container {
            background: rgba(255, 255, 255, 0.95);
            border-radius: 32px;
            box-shadow: 0 20px 35px -12px rgba(0, 0, 0, 0.2);
            overflow: hidden;
        }

        .gallery-header {
            background: #2c3e66;
            padding: 20px 28px;
            color: white;
            display: flex;
            justify-content: space-between;
            align-items: center;
            flex-wrap: wrap;
            gap: 16px;
        }

        .gallery-header h2 {
            font-size: 1.5rem;
            font-weight: 600;
        }

        .refresh-btn {
            background: rgba(255,255,255,0.2);
            border: 1px solid rgba(255,255,255,0.3);
            padding: 8px 20px;
            border-radius: 40px;
            color: white;
            cursor: pointer;
            transition: all 0.2s;
        }

        .refresh-btn:hover {
            background: rgba(255,255,255,0.3);
        }

        .image-grid {
            padding: 28px;
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
            gap: 24px;
        }

        .image-card {
            background: white;
            border-radius: 20px;
            overflow: hidden;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
            transition: transform 0.2s, box-shadow 0.2s;
            border: 1px solid #e2e8f0;
        }

        .image-card:hover {
            transform: translateY(-4px);
            box-shadow: 0 12px 24px rgba(0, 0, 0, 0.12);
        }

        .image-card-img {
            width: 100%;
            height: 200px;
            object-fit: cover;
            background: #f1f5f9;
            cursor: pointer;
        }

        .image-card-info {
            padding: 16px;
        }

        .image-filename {
            font-size: 0.85rem;
            color: #475569;
            margin-bottom: 8px;
            word-break: break-all;
            font-family: monospace;
        }

        .image-meta {
            font-size: 0.75rem;
            color: #64748b;
            margin-bottom: 12px;
        }

        .image-actions {
            display: flex;
            gap: 8px;
        }

        .image-actions .btn {
            padding: 6px 12px;
            font-size: 0.8rem;
            flex: 1;
        }

        .empty-gallery {
            text-align: center;
            padding: 60px 20px;
            color: #94a3b8;
        }

        .modal {
            display: none;
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background: rgba(0,0,0,0.9);
            z-index: 1000;
            cursor: pointer;
            align-items: center;
            justify-content: center;
        }

        .modal img {
            max-width: 90%;
            max-height: 90%;
            object-fit: contain;
        }

        .footer {
            background: #f1f5f9;
            padding: 16px 28px;
            font-size: 0.75rem;
            color: #5c6e82;
            text-align: center;
            border-top: 1px solid #e2edf7;
        }

        @media (max-width: 768px) {
            .card-body, .image-grid {
                padding: 20px;
            }
            .image-grid {
                grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
                gap: 16px;
            }
            .header h1, .gallery-header h2 {
                font-size: 1.3rem;
            }
        }
    </style>
</head>
<body>
<div class="app-container">
    <!-- 上传区域 -->
    <div class="upload-container">
        <div class="header">
            <h1>📸 图片上传服务</h1>
            <p>基于 Go + Gin 的高效上传 | 支持 JPG / PNG / GIF / WEBP</p>
        </div>
        <div class="card-body">
            <div class="upload-area" id="dropArea">
                <div class="upload-icon">🖼️⬆️</div>
                <p>点击或拖拽图片至此区域</p>
                <p class="small-hint">支持 .jpg, .png, .gif, .webp 格式，最大 10MB</p>
                <input type="file" id="fileInput" accept="image/jpeg,image/png,image/gif,image/webp">
            </div>

            <div class="preview-section" id="previewSection">
                <div class="preview-title">
                    <span>✨ 图片预览</span>
                </div>
                <img id="previewImage" class="preview-img" alt="预览图">
                <div class="file-info" id="fileInfo"></div>
            </div>

            <div class="action-buttons">
                <button class="btn btn-outline" id="clearBtn">清空</button>
                <button class="btn btn-primary" id="uploadBtn">📤 上传到服务器</button>
            </div>

            <div id="statusMessage" class="status-message">
                💡 等待操作，选择图片后即可上传
            </div>
        </div>
    </div>

    <!-- 图片列表区域 -->
    <div class="gallery-container">
        <div class="gallery-header">
            <h2>🖼️ 已上传图片列表</h2>
            <button class="refresh-btn" id="refreshBtn">🔄 刷新列表</button>
        </div>
        <div id="imageGrid" class="image-grid">
            <div class="empty-gallery">加载中...</div>
        </div>
        <div class="footer">
            ⚡ Gin 后端 · 图片存储于 ./uploads 目录 · 支持查看和删除
        </div>
    </div>
</div>

<!-- 图片预览模态框 -->
<div id="imageModal" class="modal" onclick="closeModal()">
    <img id="modalImage" alt="大图预览">
</div>

<script>
    // DOM 元素
    const dropArea = document.getElementById('dropArea');
    const fileInput = document.getElementById('fileInput');
    const previewSection = document.getElementById('previewSection');
    const previewImage = document.getElementById('previewImage');
    const fileInfo = document.getElementById('fileInfo');
    const clearBtn = document.getElementById('clearBtn');
    const uploadBtn = document.getElementById('uploadBtn');
    const statusDiv = document.getElementById('statusMessage');
    const refreshBtn = document.getElementById('refreshBtn');
    const imageGrid = document.getElementById('imageGrid');

    let currentFile = null;

    // 显示状态消息
    function setStatus(type, text) {
        statusDiv.innerHTML = text;
        statusDiv.className = 'status-message';
        if (type === 'success') {
            statusDiv.classList.add('status-success');
        } else if (type === 'error') {
            statusDiv.classList.add('status-error');
        } else if (type === 'loading') {
            statusDiv.classList.add('status-loading');
        } else {
            statusDiv.classList.remove('status-success', 'status-error', 'status-loading');
        }
    }

    // 清空所有选择和预览
    function clearAll() {
        currentFile = null;
        fileInput.value = '';
        previewSection.style.display = 'none';
        previewImage.src = '';
        fileInfo.innerText = '';
        setStatus('info', '🗑️ 已清空，可以重新选择图片');
    }

    // 预览文件
    function previewFile(file) {
        if (!file) {
            previewSection.style.display = 'none';
            return;
        }

        if (!file.type.startsWith('image/')) {
            setStatus('error', '❌ 无效文件类型，请选择图片文件');
            clearAll();
            return false;
        }

        const maxSize = 10 * 1024 * 1024;
        if (file.size > maxSize) {
            setStatus('error', '⚠️ 图片超过 10MB，请压缩后再上传');
            clearAll();
            return false;
        }

        const reader = new FileReader();
        reader.onload = function(e) {
            previewImage.src = e.target.result;
            const sizeKB = (file.size / 1024).toFixed(1);
            const sizeMB = (file.size / (1024 * 1024)).toFixed(2);
            const displaySize = file.size > 1024 * 1024 ? `${sizeMB} MB` : `${sizeKB} KB`;
            fileInfo.innerText = `文件名: ${file.name}  |  大小: ${displaySize}  |  类型: ${file.type}`;
            previewSection.style.display = 'block';
            setStatus('info', '✅ 图片已加载，可以上传');
        };
        reader.onerror = () => {
            setStatus('error', '读取图片失败，请重试');
            clearAll();
        };
        reader.readAsDataURL(file);
        return true;
    }

    // 处理文件选择
    function handleFiles(files) {
        if (!files || files.length === 0) return;
        const file = files[0];
        if (!file.type.startsWith('image/')) {
            setStatus('error', '❌ 请选择有效的图片文件');
            return;
        }
        currentFile = file;
        previewFile(currentFile);
    }

    // 获取图片列表
    async function loadImageList() {
        imageGrid.innerHTML = '<div class="empty-gallery">加载中...</div>';
        try {
            const response = await fetch('/api/images');
            const result = await response.json();
            
            if (result.code === 200 && result.data && result.data.images) {
                const images = result.data.images;
                if (images.length === 0) {
                    imageGrid.innerHTML = '<div class="empty-gallery">📭 暂无图片，上传第一张吧</div>';
                    return;
                }
                
                imageGrid.innerHTML = '';
                images.forEach(img => {
                    const card = createImageCard(img);
                    imageGrid.appendChild(card);
                });
            } else {
                imageGrid.innerHTML = '<div class="empty-gallery">❌ 加载失败: ' + (result.message || '未知错误') + '</div>';
            }
        } catch (err) {
            console.error('加载图片列表失败:', err);
            imageGrid.innerHTML = '<div class="empty-gallery">❌ 网络错误，无法加载图片列表</div>';
        }
    }

    // 创建图片卡片
    function createImageCard(img) {
        const card = document.createElement('div');
        card.className = 'image-card';
        
        const uploadDate = new Date(img.upload_time).toLocaleString('zh-CN');
        const sizeDisplay = img.size_mb > 1 ? `${img.size_mb.toFixed(2)} MB` : `${img.size_kb.toFixed(1)} KB`;
        
        card.innerHTML = `
            <img class="image-card-img" src="${img.url}" alt="${img.filename}" onclick="viewImage('${img.url}')">
            <div class="image-card-info">
                <div class="image-filename" title="${img.filename}">${img.filename.substring(0, 30)}${img.filename.length > 30 ? '...' : ''}</div>
                <div class="image-meta">
                    📅 ${uploadDate}<br>
                    📦 ${sizeDisplay} · ${img.ext.toUpperCase()}
                </div>
                <div class="image-actions">
                    <button class="btn btn-outline" onclick="viewImage('${img.url}')">🔍 查看</button>
                    <button class="btn btn-danger" onclick="deleteImage('${img.filename}')">🗑️ 删除</button>
                </div>
            </div>
        `;
        
        return card;
    }

    // 查看大图
    function viewImage(url) {
        const modal = document.getElementById('imageModal');
        const modalImg = document.getElementById('modalImage');
        modal.style.display = 'flex';
        modalImg.src = url;
    }

    // 关闭模态框
    function closeModal() {
        const modal = document.getElementById('imageModal');
        modal.style.display = 'none';
    }

    // 删除图片
    async function deleteImage(filename) {
        if (!confirm(`确定要删除图片 "${filename}" 吗？`)) {
            return;
        }
        
        try {
            const response = await fetch(`/api/images/${encodeURIComponent(filename)}`, {
                method: 'DELETE'
            });
            const result = await response.json();
            
            if (result.code === 200) {
                setStatus('success', `✅ 删除成功: ${filename}`);
                loadImageList(); // 刷新列表
            } else {
                setStatus('error', `❌ 删除失败: ${result.message}`);
            }
        } catch (err) {
            console.error('删除失败:', err);
            setStatus('error', '❌ 网络错误，删除失败');
        }
    }

    // 上传图片
    async function uploadImage() {
        if (!currentFile) {
            setStatus('error', '⚠️ 没有选择图片，请先选择或拖拽图片');
            return;
        }

        if (!currentFile.type.startsWith('image/')) {
            setStatus('error', '无效图片类型');
            return;
        }
        
        if (currentFile.size > 10 * 1024 * 1024) {
            setStatus('error', '图片超过10MB限制');
            return;
        }

        const formData = new FormData();
        formData.append('image', currentFile);

        setStatus('loading', '⏳ 上传中，请稍候...');
        uploadBtn.disabled = true;
        uploadBtn.textContent = '上传中...';

        try {
            const response = await fetch('/api/upload', {
                method: 'POST',
                body: formData
            });
            const result = await response.json();

            if (response.ok && result.code === 200) {
                let imageUrl = result.data.url;
                setStatus('success', `🎉 上传成功！<br> 文件: ${result.data.filename}<br> 访问地址: <a href="${imageUrl}" target="_blank" style="color:#2c7a4d;">${imageUrl}</a>`);
                clearAll();
                loadImageList(); // 刷新图片列表
            } else {
                setStatus('error', `❌ 上传失败: ${result.message || '未知错误'}`);
            }
        } catch (err) {
            console.error('上传异常:', err);
            setStatus('error', `❌ 网络或服务器错误: ${err.message}`);
        } finally {
            uploadBtn.disabled = false;
            uploadBtn.textContent = '📤 上传到服务器';
        }
    }

    // 事件监听
    dropArea.addEventListener('click', () => {
        fileInput.click();
    });

    fileInput.addEventListener('change', (e) => {
        if (e.target.files && e.target.files.length > 0) {
            handleFiles(e.target.files);
        }
    });

    dropArea.addEventListener('dragover', (e) => {
        e.preventDefault();
        dropArea.style.borderColor = '#2c5f8a';
        dropArea.style.backgroundColor = '#f0f7fe';
    });

    dropArea.addEventListener('dragleave', (e) => {
        e.preventDefault();
        dropArea.style.borderColor = '#b9c4d1';
        dropArea.style.backgroundColor = '#fefefe';
    });

    dropArea.addEventListener('drop', (e) => {
        e.preventDefault();
        dropArea.style.borderColor = '#b9c4d1';
        dropArea.style.backgroundColor = '#fefefe';
        if (e.dataTransfer.files && e.dataTransfer.files.length) {
            handleFiles(e.dataTransfer.files);
        } else {
            setStatus('error', '拖拽无效，请直接选择图片');
        }
    });

    clearBtn.addEventListener('click', () => {
        clearAll();
    });

    uploadBtn.addEventListener('click', uploadImage);
    refreshBtn.addEventListener('click', loadImageList);

    // 键盘ESC关闭模态框
    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape') {
            closeModal();
        }
    });

    // 初始化加载图片列表
    loadImageList();
    setStatus('info', '✨ 就绪，点击上方区域选择图片或拖拽图片至此');
</script>
</body>
</html>`
}
新增功能说明
现在您可以通过以下接口管理已上传的图片：

接口  方法  说明
/api/images GET 获取所有已上传图片的列表，按上传时间倒序排列
/api/images/:filename   DELETE  删除指定图片文件
前端界面增强：

图片列表区域：页面下方会展示所有已上传的图片，以网格形式呈现

图片卡片：每个图片显示缩略图、文件名、上传时间、文件大小和格式

查看大图：点击图片或"查看"按钮可全屏预览

删除功能：每个图片卡片配有删除按钮，支持删除操作

自动刷新：上传成功后自动刷新列表，也可手动点击刷新按钮

运行方式不变：执行 go run main.go 后访问 http://localhost:8080 即可体验完整功能。

本回答由 AI 生成，内容仅供参考，请仔细甄别。
修复 go 文件`报错
我来修复 Go 文件中的问题。主要问题是 HTML 字符串中的反引号冲突导致的语法错误。以下是修复后的代码：

go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "sort"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
)

// 上传配置常量
const (
    maxUploadSize   = 10 << 20                 // 10MB
    uploadDir       = "./uploads"              // 存储目录
    allowedExts     = ".jpg.jpeg.png.gif.webp" // 允许扩展名
    staticRoutePath = "/uploads"               // 静态访问路由
)

// UploadResponse 定义上传返回结构
type UploadResponse struct {
    Code    int       `json:"code"`
    Message string    `json:"message"`
    Data    *FileInfo `json:"data,omitempty"`
}

// FileInfo 文件信息结构
type FileInfo struct {
    Filename   string    `json:"filename"`
    URL        string    `json:"url"`
    Size       int64     `json:"size"`
    SizeKB     float64   `json:"size_kb"`
    SizeMB     float64   `json:"size_mb"`
    UploadTime time.Time `json:"upload_time"`
    Ext        string    `json:"ext"`
}

// ImageListResponse 图片列表返回结构
type ImageListResponse struct {
    Code    int            `json:"code"`
    Message string         `json:"message"`
    Data    *ImageListData `json:"data,omitempty"`
}

// ImageListData 图片列表数据
type ImageListData struct {
    Total   int         `json:"total"`
    Images  []*FileInfo `json:"images"`
    BaseURL string      `json:"base_url"`
}

func main() {
    // 设置 Gin 模式
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()

    // 添加CORS中间件（可选，方便跨域调试）
    r.Use(corsMiddleware())

    // 确保上传目录存在
    if err := ensureUploadDir(uploadDir); err != nil {
        log.Fatalf("创建上传目录失败: %v", err)
    }

    // 提供静态文件服务: 使上传的图片可以通过 /uploads/xxx 访问
    r.Static(staticRoutePath, uploadDir)

    // 处理前端页面
    r.GET("/", serveUploadPage)

    // 上传图片接口
    r.POST("/api/upload", handleImageUpload)

    // 获取图片列表接口
    r.GET("/api/images", handleGetImageList)

    // 删除图片接口（可选）
    r.DELETE("/api/images/:filename", handleDeleteImage)

    // 健康检查
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().String()})
    })

    // 启动服务
    port := ":8080"
    log.Printf("🚀 Gin 图片上传服务启动在 http://localhost%s", port)
    log.Printf("📁 上传目录: %s (绝对路径: %s)", uploadDir, getAbsPath(uploadDir))
    log.Printf("📋 图片列表接口: http://localhost%s/api/images", port)
    if err := r.Run(port); err != nil {
        log.Fatal("服务启动失败:", err)
    }
}

// corsMiddleware CORS中间件，允许前端跨域请求
func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

// serveUploadPage 返回图片上传界面HTML
func serveUploadPage(c *gin.Context) {
    htmlContent := getUploadHTML()
    c.Header("Content-Type", "text/html; charset=utf-8")
    c.String(http.StatusOK, htmlContent)
}

// handleImageUpload 处理图片上传
func handleImageUpload(c *gin.Context) {
    // 限制请求体大小
    c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)
    // 从表单获取文件，字段名为 'image'
    file, header, err := c.Request.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: "获取上传文件失败: " + err.Error(),
        })
        return
    }
    defer file.Close()

    // 校验文件大小
    if header.Size > maxUploadSize {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: fmt.Sprintf("文件超过最大限制 %dMB", maxUploadSize>>20),
        })
        return
    }

    // 获取文件名并校验扩展名
    filename := header.Filename
    ext := strings.ToLower(filepath.Ext(filename))
    if !isAllowedExtension(ext) {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: "不支持的文件类型，请上传 JPG/PNG/GIF/WEBP 格式图片",
        })
        return
    }

    // 生成安全文件名
    safeFilename := generateSafeFilename(ext)
    savePath := filepath.Join(uploadDir, safeFilename)

    // 保存文件
    if err := c.SaveUploadedFile(header, savePath); err != nil {
        log.Printf("保存文件失败: %v", err)
        c.JSON(http.StatusInternalServerError, UploadResponse{
            Code:    500,
            Message: "保存图片失败，请稍后重试",
        })
        return
    }

    // 构建可访问的URL
    fileURL := fmt.Sprintf("%s/%s", staticRoutePath, safeFilename)

    // 返回成功信息
    c.JSON(http.StatusOK, UploadResponse{
        Code:    200,
        Message: "上传成功",
        Data: &FileInfo{
            Filename:   safeFilename,
            URL:        fileURL,
            Size:       header.Size,
            SizeKB:     float64(header.Size) / 1024,
            SizeMB:     float64(header.Size) / (1024 * 1024),
            UploadTime: time.Now(),
            Ext:        ext,
        },
    })
    log.Printf("上传成功: %s -> %s (大小: %d bytes)", filename, savePath, header.Size)
}

// handleGetImageList 获取已上传图片列表
func handleGetImageList(c *gin.Context) {
    // 读取上传目录
    files, err := os.ReadDir(uploadDir)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ImageListResponse{
            Code:    500,
            Message: "读取图片列表失败: " + err.Error(),
        })
        return
    }

    var imageInfos []*FileInfo
    for _, file := range files {
        if file.IsDir() {
            continue
        }

        filename := file.Name()
        ext := strings.ToLower(filepath.Ext(filename))

        // 只显示图片文件
        if !isAllowedExtension(ext) {
            continue
        }

        // 获取文件详细信息
        filePath := filepath.Join(uploadDir, filename)
        fileInfo, err := os.Stat(filePath)
        if err != nil {
            log.Printf("获取文件信息失败 %s: %v", filename, err)
            continue
        }

        // 获取文件修改时间作为上传时间
        uploadTime := fileInfo.ModTime()

        // 构建文件信息
        imageInfo := &FileInfo{
            Filename:   filename,
            URL:        fmt.Sprintf("%s/%s", staticRoutePath, filename),
            Size:       fileInfo.Size(),
            SizeKB:     float64(fileInfo.Size()) / 1024,
            SizeMB:     float64(fileInfo.Size()) / (1024 * 1024),
            UploadTime: uploadTime,
            Ext:        ext,
        }
        imageInfos = append(imageInfos, imageInfo)
    }

    // 按上传时间倒序排序（最新的在前面）
    sort.Slice(imageInfos, func(i, j int) bool {
        return imageInfos[i].UploadTime.After(imageInfos[j].UploadTime)
    })

    c.JSON(http.StatusOK, ImageListResponse{
        Code:    200,
        Message: "获取成功",
        Data: &ImageListData{
            Total:   len(imageInfos),
            Images:  imageInfos,
            BaseURL: staticRoutePath,
        },
    })
}

// handleDeleteImage 删除图片
func handleDeleteImage(c *gin.Context) {
    filename := c.Param("filename")

    // 安全检查，防止路径遍历攻击
    if strings.Contains(filename, "..") || strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    400,
            "message": "非法文件名",
        })
        return
    }

    // 检查文件扩展名
    ext := strings.ToLower(filepath.Ext(filename))
    if !isAllowedExtension(ext) {
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    400,
            "message": "只能删除图片文件",
        })
        return
    }

    filePath := filepath.Join(uploadDir, filename)

    // 检查文件是否存在
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{
            "code":    404,
            "message": "文件不存在",
        })
        return
    }

    // 删除文件
    if err := os.Remove(filePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "code":    500,
            "message": "删除文件失败: " + err.Error(),
        })
        return
    }

    log.Printf("删除成功: %s", filename)
    c.JSON(http.StatusOK, gin.H{
        "code":    200,
        "message": "删除成功",
    })
}

// 检查扩展名是否在白名单内
func isAllowedExtension(ext string) bool {
    ext = strings.ToLower(ext)
    allowedList := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
    for _, allowed := range allowedList {
        if ext == allowed {
            return true
        }
    }
    return false
}

// 生成唯一文件名
func generateSafeFilename(ext string) string {
    timestamp := time.Now().UnixNano()
    randomPart := fmt.Sprintf("%d", time.Now().UnixNano()%1000000)
    return fmt.Sprintf("%d_%s%s", timestamp, randomPart, ext)
}

// 确保上传目录存在
func ensureUploadDir(dir string) error {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        return os.MkdirAll(dir, 0755)
    }
    return nil
}

// 获取绝对路径
func getAbsPath(dir string) string {
    abs, err := filepath.Abs(dir)
    if err != nil {
        return dir
    }
    return abs
}

// getUploadHTML 返回前端html界面（包含图片列表功能）
func getUploadHTML() string {
    return `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=yes">
    <title>图片管理 - Go Gin 服务</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: 'Segoe UI', 'Roboto', 'Helvetica Neue', sans-serif;
            background: linear-gradient(135deg, #f5f7fa 0%, #e9edf2 100%);
            min-height: 100vh;
            padding: 24px;
        }

        .app-container {
            max-width: 1400px;
            margin: 0 auto;
        }

        .upload-container {
            background: rgba(255, 255, 255, 0.95);
            border-radius: 32px;
            box-shadow: 0 20px 35px -12px rgba(0, 0, 0, 0.2);
            overflow: hidden;
            margin-bottom: 32px;
        }

        .header {
            background: #2c3e66;
            padding: 24px 28px;
            color: white;
        }

        .header h1 {
            font-size: 1.9rem;
            font-weight: 600;
            letter-spacing: -0.3px;
            margin-bottom: 8px;
        }

        .header p {
            opacity: 0.85;
            font-size: 0.95rem;
        }

        .card-body {
            padding: 32px 28px;
        }

        .upload-area {
            border: 2px dashed #b9c4d1;
            border-radius: 28px;
            background: #fefefe;
            padding: 40px 20px;
            text-align: center;
            cursor: pointer;
            transition: all 0.25s;
            margin-bottom: 28px;
        }

        .upload-area:hover {
            border-color: #3b6e9e;
            background: #f4f9ff;
            transform: scale(0.99);
        }

        .upload-icon {
            font-size: 48px;
            margin-bottom: 16px;
            color: #5a6e8a;
        }

        .upload-area p {
            color: #2c3e50;
            font-size: 1.1rem;
            margin: 8px 0;
        }

        .upload-area .small-hint {
            font-size: 0.8rem;
            color: #6c7a91;
            margin-top: 12px;
        }

        #fileInput {
            display: none;
        }

        .preview-section {
            background: #f8fafc;
            border-radius: 24px;
            padding: 16px;
            margin: 20px 0 20px;
            display: none;
            border: 1px solid #e2e8f0;
        }

        .preview-title {
            font-weight: 500;
            margin-bottom: 12px;
            color: #1e2a3e;
            display: flex;
            align-items: center;
            gap: 8px;
        }

        .preview-img {
            max-width: 100%;
            max-height: 240px;
            border-radius: 16px;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
            object-fit: contain;
            background: #fff;
            display: block;
            margin: 0 auto;
        }

        .file-info {
            font-size: 0.85rem;
            color: #3b4b62;
            margin-top: 12px;
            word-break: break-all;
            text-align: center;
        }

        .action-buttons {
            display: flex;
            gap: 16px;
            justify-content: space-between;
            margin-top: 8px;
            flex-wrap: wrap;
        }

        .btn {
            border: none;
            padding: 12px 24px;
            border-radius: 48px;
            font-weight: 600;
            font-size: 0.95rem;
            cursor: pointer;
            transition: all 0.2s ease;
            background: #edf2f7;
            color: #1e2a3e;
            flex: 1;
            text-align: center;
        }

        .btn-primary {
            background: #2c5f8a;
            color: white;
            box-shadow: 0 2px 5px rgba(0,0,0,0.1);
        }

        .btn-primary:hover {
            background: #1f496b;
            transform: translateY(-1px);
        }

        .btn-primary:active {
            transform: translateY(1px);
        }

        .btn-outline {
            background: transparent;
            border: 1px solid #cbd5e1;
        }

        .btn-outline:hover {
            background: #f1f5f9;
            border-color: #94a3b8;
        }

        .btn-danger {
            background: #dc2626;
            color: white;
        }

        .btn-danger:hover {
            background: #b91c1c;
        }

        .status-message {
            margin-top: 24px;
            padding: 14px 18px;
            border-radius: 20px;
            background: #f1f5f9;
            font-size: 0.9rem;
            border-left: 5px solid #3b6e9e;
            transition: 0.2s;
        }

        .status-success {
            background: #e0f2e9;
            border-left-color: #2c7a4d;
            color: #1e4620;
        }

        .status-error {
            background: #fee9e6;
            border-left-color: #c2412c;
            color: #9b2c1d;
        }

        .status-loading {
            background: #eef2ff;
            border-left-color: #3b6e9e;
        }

        /* 图片列表样式 */
        .gallery-container {
            background: rgba(255, 255, 255, 0.95);
            border-radius: 32px;
            box-shadow: 0 20px 35px -12px rgba(0, 0, 0, 0.2);
            overflow: hidden;
        }

        .gallery-header {
            background: #2c3e66;
            padding: 20px 28px;
            color: white;
            display: flex;
            justify-content: space-between;
            align-items: center;
            flex-wrap: wrap;
            gap: 16px;
        }

        .gallery-header h2 {
            font-size: 1.5rem;
            font-weight: 600;
        }

        .refresh-btn {
            background: rgba(255,255,255,0.2);
            border: 1px solid rgba(255,255,255,0.3);
            padding: 8px 20px;
            border-radius: 40px;
            color: white;
            cursor: pointer;
            transition: all 0.2s;
        }

        .refresh-btn:hover {
            background: rgba(255,255,255,0.3);
        }

        .image-grid {
            padding: 28px;
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
            gap: 24px;
        }

        .image-card {
            background: white;
            border-radius: 20px;
            overflow: hidden;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
            transition: transform 0.2s, box-shadow 0.2s;
            border: 1px solid #e2e8f0;
        }

        .image-card:hover {
            transform: translateY(-4px);
            box-shadow: 0 12px 24px rgba(0, 0, 0, 0.12);
        }

        .image-card-img {
            width: 100%;
            height: 200px;
            object-fit: cover;
            background: #f1f5f9;
            cursor: pointer;
        }

        .image-card-info {
            padding: 16px;
        }

        .image-filename {
            font-size: 0.85rem;
            color: #475569;
            margin-bottom: 8px;
            word-break: break-all;
            font-family: monospace;
        }

        .image-meta {
            font-size: 0.75rem;
            color: #64748b;
            margin-bottom: 12px;
        }

        .image-actions {
            display: flex;
            gap: 8px;
        }

        .image-actions .btn {
            padding: 6px 12px;
            font-size: 0.8rem;
            flex: 1;
        }

        .empty-gallery {
            text-align: center;
            padding: 60px 20px;
            color: #94a3b8;
        }

        .modal {
            display: none;
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background: rgba(0,0,0,0.9);
            z-index: 1000;
            cursor: pointer;
            align-items: center;
            justify-content: center;
        }

        .modal img {
            max-width: 90%;
            max-height: 90%;
            object-fit: contain;
        }

        .footer {
            background: #f1f5f9;
            padding: 16px 28px;
            font-size: 0.75rem;
            color: #5c6e82;
            text-align: center;
            border-top: 1px solid #e2edf7;
        }

        @media (max-width: 768px) {
            .card-body, .image-grid {
                padding: 20px;
            }
            .image-grid {
                grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
                gap: 16px;
            }
            .header h1, .gallery-header h2 {
                font-size: 1.3rem;
            }
        }
    </style>
</head>
<body>
<div class="app-container">
    <!-- 上传区域 -->
    <div class="upload-container">
        <div class="header">
            <h1>📸 图片上传服务</h1>
            <p>基于 Go + Gin 的高效上传 | 支持 JPG / PNG / GIF / WEBP</p>
        </div>
        <div class="card-body">
            <div class="upload-area" id="dropArea">
                <div class="upload-icon">🖼️⬆️</div>
                <p>点击或拖拽图片至此区域</p>
                <p class="small-hint">支持 .jpg, .png, .gif, .webp 格式，最大 10MB</p>
                <input type="file" id="fileInput" accept="image/jpeg,image/png,image/gif,image/webp">
            </div>

            <div class="preview-section" id="previewSection">
                <div class="preview-title">
                    <span>✨ 图片预览</span>
                </div>
                <img id="previewImage" class="preview-img" alt="预览图">
                <div class="file-info" id="fileInfo"></div>
            </div>

            <div class="action-buttons">
                <button class="btn btn-outline" id="clearBtn">清空</button>
                <button class="btn btn-primary" id="uploadBtn">📤 上传到服务器</button>
            </div>

            <div id="statusMessage" class="status-message">
                💡 等待操作，选择图片后即可上传
            </div>
        </div>
    </div>

    <!-- 图片列表区域 -->
    <div class="gallery-container">
        <div class="gallery-header">
            <h2>🖼️ 已上传图片列表</h2>
            <button class="refresh-btn" id="refreshBtn">🔄 刷新列表</button>
        </div>
        <div id="imageGrid" class="image-grid">
            <div class="empty-gallery">加载中...</div>
        </div>
        <div class="footer">
            ⚡ Gin 后端 · 图片存储于 ./uploads 目录 · 支持查看和删除
        </div>
    </div>
</div>

<!-- 图片预览模态框 -->
<div id="imageModal" class="modal" onclick="closeModal()">
    <img id="modalImage" alt="大图预览">
</div>

<script>
    // DOM 元素
    const dropArea = document.getElementById('dropArea');
    const fileInput = document.getElementById('fileInput');
    const previewSection = document.getElementById('previewSection');
    const previewImage = document.getElementById('previewImage');
    const fileInfo = document.getElementById('fileInfo');
    const clearBtn = document.getElementById('clearBtn');
    const uploadBtn = document.getElementById('uploadBtn');
    const statusDiv = document.getElementById('statusMessage');
    const refreshBtn = document.getElementById('refreshBtn');
    const imageGrid = document.getElementById('imageGrid');

    let currentFile = null;

    // 显示状态消息
    function setStatus(type, text) {
        statusDiv.innerHTML = text;
        statusDiv.className = 'status-message';
        if (type === 'success') {
            statusDiv.classList.add('status-success');
        } else if (type === 'error') {
            statusDiv.classList.add('status-error');
        } else if (type === 'loading') {
            statusDiv.classList.add('status-loading');
        } else {
            statusDiv.classList.remove('status-success', 'status-error', 'status-loading');
        }
    }

    // 清空所有选择和预览
    function clearAll() {
        currentFile = null;
        fileInput.value = '';
        previewSection.style.display = 'none';
        previewImage.src = '';
        fileInfo.innerText = '';
        setStatus('info', '🗑️ 已清空，可以重新选择图片');
    }

    // 预览文件
    function previewFile(file) {
        if (!file) {
            previewSection.style.display = 'none';
            return;
        }

        if (!file.type.startsWith('image/')) {
            setStatus('error', '❌ 无效文件类型，请选择图片文件');
            clearAll();
            return false;
        }

        const maxSize = 10 * 1024 * 1024;
        if (file.size > maxSize) {
            setStatus('error', '⚠️ 图片超过 10MB，请压缩后再上传');
            clearAll();
            return false;
        }

        const reader = new FileReader();
        reader.onload = function(e) {
            previewImage.src = e.target.result;
            const sizeKB = (file.size / 1024).toFixed(1);
            const sizeMB = (file.size / (1024 * 1024)).toFixed(2);
            const displaySize = file.size > 1024 * 1024 ? sizeMB + ' MB' : sizeKB + ' KB';
            fileInfo.innerText = '文件名: ' + file.name + '  |  大小: ' + displaySize + '  |  类型: ' + file.type;
            previewSection.style.display = 'block';
            setStatus('info', '✅ 图片已加载，可以上传');
        };
        reader.onerror = () => {
            setStatus('error', '读取图片失败，请重试');
            clearAll();
        };
        reader.readAsDataURL(file);
        return true;
    }

    // 处理文件选择
    function handleFiles(files) {
        if (!files || files.length === 0) return;
        const file = files[0];
        if (!file.type.startsWith('image/')) {
            setStatus('error', '❌ 请选择有效的图片文件');
            return;
        }
        currentFile = file;
        previewFile(currentFile);
    }

    // 获取图片列表
    async function loadImageList() {
        imageGrid.innerHTML = '<div class="empty-gallery">加载中...</div>';
        try {
            const response = await fetch('/api/images');
            const result = await response.json();
            
            if (result.code === 200 && result.data && result.data.images) {
                const images = result.data.images;
                if (images.length === 0) {
                    imageGrid.innerHTML = '<div class="empty-gallery">📭 暂无图片，上传第一张吧</div>';
                    return;
                }
                
                imageGrid.innerHTML = '';
                images.forEach(img => {
                    const card = createImageCard(img);
                    imageGrid.appendChild(card);
                });
            } else {
                imageGrid.innerHTML = '<div class="empty-gallery">❌ 加载失败: ' + (result.message || '未知错误') + '</div>';
            }
        } catch (err) {
            console.error('加载图片列表失败:', err);
            imageGrid.innerHTML = '<div class="empty-gallery">❌ 网络错误，无法加载图片列表</div>';
        }
    }

    // 创建图片卡片
    function createImageCard(img) {
        const card = document.createElement('div');
        card.className = 'image-card';
        
        const uploadDate = new Date(img.upload_time).toLocaleString('zh-CN');
        const sizeDisplay = img.size_mb > 1 ? img.size_mb.toFixed(2) + ' MB' : img.size_kb.toFixed(1) + ' KB';
        
        card.innerHTML = '<img class="image-card-img" src="' + img.url + '" alt="' + img.filename + '" onclick="viewImage(\'' + img.url + '\')">' +
            '<div class="image-card-info">' +
            '<div class="image-filename" title="' + img.filename + '">' + (img.filename.substring(0, 30) + (img.filename.length > 30 ? '...' : '')) + '</div>' +
            '<div class="image-meta">📅 ' + uploadDate + '<br>📦 ' + sizeDisplay + ' · ' + img.ext.toUpperCase() + '</div>' +
            '<div class="image-actions">' +
            '<button class="btn btn-outline" onclick="viewImage(\'' + img.url + '\')">🔍 查看</button>' +
            '<button class="btn btn-danger" onclick="deleteImage(\'' + img.filename + '\')">🗑️ 删除</button>' +
            '</div>' +
            '</div>';
        
        return card;
    }

    // 查看大图
    function viewImage(url) {
        const modal = document.getElementById('imageModal');
        const modalImg = document.getElementById('modalImage');
        modal.style.display = 'flex';
        modalImg.src = url;
    }

    // 关闭模态框
    function closeModal() {
        const modal = document.getElementById('imageModal');
        modal.style.display = 'none';
    }

    // 删除图片
    async function deleteImage(filename) {
        if (!confirm('确定要删除图片 "' + filename + '" 吗？')) {
            return;
        }
        
        try {
            const response = await fetch('/api/images/' + encodeURIComponent(filename), {
                method: 'DELETE'
            });
            const result = await response.json();
            
            if (result.code === 200) {
                setStatus('success', '✅ 删除成功: ' + filename);
                loadImageList();
            } else {
                setStatus('error', '❌ 删除失败: ' + result.message);
            }
        } catch (err) {
            console.error('删除失败:', err);
            setStatus('error', '❌ 网络错误，删除失败');
        }
    }

    // 上传图片
    async function uploadImage() {
        if (!currentFile) {
            setStatus('error', '⚠️ 没有选择图片，请先选择或拖拽图片');
            return;
        }

        if (!currentFile.type.startsWith('image/')) {
            setStatus('error', '无效图片类型');
            return;
        }
        
        if (currentFile.size > 10 * 1024 * 1024) {
            setStatus('error', '图片超过10MB限制');
            return;
        }

        const formData = new FormData();
        formData.append('image', currentFile);

        setStatus('loading', '⏳ 上传中，请稍候...');
        uploadBtn.disabled = true;
        uploadBtn.textContent = '上传中...';

        try {
            const response = await fetch('/api/upload', {
                method: 'POST',
                body: formData
            });
            const result = await response.json();

            if (response.ok && result.code === 200) {
                let imageUrl = result.data.url;
                setStatus('success', '🎉 上传成功！<br> 文件: ' + result.data.filename + '<br> 访问地址: <a href="' + imageUrl + '" target="_blank" style="color:#2c7a4d;">' + imageUrl + '</a>');
                clearAll();
                loadImageList();
            } else {
                setStatus('error', '❌ 上传失败: ' + (result.message || '未知错误'));
            }
        } catch (err) {
            console.error('上传异常:', err);
            setStatus('error', '❌ 网络或服务器错误: ' + err.message);
        } finally {
            uploadBtn.disabled = false;
            uploadBtn.textContent = '📤 上传到服务器';
        }
    }

    // 事件监听
    dropArea.addEventListener('click', () => {
        fileInput.click();
    });

    fileInput.addEventListener('change', (e) => {
        if (e.target.files && e.target.files.length > 0) {
            handleFiles(e.target.files);
        }
    });

    dropArea.addEventListener('dragover', (e) => {
        e.preventDefault();
        dropArea.style.borderColor = '#2c5f8a';
        dropArea.style.backgroundColor = '#f0f7fe';
    });

    dropArea.addEventListener('dragleave', (e) => {
        e.preventDefault();
        dropArea.style.borderColor = '#b9c4d1';
        dropArea.style.backgroundColor = '#fefefe';
    });

    dropArea.addEventListener('drop', (e) => {
        e.preventDefault();
        dropArea.style.borderColor = '#b9c4d1';
        dropArea.style.backgroundColor = '#fefefe';
        if (e.dataTransfer.files && e.dataTransfer.files.length) {
            handleFiles(e.dataTransfer.files);
        } else {
            setStatus('error', '拖拽无效，请直接选择图片');
        }
    });

    clearBtn.addEventListener('click', () => {
        clearAll();
    });

    uploadBtn.addEventListener('click', uploadImage);
    refreshBtn.addEventListener('click', loadImageList);

    // 键盘ESC关闭模态框
    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape') {
            closeModal();
        }
    });

    // 初始化加载图片列表
    loadImageList();
    setStatus('info', '✨ 就绪，点击上方区域选择图片或拖拽图片至此');
</script>
</body>
</html>`
}
主要修复内容：
修复了模板字符串语法：将 JavaScript 中的模板字符串（反引号）改为字符串拼接，避免与 Go 的多行字符串反引号冲突

修复了字符串拼接：

sizeMB + ' MB' 替代 `${sizeMB} MB`

使用 + 运算符进行字符串拼接

代码现在可以正常编译运行：

bash
go mod init image-uploader
go get github.com/gin-gonic/gin
go run main.go
服务将在 http://localhost:8080 启动，提供完整的上传、查看列表、删除功能。

本回答由 AI 生成，内容仅供参考，请仔细甄别。
新增支持任意类型文件上传
以下是支持任意类型文件上传的完整实现，包括前端界面和后端服务：

go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "sort"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
)

// 上传配置常量
const (
    maxUploadSize   = 100 << 20               // 100MB (支持更大的文件)
    uploadDir       = "./uploads"             // 存储目录
    staticRoutePath = "/uploads"              // 静态访问路由
)

// 允许的文件类型映射
var allowedMimeTypes = map[string]bool{
    // 图片类型
    "image/jpeg":      true,
    "image/png":       true,
    "image/gif":       true,
    "image/webp":      true,
    "image/bmp":       true,
    "image/svg+xml":   true,
    // 文档类型
    "application/pdf":               true,
    "application/msword":            true,
    "application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
    "application/vnd.ms-excel":            true,
    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": true,
    "application/vnd.ms-powerpoint":         true,
    "application/vnd.openxmlformats-officedocument.presentationml.presentation": true,
    "text/plain":                   true,
    "text/csv":                     true,
    "application/json":             true,
    "application/xml":              true,
    // 压缩文件
    "application/zip":      true,
    "application/x-rar-compressed": true,
    "application/x-tar":            true,
    "application/gzip":             true,
    // 音视频
    "audio/mpeg":        true,
    "audio/wav":         true,
    "video/mp4":         true,
    "video/mpeg":        true,
    "video/quicktime":   true,
    // 其他
    "application/octet-stream": true, // 二进制文件
}

// UploadResponse 定义上传返回结构
type UploadResponse struct {
    Code    int       `json:"code"`
    Message string    `json:"message"`
    Data    *FileInfo `json:"data,omitempty"`
}

// FileInfo 文件信息结构
type FileInfo struct {
    Filename   string    `json:"filename"`
    OriginalName string  `json:"original_name"`
    URL        string    `json:"url"`
    Size       int64     `json:"size"`
    SizeKB     float64   `json:"size_kb"`
    SizeMB     float64   `json:"size_mb"`
    UploadTime time.Time `json:"upload_time"`
    Ext        string    `json:"ext"`
    MimeType   string    `json:"mime_type"`
    FileType   string    `json:"file_type"` // image, document, video, audio, archive, other
}

// ImageListResponse 文件列表返回结构
type ImageListResponse struct {
    Code    int            `json:"code"`
    Message string         `json:"message"`
    Data    *FileListData  `json:"data,omitempty"`
}

// FileListData 文件列表数据
type FileListData struct {
    Total   int         `json:"total"`
    Files   []*FileInfo `json:"files"`
    BaseURL string      `json:"base_url"`
}

func main() {
    // 设置 Gin 模式
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()

    // 添加CORS中间件
    r.Use(corsMiddleware())

    // 确保上传目录存在
    if err := ensureUploadDir(uploadDir); err != nil {
        log.Fatalf("创建上传目录失败: %v", err)
    }

    // 提供静态文件服务
    r.Static(staticRoutePath, uploadDir)

    // 处理前端页面
    r.GET("/", serveUploadPage)

    // 上传文件接口
    r.POST("/api/upload", handleFileUpload)

    // 获取文件列表接口
    r.GET("/api/files", handleGetFileList)

    // 删除文件接口
    r.DELETE("/api/files/:filename", handleDeleteFile)

    // 按类型筛选文件列表
    r.GET("/api/files/type/:fileType", handleGetFilesByType)

    // 健康检查
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().String()})
    })

    // 启动服务
    port := ":8080"
    log.Printf("🚀 Gin 文件上传服务启动在 http://localhost%s", port)
    log.Printf("📁 上传目录: %s (绝对路径: %s)", uploadDir, getAbsPath(uploadDir))
    log.Printf("📋 文件列表接口: http://localhost%s/api/files", port)
    if err := r.Run(port); err != nil {
        log.Fatal("服务启动失败:", err)
    }
}

// corsMiddleware CORS中间件
func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

// serveUploadPage 返回上传界面HTML
func serveUploadPage(c *gin.Context) {
    htmlContent := getUploadHTML()
    c.Header("Content-Type", "text/html; charset=utf-8")
    c.String(http.StatusOK, htmlContent)
}

// handleFileUpload 处理文件上传
func handleFileUpload(c *gin.Context) {
    // 限制请求体大小
    c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)
    
    // 从表单获取文件，字段名为 'file'
    file, header, err := c.Request.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: "获取上传文件失败: " + err.Error(),
        })
        return
    }
    defer file.Close()

    // 校验文件大小
    if header.Size > maxUploadSize {
        c.JSON(http.StatusBadRequest, UploadResponse{
            Code:    400,
            Message: fmt.Sprintf("文件超过最大限制 %dMB", maxUploadSize>>20),
        })
        return
    }

    // 获取文件信息
    originalFilename := header.Filename
    ext := strings.ToLower(filepath.Ext(originalFilename))
    
    // 检测MIME类型（读取前512字节）
    buffer := make([]byte, 512)
    file.Read(buffer)
    mimeType := http.DetectContentType(buffer)
    // 重置文件指针
    file.Seek(0, 0)
    
    // 如果没有检测到MIME类型，使用默认值
    if mimeType == "application/octet-stream" && ext != "" {
        // 根据扩展名推断类型
        mimeType = getMimeTypeFromExt(ext)
    }

    // 生成安全文件名
    safeFilename := generateSafeFilename(originalFilename)
    savePath := filepath.Join(uploadDir, safeFilename)

    // 保存文件
    if err := c.SaveUploadedFile(header, savePath); err != nil {
        log.Printf("保存文件失败: %v", err)
        c.JSON(http.StatusInternalServerError, UploadResponse{
            Code:    500,
            Message: "保存文件失败，请稍后重试",
        })
        return
    }

    // 构建可访问的URL
    fileURL := fmt.Sprintf("%s/%s", staticRoutePath, safeFilename)
    
    // 确定文件类型分类
    fileType := getFileTypeCategory(mimeType, ext)

    // 返回成功信息
    c.JSON(http.StatusOK, UploadResponse{
        Code:    200,
        Message: "上传成功",
        Data: &FileInfo{
            Filename:     safeFilename,
            OriginalName: originalFilename,
            URL:          fileURL,
            Size:         header.Size,
            SizeKB:       float64(header.Size) / 1024,
            SizeMB:       float64(header.Size) / (1024 * 1024),
            UploadTime:   time.Now(),
            Ext:          ext,
            MimeType:     mimeType,
            FileType:     fileType,
        },
    })
    log.Printf("上传成功: %s -> %s (大小: %d bytes, 类型: %s)", originalFilename, savePath, header.Size, mimeType)
}

// handleGetFileList 获取已上传文件列表
func handleGetFileList(c *gin.Context) {
    files, err := os.ReadDir(uploadDir)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ImageListResponse{
            Code:    500,
            Message: "读取文件列表失败: " + err.Error(),
        })
        return
    }

    var fileInfos []*FileInfo
    for _, file := range files {
        if file.IsDir() {
            continue
        }
        
        filename := file.Name()
        
        // 获取文件详细信息
        filePath := filepath.Join(uploadDir, filename)
        fileInfo, err := os.Stat(filePath)
        if err != nil {
            log.Printf("获取文件信息失败 %s: %v", filename, err)
            continue
        }
        
        // 检测文件类型
        mimeType, fileType := detectFileType(filePath)
        ext := strings.ToLower(filepath.Ext(filename))
        
        // 构建文件信息
        info := &FileInfo{
            Filename:     filename,
            OriginalName: filename, // 可以从元数据中读取，这里简化为文件名
            URL:          fmt.Sprintf("%s/%s", staticRoutePath, filename),
            Size:         fileInfo.Size(),
            SizeKB:       float64(fileInfo.Size()) / 1024,
            SizeMB:       float64(fileInfo.Size()) / (1024 * 1024),
            UploadTime:   fileInfo.ModTime(),
            Ext:          ext,
            MimeType:     mimeType,
            FileType:     fileType,
        }
        fileInfos = append(fileInfos, info)
    }
    
    // 按上传时间倒序排序
    sort.Slice(fileInfos, func(i, j int) bool {
        return fileInfos[i].UploadTime.After(fileInfos[j].UploadTime)
    })
    
    c.JSON(http.StatusOK, ImageListResponse{
        Code:    200,
        Message: "获取成功",
        Data: &FileListData{
            Total:   len(fileInfos),
            Files:   fileInfos,
            BaseURL: staticRoutePath,
        },
    })
}

// handleGetFilesByType 按类型筛选文件列表
func handleGetFilesByType(c *gin.Context) {
    fileType := c.Param("fileType")
    
    files, err := os.ReadDir(uploadDir)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ImageListResponse{
            Code:    500,
            Message: "读取文件列表失败: " + err.Error(),
        })
        return
    }

    var fileInfos []*FileInfo
    for _, file := range files {
        if file.IsDir() {
            continue
        }
        
        filename := file.Name()
        filePath := filepath.Join(uploadDir, filename)
        fileInfo, err := os.Stat(filePath)
        if err != nil {
            continue
        }
        
        mimeType, fType := detectFileType(filePath)
        
        // 按类型筛选
        if fType != fileType && fileType != "all" {
            continue
        }
        
        ext := strings.ToLower(filepath.Ext(filename))
        info := &FileInfo{
            Filename:     filename,
            OriginalName: filename,
            URL:          fmt.Sprintf("%s/%s", staticRoutePath, filename),
            Size:         fileInfo.Size(),
            SizeKB:       float64(fileInfo.Size()) / 1024,
            SizeMB:       float64(fileInfo.Size()) / (1024 * 1024),
            UploadTime:   fileInfo.ModTime(),
            Ext:          ext,
            MimeType:     mimeType,
            FileType:     fType,
        }
        fileInfos = append(fileInfos, info)
    }
    
    sort.Slice(fileInfos, func(i, j int) bool {
        return fileInfos[i].UploadTime.After(fileInfos[j].UploadTime)
    })
    
    c.JSON(http.StatusOK, ImageListResponse{
        Code:    200,
        Message: "获取成功",
        Data: &FileListData{
            Total:   len(fileInfos),
            Files:   fileInfos,
            BaseURL: staticRoutePath,
        },
    })
}

// handleDeleteFile 删除文件
func handleDeleteFile(c *gin.Context) {
    filename := c.Param("filename")
    
    // 安全检查
    if strings.Contains(filename, "..") || strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    400,
            "message": "非法文件名",
        })
        return
    }
    
    filePath := filepath.Join(uploadDir, filename)
    
    // 检查文件是否存在
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{
            "code":    404,
            "message": "文件不存在",
        })
        return
    }
    
    // 删除文件
    if err := os.Remove(filePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "code":    500,
            "message": "删除文件失败: " + err.Error(),
        })
        return
    }
    
    log.Printf("删除成功: %s", filename)
    c.JSON(http.StatusOK, gin.H{
        "code":    200,
        "message": "删除成功",
    })
}

// 生成唯一文件名
func generateSafeFilename(originalName string) string {
    timestamp := time.Now().UnixNano()
    randomPart := fmt.Sprintf("%d", time.Now().UnixNano()%1000000)
    ext := filepath.Ext(originalName)
    return fmt.Sprintf("%d_%s%s", timestamp, randomPart, ext)
}

// 确保上传目录存在
func ensureUploadDir(dir string) error {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        return os.MkdirAll(dir, 0755)
    }
    return nil
}

// 获取绝对路径
func getAbsPath(dir string) string {
    abs, err := filepath.Abs(dir)
    if err != nil {
        return dir
    }
    return abs
}

// 根据扩展名获取MIME类型
func getMimeTypeFromExt(ext string) string {
    mimeMap := map[string]string{
        ".jpg":  "image/jpeg",
        ".jpeg": "image/jpeg",
        ".png":  "image/png",
        ".gif":  "image/gif",
        ".webp": "image/webp",
        ".pdf":  "application/pdf",
        ".doc":  "application/msword",
        ".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
        ".xls":  "application/vnd.ms-excel",
        ".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
        ".txt":  "text/plain",
        ".zip":  "application/zip",
        ".mp4":  "video/mp4",
        ".mp3":  "audio/mpeg",
    }
    
    if mime, ok := mimeMap[ext]; ok {
        return mime
    }
    return "application/octet-stream"
}

// 检测文件类型
func detectFileType(filePath string) (mimeType string, fileType string) {
    file, err := os.Open(filePath)
    if err != nil {
        return "application/octet-stream", "other"
    }
    defer file.Close()
    
    buffer := make([]byte, 512)
    _, err = file.Read(buffer)
    if err != nil {
        return "application/octet-stream", "other"
    }
    
    mimeType = http.DetectContentType(buffer)
    fileType = getFileTypeCategory(mimeType, filepath.Ext(filePath))
    return mimeType, fileType
}

// 获取文件类型分类
func getFileTypeCategory(mimeType string, ext string) string {
    if strings.HasPrefix(mimeType, "image/") {
        return "image"
    }
    if strings.HasPrefix(mimeType, "video/") {
        return "video"
    }
    if strings.HasPrefix(mimeType, "audio/") {
        return "audio"
    }
    if strings.HasPrefix(mimeType, "text/") {
        return "document"
    }
    if strings.Contains(mimeType, "pdf") || strings.Contains(mimeType, "document") || strings.Contains(mimeType, "spreadsheet") {
        return "document"
    }
    if strings.Contains(mimeType, "zip") || strings.Contains(mimeType, "compressed") || ext == ".zip" || ext == ".rar" || ext == ".7z" {
        return "archive"
    }
    return "other"
}

// getUploadHTML 返回前端界面
func getUploadHTML() string {
    return `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>通用文件上传 - Go Gin 服务</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 24px;
        }

        .app-container {
            max-width: 1400px;
            margin: 0 auto;
        }

        .upload-container, .gallery-container {
            background: white;
            border-radius: 24px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            overflow: hidden;
            margin-bottom: 32px;
        }

        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            padding: 32px;
            color: white;
        }

        .header h1 {
            font-size: 2rem;
            margin-bottom: 8px;
        }

        .header p {
            opacity: 0.9;
        }

        .card-body {
            padding: 32px;
        }

        .upload-area {
            border: 2px dashed #cbd5e1;
            border-radius: 16px;
            padding: 48px;
            text-align: center;
            cursor: pointer;
            transition: all 0.3s;
            background: #f8fafc;
        }

        .upload-area:hover {
            border-color: #667eea;
            background: #f0f4ff;
        }

        .upload-icon {
            font-size: 64px;
            margin-bottom: 16px;
        }

        .file-input {
            display: none;
        }

        .preview-section {
            margin-top: 24px;
            padding: 16px;
            background: #f1f5f9;
            border-radius: 12px;
            display: none;
        }

        .file-preview {
            display: flex;
            align-items: center;
            gap: 16px;
            flex-wrap: wrap;
        }

        .file-icon {
            font-size: 48px;
        }

        .file-details {
            flex: 1;
        }

        .file-name {
            font-weight: 600;
            margin-bottom: 4px;
        }

        .file-meta {
            font-size: 0.85rem;
            color: #64748b;
        }

        .progress-bar {
            margin-top: 16px;
            height: 8px;
            background: #e2e8f0;
            border-radius: 4px;
            overflow: hidden;
            display: none;
        }

        .progress-fill {
            height: 100%;
            background: linear-gradient(90deg, #667eea, #764ba2);
            width: 0%;
            transition: width 0.3s;
        }

        .action-buttons {
            display: flex;
            gap: 12px;
            margin-top: 24px;
        }

        .btn {
            padding: 10px 24px;
            border: none;
            border-radius: 8px;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.2s;
        }

        .btn-primary {
            background: #667eea;
            color: white;
        }

        .btn-primary:hover {
            background: #5a67d8;
            transform: translateY(-1px);
        }

        .btn-outline {
            background: transparent;
            border: 1px solid #cbd5e1;
        }

        .btn-outline:hover {
            background: #f1f5f9;
        }

        .btn-danger {
            background: #ef4444;
            color: white;
        }

        .btn-danger:hover {
            background: #dc2626;
        }

        .status-message {
            margin-top: 16px;
            padding: 12px;
            border-radius: 8px;
            font-size: 0.9rem;
        }

        .status-success {
            background: #d1fae5;
            color: #065f46;
        }

        .status-error {
            background: #fee2e2;
            color: #991b1b;
        }

        .status-loading {
            background: #e0e7ff;
            color: #3730a3;
        }

        /* 文件列表样式 */
        .gallery-header {
            padding: 24px 32px;
            background: #f8fafc;
            border-bottom: 1px solid #e2e8f0;
            display: flex;
            justify-content: space-between;
            align-items: center;
            flex-wrap: wrap;
            gap: 16px;
        }

        .filter-buttons {
            display: flex;
            gap: 8px;
            flex-wrap: wrap;
        }

        .filter-btn {
            padding: 6px 16px;
            border: 1px solid #cbd5e1;
            background: white;
            border-radius: 20px;
            cursor: pointer;
            transition: all 0.2s;
        }

        .filter-btn.active {
            background: #667eea;
            color: white;
            border-color: #667eea;
        }

        .file-grid {
            padding: 32px;
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
            gap: 24px;
        }

        .file-card {
            background: white;
            border: 1px solid #e2e8f0;
            border-radius: 12px;
            overflow: hidden;
            transition: all 0.3s;
        }

        .file-card:hover {
            transform: translateY(-4px);
            box-shadow: 0 12px 24px rgba(0,0,0,0.1);
        }

        .file-preview-img {
            width: 100%;
            height: 160px;
            object-fit: cover;
            background: #f1f5f9;
        }

        .file-preview-icon {
            width: 100%;
            height: 160px;
            display: flex;
            align-items: center;
            justify-content: center;
            background: #f1f5f9;
            font-size: 64px;
        }

        .file-info {
            padding: 16px;
        }

        .file-title {
            font-weight: 600;
            margin-bottom: 8px;
            word-break: break-all;
            font-size: 0.9rem;
        }

        .file-meta-info {
            font-size: 0.75rem;
            color: #64748b;
            margin-bottom: 12px;
        }

        .file-actions {
            display: flex;
            gap: 8px;
        }

        .file-actions .btn {
            flex: 1;
            padding: 6px 12px;
            font-size: 0.8rem;
        }

        .empty-state {
            text-align: center;
            padding: 60px;
            color: #94a3b8;
        }

        .modal {
            display: none;
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background: rgba(0,0,0,0.9);
            z-index: 1000;
            align-items: center;
            justify-content: center;
        }

        .modal-content {
            max-width: 90%;
            max-height: 90%;
        }

        @media (max-width: 768px) {
            .card-body, .file-grid {
                padding: 20px;
            }
            .header h1 {
                font-size: 1.5rem;
            }
        }
    </style>
</head>
<body>
<div class="app-container">
    <!-- 上传区域 -->
    <div class="upload-container">
        <div class="header">
            <h1>📁 通用文件上传服务</h1>
            <p>支持任意类型文件 | 最大 100MB | Go + Gin 后端</p>
        </div>
        <div class="card-body">
            <div class="upload-area" id="dropZone">
                <div class="upload-icon">📤</div>
                <p>点击或拖拽文件至此区域</p>
                <p style="font-size: 0.85rem; margin-top: 8px;">支持任意类型文件，最大 100MB</p>
                <input type="file" id="fileInput" class="file-input">
            </div>

            <div class="preview-section" id="previewSection">
                <div class="file-preview">
                    <div class="file-icon" id="fileIcon">📄</div>
                    <div class="file-details">
                        <div class="file-name" id="fileName"></div>
                        <div class="file-meta" id="fileMeta"></div>
                    </div>
                </div>
                <div class="progress-bar" id="progressBar">
                    <div class="progress-fill" id="progressFill"></div>
                </div>
            </div>

            <div class="action-buttons">
                <button class="btn btn-outline" id="clearBtn">清空</button>
                <button class="btn btn-primary" id="uploadBtn">上传文件</button>
            </div>

            <div id="statusMessage" class="status-message">✨ 就绪，选择文件后即可上传</div>
        </div>
    </div>

    <!-- 文件列表区域 -->
    <div class="gallery-container">
        <div class="gallery-header">
            <h2>📋 已上传文件列表</h2>
            <div class="filter-buttons">
                <button class="filter-btn active" data-type="all">全部</button>
                <button class="filter-btn" data-type="image">图片</button>
                <button class="filter-btn" data-type="document">文档</button>
                <button class="filter-btn" data-type="video">视频</button>
                <button class="filter-btn" data-type="audio">音频</button>
                <button class="filter-btn" data-type="archive">压缩包</button>
                <button class="filter-btn" data-type="other">其他</button>
            </div>
            <button class="btn btn-outline" id="refreshBtn">🔄 刷新</button>
        </div>
        <div id="fileGrid" class="file-grid">
            <div class="empty-state">加载中...</div>
        </div>
    </div>
</div>

<!-- 预览模态框 -->
<div id="previewModal" class="modal" onclick="closePreview()">
    <img id="previewImage" class="modal-content" style="display: none;">
    <iframe id="previewPDF" style="display: none; width: 90%; height: 90%; border: none;"></iframe>
    <div id="previewOther" class="modal-content" style="display: none; text-align: center; color: white;"></div>
</div>

<script>
    let currentFile = null;
    let currentFilter = 'all';

    // DOM 元素
    const dropZone = document.getElementById('dropZone');
    const fileInput = document.getElementById('fileInput');
    const previewSection = document.getElementById('previewSection');
    const fileIcon = document.getElementById('fileIcon');
    const fileName = document.getElementById('fileName');
    const fileMeta = document.getElementById('fileMeta');
    const progressBar = document.getElementById('progressBar');
    const progressFill = document.getElementById('progressFill');
    const clearBtn = document.getElementById('clearBtn');
    const uploadBtn = document.getElementById('uploadBtn');
    const statusDiv = document.getElementById('statusMessage');
    const refreshBtn = document.getElementById('refreshBtn');
    const fileGrid = document.getElementById('fileGrid');
    const filterBtns = document.querySelectorAll('.filter-btn');

    // 文件类型图标映射
    const iconMap = {
        image: '🖼️',
        video: '🎬',
        audio: '🎵',
        document: '📄',
        archive: '📦',
        other: '📁'
    };

    function setStatus(type, text) {
        statusDiv.innerHTML = text;
        statusDiv.className = 'status-message';
        if (type === 'success') statusDiv.classList.add('status-success');
        else if (type === 'error') statusDiv.classList.add('status-error');
        else if (type === 'loading') statusDiv.classList.add('status-loading');
    }

    function clearPreview() {
        currentFile = null;
        previewSection.style.display = 'none';
        fileInput.value = '';
        progressFill.style.width = '0%';
        progressBar.style.display = 'none';
    }

    function previewFile(file) {
        if (!file) return;
        
        const sizeMB = (file.size / (1024 * 1024)).toFixed(2);
        const sizeKB = (file.size / 1024).toFixed(1);
        const displaySize = file.size > 1024 * 1024 ? sizeMB + ' MB' : sizeKB + ' KB';
        
        fileName.textContent = file.name;
        fileMeta.textContent = '大小: ' + displaySize + ' | 类型: ' + (file.type || '未知');
        
        // 设置图标
        if (file.type.startsWith('image/')) fileIcon.textContent = '🖼️';
        else if (file.type.startsWith('video/')) fileIcon.textContent = '🎬';
        else if (file.type.startsWith('audio/')) fileIcon.textContent = '🎵';
        else if (file.type.includes('pdf')) fileIcon.textContent = '📕';
        else if (file.type.includes('document') || file.type.includes('text')) fileIcon.textContent = '📄';
        else fileIcon.textContent = '📁';
        
        previewSection.style.display = 'block';
        setStatus('info', '✅ 文件已加载，可以上传');
    }

    function handleFiles(files) {
        if (!files || files.length === 0) return;
        currentFile = files[0];
        previewFile(currentFile);
    }

    async function uploadFile() {
        if (!currentFile) {
            setStatus('error', '⚠️ 没有选择文件');
            return;
        }

        const formData = new FormData();
        formData.append('file', currentFile);

        setStatus('loading', '⏳ 上传中...');
        uploadBtn.disabled = true;
        progressBar.style.display = 'block';
        
        // 模拟进度（实际可以用 XMLHttpRequest 实现真实进度）
        let progress = 0;
        const interval = setInterval(() => {
            progress = Math.min(progress + 10, 90);
            progressFill.style.width = progress + '%';
        }, 200);

        try {
            const response = await fetch('/api/upload', {
                method: 'POST',
                body: formData
            });
            const result = await response.json();

            clearInterval(interval);
            progressFill.style.width = '100%';

            if (response.ok && result.code === 200) {
                setStatus('success', '🎉 上传成功！文件: ' + result.data.filename);
                clearPreview();
                loadFileList();
            } else {
                setStatus('error', '❌ 上传失败: ' + (result.message || '未知错误'));
            }
        } catch (err) {
            clearInterval(interval);
            setStatus('error', '❌ 上传失败: ' + err.message);
        } finally {
            uploadBtn.disabled = false;
            setTimeout(() => {
                progressBar.style.display = 'none';
                progressFill.style.width = '0%';
            }, 1000);
        }
    }

    async function loadFileList() {
        fileGrid.innerHTML = '<div class="empty-state">加载中...</div>';
        try {
            let url = '/api/files';
            if (currentFilter !== 'all') {
                url = '/api/files/type/' + currentFilter;
            }
            const response = await fetch(url);
            const result = await response.json();
            
            if (result.code === 200 && result.data && result.data.files) {
                const files = result.data.files;
                if (files.length === 0) {
                    fileGrid.innerHTML = '<div class="empty-state">📭 暂无文件，上传第一个吧</div>';
                    return;
                }
                
                fileGrid.innerHTML = '';
                files.forEach(file => {
                    const card = createFileCard(file);
                    fileGrid.appendChild(card);
                });
            } else {
                fileGrid.innerHTML = '<div class="empty-state">❌ 加载失败</div>';
            }
        } catch (err) {
            fileGrid.innerHTML = '<div class="empty-state">❌ 网络错误</div>';
        }
    }

    function createFileCard(file) {
        const card = document.createElement('div');
        card.className = 'file-card';
        
        const uploadDate = new Date(file.upload_time).toLocaleString('zh-CN');
        const sizeDisplay = file.size_mb > 1 ? file.size_mb.toFixed(2) + ' MB' : file.size_kb.toFixed(1) + ' KB';
        const icon = iconMap[file.file_type] || '📁';
        
        let previewHtml = '';
        if (file.file_type === 'image') {
            previewHtml = '<img class="file-preview-img" src="' + file.url + '" alt="' + file.filename + '" onclick="previewFileByUrl(\'' + file.url + '\', \'image\')">';
        } else {
            previewHtml = '<div class="file-preview-icon" onclick="previewFileByUrl(\'' + file.url + '\', \'' + file.file_type + '\')">' + icon + '</div>';
        }
        
        card.innerHTML = previewHtml +
            '<div class="file-info">' +
            '<div class="file-title" title="' + file.filename + '">' + (file.filename.substring(0, 40) + (file.filename.length > 40 ? '...' : '')) + '</div>' +
            '<div class="file-meta-info">📅 ' + uploadDate + '<br>📦 ' + sizeDisplay + '<br>🔤 ' + (file.ext || '未知') + '</div>' +
            '<div class="file-actions">' +
            '<button class="btn btn-outline" onclick="previewFileByUrl(\'' + file.url + '\', \'' + file.file_type + '\')">🔍 预览</button>' +
            '<button class="btn btn-danger" onclick="deleteFile(\'' + file.filename + '\')">🗑️ 删除</button>' +
            '</div>' +
            '</div>';
        
        return card;
    }

    function previewFileByUrl(url, fileType) {
        const modal = document.getElementById('previewModal');
        const imgPreview = document.getElementById('previewImage');
        const pdfPreview = document.getElementById('previewPDF');
        const otherPreview = document.getElementById('previewOther');
        
        imgPreview.style.display = 'none';
        pdfPreview.style.display = 'none';
        otherPreview.style.display = 'none';
        
        if (fileType === 'image') {
            imgPreview.src = url;
            imgPreview.style.display = 'block';
        } else if (fileType === 'document' && url.toLowerCase().endsWith('.pdf')) {
            pdfPreview.src = url;
            pdfPreview.style.display = 'block';
        } else {
            otherPreview.innerHTML = '<div><div style="font-size: 64px;">📄</div><div style="margin-top: 20px;"><a href="' + url + '" download style="color: white; font-size: 18px;">点击下载文件</a></div><div style="margin-top: 10px; font-size: 14px;">该文件类型无法直接预览</div></div>';
            otherPreview.style.display = 'block';
        }
        
        modal.style.display = 'flex';
    }

    function closePreview() {
        document.getElementById('previewModal').style.display = 'none';
    }

    async function deleteFile(filename) {
        if (!confirm('确定要删除文件 "' + filename + '" 吗？')) return;
        
        try {
            const response = await fetch('/api/files/' + encodeURIComponent(filename), {
                method: 'DELETE'
            });
            const result = await response.json();
            
            if (result.code === 200) {
                setStatus('success', '✅ 删除成功');
                loadFileList();
            } else {
                setStatus('error', '❌ 删除失败: ' + result.message);
            }
        } catch (err) {
            setStatus('error', '❌ 删除失败');
        }
    }

    // 事件监听
    dropZone.addEventListener('click', () => fileInput.click());
    fileInput.addEventListener('change', (e) => handleFiles(e.target.files));
    clearBtn.addEventListener('click', clearPreview);
    uploadBtn.addEventListener('click', uploadFile);
    refreshBtn.addEventListener('click', loadFileList);
    
    filterBtns.forEach(btn => {
        btn.addEventListener('click', () => {
            filterBtns.forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            currentFilter = btn.dataset.type;
            loadFileList();
        });
    });
    
    dropZone.addEventListener('dragover', (e) => {
        e.preventDefault();
        dropZone.style.borderColor = '#667eea';
        dropZone.style.backgroundColor = '#f0f4ff';
    });
    
    dropZone.addEventListener('dragleave', (e) => {
        e.preventDefault();
        dropZone.style.borderColor = '#cbd5e1';
        dropZone.style.backgroundColor = '#f8fafc';
    });
    
    dropZone.addEventListener('drop', (e) => {
        e.preventDefault();
        dropZone.style.borderColor = '#cbd5e1';
        dropZone.style.backgroundColor = '#f8fafc';
        if (e.dataTransfer.files && e.dataTransfer.files.length) {
            handleFiles(e.dataTransfer.files);
        }
    });
    
    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape') closePreview();
    });
    
    // 初始化
    loadFileList();
</script>
</body>
</html>`
}
主要功能特点
后端接口：

POST /api/upload - 上传任意类型文件（最大100MB）

GET /api/files - 获取所有文件列表

GET /api/files/type/:fileType - 按类型筛选文件

DELETE /api/files/:filename - 删除文件

支持的文件类型：

图片：JPG, PNG, GIF, WEBP, BMP, SVG

文档：PDF, Word, Excel, PowerPoint, TXT, CSV, JSON, XML

视频：MP4, MPEG, MOV

音频：MP3, WAV

压缩包：ZIP, RAR, TAR, GZ

其他：任意二进制文件

前端功能：

拖拽或点击上传

文件类型分类筛选

图片和PDF在线预览

文件列表网格展示

文件大小和上传时间显示

支持文件删除

运行方式：

bash
go mod init file-uploader
go get github.com/gin-gonic/gin
go run main.go
访问 http://localhost:8080 即可使用完整功能。

本回答由 AI 生成，内容仅供参考，请仔细甄别。
上述还能拓展些什么功能
基于现有的通用文件上传系统，以下是一些可以拓展的功能建议：

🚀 推荐拓展功能
1. 用户认证与权限管理
go
// 添加JWT认证
- 用户注册/登录
- 文件归属用户（每个用户独立空间）
- 公开/私密文件设置
- 分享链接生成（带过期时间）
2. 文件管理增强
go
- 文件夹/目录管理（创建、重命名、移动）
- 批量上传（多文件同时上传）
- 批量删除和批量下载（ZIP打包）
- 文件重命名和移动
- 文件搜索（按名称、类型、日期）
- 文件标签系统
- 星标/收藏功能
3. 文件处理功能
go
- 图片压缩和格式转换
- 图片缩略图生成（多尺寸）
- 图片水印添加
- 视频转码和截图
- PDF合并/分割
- 文档在线编辑（Collabora/OnlyOffice集成）
- 文件加密存储
- 病毒扫描（ClamAV集成）
4. 分享与协作
go
- 生成分享链接（带密码、有效期）
- 二维码生成
- 文件评论/备注
- 协作编辑（实时）
- 文件版本历史
- 回收站（软删除，可恢复）
5. 高级功能
go
- 断点续传（大文件分片上传）
- 秒传功能（MD5校验）
- 文件去重检测
- 离线下载（从URL下载文件）
- WebDAV协议支持
- FTP/SFTP同步
- 云存储备份（OSS/S3）
6. 分析与监控
go
- 下载统计和访问日志
- 存储空间使用统计（图表）
- 文件类型分布统计
- 用户行为分析
- 系统资源监控（CPU/内存/磁盘）
- 操作审计日志
7. 界面与体验
go
- 深色/浅色主题切换
- 网格/列表视图切换
- 拖拽排序
- 快捷键支持
- 文件拖拽到文件夹
- 右键菜单（复制、移动、删除）
- 响应式设计优化
📝 重点功能实现示例
1. 批量上传 + 断点续传
go
// 分片上传结构
type ChunkUpload struct {
    FileName   string `json:"file_name"`
    ChunkIndex int    `json:"chunk_index"`
    TotalChunks int   `json:"total_chunks"`
    UploadID   string `json:"upload_id"`
}

// API端点
POST /api/upload/chunk    // 上传分片
POST /api/upload/merge     // 合并分片
GET  /api/upload/status    // 查询上传进度
2. 文件分享功能
go
type ShareInfo struct {
    ID         string    `json:"id"`
    FileID     string    `json:"file_id"`
    Password   string    `json:"password"`
    ExpireAt   time.Time `json:"expire_at"`
    MaxDownloads int     `json:"max_downloads"`
    DownloadCount int    `json:"download_count"`
}

// API端点
POST /api/share/:fileId      // 创建分享
GET  /api/share/:shareId     // 访问分享
DELETE /api/share/:shareId   // 取消分享
3. 文件版本控制
go
type FileVersion struct {
    VersionID   string    `json:"version_id"`
    FileID      string    `json:"file_id"`
    Size        int64     `json:"size"`
    CreatedAt   time.Time `json:"created_at"`
    Comment     string    `json:"comment"`
}

// API端点
GET    /api/files/:id/versions     // 获取版本列表
POST   /api/files/:id/versions     // 创建新版本
GET    /api/files/:id/versions/:vid // 恢复版本
DELETE /api/files/:id/versions/:vid // 删除版本
4. 文件搜索功能
go
// 全文搜索 + 过滤
GET /api/files/search?q=keyword&type=image&sort=date&page=1&size=20

// 支持搜索字段
- 文件名模糊搜索
- 文件内容搜索（文本文件）
- 按标签搜索
- 按上传时间范围
- 按文件大小范围
5. 回收站功能
go
type TrashItem struct {
    ID         string    `json:"id"`
    FileID     string    `json:"file_id"`
    DeletedAt  time.Time `json:"deleted_at"`
    ExpireAt   time.Time `json:"expire_at"` // 30天后自动清理
}

// API端点
GET    /api/trash              // 获取回收站列表
POST   /api/trash/restore/:id  // 恢复文件
DELETE /api/trash/permanent/:id // 永久删除
DELETE /api/trash/empty        // 清空回收站
🎨 前端功能增强
1. 文件管理器界面
html
- 左侧目录树
- 面包屑导航
- 文件多选（Ctrl/Shift）
- 拖拽移动文件
- 右键上下文菜单
- 键盘快捷键（删除、重命名、复制）
2. 实时通知
javascript
- WebSocket连接
- 上传完成通知
- 新文件通知
- 分享访问通知
3. 高级预览
javascript
- PDF.js集成（更好的PDF预览）
- 代码高亮（文本文件）
- Markdown渲染
- 3D模型预览（Three.js）
- 音视频播放器
📊 数据库设计建议
sql
-- 用户表
users (
    id, username, password, email, 
    quota_total, quota_used, created_at
)

-- 文件表
files (
    id, user_id, parent_id, name, 
    size, mime_type, hash, path,
    is_public, download_count, created_at
)

-- 分享表
shares (
    id, file_id, code, password,
    expire_at, max_downloads, created_at
)

-- 操作日志
logs (
    id, user_id, action, file_id,
    ip_address, created_at
)
🛠️ 技术栈建议
yaml
后端:
  - 数据库: PostgreSQL / MySQL
  - 缓存: Redis (会话、限流、队列)
  - 消息队列: RabbitMQ (异步处理)
  - 对象存储: MinIO / OSS / S3

前端:
  - 框架: Vue 3 / React
  - UI库: Element Plus / Ant Design
  - 状态管理: Pinia / Redux
  - 文件上传: resumable.js / spark-md5

部署:
  - 容器化: Docker + Docker Compose
  - 反向代理: Nginx
  - 进程管理: Supervisor / systemd
🎯 优先级建议
高优先级（用户体验）：

批量上传和拖拽上传

文件搜索功能

分享链接生成

缩略图预览

中优先级（功能完善）：

用户认证系统

文件夹管理

回收站功能

文件版本历史

低优先级（高级特性）：

在线编辑

断点续传

视频转码

WebDAV支持

这些功能可以根据您的实际需求选择性实现，建议先从高优先级功能开始逐步完善。需要我帮您实现其中某个具体功能吗？
```

