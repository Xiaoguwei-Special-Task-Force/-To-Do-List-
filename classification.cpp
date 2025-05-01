#define _CRT_SECURE_NO_WARNINGS
#include <iostream>
#include <string>
using namespace std;

string classifyTask(string taskDescription) {
    // 转小写方便统一判断
    string lowerDesc = taskDescription;
    for (auto& c : lowerDesc) {
        c = tolower(c);
    }
    // 简单关键词判断
    if (lowerDesc.find("meeting") != string::npos) {
        return "work";
    }
    else if (lowerDesc.find("homework") != string::npos) {
        return "study";
    }
    return "other";
}

int main() {
    string task = "Prepare for meeting";
    cout << "Task: " << task << " Classification: " << classifyTask(task) << endl;
    return 0;
}