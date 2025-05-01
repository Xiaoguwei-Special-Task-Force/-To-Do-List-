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
    string task1 = "Finish the report";
    string task2 = "Do homework";
    string task3 = "Plan a trip";
    cout << "Task: " << task1 << " Classification: " << classifyTask(task1) << endl;
    cout << "Task: " << task2 << " Classification: " << classifyTask(task2) << endl;
    cout << "Task: " << task3 << " Classification: " << classifyTask(task3) << endl;
    return 0;
}