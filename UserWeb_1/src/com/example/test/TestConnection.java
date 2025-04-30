package com.example.test;


import com.example.util.DBUtil;
import java.sql.Connection;

public class TestConnection {
    public static void main(String[] args) {
        try {
            Connection conn = DBUtil.getConnection();
            System.out.println("数据库连接成功！");
            conn.close();
        } catch (Exception e) {
            System.out.println("数据库连接失败：");
            e.printStackTrace();
        }
    }
}
