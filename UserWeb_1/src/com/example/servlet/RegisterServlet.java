package com.example.servlet;

import com.example.Dao.UserDAO;
import com.example.user.User;

import java.io.IOException;
import jakarta.servlet.ServletException;
import jakarta.servlet.annotation.WebServlet;
import jakarta.servlet.http.HttpServlet;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import jakarta.servlet.http.HttpSession;

@WebServlet("/register")
public class RegisterServlet extends HttpServlet {
    private static final long serialVersionUID = 1L;

    protected void doPost(HttpServletRequest request, HttpServletResponse response)
            throws ServletException, IOException {

        String username = request.getParameter("username");
        String email = request.getParameter("email");
        String password = request.getParameter("password");
        String confirmPassword = request.getParameter("confirmPassword");

        // 验证密码是否匹配
        if (!password.equals(confirmPassword)) {
            request.setAttribute("errorMessage", "两次输入的密码不一致");
            request.setAttribute("username", username);
            request.setAttribute("email", email);
            request.getRequestDispatcher("User.jsp?register=true").forward(request, response);
            return;
        }

        // 验证密码长度
        if (password.length() < 6) {
            request.setAttribute("errorMessage", "密码长度至少为6位");
            request.setAttribute("username", username);
            request.setAttribute("email", email);
            request.getRequestDispatcher("User.jsp?register=true").forward(request, response);
            return;
        }

        UserDAO userDao = new UserDAO();

        // 检查邮箱是否已存在
        if (userDao.emailExists(email)) {
            request.setAttribute("errorMessage", "该邮箱已被注册");
            request.setAttribute("username", username);
            request.setAttribute("email", email);
            request.getRequestDispatcher("User.jsp?register=true").forward(request, response);
            return;
        }

        // 创建新用户
        User newUser = new User(username, email, password);
        boolean isRegistered = userDao.registerUser(newUser);

        if (isRegistered) {
            request.setAttribute("successMessage", "注册成功，请登录");
            request.getRequestDispatcher("/User.jsp").forward(request, response);
//            response.sendRedirect("User.jsp");
        } else {
            request.setAttribute("errorMessage", "注册失败，请稍后再试");
            request.setAttribute("username", username);
            request.setAttribute("email", email);
            request.getRequestDispatcher("User.jsp?register=true").forward(request, response);
        }
    }
}
