<script>
  import { CreateTask, DeleteTask, GetTasks, UpdateTask } from '../wailsjs/go/internal/App.js';
  import { writable } from 'svelte/store';
  import { onMount } from 'svelte';

  // dynamic store of tasks
  const tasks = writable([]);
  let priority = "medium";
  let newTask = ""; 
  let statusFilter = "all";

  // Reactive filtered tasks
  $: filteredTasks = $tasks.filter(task => {
    if (statusFilter === "all") return true;
    if (statusFilter === "done") return task.isCompleted;
    if (statusFilter === "todo") return !task.isCompleted;
  });

  onMount(async () => {
  const fetchedTasks = await GetTasks();
  // Map IsComplete to completed
  const mappedTasks = fetchedTasks.map(task => ({
    ...task,
    completed: task.isCompleted
  }));
  tasks.set(mappedTasks);
});

  const toggleCompletion = (task) => {
    const currentTask = $tasks.find(t => t.id === task.id);
    if (!currentTask) return;
    
    // updating local state complete/incomplete
    tasks.update(currentTasks => 
      currentTasks.map(t => 
        t.id === task.id ? { ...t, isCompleted: !currentTask.isCompleted } : t
      )
    );
    
    // updating task state in backend
    UpdateTask(task.id);
  };

  const generateUniqueID = () => {
    return Date.now().toString();
  };

  // adding task func (back to your original approach)
  const addTask = async () => {
    if (newTask.trim()) {  // input validation
      const task = {id: generateUniqueID(), title: newTask.trim(), priority, isCompleted: false};

      // Optimistic UI update: Add the task to the list immediately
      tasks.update(currentTasks => [...currentTasks, task]);

      // Send the new task to the backend for saving
      await CreateTask({ id: task.id, title: task.title, priority: task.priority });

      newTask = "";  // Clear the input field
    }
  };

  const deleteTask = async (taskId) => {
    try {
      await DeleteTask(taskId);
      tasks.update(currentTasks => currentTasks.filter(t => t.id !== taskId));
    } catch (error) {
      console.error('Failed to delete task:', error);
    }
  };
</script>

<main>
  <h1>Task List</h1>

  <!-- Input form to add a task -->
  <form on:submit|preventDefault={addTask}>
    <label>
      Task
      <input name="newTask" bind:value={newTask} placeholder="Enter task name" />
    </label>

    <label>
      Priority
      <select bind:value={priority}>
        <option value="high">High</option>
        <option value="medium">Medium</option>
        <option value="low">Low</option>
      </select>
    </label>

    <button type="submit">Add Task</button>
  </form>

  <!-- Filtering Buttons -->
  <div>
    <button class:active={statusFilter === "all"} on:click={() => statusFilter = "all"}>
      All ({$tasks.length})
    </button>
    <button class:active={statusFilter === "todo"} on:click={() => statusFilter = "todo"}>
      To Do ({$tasks.filter(t => !t.isCompleted).length})
    </button>
    <button class:active={statusFilter === "done"} on:click={() => statusFilter = "done"}>
      Done ({$tasks.filter(t => t.isCompleted).length})
    </button>
  </div>

  <!-- Display filtered tasks -->
  <ul>
    {#each filteredTasks as task (task.id)}
      <li
      class:completed={task.isCompleted}
      class={task.priority}
      style="background-color: {task.priority}"
      >
        <span class="task-content">
          {task.title}      ({task.priority})
        </span>
        <div class="task-actions">
          <button on:click={() => toggleCompletion(task)}>
            {task.isCompleted ? "Mark Incomplete" : "Mark Complete"}
          </button>
          <button class="delete" on:click={() => deleteTask(task.id)}>Delete</button>
        </div>
      </li>
    {/each}
  </ul>

  {#if filteredTasks.length === 0}
    <p>No tasks found for the current filter.</p>
  {/if}
</main>

<style>
  main {
    max-width: 600px;
    margin: 0 auto;
    padding: 20px;
  }

  .low {
    background-color: rgba(156, 202, 96, 0.623)
  }

  .medium {
    background-color: rgba(219, 190, 95, 0.781)
  }

  .high {
    background-color: rgba(230, 125, 118, 0.781)
  }

  form {
    display: flex;
    gap: 10px;
    margin-bottom: 20px;
    align-items: end;
  }

  label {
    display: flex;
    flex-direction: column;
    font-weight: bold;
  }

  input, select {
    margin-top: 5px;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }

  button {
    padding: 8px 16px;
    background-color: #f0ae22;
    color: #333;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.2s;
    font-weight: 500;
  }

  button:hover {
    background-color: #e09a1e;
  }

  button.active {
    background-color: #333;
    color: white;
  }

  button.delete {
    background-color: #e74c3c;
    color: white;
  }

  button.delete:hover {
    background-color: #746043;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin: 10px 0;
    padding: 10px;
    border: 1px solid #eee;
    border-radius: 4px;
    background-color: #f9f9f9;
    color: #333;
  }

  li.completed {
    opacity: 0.7;
    background-color: #8b8b8a;
  }

  li.completed .task-content {
    text-decoration: line-through;
    color: #666;
  }

  .task-content {
    color: #333;
    font-weight: 500;
  }

  .task-actions {
    display: flex;
    gap: 5px;
  }

  .task-actions button {
    font-size: 12px;
    padding: 4px 8px;
  }

  .status {
    font-style: italic;
    font-size: 0.9em;
    margin-left: 8px;
  }
</style>