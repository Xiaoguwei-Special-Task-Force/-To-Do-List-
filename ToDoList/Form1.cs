using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace ToDoList
{
    public class Task
    {
        public string Name { get; set; } = "";
        public DateTime Deadline { get; set; } = DateTime.Now;
        public string Description { get; set; } = "";

        public override string ToString()
        {
            return $"{Name} ({Deadline:yyyy-MM-dd})";
        }
    }

    public partial class Form1 : Form
    {
        private List<Task> tasks = new List<Task>();
        private ListBox taskListBox = new ListBox();
        private TextBox taskTextBox = new TextBox();
        private Button addButton = new Button { Text = "添加" };
        private Button deleteButton = new Button { Text = "删除" };

        public Form1()
        {
            InitializeComponent();
            InitializeUI();
            addButton.Click += AddButton_Click;
            deleteButton.Click += DeleteButton_Click;
        }

        private void InitializeComponent()
        {
            this.ClientSize = new System.Drawing.Size(400, 300);
        this.MinimumSize = new System.Drawing.Size(400, 300);
            this.Text = "ToDo List";
        }

        private DateTimePicker deadlinePicker;
        private TextBox descriptionTextBox;

        private void InitializeUI()
        {
            // 任务列表
            taskListBox = new ListBox();
            taskListBox.Location = new Point(10, 10);
            taskListBox.Size = new Size(380, 150);
            taskListBox.Anchor = AnchorStyles.Top | AnchorStyles.Bottom | AnchorStyles.Left | AnchorStyles.Right;

            // 时间选择器
            deadlinePicker = new DateTimePicker();
            deadlinePicker.Location = new Point(10, 170);
            deadlinePicker.Size = new Size(150, 20);
            deadlinePicker.Anchor = AnchorStyles.Top | AnchorStyles.Left | AnchorStyles.Right;

            // 描述输入框
            descriptionTextBox = new TextBox();
            descriptionTextBox.Location = new Point(170, 170);
            descriptionTextBox.Size = new Size(220, 20);
            descriptionTextBox.Anchor = AnchorStyles.Top | AnchorStyles.Left | AnchorStyles.Right;
            descriptionTextBox.PlaceholderText = "任务描述";

            // 原有控件调整位置
            taskTextBox.Location = new Point(10, 200);
            taskTextBox.Anchor = AnchorStyles.Bottom | AnchorStyles.Left | AnchorStyles.Right;
            addButton.Location = new Point(230, 200);
            addButton.Size = new Size(80, 30);
            addButton.Anchor = AnchorStyles.Bottom | AnchorStyles.Right;
            deleteButton.Location = new Point(320, 200);
            deleteButton.Size = new Size(80, 30);
            deleteButton.Anchor = AnchorStyles.Bottom | AnchorStyles.Right;

            this.Controls.AddRange(new Control[] {
                taskListBox,
                deadlinePicker,
                descriptionTextBox,
                taskTextBox,
                addButton,
                deleteButton
            });
        }

        private void AddButton_Click(object sender, EventArgs e)
        {
            var newTask = new Task {
                Name = taskTextBox.Text.Trim(),
                Deadline = deadlinePicker.Value,
                Description = descriptionTextBox.Text.Trim()
            };

            if (!string.IsNullOrEmpty(newTask.Name))
            {
                tasks.Add(newTask);
                taskListBox.Items.Add(newTask);
                taskTextBox.Clear();
                descriptionTextBox.Clear();
            }
        }

        private void DeleteButton_Click(object sender, EventArgs e)
        {
            if (taskListBox.SelectedIndex != -1)
            {
                tasks.RemoveAt(taskListBox.SelectedIndex);
                taskListBox.Items.RemoveAt(taskListBox.SelectedIndex);
            }
        }
    }
}