<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%--<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>--%>
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>用户登录/注册</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: 'Arial', sans-serif;
        }

        body {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            color: #333;
        }

        .container {
            background-color: white;
            border-radius: 10px;
            box-shadow: 0 15px 30px rgba(0, 0, 0, 0.2);
            width: 400px;
            max-width: 90%;
            overflow: hidden;
        }

        .form-container {
            padding: 40px;
            transition: all 0.6s ease-in-out;
        }

        .header {
            text-align: center;
            margin-bottom: 30px;
        }

        .header h1 {
            font-size: 28px;
            margin-bottom: 10px;
            color: #444;
        }

        .toggle-container {
            display: flex;
            margin-bottom: 20px;
            border-radius: 30px;
            overflow: hidden;
            background-color: #f0f0f0;
        }

        .toggle {
            flex: 1;
            padding: 10px;
            text-align: center;
            cursor: pointer;
            transition: all 0.3s;
        }

        .toggle.active {
            background-color: #667eea;
            color: white;
        }

        .form-group {
            margin-bottom: 20px;
        }

        .form-group label {
            display: block;
            margin-bottom: 8px;
            font-weight: 600;
            color: #555;
        }

        .form-group input {
            width: 100%;
            padding: 12px 15px;
            border: 1px solid #ddd;
            border-radius: 5px;
            font-size: 16px;
            transition: border 0.3s;
        }

        .form-group input:focus {
            border-color: #667eea;
            outline: none;
        }

        .btn {
            width: 100%;
            padding: 12px;
            background-color: #667eea;
            color: white;
            border: none;
            border-radius: 5px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            transition: background-color 0.3s;
        }

        .btn:hover {
            background-color: #5a6fd1;
        }

        .forgot-password {
            text-align: center;
            margin-top: 15px;
        }

        .forgot-password a {
            color: #667eea;
            text-decoration: none;
            font-size: 14px;
        }

        .forgot-password a:hover {
            text-decoration: underline;
        }

        #register-form {
            display: none;
        }

        .social-login {
            margin-top: 30px;
            text-align: center;
        }

        .social-login p {
            color: #777;
            margin-bottom: 15px;
            position: relative;
        }

        .social-login p::before,
        .social-login p::after {
            content: "";
            flex: 1;
            height: 1px;
            background: #ddd;
            position: absolute;
            top: 50%;
            width: 30%;
        }

        .social-login p::before {
            left: 0;
        }

        .social-login p::after {
            right: 0;
        }

        .social-icons {
            display: flex;
            justify-content: center;
            gap: 15px;
        }

        .social-icon {
            width: 40px;
            height: 40px;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            cursor: pointer;
            transition: all 0.3s;
        }

        .social-icon:hover {
            transform: translateY(-3px);
        }

        .facebook {
            background-color: #3b5998;
            color: white;
        }

        .google {
            background-color: #dd4b39;
            color: white;
        }

        .twitter {
            background-color: #1da1f2;
            color: white;
        }

        .error-message {
            color: #e74c3c;
            font-size: 14px;
            margin-top: 5px;
        }

        .success-message {
            color: #2ecc71;
            font-size: 14px;
            margin-top: 5px;
        }
    </style>
</head>
<body>
<div class="container">
    <div class="form-container">
        <div class="header">
            <h1>${not empty param.register ? '创建账户' : '欢迎回来'}</h1>
            <p>${not empty param.register ? '注册新用户' : '请登录您的账户'}</p>
        </div>

        <div class="toggle-container">
            <div class="toggle ${empty param.register ? 'active' : ''}" id="login-toggle">登录</div>
            <div class="toggle ${not empty param.register ? 'active' : ''}" id="register-toggle">注册</div>
        </div>

        <%-- 显示错误消息 --%>
        <c:if test="${not empty errorMessage}">
            <div class="error-message" style="text-align: center; margin-bottom: 15px;">
                    ${errorMessage}
            </div>
        </c:if>

        <%-- 显示成功消息 --%>
        <c:if test="${not empty successMessage}">
            <div class="success-message" style="text-align: center; margin-bottom: 15px;">
                    ${successMessage}
            </div>
        </c:if>

        <form id="login-form" action="login" method="post" ${not empty param.register ? 'style="display: none;"' : ''}>
            <div class="form-group">
                <label for="login-email">邮箱</label>
                <input type="email" id="login-email" name="email" placeholder="请输入您的邮箱" required value="${param.email}">
            </div>

            <div class="form-group">
                <label for="login-password">密码</label>
                <input type="password" id="login-password" name="password" placeholder="请输入您的密码" required>
            </div>

            <button type="submit" class="btn">登录</button>

            <%--            <div class="forgot-password">--%>
            <%--                <a href="forgotPassword.jsp">忘记密码?</a>--%>
            <%--            </div>--%>

        </form>

        <form id="register-form" action="register" method="post" ${empty param.register ? 'style="display: none;"' : ''}>
            <div class="form-group">
                <label for="register-name">用户名</label>
                <input type="text" id="register-name" name="username" placeholder="请输入您的用户名" required value="${param.username}">
            </div>

            <div class="form-group">
                <label for="register-email">邮箱</label>
                <input type="email" id="register-email" name="email" placeholder="请输入您的邮箱" required value="${param.email}">
            </div>

            <div class="form-group">
                <label for="register-password">密码</label>
                <input type="password" id="register-password" name="password" placeholder="请输入至少6位密码" required>
            </div>

            <div class="form-group">
                <label for="register-confirm-password">确认密码</label>
                <input type="password" id="register-confirm-password" name="confirmPassword" placeholder="请再次输入密码" required>
            </div>

            <button type="submit" class="btn">注册</button>

        </form>
    </div>
</div>

<!-- JQuery -->
<script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>

<!-- Font Awesome 图标库 -->
<script src="https://kit.fontawesome.com/a076d05399.js" crossorigin="anonymous"></script>

<script>
    $(document).ready(function() {
        // 根据URL参数决定显示哪个表单
        const urlParams = new URLSearchParams(window.location.search);
        const registerParam = urlParams.get('register');

        // 切换登录/注册表单
        $('#login-toggle').click(function() {
            $(this).addClass('active');
            $('#register-toggle').removeClass('active');
            $('#login-form').show();
            $('#register-form').hide();
            $('.header h1').text('欢迎回来');
            $('.header p').text('请登录您的账户');
            // 更新URL参数
            window.history.pushState({}, '', window.location.pathname);
        });

        $('#register-toggle').click(function() {
            $(this).addClass('active');
            $('#login-toggle').removeClass('active');
            $('#login-form').hide();
            $('#register-form').show();
            $('.header h1').text('创建账户');
            $('.header p').text('注册新用户');
            // 更新URL参数
            window.history.pushState({}, '', window.location.pathname + '?register=true');
        });

        // 注册表单验证
        $('#register-form').submit(function(e) {
            const password = $('#register-password').val();
            const confirmPassword = $('#register-confirm-password').val();

            if (password !== confirmPassword) {
                e.preventDefault();
                alert('两次输入的密码不一致');
                return false;
            }

            if (password.length < 6) {
                e.preventDefault();
                alert('密码长度至少为6位');
                return false;
            }

            return true;
        });
    });
</script>
</body>
</html>