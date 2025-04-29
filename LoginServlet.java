package com.example.servlet;

import com.example.Dao.UserDAO;
import com.example.user.User;

import java.io.IOException;
import jakarta.servlet.ServletException;
import jakarta.servlet.annotation.WebServlet;
import jakarta.servlet.http.*;

@WebServlet("/login")
public class LoginServlet extends HttpServlet {
    private static final long serialVersionUID = 1L;


    protected void doPost(HttpServletRequest request, HttpServletResponse response)
            throws ServletException, IOException {

        String email = request.getParameter("email");
        String password = request.getParameter("password");

        UserDAO userDao = new UserDAO();
        User user = userDao.loginUser(email, password);

        if (user != null) {
            // 登录成功，创建session
            HttpSession session = request.getSession();
            session.setAttribute("user", user);

            // 设置 Cookie（用于前端检查）
            Cookie loginCookie = new Cookie("isLoggedIn", "true");
            loginCookie.setMaxAge(24 * 60 * 60); // 1天有效期
            loginCookie.setPath("/"); // 整个域名可用
            response.addCookie(loginCookie);

            // 重定向到主页或用户仪表板
            response.sendRedirect("task.html");
        } else {
            // 登录失败，返回错误信息
            request.setAttribute("errorMessage", "邮箱或密码错误");
            request.getRequestDispatcher("User.jsp").forward(request, response);
        }
    }
}
