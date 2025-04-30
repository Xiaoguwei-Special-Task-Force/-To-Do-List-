from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import List
from uuid import uuid4

app = FastAPI()

# 定义任务模型
class Task(BaseModel):
    id: str
    title: str
    description: str = ""
    deadline: str = ""
    status: str = "未开始"  # 默认状态

# 内存数据库（后续可以替换为真正的数据库）
tasks_db: List[Task] = []

# 路由：获取所有任务
@app.get("/tasks")
def read_tasks():
    return {"tasks": tasks_db}

# 路由：添加新任务
@app.post("/tasks")
def create_task(task: Task):
    new_task = Task(
        id=str(uuid4()),
        title=task.title,
        description=task.description,
        deadline=task.deadline,
        status=task.status
    )
    tasks_db.append(new_task)
    return {"message": "任务添加成功", "task": new_task}

# 路由：根据任务 ID 删除任务
@app.delete("/tasks/{id}")
def delete_task(id: str):
    global tasks_db
    tasks_db = [task for task in tasks_db if task.id != id]
    return {"message": f"任务 {id} 删除成功"}
