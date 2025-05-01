#define _CRT_SECURE_NO_WARNINGS
#include <iostream>
#include <string>
using namespace std;

string classifyTask(string taskDescription) {
    string lowerDesc = taskDescription;
    for (auto& c : lowerDesc) {
        c = tolower(c);
    }
    if (lowerDesc.find("meeting") != string::npos || lowerDesc.find("report") != string::npos) {
        return "work";
    }
    else if (lowerDesc.find("homework") != string::npos || lowerDesc.find("thesis") != string::npos) {
        return "study";
    }
    else if (lowerDesc.find("shopping") != string::npos || lowerDesc.find("trip") != string::npos) {
        return "life";
    }
    return "other";
}

int main() {
    string taskDescription;
    getline(cin, taskDescription); // 从输入获取任务描述
    cout << classifyTask(taskDescription) << endl; // 输出分类结果
    return 0;
}