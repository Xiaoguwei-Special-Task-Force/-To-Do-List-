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
            taskListView.Size = new Size(leftPanelWidth, this.ClientSize.Height - 130);
            taskListView.Anchor = AnchorStyles.Top | AnchorStyles.Bottom | AnchorStyles.Left;
            taskListView.View = View.Details;
            taskListView.FullRowSelect = true;
            
            // 添加列标题
            taskListView.Columns.Add("任务名称", 120);
            taskListView.Columns.Add("任务描述", 200);
            taskListView.Columns.Add("截止时间", 120);

            // 右侧预留区域
            Panel rightPanel = new Panel();
            rightPanel.Location = new Point(leftPanelWidth + 20, 10);
            rightPanel.Size = new Size(this.ClientSize.Width - leftPanelWidth - 30, this.ClientSize.Height - 20);
            rightPanel.Anchor = AnchorStyles.Top | AnchorStyles.Bottom | AnchorStyles.Left | AnchorStyles.Right;
            rightPanel.BackColor = Color.LightGray;

            // 时间选择器
            deadlinePicker = new DateTimePicker();
            deadlinePicker.Location = new Point(10, taskListView.Bottom + 10);
            deadlinePicker.Size = new Size(150, 20);
            deadlinePicker.Anchor = AnchorStyles.Bottom | AnchorStyles.Left;

            // 描述输入框
            descriptionTextBox = new TextBox();
            descriptionTextBox.Location = new Point(deadlinePicker.Right + 10, taskListView.Bottom + 10);
            descriptionTextBox.Size = new Size(leftPanelWidth - deadlinePicker.Width - 20, 20);
            descriptionTextBox.Anchor = AnchorStyles.Bottom | AnchorStyles.Left;
            descriptionTextBox.PlaceholderText = "任务描述";

            // 原有控件调整位置
            taskTextBox.Location = new Point(10, descriptionTextBox.Bottom + 10);
            taskTextBox.Size = new Size(leftPanelWidth - 200, 20);
            taskTextBox.Anchor = AnchorStyles.Bottom | AnchorStyles.Left;
            addButton.Location = new Point(taskTextBox.Right + 10, descriptionTextBox.Bottom + 10);
            addButton.Size = new Size(80, 20);
            addButton.Anchor = AnchorStyles.Bottom | AnchorStyles.Left;
            deleteButton.Location = new Point(addButton.Right + 10, descriptionTextBox.Bottom + 10);
            deleteButton.Size = new Size(80, 20);
            deleteButton.Anchor = AnchorStyles.Bottom | AnchorStyles.Left;

            this.Controls.AddRange(new Control[] {
                taskListView,
                rightPanel,
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