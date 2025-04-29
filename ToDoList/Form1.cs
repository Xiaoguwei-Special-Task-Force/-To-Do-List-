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
        // 将 ListBox 替换为 ListView
        private ListView taskListView = new ListView();
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
            this.ClientSize = new System.Drawing.Size(800, 600);
            this.MinimumSize = new System.Drawing.Size(800, 600);
            this.Text = "ToDo List";
        }

        private DateTimePicker deadlinePicker;
        private TextBox descriptionTextBox;

        private void InitializeUI()
        {
            // 左侧任务列表区域
            int leftPanelWidth = this.ClientSize.Width / 2 - 10;
            taskListView = new ListView();
            taskListView.Location = new Point(10, 10);
            taskListView.Size = new Size(leftPanelWidth, this.ClientSize.Height - 20);
            taskListView.Anchor = AnchorStyles.Top | AnchorStyles.Bottom | AnchorStyles.Left;
            taskListView.View = View.Details;
            taskListView.FullRowSelect = true;
            
            // 添加列标题
            taskListView.Columns.Add("任务名称", 120);
            taskListView.Columns.Add("任务描述", 200);
            taskListView.Columns.Add("截止时间", 120);

            // 右侧控件区域
            int rightPanelX = leftPanelWidth + 20;
            int rightPanelWidth = this.ClientSize.Width - leftPanelWidth - 30;
            int verticalSpacing = 10;
            int currentY = 10;

            // 任务名称输入框
            taskTextBox.Location = new Point(rightPanelX, currentY);
            taskTextBox.Size = new Size(rightPanelWidth, 20);
            taskTextBox.Anchor = AnchorStyles.Top | AnchorStyles.Right;
            currentY += taskTextBox.Height + verticalSpacing;

            // 截止时间选择器
            deadlinePicker = new DateTimePicker();
            deadlinePicker.Location = new Point(rightPanelX, currentY);
            deadlinePicker.Size = new Size(rightPanelWidth, 20);
            deadlinePicker.Anchor = AnchorStyles.Top | AnchorStyles.Right;
            currentY += deadlinePicker.Height + verticalSpacing;

            // 任务描述输入框
            descriptionTextBox = new TextBox();
            descriptionTextBox.Location = new Point(rightPanelX, currentY);
            descriptionTextBox.Size = new Size(rightPanelWidth, 100);
            descriptionTextBox.Anchor = AnchorStyles.Top | AnchorStyles.Right;
            descriptionTextBox.Multiline = true;
            currentY += descriptionTextBox.Height + verticalSpacing;

            // 添加按钮
            addButton.Location = new Point(rightPanelX, currentY);
            addButton.Size = new Size(rightPanelWidth / 2 - 5, 20);
            addButton.Anchor = AnchorStyles.Top | AnchorStyles.Right;

            // 删除按钮
            deleteButton.Location = new Point(rightPanelX + addButton.Width + 10, currentY);
            deleteButton.Size = new Size(rightPanelWidth / 2 - 5, 20);
            deleteButton.Anchor = AnchorStyles.Top | AnchorStyles.Right;
            
            this.Controls.AddRange(new Control[] {
                taskListView,
                taskTextBox,
                deadlinePicker,
                descriptionTextBox,
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
                // 添加任务到 ListView
                ListViewItem item = new ListViewItem(newTask.Name);
                item.SubItems.Add(newTask.Description);
                item.SubItems.Add(newTask.Deadline.ToString("yyyy-MM-dd"));
                taskListView.Items.Add(item);
                
                taskTextBox.Clear();
                descriptionTextBox.Clear();
            }
        }

        private void DeleteButton_Click(object sender, EventArgs e)
        {
            if (taskListView.SelectedItems.Count > 0)
            {
                int selectedIndex = taskListView.SelectedIndices[0];
                tasks.RemoveAt(selectedIndex);
                taskListView.Items.RemoveAt(selectedIndex);
            }
        }
    }
}